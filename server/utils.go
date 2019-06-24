package server

import (
	"math/rand"
	"time"

	"github.com/zcgeng/aeilos/pb"
)

func genServerMessage(msg string) *pb.ServerToClient {
	rpl := &pb.ServerToClient{Response: &pb.ServerToClient_Msg{Msg: &pb.ChatMsg{
		Msg:      msg,
		UserName: "System",
		NickName: "System",
		Time:     time.Now().Unix(),
	}}}

	return rpl
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandString gives you a random string which has a size of n
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
