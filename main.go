package main

import (
	"math/rand"
	"time"

	"github.com/zcgeng/aeilos/server"
)

func main() {
	rand.Seed(time.Now().Unix())
	s := server.NewMineServer()
	s.Start()
	return
}
