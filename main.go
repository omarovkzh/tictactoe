package main

import (
	"fmt"
	"os"

	"github.com/team/tic-tac-toe-arena/internal/cli"
	"github.com/team/tic-tac-toe-arena/internal/game"
)

func main() {
	cfg, err := cli.Parse(os.Args[1:])
	if err != nil {
		if err.Message != "" {
			fmt.Fprintln(os.Stderr, err.Message)
		}
		fmt.Fprint(os.Stderr, cli.Usage())
		os.Exit(err.Code)
	}

	if cfg.ShowHelp {
		fmt.Print(cli.Usage())
		return
	}

	runner := game.New(cfg, os.Stdin, os.Stdout)
	runner.Run()
}
