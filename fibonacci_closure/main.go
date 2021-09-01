package main

import "fmt"

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		buf := first
		first, second = second, buf + second

		return buf
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
