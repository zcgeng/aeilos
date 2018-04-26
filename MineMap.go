package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	bombRate = 25
)

const (
	hidden = iota
	show
	flag
)

func block2area(x, y int) (int, int) {
	ax := x / 100
	ay := y / 100
	if x < 0 {
		ax = (x+1)/100 - 1
	}
	if y < 0 {
		ay = (y+1)/100 - 1
	}
	return ax, ay
}

func area2block(x, y int) (int, int) {
	return x * 100, y * 100
}

// MineBlock ...
type MineBlock struct {
	value  uint8 // 9 as a bumb, 11 as unknown(border condition)
	status uint8
	user   string
}

// MineArea ...
type MineArea struct {
	x      int
	y      int
	blocks [100][100]MineBlock
	// m      *MineMap
}

// MineMap ...
type MineMap struct {
	areas map[string]*MineArea
}

// GetBlock2 ...
func (area *MineArea) GetBlock2(x, y int) *MineBlock {
	if x > 99 || x < 0 || y > 99 || y < 0 {
		panic("out of area")
	}
	return &area.blocks[x][y]
}
func (area *MineArea) calcBombs(x, y int) uint8 {
	if x >= 99 || x <= 0 || y >= 99 || y <= 0 {
		panic("out of area")
	}

	// TODO: use DP instead of this stupid method
	count := uint8(0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if area.GetBlock2(x+j, y+i).value == 9 {
				count++
			}
		}
	}
	return uint8(count)
}

// rate is an integer in 0..100, which is the bomb rate out of 100
func (area *MineArea) shuffleArea(rate int) {
	// generate the bombs
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if rand.Intn(100)+1 <= rate {
				area.GetBlock2(i, j).value = 9
			}
		}
	}
	// generate numbers
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ { // x : j, y : i
			if area.GetBlock2(j, i).value == 9 {
				continue
			} else if i == 0 || i == 99 || j == 0 || j == 99 {
				area.GetBlock2(j, i).value = 11
			} else {
				area.GetBlock2(j, i).value = area.calcBombs(j, i)
			}
		}
	}
}

func newMineArea(x, y int) *MineArea {
	area := new(MineArea)
	area.x = x
	area.y = y
	// area.m = m
	return area
}

func newMineMap() *MineMap {
	m := new(MineMap)
	m.areas = make(map[string]*MineArea)
	return m
}

func (area *MineArea) getBorder() (l int, r int, d int, u int) {
	x := area.x
	y := area.y
	return x * 100, x*100 + 100, y * 100, y*100 + 100
}

func (area *MineArea) checkInside(x, y int) bool {
	l, r, d, u := area.getBorder()
	return (l <= x && x < r && d <= y && y < u)
}

// GetBlock ...
func (area *MineArea) GetBlock(x, y int) *MineBlock {
	if area.checkInside(x, y) {
		ax, ay := area2block(area.x, area.y)
		ax = x - ax
		ay = y - ay
		return &area.blocks[ax][ay]
	}
	fmt.Println(x, y)
	fmt.Println(area.x, area.y)
	fmt.Println(area.getBorder())
	panic("out of area")
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
	if b.status == show {
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

func printMap(mapa *MineMap) {
	for i := -20; i < 20; i++ {
		for j := -20; j < 20; j++ {
			v := strconv.Itoa(int(mapa.GetBlock(i, j).value))
			status := mapa.GetBlock(i, j).status
			switch status {
			case hidden:
				v = "@"
			case show:
				switch v {
				case "11":
					v = "."
				case "9":
					v = "*"
				}
			case flag:
				v = "P"
			}

			fmt.Print(v, " ")
		}
		fmt.Print("\n")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	mapa := newMineMap()
	for i := 0; i < 10; i++ {
		mapa.ShowBlock(i, i)
	}
	printMap(mapa)
	return
}
