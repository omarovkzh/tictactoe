package ai

import "github.com/omarov-kzh/tictactoe/internal/board"

type Move struct {
	Cell   int
	Reason string
}

func Choose(b *board.Board) Move {
	if cell, ok := completingMove(b, board.O); ok {
		return Move{Cell: cell, Reason: "win"}
	}
	if cell, ok := completingMove(b, board.X); ok {
		return Move{Cell: cell, Reason: "block"}
	}
	if b.IsFree(5) {
		return Move{Cell: 5, Reason: "center"}
	}
	for _, cell := range []int{1, 3, 7, 9} {
		if b.IsFree(cell) {
			return Move{Cell: cell, Reason: "corner"}
		}
	}
	for _, cell := range []int{2, 4, 6, 8} {
		if b.IsFree(cell) {
			return Move{Cell: cell, Reason: "side"}
		}
	}
	return Move{}
}

func completingMove(b *board.Board, mark string) (int, bool) {
	for cell := 1; cell <= len(b.Cells); cell++ {
		if !b.IsFree(cell) {
			continue
		}
		b.Set(cell, mark)
		_, win := b.CheckWin(mark)
		b.Clear(cell)
		if win {
			return cell, true
		}
	}
	return 0, false
}
