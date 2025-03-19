package main

import (
	"fmt"
	"math"
)

func main() {
	//
	// 2.1 hello world
	fmt.Println("Hello World!")

	//
	// 2.2 变量
	var a = "initial"
	var b, c int = 1, 2
	var d = true

	var e float64
	f := float32(e)

	g := a + "foo"
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println(g)

	const s string = "constant"
	const h = 500000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Cos(h))
}
