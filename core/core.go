package core

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
)

func ComputeHashOfFiles(path string) (map[string]string, error) {
	res := make(map[string]string)
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error while going through path %s", path)
			return err
		}
		if info.IsDir() {
			return nil
		}
		currentFile, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("error while reading the file %s : %s", path, err)
			return err
		}
		res[GetHash(currentFile)] = path
		return nil
	})
	if err != nil {
		log.Printf("could not walk throug the path %s ", path)
		return nil, errors.New("problem during walk through path")
	}
	return res, nil
}
