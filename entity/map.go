// Package entity
package entity

import (
	"fmt"
	"log/slog"
	"strconv"
)

// Коды полей
// public data
const (
	water    = 0
	hitwater = 1
)

// private data
const (
	aliveLinkor   = 7 // 7 - линкор
	aliveCruiser  = 6 // 6 - крейсер
	aliveEsminets = 5 // 5 - эсминец
	aliveBoat     = 4 // 4 - катер
)

// public data
const (
	hitLinkor   = 17 // 17 - попадание в линкор
	hitCruiser  = 16 // 16 - попадание в крейсер
	hitEsminets = 15 // 15 - попадание в эсминец
	hitBoat     = 14 // 14 - попадание в катер
)

// public data
const (
	destroyedLinkor   = 27 // 27 - уничтожение линкора
	destroyedCruiser  = 26 // 26 - уничтожение крейсера
	destroyedEsminets = 25 // 25 - уничтожение эсминца
	destroyedBoat     = 24 // 24 - уничтожение катера
)

// представление поля
// ╲ 0 1 2 3 4 5 6 7 8 9
// 0 _ _ _ _ _ _ _ _ _ _
// 1 _ _ _ _ _ _ _ _ _ _
// 2 _ _ _ _ _ _ _ _ _ _
// 3 _ _ _ _ _ _ _ _ _ _
// 4 _ _ _ _ _ _ _ _ _ _
// 5 _ _ _ _ _ _ _ _ _ _
// 6 _ _ _ _ _ _ _ _ _ _
// 7 _ _ _ _ _ _ _ _ _ _
// 8 _ _ _ _ _ _ _ _ _ _
// 9 _ _ _ _ _ _ _ _ _ _

type MapField [10][10]int

const (
	ally MapType = iota
	enemy
	spectator
)

type MapType int

type BattleMap struct {
	Type  MapType
	Field MapField
}

func (bm BattleMap) InitBattleMap(ships [10]Ship) BattleMap {
	for _, sh := range ships {
		bm.Field.placeShip(sh)
	}
	err := bm.Field.validateInitMap()
	if err != nil {
		slog.Error("InitBattleMap error", "err=", err)
	}
	return bm
}

func (mf *MapField) placeShip(sh Ship) {
	value, err := sh.Type.shipTypeToMapCode()
	if err != nil {
		slog.Error("placeShip error", "err=", err)
	}
	if sh.Position.Horizontal {
		for j := sh.Position.StartLon; j <= sh.Position.EndLon; j++ {
			mf[sh.Position.EndLat][j] = value
		}
	} else {
		for i := sh.Position.StartLat; i <= sh.Position.EndLat; i++ {
			mf[i][sh.Position.StartLat] = value
		}

	}
}

// validate ship checksum
func (mf MapField) validateInitMap() error {
	sum := 0
	for i := range 10 {
		for j := range 10 {
			sum += mf[i][j]
		}
	}
	linkorSum := 1 * 4 * 7
	cruiserSum := 2 * 3 * 6
	esminetsSum := 3 * 2 * 5
	boatSum := 4 * 1 * 4
	if sum != linkorSum+cruiserSum+esminetsSum+boatSum {
		return fmt.Errorf("validateInitMap error")
	}
	return nil
}

func (mf MapField) String() string {
	res := ""
	for i := range 10 {
		tmp := ""
		for j := range 10 {
			tmp += strconv.Itoa(mf[i][j]) + " "
		}
		tmp += "\n"
		res += tmp
	}
	return res
}

func (st ShipType) shipTypeToMapCode() (int, error) {
	switch st {
	case linkor:
		return aliveLinkor, nil
	case cruiser:
		return aliveCruiser, nil
	case esminets:
		return aliveEsminets, nil
	case boat:
		return aliveBoat, nil
	}
	slog.Error("ShipTypeToMapCode error", "ShipType=", st)
	return 0, fmt.Errorf("validateInitMap error")
}
