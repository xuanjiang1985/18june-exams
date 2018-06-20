package main

import "fmt"

const (
	monday = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

const (
	a = iota
	b
)

func calc() (int, int) {
	return 1, 2
}

func main() {
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(wednesday)
	fmt.Println(thursday)
	fmt.Println(friday)
	fmt.Println(saturday)
	fmt.Println(sunday)
	fmt.Println(a)
	fmt.Println(b)

}
