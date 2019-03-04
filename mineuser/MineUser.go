package mineuser

import (
	"encoding/json"
)

type MineUser struct {
	UserName string
	Password string
	FlagURL  string
	Email    string
	Score    int
}

func (u *MineUser) ToString() string {
	return string(u.Marshal())
}

func (u *MineUser) Marshal() []byte {
	res, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return res
}

func UnMarshalUser(b []byte) *MineUser {
	u := &MineUser{}
	err := json.Unmarshal(b, u)
	if err != nil {
		panic(err)
	}
	return u
}
