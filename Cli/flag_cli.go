package main

import (
	"flag"
	//Для работы с короткими или слитными аругменатми использовать другой пакет flag
	//launchpad.net/gnuflag
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to .")
var spanish bool

func init () {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
}
//
//func VisitAll(func(flag *flag.Flag))  {
//	format := "\t-%s: %s (Default: '%s')\n"
//	fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
//}

func main(){
		//Если используется пакет стандартный
		flag.Parse()
		//Если используется пакет launchpad.net/gnuflag
		//flag.Parse(bool) - где bool: false or true
		if spanish == true {
			fmt.Printf("Hola %s!\n", *name)
		} else {
			fmt.Printf("Hello %s!\n", *name)
		}
}