package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"math"
	"os"
)

func main() {
	// Read image from file that already existsqss
	existingImageFile, err := os.Open("edge.png")
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
	var maxG float32 = 0
	gris :=  make([][]uint8, 2)
	gris[0] = make([]uint8, imgWidth)
	gris[1] = make([]uint8, imgHeight)
	gradient :=  make([][]float32, 2)
	gradient[0] = make([]float32, imgWidth)
	gradient[1] = make([]float32, imgHeight)

	for cpt := 0; cpt < imgWidth; cpt++ {
		for cpt2 := 0; cpt2 < imgHeight; cpt2++ {
			red, gr, blue, _ := loadedImage.At(cpt, cpt2).RGBA()
			gris[cpt][cpt2] = uint8(0.2125*float32(red*255/65535) + 0.7154*float32(gr*255/65535) + 0.0721*float32(blue*255/65535))

		}
	}


	for cpt := 0; cpt < imgWidth; cpt+=3 {
		for cpt2 := 0; cpt2 < imgHeight; cpt2+=3 {
			if cpt==0 || cpt== imgWidth-1 || cpt2==0 || cpt2==imgHeight-1 {
				gradient[cpt][cpt2] = 0
			} else {
				gx := gris[cpt+1][cpt2-1] + 2*gris[cpt+1][cpt2] + gris[cpt+1][cpt2+1] - gris[cpt-1][cpt2-1] - 2*gris[cpt-1][cpt2] - gris[cpt-1][cpt2+1]
				gy := gris[cpt-1][cpt2+1] + 2* gris[cpt][cpt2+1] + gris[cpt+1][cpt2+1] - gris[cpt-1][cpt2-1] - 2*gris[cpt][cpt2-1] - gris[cpt+1][cpt2-1]
				gradient[cpt][cpt2] = float32(math.Abs(float64(gx))+math.Abs(float64(gy)))
				if gradient[cpt][cpt2] > maxG {
					maxG = gradient[cpt][cpt2]
				}

			}
		}
	}
	for cpt := 0; cpt < imgWidth; cpt+=3 {
		for cpt2 := 0; cpt2 < imgHeight; cpt2 += 3 {
			valsobel := uint8(gradient[cpt][cpt2] * 255 / maxG)
			myImage.Set(cpt, cpt2, color.RGBA{valsobel, valsobel, valsobel, 255})
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
