package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var (
		d1 = `true`
		v1 bool
	)
	err := json.Unmarshal([]byte(d1), &v1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(v1)
}
