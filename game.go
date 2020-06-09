package main

import (
	"math/rand"
)

type ChessGame struct {
	Players       [2]Player
	Sets          [2]Set
	CurrentPlayer int
}

func (g *ChessGame) InitGame() {
	g.CurrentPlayer = rand.Intn(1<<32) % 2
}

func (g *ChessGame) Restart() {
	g.InitGame()
}

func (g *ChessGame) IsFinished() bool {
	return false
}

func (g *ChessGame) TakeTurn() {
	// TODO: implement what a turn does...
	g.CurrentPlayer ^= 1 // toggle the bit
}
