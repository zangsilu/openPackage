package main

import (
	. "fmt"
	"zhangshilu/openPackage/infra/algo"
)

func main() {
	count, amount := int64(10), int64(100)

	for i := int64(0); i < count; i++ {
		x := algo.SimpleRand(count, amount*100)
		Print(float64(x)/float64(100), "--")
	}
	Println()
}
