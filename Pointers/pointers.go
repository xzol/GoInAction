package main

import "fmt"

func te(f *int){
	fmt.Println(f)
}

func main()  {
	i := 12
	//u := &i
	var z* int
	z = &i
	te(z)
	fmt.Println(*z)
}

