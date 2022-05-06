package core

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func FindDuplicates(mainFolder, sourceFolder string, delete bool) error {
	logDuplicatesToFile := true
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
	fmt.Printf("found %d duplicates :\n", len(duplicates))
	duplicatesFile, err := os.OpenFile("doublons.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("could not create file to list doublons : %s", err)
		logDuplicatesToFile = false
	}
	defer func(duplicatesFile *os.File) {
		err := duplicatesFile.Close()
		if err != nil {

		}
	}(duplicatesFile)
	fileWriter := bufio.NewWriter(duplicatesFile)
	for _, path := range duplicates {
		fmt.Printf("%s\n", path)
		if logDuplicatesToFile {
			//duplicatesFile.Write([]byte(path))
			_, _ = fileWriter.Write([]byte(fmt.Sprintf("%s\n", path)))
		}
	}
	err = fileWriter.Flush()
	if err != nil {
		return err
	}
	if delete {
		err := deleteDuplicates(duplicates)
		if err != nil {
			return err
		}
	}
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

func deleteDuplicates(pathsOfFilesToDelete []string) error {
	for _, path := range pathsOfFilesToDelete {
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}
	return nil
}
