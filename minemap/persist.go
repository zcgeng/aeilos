package minemap

import (
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"
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

func (p *Persister) GetScore(user string) string {
	return p.get("[score]" + user)
}

func (p *Persister) AddScore(user string, score int) {
	p.incrby("[score]"+user, score)
}

func (p *Persister) RecordByIP(ip string, value string) {
	p.lpush("[ip]"+ip, value)
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

func (p *Persister) set(key, value string) {
	_, err := p.conn.Do("SET", key, value)
	if err != nil {
		panic(err)
	}
}

func (p *Persister) get(key string) string {
	val, err := redis.String(p.conn.Do("GET", key))
	if err == redis.ErrNil {
		// fmt.Println("received (nil)")
		return "(nil)"
	} else if err != nil {
		panic(err)
	} else {
		return val
	}
	return ""
}
