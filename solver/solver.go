package main

import (
    "sarah/solver"
)

func main() {
	board := solver.CreateBoard()
	board.Solve()
	board.PrintBoard()
}