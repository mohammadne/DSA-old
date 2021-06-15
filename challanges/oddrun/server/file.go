package main

import (
	"os"
)

func createDirIfMissed(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		return os.Mkdir(dirName, os.ModePerm)
	}

	return nil
}

func getFile(filename string) (*os.File, error) {
	if isFileExists(filename) {
		return openFile(filename)
	}

	return os.Create(filename)
}

func isFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	if info.IsDir() {
		return false
	}

	return true
}

func openFile(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
}

func createFile(filename string) (*os.File, error) {
	return os.Create(filename)
}
