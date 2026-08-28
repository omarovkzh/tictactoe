# Tic-Tac-Toe Arena

Tic-Tac-Toe Arena is a terminal tic-tac-toe game written in Go. It supports two human players, a deterministic computer opponent, colored output, large board rendering, custom names, larger human-only boards, and session statistics.

## How to Run

```sh
go run . --players
go run . --ai
```

Available flags:

- `--players` - two human players take turns.
- `--ai` - play against the computer. You are always `X`; the computer is always `O`.
- `--color` - enable ANSI colors.
- `--big` - render the board with large glyphs.
- `--verbose` - show extended statistics and AI decision traces.
- `--first X|O` - choose who moves first. Default: `X`.
- `--name A,B` - set custom names for `X` and `O`.
- `--size N` - use an `N x N` board in players mode. Default: `3`.
- `--help` or `-h` - print help.

## Rules

Cells are numbered from left to right, top to bottom. Players enter a cell number to place their mark. Empty cells show their number as a hint.

A player wins by filling a complete row, column, or diagonal. If the board is full and there is no winner, the game is a draw. After each game, the session score is shown and the player can start another round.

In AI mode, `X` is the human and `O` is the computer. The AI is rule-based and deterministic:

1. Win if `O` can complete a line.
2. Block if `X` can complete a line.
3. Take the center.
4. Take the first free corner in order `1, 3, 7, 9`.
5. Take the first free side in order `2, 4, 6, 8`.

AI mode is limited to the default `3 x 3` board.

## Example

```text
$ go run . --ai
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9
X move: 5
 1 | 2 | 3
---+---+---
 4 | X | 6
---+---+---
 7 | 8 | 9
O plays 1
 O | 2 | 3
---+---+---
 4 | X | 6
---+---+---
 7 | 8 | 9
```

## Team

- Team member: `github.com/team`
