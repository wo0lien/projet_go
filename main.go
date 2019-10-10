package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "hubble_lite.png"
	f, err := os.Open(fileName)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Image opened")
	defer f.Close()
}
