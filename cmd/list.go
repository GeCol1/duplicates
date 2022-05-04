package cmd

import (
	"context"
	"duplicates/core"
	"flag"
	"github.com/google/subcommands"
	"log"
)

type listCmd struct {
	main   string
	source string
}

func (*listCmd) Name() string     { return "list" }
func (*listCmd) Synopsis() string { return "list the duplicated file from the second folder" }
func (*listCmd) Usage() string {
	return `list -main <folder1> --source <folder2>
	list the duplicates found from folder2 into folder1 and print in stdout.`
}

func (l *listCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&l.main, "main", "", "add the main folder path")
	f.StringVar(&l.source, "source", "", "add the source folder path")
}

func (l *listCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	log.Printf("execute command with main : %s and source : %s", l.main, l.source)
	err := core.FindDuplicates(l.main, l.source)
	if err != nil {
		log.Printf("error while finding duplicates : %s", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
