package board

import "testing"

func TestCheckWinRowsColumnsAndDiagonals(t *testing.T) {
	tests := []struct {
		name  string
		cells []string
		mark  string
	}{
		{name: "row", cells: []string{X, X, X, "", "", "", "", "", ""}, mark: X},
		{name: "column", cells: []string{O, "", "", O, "", "", O, "", ""}, mark: O},
		{name: "main diagonal", cells: []string{X, "", "", "", X, "", "", "", X}, mark: X},
		{name: "anti diagonal", cells: []string{"", "", O, "", O, "", O, "", ""}, mark: O},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := New(3)
			copy(b.Cells, tt.cells)
			if _, ok := b.CheckWin(tt.mark); !ok {
				t.Fatalf("expected %s win", tt.mark)
			}
		})
	}
}

func TestFullBoardWithoutLineIsDrawCandidate(t *testing.T) {
	b := New(3)
	b.Cells = []string{X, O, X, X, O, O, O, X, X}

	if _, ok := b.CheckWin(X); ok {
		t.Fatal("unexpected X win")
	}
	if _, ok := b.CheckWin(O); ok {
		t.Fatal("unexpected O win")
	}
	if !b.Full() {
		t.Fatal("expected full board")
	}
}
