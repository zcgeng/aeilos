package minemap

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	pb "github.com/zcgeng/aeilos/pb"
)

const (
	bombRate = 20
)

//
const (
	ShowBlock = iota
	PutFlag
)

// MineMap ...
type MineMap struct {
	areas    map[string]*MineArea
	CCommand chan *ServerToMMap
	CReply   chan *MMapToServer

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

// NewMineMap ...
func NewMineMap(persis *Persister) *MineMap {
	m := new(MineMap)
	m.areas = make(map[string]*MineArea)
	m.CCommand = make(chan *ServerToMMap, 1000)
	m.CReply = make(chan *MMapToServer, 1000)
	m.persister = persis

	m.run()
	return m
}

// GetArea ...
func (m *MineMap) GetArea(ax, ay int) *MineArea {
	res, ok := m.areas[GetKey(ax, ay)]
	if !ok {
		fmt.Printf("load area (%v,%v)\n", ax, ay)
		area := m.persister.LoadArea(GetKey(ax, ay))
		if area == nil {
			fmt.Printf("new area (%v,%v)\n", ax, ay)
			area = newMineArea(ax, ay)
			area.shuffleArea(bombRate)
		}
		m.putArea(area)
		return area
	}
	return res
}

// PutArea ...
func (m *MineMap) putArea(area *MineArea) {
	m.areas[GetKey(area.X, area.Y)] = area
}

// persist the cache entry (and remove it from memory)
func (m *MineMap) AreaEntryWriteBack(key string, keepInCache bool) {
	res, ok := m.areas[key]
	if !ok {
		return
	}
	m.persister.PersistArea(res)
	if !keepInCache {
		delete(m.areas, key)
	}
}

func (m *MineMap) PersistAreaCache(keeyInCache bool) {
	fmt.Printf("persist cache: currently have %v areas in memory\n", len(m.areas))
	for _, area := range m.areas {
		m.persister.PersistArea(area)
	}
	// clear the cache
	if !keeyInCache {
		m.areas = make(map[string]*MineArea)
	}
}

// GetBlock ...
func (m *MineMap) GetBlock(x, y int) *MineBlock {
	ax, ay := block2area(x, y)
	area := m.GetArea(ax, ay)
	if area.X != ax || area.Y != ay {
		panic("got wrong area")
	}
	res := area.GetBlock(x, y)
	return res
}

// PutBlock ...
func (m *MineMap) PutBlock(x, y int, b *MineBlock) {
	ax, ay := block2area(x, y)
	block := m.GetArea(ax, ay).GetBlock(x, y)
	*block = *b
}

// starts from a zero point, DFS all zeros and broadcast
// returns cumulative scores: 0 for zeros, 1 for each numbered cell
func (m *MineMap) ExploreZeros(x, y int) int {
	score := 0
	b := m.GetBlock(x, y)

	if b.Status != hidden {
		return score
	}

	if b.Value == 11 {
		b.Value = m.calcBombs(x, y)
	}

	b.Status = show
	if b.Value != 0 {
		score++
	}

	reply := &MMapToServer{
		Reply: &pb.ServerToClient{
			Response: &pb.ServerToClient_Update{
				Update: m.getCellPB(int64(x), int64(y)),
			},
		},
		Client: nil, // TODO: use pb.UpdateZeros
		Bcast:  true,
	}
	m.CReply <- reply

	if b.Value == 0 {
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				score += m.ExploreZeros(x+i, y+j)
			}
		}
	}
	return score
}

// ShowBlock returns the score that the player got
func (m *MineMap) ShowBlock(x, y int) int {
	b := m.GetBlock(x, y)
	if b.Status != hidden {
		return 0
	}

	if b.Value == 11 {
		b.Value = m.calcBombs(x, y)
	}

	score := 0
	if b.Value == 0 {
		score += m.ExploreZeros(x, y)
	}
	b.Status = show
	switch b.Value {
	case 0:
		return score
	case 9:
		return -50
	case 11:
		fmt.Println("impossible!! code:24u89kejnw9")
		return 0
	default:
		return score + 1
	}
}

// return the score
func (m *MineMap) putFlag(x, y int, user string) int {
	b := m.GetBlock(x, y)
	if b.Status != hidden {
		return 0
	}

	if b.Value == 11 {
		b.Value = m.calcBombs(x, y)
	}

	switch b.Value {
	case 9:
		b.Status = flag
		b.User = user
		return 10

	default:
		b.Status = show
		b.User = user
		return -1
	}
}

func (m *MineMap) calcBombs(x, y int) uint8 {
	// TODO: use DP instead of this stupid method
	count := uint8(0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if m.GetBlock(x+j, y+i).Value == 9 {
				count++
			}
		}
	}
	return uint8(count)
}

// PrintMap ..
func PrintMap(mapa *MineMap) string {
	s := ""
	for i := -20; i < 20; i++ {
		for j := -20; j < 20; j++ {
			v := strconv.Itoa(int(mapa.GetBlock(i, j).Value))
			status := mapa.GetBlock(i, j).Status
			switch status {
			case hidden:
				v = "+"
			case show:
				switch v {
				case "11":
					v = "."
				case "9":
					v = "*"
				case "0":
					v = " "
				}
			case flag:
				v = "P"
			}

			if j == -20 {
				s += fmt.Sprintf("%v", v)
			} else {
				s += fmt.Sprintf(" %v", v)
			}
		}
		if i != 19 {
			s += fmt.Sprintf("\n")
		}

	}
	return s
}

func (m *MineMap) getCellPB(x, y int64) *pb.Cell {
	block := m.GetBlock(int(x), int(y))
	ret := &pb.Cell{}
	ret.X = x
	ret.Y = y
	switch block.Status {
	case hidden:
		ret.CellType = &pb.Cell_UnTouched{UnTouched: true}
	case show:
		ret.CellType = &pb.Cell_Bombs{Bombs: int32(block.Value)}
	case flag:
		ret.CellType = &pb.Cell_FlagURL{FlagURL: ""}
	default:
		fmt.Printf("ERROR: non existing cell type \n")
		ret.CellType = nil
	}
	return ret
}

func (m *MineMap) handleTouchRequest(v *pb.ClientToServer_Touch) *pb.ServerToClient {
	var score int
	if v.Touch.GetTouchType() == pb.TouchType_FLAG {
		score = m.putFlag(int(v.Touch.GetX()), int(v.Touch.GetY()), "")
	} else if v.Touch.GetTouchType() == pb.TouchType_FLIP {
		score = m.ShowBlock(int(v.Touch.GetX()), int(v.Touch.GetY()))
	}

	resp := &pb.ServerToClient_Touch{Touch: &pb.TouchResponse{
		Score: int32(score),
		Cell:  m.getCellPB(v.Touch.GetX(), v.Touch.GetY()),
	}}

	return &pb.ServerToClient{Response: resp}
}

func (m *MineMap) handleGetAreaRequest(v *pb.ClientToServer_GetArea) *pb.ServerToClient {
	area := &pb.Area{
		X:     v.GetArea.GetX(),
		Y:     v.GetArea.GetY(),
		Cells: make([]*pb.Cell, 0),
	}

	for xx := int64(0); xx < ROW_HEIGHT; xx++ {
		for yy := int64(0); yy < ROW_LENGTH; yy++ {
			area.Cells = append(area.Cells, m.getCellPB(xx+v.GetArea.GetX(), yy+v.GetArea.GetY()))
		}
	}

	return &pb.ServerToClient{Response: &pb.ServerToClient_Area{Area: area}}
}

func (m *MineMap) operationLoop() {
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
				reply.Reply = m.handleTouchRequest(v)
				reply.Bcast = true
				m.persister.AddScore(strings.Split(msg.Client.RemoteAddr().String(), ":")[0], int(reply.Reply.GetTouch().GetScore()))

			case *pb.ClientToServer_GetArea:
				fmt.Printf("received GetArea request: %v\n", v)
				reply.Reply = m.handleGetAreaRequest(v)
				reply.Bcast = false

			default:
				fmt.Printf("wrong type of request: %v\n", v)
			}
			m.CReply <- reply

		case <-m.persister.pTicker.C:
			m.PersistAreaCache(true)
		}
	}
}

func (m *MineMap) run() {
	go m.operationLoop()
}
