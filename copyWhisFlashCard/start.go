package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//uuid идентификатор диска
const UUID string = "29ED-2200"
const DEVDISK = "/dev/disk/by-uuid"

type Card struct {
	UUID    string          //uuid
	found   map[string]bool //подключенные устройства
	mounted bool            //смонтирован?
}

type Mounter interface {
	Check() bool
	IsNeedCard(file os.FileInfo) bool
	IsMount(file os.FileInfo) bool
}

func main() {
	var Card Mounter = Card{
		UUID,
		map[string]bool{},
		false}
	//}
	for {
		Card.Check()
		println("Exit.")
	}
}

func (c Card) Check() bool {
	files, err := ioutil.ReadDir(DEVDISK)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if value, ok := c.found[file.Name()]; ok {
			if value {
				if c.IsMount(file) {

				} else {

				}
			} else {
				c.IsNeedCard(file)
			}
		} else {
			c.IsNeedCard(file)
		}
		fmt.Println(file.Name())
	}
	return true
}

func (c Card) IsMount(file os.FileInfo) bool {
	var allLink string = DEVDISK + "/" + c.UUID
	f, _ := os.Readlink(allLink)
	println(f)

	return true
}

func (c Card) IsNeedCard(file os.FileInfo) bool {
	if file.Name() == c.UUID {
		c.found[file.Name()] = true
		return true
	} else {
		return false
	}
}

func (c Card) Copy() {

}
