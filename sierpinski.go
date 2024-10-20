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

