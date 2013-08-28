package sum

import (
	"flag"

	"bazil.org/bazil/cliutil/subcommands"
)

type sum struct {
	subcommands.Description
	flag.FlagSet
	Arguments struct {
		A int
		B int
	}
	Config struct {
		Frob bool
	}
}

func (a *sum) Run() int {
	return int(a.Arguments.A + a.Arguments.B)
}

// Calc is exported so the unit tests can inspect it.
var Sum = sum{
	Description: "sum two numbers",
}

var _ = subcommands.FlagParser(&Sum)

func init() {
	subcommands.Register(&Sum)
	Sum.BoolVar(&Sum.Config.Frob, "frobnicate", false, "frobnicate the qubbitz")
}
