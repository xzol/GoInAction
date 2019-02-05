package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct { //Определяем стрктуры, ссоответствующей значнеиям в JSON файле
	Enabled bool
	Path string
}

func main()  {
	file, _ := os.Open("config.json")	//открытие кофигурационого файла
	defer file.Close()

	decoder := json.NewDecoder(file)  //Извлечение JSON-значений в переменные
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil{
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)


}
