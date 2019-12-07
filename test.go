package main

import "fmt"

func main() {
	dx := 3
	dy := 4
	x := make([][]uint8, dx)
	for i := range x {
		x[i] = make([]uint8, dy+i)
	}

	for cptx := range x {
		for cpty := range x[cptx] {
			if cptx != 0 {
				x[cptx][cpty] = uint8(cptx*len(x[cptx-1]) + cpty)
			} else {
				x[cptx][cpty] = uint8(cpty)
			}
		}

	}
	fmt.Println(x)
}
