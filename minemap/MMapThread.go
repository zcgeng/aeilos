package minemap

import (
	"fmt"
	"os"

	"github.com/gorilla/websocket"
	pb "github.com/zcgeng/aeilos/pb"
)

type MMapThread struct {
	mine *MineMap

	CCommand  chan *ServerToMMap
	CReply    chan *MMapToServer
	persister *Persister
}

type ServerToMMap struct {
	Cmd    *pb.ClientToServer
	Client *websocket.Conn
}

type MMapToServer struct {
	Reply  *pb.ServerToClient
	Client *websocket.Conn
	Bcast  bool
}

func NewMMapThread() *MMapThread {
	m := &MMapThread{}

	m.mine = NewMineMap()
	m.persister = NewPersister(
		os.Getenv("REDIS_ADDRESS"),
		os.Getenv("REDIS_PASSWORD"),
	)
	m.CCommand = make(chan *ServerToMMap, 1000)
	m.CReply = make(chan *MMapToServer, 1000)
	m.run()

	return m
}

func (m *MMapThread) handleTouchRequest(v *pb.ClientToServer_Touch) []*pb.ServerToClient {
	res := make([]*pb.ServerToClient, 0)

	score := 0
	if v.Touch.GetTouchType() == pb.TouchType_FLAG {
		score = m.mine.putFlag(int(v.Touch.GetX()), int(v.Touch.GetY()), "")
	} else if v.Touch.GetTouchType() == pb.TouchType_FLIP {
		score1, updates := m.mine.ShowBlock(int(v.Touch.GetX()), int(v.Touch.GetY()))
		score += score1
		res = append(res, updates...)
	}

	resp := &pb.ServerToClient_Touch{Touch: &pb.TouchResponse{
		Score: int32(score),
		Cell:  m.mine.getCellPB(v.Touch.GetX(), v.Touch.GetY()),
	}}
	res = append(res, &pb.ServerToClient{Response: resp})

	return res
}

func (m *MMapThread) handleGetAreaRequest(v *pb.ClientToServer_GetArea) *pb.ServerToClient {
	area := &pb.Area{
		X:     v.GetArea.GetX(),
		Y:     v.GetArea.GetY(),
		Cells: make([]*pb.Cell, 0),
	}

	for xx := int64(0); xx < ROW_HEIGHT; xx++ {
		for yy := int64(0); yy < ROW_LENGTH; yy++ {
			area.Cells = append(area.Cells, m.mine.getCellPB(xx+v.GetArea.GetX(), yy+v.GetArea.GetY()))
		}
	}

	return &pb.ServerToClient{Response: &pb.ServerToClient_Area{Area: area}}
}

func (m *MMapThread) operationLoop() {
	fmt.Println("MineMap: operation loop begins")
	for {
		select {
		case msg := <-m.CCommand:
			cmd := msg.Cmd

			reply := &MMapToServer{
				Reply:  nil,
				Client: msg.Client,
				Bcast:  false,
			}

			switch v := cmd.GetRequest().(type) {

			case *pb.ClientToServer_Touch:
				fmt.Printf("received Touch request: %v\n", v)
				replies := m.handleTouchRequest(v)
				for _, repl := range replies {
					m2s := &MMapToServer{
						Reply:  repl,
						Client: nil,
						Bcast:  true,
					}
					touch, ok := m2s.Reply.Response.(*pb.ServerToClient_Touch)
					if ok {
						m.persister.AddScore(v.Touch.GetUser(), int(touch.Touch.GetScore()))
					}
					m.CReply <- m2s
				}

			case *pb.ClientToServer_GetArea:
				// fmt.Printf("received GetArea request: %v\n", v)
				reply.Reply = m.handleGetAreaRequest(v)
				reply.Bcast = false
				m.CReply <- reply

			default:
				fmt.Printf("wrong type of request: %v\n", v)
			}

		case <-m.persister.pTicker.C:
			m.mine.PersistAreaCache(false)
		}
	}
}

func (m *MMapThread) run() {
	go m.operationLoop()
}
