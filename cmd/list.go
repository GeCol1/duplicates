package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
)

type ListCmd struct {
	main   string
	source string
}

func (*ListCmd) Name() string     { return "list" }
func (*ListCmd) Synopsis() string { return "list the duplicated file from the second folder" }
func (*ListCmd) Usage() string {
	return `list -main <folder1> --source <folder2>
	list the duplicates found from folder2 into folder1 and print in stdoud.`
}

func (l *ListCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&l.main, "main", "", "add the main folder path")
	f.StringVar(&l.source, "source", "", "add the source folder path")
}

func (l *ListCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	fmt.Printf("args : %s et %s", l.main, l.source)
	return subcommands.ExitSuccess
}
