package main

import (
	"compress/gzip"
	"io"
	"os"
)

// go run simple.go testDir/*

func main() {
	for _, file := range os.Args[1:] {
		compres(file)
	}
}

func compres(fileName string) error {
	fileIn, error := os.Open(fileName)
	if error != nil {
		return error
	}
	defer fileIn.Close()
	out, error := os.Create(fileName + ".gz")
	if error != nil {
		return error
	}
	defer out.Close()
	gzout := gzip.NewWriter(out)
	_, error = io.Copy(gzout, fileIn)
	defer gzout.Close()
	return error
}
