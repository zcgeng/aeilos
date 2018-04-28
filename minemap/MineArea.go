package minemap

import (
	"fmt"
	"math/rand"
)

// MineArea ...
type MineArea struct {
	x      int
	y      int
	blocks [100][100]MineBlock
	// m      *MineMap
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
