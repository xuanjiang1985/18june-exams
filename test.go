package main

import (
	"18june-exam/simpleFactory"
	"fmt"
)

func main() {
	helloApiStruct := simplefactory.NewAPI(2)
	str := helloApiStruct.Say("zhou gang")
	fmt.Println(str)
}
