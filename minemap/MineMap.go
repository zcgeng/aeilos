package minemap

import (
	"fmt"
	"strconv"

	pb "github.com/zcgeng/aeilos/pb"
)

const (
	bombRate = 30
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

// ShowBlock ...
func (m *MineMap) ShowBlock(x, y int) {
	b := m.GetBlock(x, y)
	if b.status != hidden {
		return
	}

	if b.value == 11 {
		b.value = m.calcBombs(x, y)
	}

	b.status = show
	if b.value == 0 {
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				m.ShowBlock(x+i, y+j)
			}
		}
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

// return -1: already has owner; 0 : success; 1..8 : wrong flag
func (m *MineMap) putFlag(x, y int, user string) int {
	b := m.GetBlock(x, y)
	if b.status != hidden {
		return -1
	}

	if b.value == 11 {
		b.value = m.calcBombs(x, y)
	}

	switch b.value {
	case 9:
		b.status = flag
		b.user = user
		return 0

	default:
		b.status = show
		b.user = user
		return int(b.value)
	}
}

// PrintMap ..
func PrintMap(mapa *MineMap) {
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

			fmt.Print(v, " ")
		}
		fmt.Print("\n")
	}
}

func (m *MineMap) operationLoop() {
	fmt.Println("MineMap: operation loop begin")
	for {
		cmd := <-m.CCommand
		fmt.Printf("received command: %v\n", cmd)
		m.CReply <- &pb.ServerToClient{Msg: "hello world"}
		for i := 0; i < 100; i++ {
			m.CReply <- &pb.ServerToClient{Msg: "hello world " + strconv.Itoa(i)}
		}
	}
}

func (m *MineMap) run() {
	go m.operationLoop()
}
