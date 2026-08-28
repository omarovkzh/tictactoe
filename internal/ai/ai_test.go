package ai

import (
	"testing"

	"github.com/omarov.kzh/tictactoe/internal/board"
)

func TestChooseMatchesSpecificationExamples(t *testing.T) {
	tests := []struct {
		name string
		pos  string
		want int
	}{
		{name: "win beats block", pos: "OO.XX....", want: 3},
		{name: "block top row", pos: "XX......O", want: 3},
		{name: "win ignores threat", pos: "OO.XX..X.", want: 3},
		{name: "empty board center", pos: ".........", want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := fromPosition(tt.pos)
			got := Choose(b)
			if got.Cell != tt.want {
				t.Fatalf("Choose(%q) = %d, want %d", tt.pos, got.Cell, tt.want)
			}
		})
	}
}

func fromPosition(pos string) *board.Board {
	b := board.New(3)
	for i, mark := range pos {
		switch mark {
		case 'X':
			b.Cells[i] = board.X
		case 'O':
			b.Cells[i] = board.O
		}
	}
	return b
}
