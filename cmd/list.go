package cmd

import (
	"context"
	"duplicates/core"
	"flag"
	"github.com/google/subcommands"
	"log"
)

type ListCmd struct {
	main   string
	source string
}

func (*ListCmd) Name() string     { return "list" }
func (*ListCmd) Synopsis() string { return "list the duplicated file from the second folder" }
func (*ListCmd) Usage() string {
	return `list -main <folder1> --source <folder2>
	list the duplicates found from folder2 into folder1 and print in stdout.`
}

func (l *ListCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&l.main, "main", "", "add the main folder path")
	f.StringVar(&l.source, "source", "", "add the source folder path")
}

func (l *ListCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	log.Printf("execute command with main : %s and source : %s", l.main, l.source)
	err := core.FindDuplicates(l.main, l.source)
	if err != nil {
		log.Printf("error while finding duplicates : %s", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
