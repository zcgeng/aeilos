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
	return p.zscore("score", user)
}

func (p *Persister) AddScore(user string, score int) {
	p.zincrby("score", user, score)
}

func (p *Persister) GetRank(user string) int {
	return int(p.zrevrank("score", user)) + 1
}

func (p *Persister) GetTopScores(length int) ([]string, []int) {
	return p.zrevrange("score", 0, length)
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
	p.set("[user]"+user.Email, user.ToString())
	return true
}

func (p *Persister) GetUser(email string) *mineuser.MineUser {
	res := p.get("[user]" + email)
	if res == "(nil)" {
		return nil
	}
	user := mineuser.UnMarshalUser([]byte(res))
	return user
}

func (p *Persister) UserExists(email string) bool {
	res := p.get("[user]" + email)
	return res != "(nil)"
}

func (p *Persister) CheckAuthToken(token string) string {
	res := p.get("[token]" + token)
	if res == "(nil)" {
		return ""
	}
	return res
}

func (p *Persister) SetAuthToken(token, email string) {
	p.set("[token]"+token, email)
}

func (p *Persister) ClearAuthToken(token string) {
	p.del("[token]" + token)
}

// -------------------- method abstractions ----------------

func (p *Persister) zincrby(setname string, key string, value int) {
	_, err := p.conn.Do("ZINCRBY", setname, value, key)
	if err != nil {
		panic(err)
	}
}

// return 0 if nil
func (p *Persister) zscore(setname string, key string) int64 {
	val, err := redis.String(p.conn.Do("ZSCORE", setname, key))
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

func (p *Persister) zrevrange(setname string, start int, stop int) ([]string, []int) {
	val, err := redis.Strings(p.conn.Do("ZREVRANGE", setname, start, stop, "WITHSCORES"))
	if err != nil {
		panic(err)
	}

	users := make([]string, 0)
	scores := make([]int, 0)
	for i, str := range val {
		if i%2 == 0 {
			// username
			users = append(users, str)
		} else {
			// score
			scoreInt, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				log.Fatalf("parseInt failed: %v, %v\n", str, err)
			}
			scores = append(scores, int(scoreInt))
		}
	}
	return users, scores
}

func (p *Persister) zrevrank(setname string, key string) int64 {
	val, err := redis.Int(p.conn.Do("ZREVRANK", setname, key))
	if err == redis.ErrNil {
		return 0
	} else if err != nil {
		panic(err)
	}
	return int64(val)
}

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

// return 0 if nil
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

func (p *Persister) del(key string) bool {
	val, err := redis.Int(p.conn.Do("del", key))
	if err != nil {
		panic(err)
	}
	return val == 1
}
