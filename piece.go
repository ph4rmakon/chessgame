package main

type PieceClass int

const (
	Pawn PieceClass = iota
	Rook
	Knight
	Bishop
	Queen
	King
)

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) IsValid() bool {
	return ((c.X > -1) && (c.Y > -1) && (c.X < 8) && (c.Y < 8))
}

type Piece struct {
	Color    Color
	Location Coordinate
	Class    PieceClass
}
