package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	strArr := strings.Split(s, " ")

	for _, val := range strArr {
		m[strings.Trim(val, " ")]++
	}

	return  m
}

func main() {
	wc.Test(WordCount)
}
