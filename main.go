package main

import (
	"duplicates/core"
	"fmt"
	"log"
)

func main() {
	fmt.Printf("find duplicates starting...\n")
	//path1 := "D:\\Backup Photos\\fold1"
	path1 := "/home/djay/Work/01_PROJECTS/42_GO/duplicates"
	map1, err := core.ComputeHashOfFiles(path1)
	if err != nil {
		log.Fatalf("error")
	}
	fmt.Printf("count : %d", len(map1))
}
