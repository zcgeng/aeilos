package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/zcgeng/aeilos/minemap"
	"github.com/zcgeng/aeilos/mineuser"
	"github.com/zcgeng/aeilos/pb"
)

// MineServer ...
type MineServer struct {
	mmap      *minemap.MMapThread
	clients   map[*websocket.Conn]bool
	upgrader  websocket.Upgrader
	persister *minemap.Persister
}

// NewMineServer ...
func NewMineServer() *MineServer {
	ms := new(MineServer)
	ms.persister = minemap.NewPersister(
		os.Getenv("REDIS_ADDRESS"),
		os.Getenv("REDIS_PASSWORD"),
	)
	ms.mmap = minemap.NewMMapThread()
	ms.clients = make(map[*websocket.Conn]bool)
	ms.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return ms
}

// Start ..
func (s *MineServer) Start() {
	// Configure websocket route
	http.HandleFunc("/ws/", s.handleConnections)

	// handle user register or login
	http.HandleFunc("/aeilos/register", s.handleRegister)

	// start a file server
	fs := http.FileServer(http.Dir("www/"))
	http.Handle("/", http.StripPrefix("/", fs))

	// start a thread to response to clients
	go s.handleResponses()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (s *MineServer) handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", 405)
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	u := &mineuser.MineUser{}
	u.Email = email
	u.UserName = username
	u.Password = password

	res := s.persister.NewUser(u)
	if res {
		fmt.Fprintf(w, "Success!\n")
	} else {
		fmt.Fprintf(w, "Failed: email already exists\n")
	}
}

func (m *MineServer) handleGetStatsRequest(v *pb.ClientToServer_GetStats) *pb.ServerToClient {
	user := v.GetStats.GetUserName()
	return m.handleGetStats(user)
}

func (m *MineServer) handleGetStats(user string) *pb.ServerToClient {
	stats := &pb.Stats{
		UserName: user,
		Score:    m.persister.GetScore(user),
	}

	return &pb.ServerToClient{Response: &pb.ServerToClient_Stats{Stats: stats}}
}

func (s *MineServer) handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	record := "{login: " + time.Now().Local().String() + "}"
	fmt.Printf("on connection: %v %v\n",
		strings.Split(ws.RemoteAddr().String(), ":")[0], record)

	// Register our new client
	s.clients[ws] = true

	s.persister.RecordByIP(strings.Split(ws.RemoteAddr().String(), ":")[0], record)

	// send the message history
	msgs := s.persister.GetChatMsg(0, -1)
	for _, msg := range msgs {
		rpl := &pb.ServerToClient{Response: &pb.ServerToClient_Msg{Msg: msg}}
		data, err := proto.Marshal(rpl)
		if err != nil {
			log.Fatalf("Marshal error: %v", err)
			return
		}
		s.sendMsg(data, ws)
	}

	for {
		// Read in a new message as pb and map it to a Message object
		var msg pb.ClientToServer
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Read message error: %v", err)
			ws.Close()
			delete(s.clients, ws)
			break
		}

		// fmt.Printf("received data: %v\n", data)
		err = proto.Unmarshal(data, &msg)
		if err != nil {
			log.Printf("unmarshal error: %v", err)
			break
		}

		// deal with misc messages
		switch v := msg.GetRequest().(type) {
		case *pb.ClientToServer_ChatMsg:
			fmt.Printf("received chat message: %v\n", v.ChatMsg.Msg)
			s.persister.RecordChatMsg(v.ChatMsg)
			rpl := &pb.ServerToClient{Response: &pb.ServerToClient_Msg{Msg: v.ChatMsg}}
			reply := &minemap.MMapToServer{
				Reply:  rpl,
				Client: ws,
				Bcast:  true,
			}
			s.mmap.CReply <- reply

		case *pb.ClientToServer_GetStats:
			fmt.Printf("received GetStats request: %v\n", v)
			reply := &minemap.MMapToServer{
				Reply:  s.handleGetStatsRequest(v),
				Client: ws,
				Bcast:  false,
			}
			s.mmap.CReply <- reply

		case *pb.ClientToServer_Login:
			fmt.Printf("received login: %v\n", v.Login)
			user := s.persister.GetUser(v.Login.Email)

			if user == nil || v.Login.Password != user.Password {
				rpl := &pb.ServerToClient{Response: &pb.ServerToClient_Msg{Msg: &pb.ChatMsg{
					Msg:      "User doesn't exist or wrong password",
					UserName: "System",
					Time:     time.Now().Unix(),
				}}}
				reply := &minemap.MMapToServer{
					Reply:  rpl,
					Client: ws,
					Bcast:  false,
				}
				s.mmap.CReply <- reply
				break
			}

			rpl := &pb.ServerToClient{Response: &pb.ServerToClient_Msg{Msg: &pb.ChatMsg{
				Msg:      "Login success!",
				UserName: "System",
				Time:     time.Now().Unix(),
			}}}
			reply := &minemap.MMapToServer{
				Reply:  rpl,
				Client: ws,
				Bcast:  false,
			}
			s.mmap.CReply <- reply

			reply = &minemap.MMapToServer{
				Reply:  s.handleGetStats(user.Email),
				Client: ws,
				Bcast:  false,
			}
			s.mmap.CReply <- reply

		default:
			// Send the newly received message to mine engine
			cmd := &minemap.ServerToMMap{
				Cmd:    &msg,
				Client: ws,
			}
			s.mmap.CCommand <- cmd
		}
	}
}

func (s *MineServer) handleResponses() {
	for {
		rmsg := <-s.mmap.CReply
		data, err := proto.Marshal(rmsg.Reply)
		if err != nil {
			log.Fatalf("Marshal error: %v", err)
			return
		}
		// fmt.Printf("broadcasting: [%v]\n", rmsg)
		if rmsg.Bcast {
			s.bcastMsg(data)
		} else {
			s.sendMsg(data, rmsg.Client)
		}
	}
}

func (s *MineServer) sendMsg(data []byte, client *websocket.Conn) {
	err := client.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		log.Printf("error: %v", err)
		client.Close()
		delete(s.clients, client)
	}
}

func (s *MineServer) bcastMsg(data []byte) {
	for client := range s.clients {
		err := client.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(s.clients, client)
		}
	}
}
