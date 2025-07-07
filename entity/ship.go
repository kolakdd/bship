package entity

// 1 корабль — ряд из 4 клеток («четырёхпалубный»; линкор)
// 2 корабля — ряд из 3 клеток («трёхпалубные»; 	крейсера)
// 3 корабля — ряд из 2 клеток («двухпалубные»; 	эсминцы)
// 4 корабля — 1 клетка («однопалубные»; торпедные катера)

type Ship struct {
	Position ShipPosition
	Type     ShipType
}

func (sh Ship) Size() int {
	switch sh.Type {
	case linkor:
		return 4
	case cruiser:
		return 3
	case esminets:
		return 2
	case boat:
		return 1
	}
	panic("No ship size!")
}

type ShipPosition struct {
	Horizontal bool
	StartLat   int
	StartLon   int
	EndLat     int
	EndLon     int
}

type ShipType int

const (
	linkor ShipType = iota
	cruiser
	esminets
	boat
)
