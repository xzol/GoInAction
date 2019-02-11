package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup // снова используется группа ожидания для мониторинга группы сопрограмм
	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Println("Words that appear that once:")
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}

type words struct {
	found map[string]int
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) {
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scaner := bufio.NewScanner(file) //disassemble //scaner - полезный инструмент для подбного анализа файлов
	scaner.Split(bufio.ScanWords)
	for scaner.Scan() {
		word := strings.ToLower(scaner.Text())
		dict.add(word, 1)
	}
	return scaner.Err()

}
