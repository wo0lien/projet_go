package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	// Read image from file that already existsqss
	existingImageFile, err := os.Open("test.png")
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, imageType, err := image.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}
	fmt.Println(imageData)
	fmt.Println(imageType)

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0, 0)

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		// Handle error
		//add com
	}
	fmt.Println(loadedImage)
	//Gris = 0,2125 * Rouge + 0,7154 * Vert + 0,0721 * Bleu
	b := loadedImage.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y
	imgNbPix := imgHeight * imgWidth
	for cpt := 0; cpt < imgNbPix; cpt++ {
		grisVal := loadedImage
		loadedImage.Pix[0+cpt*4] = grisVal // 1st pixel red
		loadedImage.Pix[1+cpt*4] = grisVal // 1st pixel green
		loadedImage.Pix[2+cpt*4] = grisVal // 1st pixel blue
		loadedImage.Pix[3+cpt*4] = 255     // 1st pixel alpha
	}

}
