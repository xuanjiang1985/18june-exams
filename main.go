package main

import (
	"fmt"
	"github.com/issue9/unique"
	"github.com/json-iterator/go"
	"github.com/snluu/uuid"
)

func main() {
	fmt.Println(unique.Date().String())
	var id uuid.UUID = uuid.Rand()
	fmt.Println(id.Hex())

	data := []string{"zhoug", "gang", "31"}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, _ := jsoniter.Marshal(&data)
	fmt.Println(string(b))

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b2, _ := jsoniter.Marshal(&group)
	fmt.Println(string(b2))
}
