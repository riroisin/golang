package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = int(b)
	d := float64(a * c)

	var e int = 7
	f := int(d) * e

	var g int = int(b * 3)
	var h int = int(b) * 3
	fmt.Print(g, h, f)
}
