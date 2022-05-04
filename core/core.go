package core

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
)

func FindDuplicates(mainFolder, sourceFolder string) error {
	mapMain, err := ComputeHashOfFiles(mainFolder)
	if err != nil {
		return err
	}
	mapSource, err := ComputeHashOfFiles(sourceFolder)
	if err != nil {
		return err
	}

	duplicates, err := CompareHash(mapMain, mapSource)
	if err != nil {
		return err
	}
	fmt.Printf("found %d duplicates", len(duplicates))
	return nil
}

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

func CompareHash(main, source map[string]string) ([]string, error) {
	var res []string
	if main == nil || source == nil {
		return res, errors.New("nil map not allowed")
	}
	for keyFromSource, pathInSource := range source {
		if _, ok := main[keyFromSource]; ok {
			res = append(res, pathInSource)
		}
	}
	return res, nil
}
