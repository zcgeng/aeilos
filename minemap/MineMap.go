package minemap

import (
	"fmt"
	"strconv"

	pb "github.com/zcgeng/aeilos/pb"
)

const (
	bombRate = 15
)

//
const (
	ShowBlock = iota
	PutFlag
)

// MineMap ...
type MineMap struct {
	areas    map[string]*MineArea
	CCommand chan *pb.ClientToServer
	CReply   chan *pb.ServerToClient
}

// NewMineMap ...
func NewMineMap() *MineMap {
	m := new(MineMap)
	m.areas = make(map[string]*MineArea)
	m.CCommand = make(chan *pb.ClientToServer, 100)
	m.CReply = make(chan *pb.ServerToClient, 100)
	m.run()
	return m
}

// GetArea ...
func (m *MineMap) GetArea(ax, ay int) *MineArea {
	res, ok := m.areas[strconv.Itoa(ax)+","+strconv.Itoa(ay)]
	if !ok {
		area := newMineArea(ax, ay)
		area.shuffleArea(bombRate)
		m.PutArea(area)
		return area
	}
	return res
}

// PutArea ...
func (m *MineMap) PutArea(area *MineArea) {
	m.areas[strconv.Itoa(area.x)+","+strconv.Itoa(area.y)] = area
}

// GetBlock ...
func (m *MineMap) GetBlock(x, y int) *MineBlock {
	ax, ay := block2area(x, y)
	area := m.GetArea(ax, ay)
	res := area.GetBlock(x, y)
	return res
}

// PutBlock ...
func (m *MineMap) PutBlock(x, y int, b *MineBlock) {
	ax, ay := block2area(x, y)
	block := m.GetArea(ax, ay).GetBlock(x, y)
	*block = *b
}

// start from a zero point, dfs all zeros and broadcast
// reture cumulative scores: 0 for zeros, 1 for each numbered cell
func (m *MineMap) ExploreZeros(x, y int) int {
	score := 0
	b := m.GetBlock(x, y)

	if b.status != hidden {
		return score
	}

	if b.value == 11 {
		b.value = m.calcBombs(x, y)
	}

	b.status = show
	m.CReply <- &pb.ServerToClient{Response: &pb.ServerToClient_Update{
		Update: m.getCellPB(int64(x), int64(y)),
	}}
	if b.value == 0 {
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				score += m.ExploreZeros(x+i, y+j)
			}
		}
	}
	return score + 1
}

// ShowBlock returns the score that the player got
func (m *MineMap) ShowBlock(x, y int) int {
	b := m.GetBlock(x, y)
	if b.status != hidden {
		return 0
	}

	if b.value == 11 {
		b.value = m.calcBombs(x, y)
	}

	score := 0
	if b.value == 0 {
		score += m.ExploreZeros(x, y)
	}
	b.status = show
	switch b.value {
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
	if b.status != hidden {
		return 0
	}

	if b.value == 11 {
		b.value = m.calcBombs(x, y)
	}

	switch b.value {
	case 9:
		b.status = flag
		b.user = user
		return 1

	default:
		b.status = show
		b.user = user
		return -1
	}
}

func (m *MineMap) calcBombs(x, y int) uint8 {
	// TODO: use DP instead of this stupid method
	count := uint8(0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if m.GetBlock(x+j, y+i).value == 9 {
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
			v := strconv.Itoa(int(mapa.GetBlock(i, j).value))
			status := mapa.GetBlock(i, j).status
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
	switch block.status {
	case hidden:
		ret.CellType = &pb.Cell_UnTouched{UnTouched: true}
	case show:
		ret.CellType = &pb.Cell_Bombs{Bombs: int32(block.value)}
	case flag:
		ret.CellType = &pb.Cell_FlagURL{FlagURL: ""}
	default:
		fmt.Printf("ERROR: non existing cell type \n")
		ret.CellType = nil
	}
	return ret
}

func (m *MineMap) handleTouchRequest(v *pb.ClientToServer_Touch) {
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

	m.CReply <- &pb.ServerToClient{Response: resp}
}

func (m *MineMap) handleGetAreaRequest(v *pb.ClientToServer_GetArea) {
	area := &pb.Area{
		X:     v.GetArea.GetX(),
		Y:     v.GetArea.GetY(),
		Cells: make([]*pb.Cell, 0),
	}

	for xx := int64(0); xx < 10; xx++ {
		for yy := int64(0); yy < 10; yy++ {
			area.Cells = append(area.Cells, m.getCellPB(xx+v.GetArea.GetX(), yy+v.GetArea.GetY()))
		}
	}

	m.CReply <- &pb.ServerToClient{Response: &pb.ServerToClient_Area{Area: area}}
}

func (m *MineMap) operationLoop() {
	fmt.Println("MineMap: operation loop begins")
	for {
		cmd := <-m.CCommand

		switch v := cmd.GetRequest().(type) {

		case *pb.ClientToServer_Touch:
			fmt.Printf("received Touch request: %v\n", v)
			m.handleTouchRequest(v)

		case *pb.ClientToServer_GetArea:
			fmt.Printf("received GetArea request: %v\n", v)
			m.handleGetAreaRequest(v)

		default:
			fmt.Printf("wrong type of request: %v\n", v)
		}

	}
}

func (m *MineMap) run() {
	go m.operationLoop()
}
