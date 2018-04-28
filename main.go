package main

import (
	"math/rand"
	"time"

	"github.com/zcgeng/aeilos/minemap"
	"github.com/zcgeng/aeilos/server"
)

func main() {
	rand.Seed(time.Now().Unix())
	mapa := minemap.NewMineMap()
	for i := -10; i < 10; i++ {
		mapa.ShowBlock(i, i)
		mapa.ShowBlock(-i, i)
	}
	minemap.PrintMap(mapa)
	server.Start()
	return
}
