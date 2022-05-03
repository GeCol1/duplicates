package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func main() {
	fmt.Printf("find duplicates\n")
	//path1 := "D:\\Backup Photos\\fold1"
	path1 := "/home/djay/Work/01_PROJECTS/42_GO/duplicates"
	map1, err := ComputeHashOfFiles(path1)
	if err != nil {
		log.Fatalf("error")
	}
	fmt.Printf("count : %d", len(map1))
}

func ComputeHashOfFiles(path string) (map[string]string, error) {
	var res map[string]string
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		//fmt.Printf("Hash %s of file %s\n", GetHash([]byte(info.Name())), info.Name())
		res[GetHash([]byte(info.Name()))] = info.Name()
		return nil
	})
	if err != nil {
		log.Printf("could not walk throug the path %s ", path)
		return nil, errors.New("problem during walk through path")
	}
}
