// Package storage
package storage

import (
	"fmt"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/kolakdd/bship/entity"
)

type Storage struct {
	TokenPlayerMu sync.Mutex
	TokenPlayer   map[string]*entity.Player // ws_token -> player

	ClientsWs    map[string]*websocket.Conn // client_id -> ws
	ClientsPairs map[string]string          // client_id -> client_id
	// session storage
	Sessions      map[string]*entity.Session // session_id -> session
	InviteSession map[string]*entity.Session // invite_code -> session

	ClientSesions map[*entity.Player]*entity.Session // player_id -> session

	ClientsMu  sync.Mutex
	SessionsMu sync.Mutex
}

func New() *Storage {
	return &Storage{
		TokenPlayer: make(map[string]*entity.Player),

		ClientsWs:    make(map[string]*websocket.Conn),
		ClientsPairs: make(map[string]string),
		// session storage
		Sessions:      make(map[string]*entity.Session),
		InviteSession: make(map[string]*entity.Session),
		ClientSesions: make(map[*entity.Player]*entity.Session),
	}
}

func (s *Storage) AddTokenPlayer(player *entity.Player) {
	s.TokenPlayerMu.Lock()
	defer s.TokenPlayerMu.Unlock()
	s.TokenPlayer[player.WsToken] = player
}

func (s *Storage) CreateSession(player *entity.Player) *entity.Session {
	s.SessionsMu.Lock()
	s.ClientsMu.Lock()
	defer s.ClientsMu.Unlock()
	defer s.SessionsMu.Unlock()

	session := entity.InitGame(player)

	s.ClientSesions[player] = &session
	s.InviteSession[session.InviteCode] = &session
	s.Sessions[session.ID] = &session

	return &session
}

// JoinToSession join second player room
func (s *Storage) JoinToSession(player *entity.Player, inviteCode string) error {
	s.SessionsMu.Lock()
	s.ClientsMu.Lock()

	defer s.ClientsMu.Unlock()
	defer s.SessionsMu.Unlock()

	session, exist := s.InviteSession[inviteCode]
	if !exist {
		return fmt.Errorf("session not found by inviteCode, = %s", inviteCode)
	}
	session.PrintInfo()

	session.RightPlayer = player
	s.ClientsPairs[session.LeftPlayer.ID] = player.ID
	s.ClientsPairs[player.ID] = session.LeftPlayer.ID
	s.ClientSesions[player] = session
	fmt.Println("pair = ", s.ClientsPairs)

	// add players pair
	return nil
}

func (s *Storage) PrintInfo() {
	s.SessionsMu.Lock()
	s.ClientsMu.Lock()
	s.TokenPlayerMu.Lock()
	defer s.TokenPlayerMu.Unlock()
	defer s.ClientsMu.Unlock()
	defer s.SessionsMu.Unlock()
	fmt.Println("storage.MainStorage.TokenPlayer", s.TokenPlayer)
	fmt.Println("storage.MainStorage.ClientsWs", s.ClientsWs)
	fmt.Println("storage.MainStorage.ClientsPairs", s.ClientsPairs)
	fmt.Println("storage.MainStorage.Sessions", s.Sessions)
	fmt.Println("storage.MainStorage.InviteSession", s.InviteSession)

}

func (s *Storage) GetPlayer(token string) (*entity.Player, bool) {
	s.SessionsMu.Lock()
	s.ClientsMu.Lock()
	s.TokenPlayerMu.Lock()
	defer s.TokenPlayerMu.Unlock()
	defer s.ClientsMu.Unlock()
	defer s.SessionsMu.Unlock()

	player, exist := s.TokenPlayer[token]
	return player, exist
}

func (s *Storage) GetAllTokens() []string {
	s.SessionsMu.Lock()
	s.ClientsMu.Lock()
	s.TokenPlayerMu.Lock()
	defer s.TokenPlayerMu.Unlock()
	defer s.ClientsMu.Unlock()
	defer s.SessionsMu.Unlock()

	keys := make([]string, 0, len(s.TokenPlayer))
	for k := range s.TokenPlayer {
		keys = append(keys, k)
	}
	return keys
}
