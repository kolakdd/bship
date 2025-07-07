// Package websocket
package websocket

import (
	"fmt"
	"log"
	"log/slog"
	"strings"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/kolakdd/bship/storage"
)

type WsHandler struct {
	storage *storage.Storage
}

func NewWsHandler(storage *storage.Storage) *WsHandler {
	return &WsHandler{storage: storage}
}

func (h *WsHandler) GameHandler(c *websocket.Conn) {
	token := c.Params("token")
	player, exist := h.storage.GetPlayer(token)
	if !exist {
		slog.Info("Player not found by wsToken", "token", token)
		return
	}

	registerWsCon(h.storage, player.ID, c)

	// wait con with peer
	var peerConn *websocket.Conn
	for okIn := false; !okIn; {
		time.Sleep(time.Second)
		slog.Debug("in loop", "player.ID", player.ID)
		pairID, exist := h.storage.ClientsPairs[player.ID]
		if !exist {
			slog.Debug("CLIENT PAIR NOT EXIST in loop", "player.ID", player.ID)
			continue
		}
		con, err := getWs(h.storage, pairID)
		if err != nil {
			slog.Debug("WS NOT EXIST in loop", "player.ID", player.ID)
			continue
		}
		peerConn = con
		okIn = true
	}

	// close con, delete session and all links
	defer func() {
		h.storage.ClientsMu.Lock()
		delete(h.storage.ClientsWs, player.ID)
		if peerID, ok := h.storage.ClientsPairs[player.ID]; ok {
			delete(h.storage.ClientsPairs, peerID)
			delete(h.storage.ClientsPairs, player.ID)
		}
		// TODO: add close all storage
		h.storage.ClientsMu.Unlock()
	}()

	for {
		mt, msg, err := c.ReadMessage()
		strMsg := string(msg)
		fmt.Println(strMsg)

		if err != nil {
			log.Println("read:", err)
			break
		}

		if strings.HasPrefix(string(msg), "/place_ship") {
			fmt.Println("place_ship worked")
			continue
		}
		if strings.HasPrefix(string(msg), "/ping") {
			fmt.Println("place_ship worked")
			continue
		}

		h.storage.ClientsMu.Lock()
		if err := peerConn.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
		}

		h.storage.ClientsMu.Unlock()

	}
}

func registerWsCon(storage *storage.Storage, id string, c *websocket.Conn) {
	storage.ClientsMu.Lock()
	storage.SessionsMu.Lock()
	defer storage.ClientsMu.Unlock()
	defer storage.SessionsMu.Unlock()

	storage.ClientsWs[id] = c
}

func getWs(storage *storage.Storage, id string) (*websocket.Conn, error) {
	storage.ClientsMu.Lock()
	storage.SessionsMu.Lock()
	defer storage.ClientsMu.Unlock()
	defer storage.SessionsMu.Unlock()

	res, ok := storage.ClientsWs[id]
	if !ok {
		return nil, fmt.Errorf("ws con not found ")
	}
	return res, nil
}
