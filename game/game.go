package game

import (
	"fmt"
	"io"

	"github.com/team/tic-tac-toe-arena/internal/ai"
	"github.com/team/tic-tac-toe-arena/internal/board"
	"github.com/team/tic-tac-toe-arena/internal/cli"
)

type Game struct {
	cfg   cli.Config
	in    io.Reader
	out   io.Writer
	board *board.Board
	stats Stats
}

type Stats struct {
	Games int
	XWins int
	OWins int
	Draws int
}

func New(cfg cli.Config, in io.Reader, out io.Writer) *Game {
	return &Game{
		cfg:   cfg,
		in:    in,
		out:   out,
		board: board.New(cfg.Size),
	}
}

func (g *Game) Run() {
	for {
		moves, completed := g.playOne()
		if !completed {
			return
		}
		g.printStats(moves)
		if !g.askAgain() {
			return
		}
		g.board.Reset()
	}
}

func (g *Game) playOne() (int, bool) {
	current := g.cfg.First
	moves := 0
	fmt.Fprint(g.out, g.board.Render(board.RenderOptions{Color: g.cfg.Color, Big: g.cfg.Big}))

	for {
		if g.cfg.Mode == cli.ModeAI && current == board.O {
			move := ai.Choose(g.board)
			if move.Cell == 0 {
				return moves, false
			}
			g.board.Set(move.Cell, board.O)
			moves++
			if g.cfg.Verbose {
				fmt.Fprintf(g.out, "AI: %s at %d\n", move.Reason, move.Cell)
			}
			fmt.Fprintf(g.out, "O plays %d\n", move.Cell)
		} else {
			cell, ok := g.readMove(current)
			if !ok {
				return moves, false
			}
			g.board.Set(cell, current)
			moves++
		}

		if line, ok := g.board.CheckWin(current); ok {
			winLine := make(map[int]bool, len(line))
			for _, idx := range line {
				winLine[idx] = true
			}
			fmt.Fprint(g.out, g.board.Render(board.RenderOptions{Color: g.cfg.Color, Big: g.cfg.Big, WinLine: winLine}))
			if current == board.X {
				g.stats.XWins++
			} else {
				g.stats.OWins++
			}
			g.stats.Games++
			fmt.Fprintf(g.out, "%s wins!\n", g.name(current))
			return moves, true
		}
		fmt.Fprint(g.out, g.board.Render(board.RenderOptions{Color: g.cfg.Color, Big: g.cfg.Big}))
		if g.board.Full() {
			g.stats.Draws++
			g.stats.Games++
			fmt.Fprintln(g.out, "Draw!")
			return moves, true
		}
		current = other(current)
	}
}

func (g *Game) readMove(mark string) (int, bool) {
	for {
		fmt.Fprintf(g.out, "%s move: ", g.name(mark))
		var cell int
		if _, err := fmt.Fscan(g.in, &cell); err != nil {
			var bad string
			if _, badErr := fmt.Fscan(g.in, &bad); badErr != nil {
				return 0, false
			}
			fmt.Fprintf(g.out, "%sError: enter a number 1-%d%s\n", red(g.cfg.Color), len(g.board.Cells), reset(g.cfg.Color))
			continue
		}
		if cell < 1 || cell > len(g.board.Cells) {
			fmt.Fprintf(g.out, "%sError: enter a number 1-%d%s\n", red(g.cfg.Color), len(g.board.Cells), reset(g.cfg.Color))
			continue
		}
		if !g.board.IsFree(cell) {
			fmt.Fprintf(g.out, "%sError: cell %d is taken%s\n", red(g.cfg.Color), cell, reset(g.cfg.Color))
			continue
		}
		return cell, true
	}
}

func (g *Game) askAgain() bool {
	for {
		fmt.Fprint(g.out, "Play again? (y/n): ")
		var answer string
		if _, err := fmt.Fscan(g.in, &answer); err != nil {
			return false
		}
		if answer == "y" {
			return true
		}
		if answer == "n" {
			return false
		}
	}
}

func (g *Game) printStats(moves int) {
	fmt.Fprintln(g.out, "=== Stats ===")
	fmt.Fprintf(g.out, "Games: %d   X: %d   O: %d   Draws: %d\n", g.stats.Games, g.stats.XWins, g.stats.OWins, g.stats.Draws)
	if g.cfg.Verbose {
		xRate := 0
		oRate := 0
		if g.stats.Games > 0 {
			xRate = g.stats.XWins * 100 / g.stats.Games
			oRate = g.stats.OWins * 100 / g.stats.Games
		}
		fmt.Fprintf(g.out, "Moves this game: %d   Win rate — X: %d%%   O: %d%%\n", moves, xRate, oRate)
	}
}

func (g *Game) name(mark string) string {
	if mark == board.X {
		return g.cfg.NameX
	}
	return g.cfg.NameO
}

func other(mark string) string {
	if mark == board.X {
		return board.O
	}
	return board.X
}

func red(color bool) string {
	if color {
		return "\033[91m"
	}
	return ""
}

func reset(color bool) string {
	if color {
		return "\033[0m"
	}
	return ""
}
