package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		pic[i] = make([]uint8, dx)
	}

	for x := 0; x < dy; x++ {
		for y := 0; y < dx; y++ {
			pic[x][y] =  uint8((x+y)/2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
