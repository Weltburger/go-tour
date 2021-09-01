package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	var ch1, ch2 = make(chan int), make(chan int)
	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()
	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()

	for {
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2
		if ok1 != ok2 || n1 != n2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}

	return true
}

func main() {
	ch := make(chan int)
	go func () {
		defer close(ch)
		Walk(tree.New(5), ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(3), tree.New(3)))
	fmt.Println(Same(tree.New(4), tree.New(5)))
	fmt.Println(Same(tree.New(7), tree.New(7)))
}
