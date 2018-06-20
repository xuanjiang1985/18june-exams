package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-123456))
}

func reverse(x int) (num int) {
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
		fmt.Println(x)
	}
	// 使用 math 包中定义好的最大最小值
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return
}
