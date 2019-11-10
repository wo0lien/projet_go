package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
)

func main() {
	// Read image from file that already existsqss
	existingImageFile, err := os.Open("img.jpg")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Succes")
	defer existingImageFile.Close()
	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	loadedImage, _, err := image.Decode(existingImageFile)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}

	//Gris = 0,2125 * Rouge + 0,7154 * Vert + 0,0721 * Bleu
	b := loadedImage.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y
	myImage := image.NewRGBA(loadedImage.Bounds())
	for cpt := 0; cpt < imgWidth; cpt++ {
		for cpt2 := 0; cpt2 < imgHeight; cpt2++ {
			red, gr, blue, _ := loadedImage.At(cpt, cpt2).RGBA()
			myImage.Set(cpt, cpt2, color.RGBA{uint8(255 - (red * 255 / 65535)), uint8(255 - (gr * 255 / 65535)), uint8(255 - (blue * 255 / 65535)), 255})
		}
	}
	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create("imgneg.png")
	if err != nil {
		// Handle error
	}
	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, myImage)

	// Don't forget to close files
	outputFile.Close()

}
