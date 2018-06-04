package main

import (
	"fmt"
	"github.com/issue9/unique"
	"github.com/snluu/uuid"
	"github.com/json-iterator/go"
)

func main() {
	fmt.Println(unique.Date().String())
	var id uuid.UUID = uuid.Rand()
	fmt.Println(id.Hex())

	data := []string{"zhoug", "gang", "31"}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, _ := json.Marshal(&data)
	fmt.Println(string(b))
}
