package main

import _ "main/filtre"

func main() {
	filtre.sobel("edge.png", "result.png")
}
