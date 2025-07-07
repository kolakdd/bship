package entity

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"
)

type Session struct {
	ID         string
	InviteCode string

	LeftPlayer  *Player
	RightPlayer *Player

	LeftMap  BattleMap
	RightMap BattleMap

	LeftTurnToShoot bool
}

func InitGame(player1 *Player) Session {
	id := generateGameID()
	inviteCode := generateInviteCode()

	return Session{
		ID:              id,
		InviteCode:      inviteCode,
		LeftPlayer:      player1,
		LeftTurnToShoot: true,
	}
}

func (s Session) JoinPlayer(player2 Player) Session {
	return s
}

func generateGameID() string {
	source := rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano()))
	r := rand.New(source)
	randomID := r.IntN(100)
	return strconv.Itoa(randomID)
}

func generateInviteCode() string {
	source := rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano()))
	r := rand.New(source)
	randomCode := r.IntN(100)
	return strconv.Itoa(randomCode)
}

func (s Session) PrintInfo() {
	fmt.Println(s.ID)
	fmt.Println(s.InviteCode)
	fmt.Println(s.LeftPlayer)
	fmt.Println(s.RightMap)

}
