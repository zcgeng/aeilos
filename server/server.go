package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zcgeng/aeilos/minemap"
)

// Message ...
type Message struct {
	x         int
	y         int
	user      string
	operation string
}

// ReplyMsg ...
type ReplyMsg struct {
	Success bool
	X       int
	Y       int
	Value   uint8
	Status  string
	User    string
}

// MineServer ...
type MineServer struct {
	mmap     *minemap.MineMap
	clients  map[*websocket.Conn]bool
	upgrader websocket.Upgrader
	cmsgs    chan Message
}

// NewMineServer ...
func NewMineServer() *MineServer {
	ms := new(MineServer)
	ms.mmap = minemap.NewMineMap()
	ms.clients = make(map[*websocket.Conn]bool)
	ms.cmsgs = make(chan Message, 100)
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

	// Start listening for incoming chat messages
	go s.handleMessages()

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

	// Register our new client
	s.clients[ws] = true

	for {
		var msg Message // Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(s.clients, ws)
			break
		}
		// Send the newly received message to message handler
		s.cmsgs <- msg
	}
}

func (s *MineServer) handleMessages() {
	for {
		// Grab the next message from the messages channel
		msg := <-s.cmsgs
		fmt.Println("received message: ", msg)
		// Send it out to every client that is currently connected
		switch msg.operation {
		case "clickBlock":
			s.mmap.CCommand <- minemap.Command{msg.x, msg.y, msg.user, minemap.ShowBlock}
			<-s.mmap.CReply
			s.bcastMsg(ReplyMsg{})
		case "putFlag":
			s.mmap.CCommand <- minemap.Command{msg.x, msg.y, msg.user, minemap.PutFlag}
			<-s.mmap.CReply
			s.bcastMsg(ReplyMsg{})
		case "getArea":
			return
		default:
			return
		}

	}
}

func (s *MineServer) bcastMsg(rmsg ReplyMsg) {
	for client := range s.clients {
		err := client.WriteJSON(rmsg)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(s.clients, client)
		}
	}
}
