package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

// go run simple.go testDir/*

func main() {
	var wg sync.WaitGroup //Нет необходимости инициализировать WaitGroup
	var i int = -1        //Так как переменная i необходима за пределами цикла, она обьявлена именно здесь
	var file string

	for i, file = range os.Args[1:] {
		wg.Add(1)                  //Для каждого файла сообщить группе, что ожидается выполнение еще одной оперции сжатия
		go func(filename string) { //Эта функция вызывает функцию сжатия и уведомляет группу ожидания о ее завершении
			compress(filename)
			wg.Done()
		}(file) //поскольку вызов сопрограммы происходит в цикле for, требуется небольшая хитрость, чтобы передать параметр
	}
	wg.Wait() //Внешняя сопрограмма (main) ожидает, пока все сопрограммы, выполняющее сжатие, вызовут wg.Done
	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(fileName string) error {
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
