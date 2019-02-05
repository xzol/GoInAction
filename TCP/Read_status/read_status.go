package main

import (
	"fmt"
	"bufio"
	"net"
)

func main(){
	//Соединение по TCP
	conn, _ := net.Dial("tcp", "golang.org:80")
	//для входящих соединений можно использовать net.Listen

	//Отправка строки через соединение
	//fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")
	//fmt.Fprintf(conn, "GET / HTTP/2.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n') //Вывод первой строки ответа

	fmt.Println(status)
}
