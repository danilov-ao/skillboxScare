package main

import "fmt"

func reverse (a [10] int) [10]int{
	for i:= 0; i < 5; i++ {
		a[i], a[9-i] = a[ 9 - i ], a[i]
	}
	return a
}

func main() {
	var numbers [10] int
	numbers = enterArray()
	for _, number := range reverse(numbers) {
		fmt.Println(number)
	}
}

func enterArray()(numbers [10]int) {
	for i, _ := range numbers {
		fmt.Printf("Введите %v число ", i+1)
		fmt.Scan(&numbers[i])
	}
	return
}