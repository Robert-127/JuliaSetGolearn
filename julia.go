package main



import (

	"image"

	"image/color"

	"image/png"

	"math/cmplx"

	"os"

)



func GenerateJuliaSet(c complex128, maxIter int) (image.Image, error) {

	const (

		width, height = 800, 800

		zoom          = 1.0

	)



	img := image.NewRGBA(image.Rect(0, 0, width, height))



	for x := 0; x < width; x++ {

		for y := 0; y < height; y++ {

			z := complex(float64(x)/zoom-zoom/2, float64(y)/zoom-zoom/2)

			iter := 0



			for cmplx.Abs(z) < 2 && iter < maxIter {

				z = z*z + c

				iter++

			}



			colorValue := uint8(255 * iter / maxIter)

			img.Set(x, y, color.RGBA{colorValue, 0, 0, 255})

		}

	}



	return img, nil

}



func SaveImage(filename string, img image.Image) error {

	file, err := os.Create(filename)

	if err != nil {

		return err

	}

	defer file.Close()



	return png.Encode(file, img)

}



