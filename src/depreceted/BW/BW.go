package BW

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

func FBW(in image.Image) image.Image{
	loadedImage := in
	b := loadedImage.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y
	myImage := image.NewRGBA(loadedImage.Bounds())
	for cpt := 0; cpt < imgWidth; cpt++ {
		for cpt2 := 0; cpt2 < imgHeight; cpt2++ {
			red, gr, blue, _ := loadedImage.At(cpt, cpt2).RGBA()
			gris := uint8(0.2125*float32(red*255/65535) + 0.7154*float32(gr*255/65535) + 0.0721*float32(blue*255/65535))
			myImage.Set(cpt, cpt2, color.RGBA{gris, gris, gris, 255})
		}
	}
	return myImage

}
