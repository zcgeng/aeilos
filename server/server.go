package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/zcgeng/aeilos/minemap"
	"github.com/zcgeng/aeilos/pb"
)

// MineServer ...
type MineServer struct {
	mmap     *minemap.MineMap
	clients  map[*websocket.Conn]bool
	upgrader websocket.Upgrader
}

// NewMineServer ...
func NewMineServer() *MineServer {
	ms := new(MineServer)
	ms.mmap = minemap.NewMineMap()
	ms.mmap.ShowBlock(2, 8)
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
	http.HandleFunc("/ws", s.handleConnections)

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

func (s *MineServer) handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	fmt.Printf("on connection\n")

	// Register our new client
	s.clients[ws] = true

	for {
		var msg pb.ClientToServer // Read in a new message as pb and map it to a Message object
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

		// Send the newly received message to mine engine
		s.mmap.CCommand <- &msg
	}
}

func (s *MineServer) handleResponses() {
	for {
		rmsg := <-s.mmap.CReply
		data, err := proto.Marshal(rmsg)
		if err != nil {
			log.Fatalf("Marshal error: %v", err)
			return
		}
		// fmt.Printf("broadcasting: [%v]\n", rmsg)
		s.bcastMsg(data)
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
