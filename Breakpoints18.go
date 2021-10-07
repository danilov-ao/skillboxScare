package main

import "fmt"

func doPanic(a int) {
	if a > 1 {
		doPanic(a - 1)
	}
	panic(a)
}

func main() {
	var s string

	s = "Hello world"
	fmt.Println(s)
	var ii int
	ii = 10
	fmt.Println(ii)
	fmt.Println("I am debugging now")
	doPanic(5)

}
