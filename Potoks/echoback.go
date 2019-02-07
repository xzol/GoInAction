package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout) //Вызов функции echo как Go подпрограмму
	time.Sleep(30 * time.Second) //30 сек. пауза
	fmt.Println("Timeout out.")  //Вывод сообщения о завршении ожидания
	os.Exit(0)                   //Выход из программы. При этом сопрограмма будет остановлена.
}

func echo(in io.Reader, out io.Writer) { //Функция echo является обыной функцией
	io.Copy(out, in) //io.Copy Скоприует дынные из os.Reader в os.Writer
}
