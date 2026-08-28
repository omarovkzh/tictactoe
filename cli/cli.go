package cli

import (
	"fmt"
	"strconv"
	"strings"
)

type Mode int

const (
	ModePlayers Mode = iota
	ModeAI
)

type Config struct {
	Mode     Mode
	Color    bool
	Big      bool
	Verbose  bool
	First    string
	NameX    string
	NameO    string
	Size     int
	ShowHelp bool
}

type Error struct {
	Message string
	Code    int
}

func (e *Error) Error() string {
	return e.Message
}

func Usage() string {
	return `Usage: go run main.go (--players | --ai) [options]

Modes (exactly one required):
  --players       two human players take turns
  --ai            play against the computer (you are X)

Options:
  --color         enable colored output (default: plain)
  --big           render the board with large glyphs
  --verbose       show extended statistics
  --first X|O     who moves first (default: X)
  --name A,B      custom names: X=A, O=B (e.g. --name Alice,Bob)
  --size N        board is NxN, win = N in a row (default: 3)
  --help, -h      print this help and exit 0
`
}

func Parse(args []string) (Config, *Error) {
	cfg := Config{
		First: "X",
		NameX: "X",
		NameO: "O",
		Size:  3,
	}
	players := false
	ai := false
	sizeProvided := false

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "--help", "-h":
			cfg.ShowHelp = true
			return cfg, nil
		case "--players":
			players = true
		case "--ai":
			ai = true
		case "--color":
			cfg.Color = true
		case "--big":
			cfg.Big = true
		case "--verbose":
			cfg.Verbose = true
		case "--first":
			val, ok := value(args, &i, "--first")
			if !ok {
				return cfg, fail("Error: --first requires a value")
			}
			if val != "X" && val != "O" {
				return cfg, fail("Error: --first must be X or O")
			}
			cfg.First = val
		case "--name":
			val, ok := value(args, &i, "--name")
			if !ok {
				return cfg, fail("Error: --name requires a value")
			}
			parts := strings.SplitN(val, ",", 2)
			if len(parts) != 2 || strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" {
				return cfg, fail("Error: --name must be A,B with two non-empty names")
			}
			cfg.NameX = strings.TrimSpace(parts[0])
			cfg.NameO = strings.TrimSpace(parts[1])
		case "--size":
			val, ok := value(args, &i, "--size")
			if !ok {
				return cfg, fail("Error: --size requires a value")
			}
			n, err := strconv.Atoi(val)
			if err != nil || n < 3 {
				return cfg, fail("Error: --size must be an integer >= 3")
			}
			cfg.Size = n
			sizeProvided = true
		default:
			return cfg, fail(fmt.Sprintf("Error: unknown flag %s", arg))
		}
	}

	if players == ai {
		return cfg, fail("Error: choose exactly one of --players or --ai")
	}
	if ai && sizeProvided {
		return cfg, fail("Error: --ai and --size cannot be combined (AI is 3x3 only)")
	}
	if ai {
		cfg.Mode = ModeAI
	} else {
		cfg.Mode = ModePlayers
	}
	return cfg, nil
}

func value(args []string, index *int, flag string) (string, bool) {
	next := *index + 1
	if next >= len(args) || strings.HasPrefix(args[next], "--") {
		return "", false
	}
	*index = next
	return args[next], true
}

func fail(message string) *Error {
	return &Error{Message: message, Code: 1}
}
