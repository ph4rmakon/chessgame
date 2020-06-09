package main

var board Board
var game ChessGame

func main() {
	for {
		game.Restart()
		for !game.IsFinished() {
			game.TakeTurn()
		}
	}
}
