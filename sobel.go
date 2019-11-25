package main

import (
	"./uint32slice"
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
	existingImageFile, err := os.Open("bruit2.jpg")
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
	var red = make([]uint32, 49)
	var green = make([]uint32, 49)
	var blue = make([]uint32, 49)

	for cpt := 3; cpt < imgWidth-3; cpt++ {
		for cpt2 := 3; cpt2 < imgHeight-3; cpt2++ {
			i := 0
			for cptwi := -3; cptwi < 4; cptwi++ {
				for cpthe := -3; cpthe < 4; cpthe++ {
					red[i], green[i], blue[i], _ = loadedImage.At(cpt+cptwi, cpt2+cpthe).RGBA()
					i++
				}
			}
			uint32slice.SortUint32s(red)
			uint32slice.SortUint32s(green)
			uint32slice.SortUint32s(blue)
			//fmt.Println(red[4], uint8(green[4]*255/65535), uint8(blue[4]*255/65535))
			valrouge,valvert,valbleu:=uint8(red[4] * 255 / 65535), uint8(green[4] * 255 / 65535), uint8(blue[4] * 255 / 65535)
			myImage.Set(cpt, cpt2, color.RGBA{valrouge,valvert,valbleu, 255})
		}
	}
	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create("imgmed.png")
	if err != nil {
		// Handle error
	}
	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, myImage)

	// Don't forget to close files
	outputFile.Close()

}
