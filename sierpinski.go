package main



import (

	"image"

	"image/color"

)



func GenerateSierpinskiTriangle(size int) (image.Image, error) {

	img := image.NewRGBA(image.Rect(0, 0, size, size))



	// Fill the background with black

	for x := 0; x < size; x++ {

		for y := 0; y < size; y++ {

			img.Set(x, y, color.Black)

		}

	}



	// Draw the Sierpinski triangle

	drawSierpinski(img, 0, 0, size)



	return img, nil

}



func drawSierpinski(img *image.RGBA, x, y, size int) {

	if size < 1 {

		return

	}



	// Draw the triangle

	for i := 0; i < size; i++ {

		for j := 0; j <= i; j++ {

			img.Set(x+j, y+i, color.White)

			img.Set(x+size-i+j, y+i, color.White)

		}

	}



	// Recursively draw the smaller triangles

	drawSierpinski(img, x, y, size/2)

	drawSierpinski(img, x+size/2, y, size/2)

	drawSierpinski(img, x+size/4, y+size/2, size/2)

}



