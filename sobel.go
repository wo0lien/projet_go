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
	var maxG float64 = 0
	gris := make([][]int16, imgHeight)
	for i := range gris {
		gris[i] = make([]int16, imgWidth)
	}

	gradient := make([][]float64, imgHeight)
	for i := range gradient {
		gradient[i] = make([]float64, imgWidth)
	}

	for cpt := 0; cpt < imgHeight; cpt++ {
		for cpt2 := 0; cpt2 < imgWidth; cpt2++ {
			red, gr, blue, _ := loadedImage.At(cpt2, cpt).RGBA()
			gris[cpt][cpt2] = int16(0.2125*float32(red*255/65535) + 0.7154*float32(gr*255/65535) + 0.0721*float32(blue*255/65535))

		}
	}
	/*int pixval_x =
	  ( -1* (int)img.at<uchar>(j,i)) + (0* (int)img.at<uchar>(j+1,i)) + (1 * (int)img.at<uchar>(j+2,i)) +
	  ( -2* (int)img.at<uchar>(j,i+1)) + (0* (int)img.at<uchar>(j+1,i+1)) + (2 * (int)img.at<uchar>(j+2,i+1)) +
	  ( -1 * (int)img.at<uchar>(j,i+2)) + (0 * (int)img.at<uchar>(j+1,i+2)) + (1 * (int)img.at<uchar>(j+2,i+2));

	  int pixval_y =
	  (sobel_y[0][0] * (int)newimg.at<uchar>(j,i)) + (sobel_y[0][1] * (int)newimg.at<uchar>(j+1,i)) + (sobel_y[0][2] * (int)newimg.at<uchar>(j+2,i)) +
	  (sobel_y[1][0] * (int)newimg.at<uchar>(j,i+1)) + (sobel_y[1][1] * (int)newimg.at<uchar>(j+1,i+1)) + (sobel_y[1][2] * (int)newimg.at<uchar>(j+2,i+1)) +
	  (sobel_y[2][0] * (int)newimg.at<uchar>(j,i+2)) + (sobel_y[2][1] * (int)newimg.at<uchar>(j+1,i+2)) + (sobel_y[2][2] * (int)newimg.at<uchar>(j+2,i+2));*/

	for cpt := 0; cpt < imgHeight-2; cpt += 1 {
		for cpt2 := 0; cpt2 < imgWidth-2; cpt2 += 1 {
			var gx float64
			gx = float64(-1*gris[cpt][cpt2] + 1*gris[cpt+2][cpt2] + 1*gris[cpt][cpt2] - 2*gris[cpt][cpt2+1] + 2*gris[cpt+2][cpt2+1] - 1*gris[cpt][cpt2+2] + 1*gris[cpt+2][cpt2+2])
			//gy := gris[cpt-1][cpt2+1] + 2*gris[cpt][cpt2+1] + gris[cpt+1][cpt2+1] - gris[cpt-1][cpt2-1] - 2*gris[cpt][cpt2-1] - gris[cpt+1][cpt2-1]
			gradient[cpt][cpt2] = math.Abs(gx) //float32(math.Abs(float64(gx)) + math.Abs(float64(gy)))
			if gradient[cpt][cpt2] > maxG {
				maxG = gradient[cpt][cpt2]
			}

		}
	}
	for cpt := 0; cpt < imgHeight-2; cpt += 1 {
		for cpt2 := 0; cpt2 < imgWidth-2; cpt2 += 1 {
			var valsobel uint8
			if gradient[cpt][cpt2] > 255 {
				valsobel = 255
			}
			valsobel = uint8(gradient[cpt][cpt2] * 255 / maxG)
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
