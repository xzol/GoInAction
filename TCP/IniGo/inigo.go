package main

import (
	"net/http"
	"fmt"
)

//res - запрос
//req - ответ
func hello(res http.ResponseWriter, req *http.Request)  {
	fmt.Fprint(res, "Hello, my name is Inigo Montoya")
}

func main()  {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
}