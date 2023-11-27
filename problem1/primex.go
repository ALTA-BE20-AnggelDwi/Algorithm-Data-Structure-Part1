package main

import "fmt"

func PrimeX(number int) int {
	count := 0
	i := 2
	for {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
		if count == number {
			return i
		}
		i++
	}

}

func main() {
	fmt.Println(PrimeX(1))  //2
	fmt.Println(PrimeX(5))  //11
	fmt.Println(PrimeX(8))  //19
	fmt.Println(PrimeX(9))  //23
	fmt.Println(PrimeX(10)) //29
}
