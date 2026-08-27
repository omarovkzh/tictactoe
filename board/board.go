package board

import (
	"fmt"
	"strings"
)

const (
	Empty = ""
	X     = "X"
	O     = "O"
)

type Board struct {
	Size  int
	Cells []string
}

type RenderOptions struct {
	Color   bool
	Big     bool
	WinLine map[int]bool
}

func New(size int) *Board {
	return &Board{Size: size, Cells: make([]string, size*size)}
}

func (b *Board) Reset() {
	for i := range b.Cells {
		b.Cells[i] = Empty
	}
}

func (b *Board) Set(cell int, mark string) bool {
	if cell < 1 || cell > len(b.Cells) || b.Cells[cell-1] != Empty {
		return false
	}
	b.Cells[cell-1] = mark
	return true
}

func (b *Board) Clear(cell int) {
	if cell >= 1 && cell <= len(b.Cells) {
		b.Cells[cell-1] = Empty
	}
}

func (b *Board) IsFree(cell int) bool {
	return cell >= 1 && cell <= len(b.Cells) && b.Cells[cell-1] == Empty
}

func (b *Board) Full() bool {
	for _, cell := range b.Cells {
		if cell == Empty {
			return false
		}
	}
	return true
}

func (b *Board) CheckWin(mark string) ([]int, bool) {
	n := b.Size
	for r := 0; r < n; r++ {
		line := make([]int, 0, n)
		ok := true
		for c := 0; c < n; c++ {
			idx := r*n + c
			if b.Cells[idx] != mark {
				ok = false
				break
			}
			line = append(line, idx)
		}
		if ok {
			return line, true
		}
	}

	for c := 0; c < n; c++ {
		line := make([]int, 0, n)
		ok := true
		for r := 0; r < n; r++ {
			idx := r*n + c
			if b.Cells[idx] != mark {
				ok = false
				break
			}
			line = append(line, idx)
		}
		if ok {
			return line, true
		}
	}

	line := make([]int, 0, n)
	ok := true
	for i := 0; i < n; i++ {
		idx := i*n + i
		if b.Cells[idx] != mark {
			ok = false
			break
		}
		line = append(line, idx)
	}
	if ok {
		return line, true
	}

	line = make([]int, 0, n)
	ok = true
	for i := 0; i < n; i++ {
		idx := i*n + (n - 1 - i)
		if b.Cells[idx] != mark {
			ok = false
			break
		}
		line = append(line, idx)
	}
	if ok {
		return line, true
	}
	return nil, false
}

func (b *Board) Render(opts RenderOptions) string {
	if opts.Big {
		return b.renderBig(opts)
	}
	return b.renderPlain(opts)
}

func (b *Board) renderPlain(opts RenderOptions) string {
	var out strings.Builder
	width := len(fmt.Sprint(len(b.Cells)))
	sepCell := strings.Repeat("-", width+2)
	sep := strings.Join(repeat(sepCell, b.Size), "+")
	for r := 0; r < b.Size; r++ {
		if r > 0 {
			out.WriteString(sep)
			out.WriteByte('\n')
		}
		for c := 0; c < b.Size; c++ {
			if c > 0 {
				out.WriteString("|")
			}
			idx := r*b.Size + c
			out.WriteString(" ")
			out.WriteString(b.renderCell(idx, width, opts))
			if c < b.Size-1 {
				out.WriteString(" ")
			}
		}
		out.WriteByte('\n')
	}
	return out.String()
}

func (b *Board) renderBig(opts RenderOptions) string {
	var out strings.Builder
	sep := strings.Join(repeat("-----", b.Size), "+")
	for r := 0; r < b.Size; r++ {
		if r > 0 {
			out.WriteString(sep)
			out.WriteByte('\n')
		}
		lines := []strings.Builder{{}, {}, {}}
		for c := 0; c < b.Size; c++ {
			idx := r*b.Size + c
			glyph := b.bigGlyph(idx, opts)
			for line := 0; line < 3; line++ {
				if c > 0 {
					lines[line].WriteByte('|')
				}
				lines[line].WriteString(glyph[line])
			}
		}
		for line := 0; line < 3; line++ {
			out.WriteString(strings.TrimRight(lines[line].String(), " "))
			out.WriteByte('\n')
		}
	}
	return out.String()
}

func (b *Board) bigGlyph(idx int, opts RenderOptions) [3]string {
	cell := b.Cells[idx]
	switch cell {
	case X:
		return [3]string{
			colorize("X   X", X, opts, idx),
			colorize("  X  ", X, opts, idx),
			colorize("X   X", X, opts, idx),
		}
	case O:
		return [3]string{
			colorize(" OOO ", O, opts, idx),
			colorize("O   O", O, opts, idx),
			colorize(" OOO ", O, opts, idx),
		}
	default:
		label := fmt.Sprint(idx + 1)
		if len(label) > 5 {
			label = label[:5]
		}
		padded := center(label, 5)
		return [3]string{"     ", dim(padded, opts.Color), "     "}
	}
}

func (b *Board) renderCell(idx, width int, opts RenderOptions) string {
	cell := b.Cells[idx]
	if cell == Empty {
		return dim(fmt.Sprintf("%*d", width, idx+1), opts.Color)
	}
	return colorize(fmt.Sprintf("%*s", width, cell), cell, opts, idx)
}

func colorize(text, mark string, opts RenderOptions, idx int) string {
	if !opts.Color {
		return text
	}
	if opts.WinLine != nil && opts.WinLine[idx] {
		return "\033[1;32m" + text + "\033[0m"
	}
	if mark == X {
		return "\033[91m" + text + "\033[0m"
	}
	if mark == O {
		return "\033[94m" + text + "\033[0m"
	}
	return text
}

func dim(text string, color bool) string {
	if !color {
		return text
	}
	return "\033[2m" + text + "\033[0m"
}

func center(text string, width int) string {
	if len(text) >= width {
		return text
	}
	left := (width - len(text)) / 2
	right := width - len(text) - left
	return strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
}

func repeat(s string, n int) []string {
	items := make([]string, n)
	for i := range items {
		items[i] = s
	}
	return items
}
