package main

import "fmt"

func someFunction(x int, y int, A func(int, int) int) {
	defer func() {
		x *= 2
		y /= 2
		fmt.Println(A(x, y))
	}()
	fmt.Println("Результат расчетов:")
}

func main() {
	f := func(x int, y int) int { return (x + y) * (x + y) }
	someFunction(5, 10, func(x int, y int) int { return x + y })
	someFunction(2, 6, f)
	someFunction(5, 8, func(x int, y int) int { return x*2 + y/2 })
}
