package main

import "fmt"

var x int

const a  = 5
const b int  = 10

const (
	 c = 25
	 d float32 = 12.25
)

const  (
		i = 2017 + iota
		j = 2017 + iota
		k = 2017 + iota
	)

func main() {
	fmt.Println("Exercise 1")
	x = 42
	fmt.Printf("%b\t%#x\t%d", x,x,x)

	fmt.Println("")
	fmt.Println("Exercise 2")
	z := 32
	y := 25
	fmt.Println(z > y)
	fmt.Println(z == y)
	fmt.Println(z >= y)
	fmt.Println(z < y)
	fmt.Println(z <= y)
	fmt.Println(z != y)

	fmt.Println("")
	fmt.Println("Exercise 3")
	fmt.Printf("%v\t%v\t%v\t%v\t", a, b, c, d)

	fmt.Println()
	fmt.Println()
	fmt.Println("Exercise 4")
	var num int = 42
	fmt.Printf("%b\t%#x\t%d", num, num ,num)

	num2 := num << 1
	fmt.Println("\nShifted", num2)
	fmt.Printf("%b\t%#x\t%d", num2, num2, num2)

	fmt.Println()
	fmt.Println()
	fmt.Println("Exercise 5")
	name := "James Bond"
	fmt.Printf(name)

	fmt.Println()
	fmt.Println()
	fmt.Println("Exercise 6")
	fmt.Println(i, j, k)

}
