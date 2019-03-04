package minemap

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
	"github.com/zcgeng/aeilos/mineuser"
	"github.com/zcgeng/aeilos/pb"
)

// https://itnext.io/storing-go-structs-in-redis-using-rejson-dab7f8fc0053
// https://github.com/nitishm/rejson-struct

const (
	ClearCacheTimeSeconds = 600
)

type Persister struct {
	conn    redis.Conn
	pTicker *time.Ticker
}

func NewPersister(addr, passwd string) *Persister {
	conn, err := redis.Dial("tcp", addr, redis.DialPassword(passwd))
	if err != nil {
		panic(err)
	}

	return &Persister{
		conn:    conn,
		pTicker: time.NewTicker(ClearCacheTimeSeconds * time.Second),
	}
}

func (p *Persister) close() {
	p.conn.Close()
}

func (p *Persister) PersistArea(area *MineArea) {
	b, err := json.Marshal(*area)
	if err != nil {
		panic(err)
	}
	p.set(area.GetKey(), string(b))
}

func (p *Persister) LoadArea(key string) *MineArea {
	objStr := p.get(key)
	if objStr == "(nil)" {
		return nil
	}
	b := []byte(objStr)
	area := &MineArea{}
	err := json.Unmarshal(b, area)
	if err != nil {
		return nil
	}
	return area
}

func (p *Persister) GetScore(user string) int64 {
	return p.getInt64("[score]" + user)
}

func (p *Persister) AddScore(user string, score int) {
	p.incrby("[score]"+user, score)
}

func (p *Persister) RecordByIP(ip string, value string) {
	p.lpush("[ip]"+ip, value)
}

func (p *Persister) RecordChatMsg(v *pb.ChatMsg) {
	marshaled, err := proto.Marshal(v)
	if err != nil {
		return
	}
	p.rpush("[chatMsg]", string(marshaled))
}

func (p *Persister) GetChatMsg(start, stop int) []*pb.ChatMsg {
	res := make([]*pb.ChatMsg, 0)
	msgs := p.lrange("[chatMsg]", start, stop)
	for _, data := range msgs {
		var msg pb.ChatMsg
		err := proto.Unmarshal([]byte(data), &msg)
		if err != nil {
			log.Printf("unmarshal error: %v", err)
			break
		}
		res = append(res, &msg)
	}
	return res
}

func (p *Persister) NewUser(user *mineuser.MineUser) bool {
	if p.UserExists(user.Email) {
		return false
	}
	p.set(user.Email, "[user]"+user.ToString())
	return true
}

func (p *Persister) GetUser(email string) *mineuser.MineUser {
	res := p.get(email)
	return mineuser.UnMarshalUser([]byte(res))
}

func (p *Persister) UserExists(email string) bool {
	res := p.get(email)
	return res != "(nil)"
}

// -------------------- method abstractions ----------------

func (p *Persister) incrby(key string, value int) {
	// fmt.Printf("incrby %v %v\n", key, value)
	_, err := p.conn.Do("INCRBY", key, value)
	if err != nil {
		panic(err)
	}
}

func (p *Persister) lpush(key, value string) {
	_, err := p.conn.Do("LPUSH", key, value)
	if err != nil {
		panic(err)
	}
}

func (p *Persister) rpush(key, value string) {
	_, err := p.conn.Do("RPUSH", key, value)
	if err != nil {
		panic(err)
	}
}

func (p *Persister) lrange(key string, start, stop int) []string {
	val, err := redis.Strings(p.conn.Do("LRANGE", key, start, stop))
	if err != nil {
		panic(err)
	}
	return val
}

func (p *Persister) set(key, value string) {
	_, err := p.conn.Do("SET", key, value)
	if err != nil {
		panic(err)
	}
}

func (p *Persister) getInt64(key string) int64 {
	val, err := redis.String(p.conn.Do("GET", key))
	if err == redis.ErrNil {
		return 0
	} else if err != nil {
		panic(err)
	}

	valInt, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Fatalf("parseInt failed: %v, %v\n", val, err)
	}
	return valInt

}

func (p *Persister) get(key string) string {
	val, err := redis.String(p.conn.Do("GET", key))
	if err == redis.ErrNil {
		// fmt.Println("received (nil)")
		return "(nil)"
	} else if err != nil {
		panic(err)
	}
	return val
}
