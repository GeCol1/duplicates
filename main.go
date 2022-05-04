package main

import (
	"context"
	"duplicates/cmd"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("find duplicates from source into main...\n\n")
	ctx := context.Background()
	sts := cmd.Execute(ctx)
	os.Exit(int(sts))
}
