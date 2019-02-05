package main

import (
	"fmt"
	"gopkg.in/gcfg.v1" //импорт пакета для работы с ini-файлами
)

func main()  {
	config := struct { //Создание структуры для хранения конф значений
		Section struct{
			Enables bool
			Path string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "conf.ini") //Извлечение даных из INI-файла в структуру с обрабоктой ошибок
	if err != nil{
		fmt.Println("failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Enables)
	fmt.Println(config.Section.Path)
}
