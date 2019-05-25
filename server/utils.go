package server

import (
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
