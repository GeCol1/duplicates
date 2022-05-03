package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func main() {
	fmt.Printf("find duplicates\n")
	path1 := "D:\\Backup Photos\\fold1"
	err := filepath.Walk(path1, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Printf("file %s\n", info.Name())
		return nil
	})
	if err != nil {
		log.Fatalf("could not walk throug the path %s ", path1)
	}
}
