package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	// Read image from file that already existsqss
	existingImageFile, err := os.Open("couleurs.jpeg")
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()
	/*
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
		existingImageFile.Seek(0, 0)*/

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := jpeg.Decode(existingImageFile)

	//loadedImage.Set(0,0, color.RGBA{125,125, 125, 1})
	if err != nil {
		// Handle error
	}
	//fmt.Println(loadedImage)
	//Gris = 0,2125 * Rouge + 0,7154 * Vert + 0,0721 * Bleu
	b := loadedImage.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y
	myImage := image.NewRGBA(loadedImage.Bounds())
	//imgNbPix := imgHeight * imgWidth
	for cpt := 0; cpt < imgWidth; cpt++ {
		for cpt2 := 0; cpt2 < imgHeight; cpt2++ {
			red, gr, bck, alpha := loadedImage.At(cpt, cpt2).RGBA()
			gris := 0.2125*red + 0.7154*gr + 0.0721*bck
			myImage.Set(cpt, cpt2, color.RGBA{gris, gris, gris, 1})
		}
	}
	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create("test2.png")
	if err != nil {
		// Handle error
	}
	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, myImage)

	// Don't forget to close files
	outputFile.Close()

}
