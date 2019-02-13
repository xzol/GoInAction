package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//uuid идентификатор диска
const UUID string = "29ED-2200"

type Card struct {
	UUID  string          //uuid
	found map[string]bool //подключенные устройства
}

type Mounter interface {
	Mnt() bool
	isNeedCard(file os.FileInfo) bool
}

func main() {
	for {
		var Card Mounter = Card{UUID, map[string]bool{}}
		Card.Mnt()
		println("Exit.")
	}
}

func (c Card) Mnt() bool {
	files, err := ioutil.ReadDir("/dev/disk/by-uuid")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if value, ok := c.found[file.Name()]; ok {
			if value {
				//выполняем копирование
			} else {
				c.isNeedCard(file)
			}
		} else {
			c.isNeedCard(file)
		}
		fmt.Println(file.Name())
	}
	return true
}

func (c Card) isNeedCard(file os.FileInfo) bool {
	if file.Name() == c.UUID {
		c.found[file.Name()] = true
		return true
	} else {
		return false
	}
}

func (c Card) Copy() {

}
