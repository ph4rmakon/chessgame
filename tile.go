package main

type Color int

type Tile struct {
	Color Color
	Piece Piece
	Coord Coordinate
}

const (
	BLACK Color = iota
	WHITE
)
