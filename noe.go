package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	// Create a blank image 10 pixels wide by 4 pixels tall
	myImage := image.NewRGBA(image.Rect(0, 0, 10, 4))

	for cpt := 0; cpt < 10; cpt++ {
		// You can access the pixels through myImage.Pix[i]
		// One pixel takes up four bytes/uint8. One for each: RGBA
		// So the first pixel is controlled by the first 4 elements
		// Values for color are 0 black - 255 full color
		// Alpha value is 0 transparent - 255 opaque
		myImage.Pix[0+cpt*4] = 255 // 1st pixel red
		myImage.Pix[1+cpt*4] = 0   // 1st pixel green
		myImage.Pix[2+cpt*4] = 255 // 1st pixel blue
		myImage.Pix[3+cpt*4] = 255 // 1st pixel alpha
	}
	for cpt := 10; cpt < 40; cpt++ {
		// You can access the pixels through myImage.Pix[i]
		// One pixel takes up four bytes/uint8. One for each: RGBA
		// So the first pixel is controlled by the first 4 elements
		// Values for color are 0 black - 255 full color
		// Alpha value is 0 transparent - 255 opaque
		myImage.Pix[0+cpt*4] = 0   // 1st pixel red
		myImage.Pix[1+cpt*4] = 0   // 1st pixel green
		myImage.Pix[2+cpt*4] = 0   // 1st pixel blue
		myImage.Pix[3+cpt*4] = 255 // 1st pixel alpha
	}

	// myImage.Pix contains all the pixels
	// in a one-dimensional slice
	fmt.Println(myImage.Pix)

	// Stride is how many bytes take up 1 row of the image
	// Since 4 bytes are used for each pixel, the stride is
	// equal to 4 times the width of the image
	// Since all the pixels are stored in a 1D slice,
	// we need this to calculate where pixels are on different rows.
	fmt.Println(myImage.Stride) // 40 for an image 10 pixels wide

	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create("test.png")
	if err != nil {
		// Handle error
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, myImage)

	// Don't forget to close files
	outputFile.Close()
}
