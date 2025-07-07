package entity

import (
	"math/rand/v2"
	"strconv"
	"time"
)

type Player struct {
	ID             string
	WsToken        string
	Name           string
	Ships          [10]Ship
	Prepared       bool
	CreatedAt      time.Time
	LastActivityAt time.Time
	Pair           string
	ItLeftPlayer   bool
}

func NewPlayer(name string) Player {

	return Player{generateID(), generateToken(), name, [10]Ship{}, false, time.Now(), time.Now(), "", false}
}

func generateID() string {
	source := rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano()))
	r := rand.New(source)
	randomID := r.IntN(100)
	return strconv.Itoa(randomID)
}

func generateToken() string {
	source := rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano()))
	r := rand.New(source)
	randomToken := r.IntN(100)
	return strconv.Itoa(randomToken)
}
