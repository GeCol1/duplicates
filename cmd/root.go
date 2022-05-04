package cmd

import (
	"context"
	"flag"
	"github.com/google/subcommands"
)

func Execute(ctx context.Context) subcommands.ExitStatus {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&listCmd{}, "")

	flag.Parse()
	return subcommands.Execute(ctx)
}
