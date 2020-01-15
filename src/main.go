package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"tcpServer/src/filtres/noise"

	_ "./filtres/colors"
	_ "./filtres/edge"
	_ "./filtres/noise"
)

func main() {
	name := "./src/imgbruit.png"
	out := "./src/out.png"
	//Chregement de l'image
	img := loadPic(name)
	//Filtre de Sobel
	//writePic(edge.FSobel(img), out)
	//Filtre Noir et Blanc
	//writePic(colors.FBW(img), out)
	//Filtre Negatif
	//writePic(colors.Fnegative(img),out)
	//Filtre reduction du bruit par mediane
	writePic(noise.Fmediane(img, 3), out)
	//Filtre reduction du bruit par moyennage
	//writePic(noise.Fmean(img, 1),out)

}

func loadPic(name string) image.Image {
	// Read image from file that already exists
	existingImageFile, err := os.Open(name)
	if err != nil {
		//Handle error
		fmt.Print(err)
	}
	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string
	loadedImage, _, err := image.Decode(existingImageFile)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	return loadedImage
}

func writePic(file image.Image, name string) {
	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	err2 := png.Encode(outputFile, file)
	if err2 != nil {
		fmt.Println(err2)
	}
	outputFile.Close()
}
