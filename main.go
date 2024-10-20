package main



import (

	"fmt"

)



func main() {

	// Generate Julia set

	juliaSetImage, err := GenerateJuliaSet(-0.7+0.27015i, 255)

	if err != nil {

		fmt.Println("Error generating Julia set:", err)

		return

	}

	if err := SaveImage("julia.png", juliaSetImage); err != nil {

		fmt.Println("Error saving Julia set image:", err)

		return

	}

	fmt.Println("Julia set image saved as 'julia.png'")



	// Generate Sierpinski triangle

	sierpinskiImage, err := GenerateSierpinskiTriangle(800)

	if err != nil {

		fmt.Println("Error generating Sierpinski triangle:", err)

		return

	}

	if err := SaveImage("sierpinski.png", sierpinskiImage); err != nil {

		fmt.Println("Error saving Sierpinski triangle image:", err)

		return

	}

	fmt.Println("Sierpinski triangle image saved as 'sierpinski.png'")

}



