package main

import (
	"context"
	"duplicates/cmd"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"os"
)

func main() {
	fmt.Printf("find duplicates from source into main...\n\n")
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&cmd.ListCmd{}, "")
	//subcommands.Register(&printCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
	//path1 := "D:\\Backup Photos\\fold1"
	//path1 := "/home/djay/Work/01_PROJECTS/42_GO/duplicates"
	//map1, err := core.ComputeHashOfFiles(path1)
	//if err != nil {
	//	log.Fatalf("error")
	//}
	//fmt.Printf("count : %d", len(map1))
}
