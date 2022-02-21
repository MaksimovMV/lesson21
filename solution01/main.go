package main

import "fmt"

func calcByFormula(x int16, y uint8, z float32) float32 {
	var S float32
	S = 2*float32(x) + float32(y)*float32(y) - 3/z
	return S
}

func main() {
	var x int16 = 5
	var y uint8 = 4
	var z float32 = 3
	fmt.Println(calcByFormula(x, y, z))
}
