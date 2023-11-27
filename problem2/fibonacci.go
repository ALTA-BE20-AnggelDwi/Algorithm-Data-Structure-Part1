package main

import "fmt"

func Fibonacci(number int) int {
	if number <= 1 {
		return number
	}
	return Fibonacci(number-1) + Fibonacci(number-2)

	// if number <= 1 {
	// 	return number
	// }

	// var first int = 0
	// var second int = 1
	// var result int = 0

	// for i := 2; i <= number; i++ {
	// 	result = first + second
	// 	first, second = second, result
	// }

	// return result
}

func main() {

	fmt.Println(Fibonacci(0))  // 0
	fmt.Println(Fibonacci(2))  // 1
	fmt.Println(Fibonacci(9))  // 34
	fmt.Println(Fibonacci(10)) // 55
	fmt.Println(Fibonacci(12)) // 144

}
