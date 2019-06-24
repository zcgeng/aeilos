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
	clients   map[*websocket.Conn]string // map Connections to emails
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
	ms.clients = make(map[*websocket.Conn]string)
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

	// handle user register
	http.HandleFunc("/aeilos/register", s.handleRegister)

	// start a file server
	fs := http.FileServer(http.Dir("www/"))
	http.Handle("/aeilos/", http.StripPrefix("/aeilos/", fs))

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
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Printf("on register: %v, %v\n", email, username)

	u := &mineuser.MineUser{}
	u.Email = email
	u.UserName = username
	u.Password = password

	res := s.persister.NewUser(u)
	if res {
		fmt.Fprintf(w, "Success!\nPlease go back to login")
	} else {
		fmt.Fprintf(w, "Failed: email already exists\n")
	}
}

func (s *MineServer) handleGetStatsRequest(email string) *pb.ServerToClient {
	return s.handleGetStats(email)
}

func (s *MineServer) handleGetStats(email string) *pb.ServerToClient {
	fmt.Printf("user: %v\n", s.persister.GetUser(email))
	user := s.persister.GetUser(email)
	if user == nil {
		panic("user doesn't exist in the db: " + email)
	}
	stats := &pb.Stats{
		UserName: email,
		NickName: s.persister.GetUser(email).UserName,
		Score:    s.persister.GetScore(email),
	}

	return &pb.ServerToClient{Response: &pb.ServerToClient_Stats{Stats: stats}}
}

func (s *MineServer) handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	mytoken, _ := r.Cookie("aeilos_token")
	myemail := s.persister.CheckAuthToken(mytoken.Value)
	fmt.Println(mytoken.Value)
	fmt.Println(myemail)

	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	record := "{connect: " + time.Now().Local().String() + "}"
	fmt.Printf("on connection: %v %v\n",
		strings.Split(ws.RemoteAddr().String(), ":")[0], record)

	// Register our new client, empty string means un-logined
	s.clients[ws] = myemail
	if myemail != "" {
		reply := &minemap.MMapToServer{
			Reply:  s.handleGetStatsRequest(myemail),
			Client: ws,
			Bcast:  false,
		}
		s.mmap.CReply <- reply
	}

	s.persister.RecordByIP(strings.Split(ws.RemoteAddr().String(), ":")[0], record)

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

			if s.clients[ws] == "" || s.clients[ws] != v.ChatMsg.UserName { // check login status
				reply := &minemap.MMapToServer{
					Reply:  genServerMessage("Please login first"),
					Client: ws,
					Bcast:  true,
				}
				s.mmap.CReply <- reply
				break
			}
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
			if s.clients[ws] == "" || s.clients[ws] != v.GetStats.UserName { // check login status
				reply := &minemap.MMapToServer{
					Reply:  genServerMessage("Please login first"),
					Client: ws,
					Bcast:  true,
				}
				s.mmap.CReply <- reply
				break
			}
			reply := &minemap.MMapToServer{
				Reply:  s.handleGetStatsRequest(v.GetStats.GetUserName()),
				Client: ws,
				Bcast:  false,
			}
			s.mmap.CReply <- reply

		case *pb.ClientToServer_GetChatHistory:
			fmt.Printf("received GetChatHistory request: %v\n", v)
			// send the message history
			msgs := s.persister.GetChatMsg(0, -1)
			for _, msg := range msgs {
				rpl := &pb.ServerToClient{Response: &pb.ServerToClient_Msg{Msg: msg}}
				reply := &minemap.MMapToServer{
					Reply:  rpl,
					Client: ws,
					Bcast:  false,
				}
				s.mmap.CReply <- reply
			}

		case *pb.ClientToServer_Login:
			fmt.Printf("received login: %v\n", v.Login)
			user := s.persister.GetUser(v.Login.Email)

			if user == nil || v.Login.Password != user.Password {
				rpl := &pb.ServerToClient{Response: &pb.ServerToClient_LoginResult{LoginResult: &pb.LoginResult{
					Success: false,
					Msg:     "User doesn't exist or wrong password",
					Token:   "",
				}}}
				reply := &minemap.MMapToServer{
					Reply:  rpl,
					Client: ws,
					Bcast:  false,
				}
				s.mmap.CReply <- reply
				break
			}

			s.clients[ws] = v.Login.Email // record this connection as logged in
			token := RandString(32)
			s.persister.SetAuthToken(token, v.Login.Email)

			// reply login success message
			rpl := &pb.ServerToClient{Response: &pb.ServerToClient_LoginResult{LoginResult: &pb.LoginResult{
				Success: true,
				Msg:     "Login Success!",
				Token:   token,
			}}}
			reply := &minemap.MMapToServer{
				Reply:  rpl,
				Client: ws,
				Bcast:  false,
			}
			s.mmap.CReply <- reply

			// send back user stats
			reply = &minemap.MMapToServer{
				Reply:  s.handleGetStats(user.Email),
				Client: ws,
				Bcast:  false,
			}
			s.mmap.CReply <- reply

		case *pb.ClientToServer_Logout:
			if s.clients[ws] == "" { // check login status
				reply := &minemap.MMapToServer{
					Reply:  genServerMessage("you have not logged in"),
					Client: ws,
					Bcast:  true,
				}
				s.mmap.CReply <- reply
				break
			}

			s.clients[ws] = ""
			s.persister.ClearAuthToken(v.Logout.Token)
			reply := &minemap.MMapToServer{
				Reply:  &pb.ServerToClient{Response: &pb.ServerToClient_LogoutResult{LogoutResult: &pb.Empty{}}},
				Client: ws,
				Bcast:  false,
			}
			s.mmap.CReply <- reply

		case *pb.ClientToServer_Touch:
			if s.clients[ws] == "" || s.clients[ws] != v.Touch.GetUser() { // check login status
				reply := &minemap.MMapToServer{
					Reply:  genServerMessage("Please login first"),
					Client: ws,
					Bcast:  true,
				}
				s.mmap.CReply <- reply
				break
			}
			// hand it over to minemap thread
			cmd := &minemap.ServerToMMap{
				Cmd:    &msg,
				Client: ws,
			}
			s.mmap.CCommand <- cmd

		case *pb.ClientToServer_GetLeaderBoard:
			// fmt.Printf("received getLeaderBoard\n")

			ranklist := make([]*pb.RankInfo, 0)
			// get top 10 players
			names, scores := s.persister.GetTopScores(10)
			for i, name := range names {
				ranklist = append(ranklist, &pb.RankInfo{
					NickName: s.persister.GetUser(name).UserName,
					Score:    int64(scores[i]),
					Rank:     int64(i + 1),
				})
			}
			// add the user's rank to the end of the list
			if s.clients[ws] != "" {
				myemail := s.clients[ws]
				ranklist = append(ranklist, &pb.RankInfo{
					NickName: s.persister.GetUser(myemail).UserName,
					Score:    int64(s.persister.GetScore(myemail)),
					Rank:     int64(s.persister.GetRank(myemail)),
				})
			}

			// generate reply
			rpl := &pb.ServerToClient{Response: &pb.ServerToClient_LeaderBoard{LeaderBoard: &pb.LeaderBoard{
				Ranklist: ranklist,
			}}}
			reply := &minemap.MMapToServer{
				Reply:  rpl,
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
