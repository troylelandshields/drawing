package main

import (
	"fmt"

	"github.com/troylelandshields/drawing"
)

func main() {
	// create a new canvas on which to draw
	canvas := drawing.NewCanvas(1000, 1000)

	// draw the other 4 lines
	for i := 0; i < 3; i++ {

		canvas.Move(0, 0)
		canvas.DrawPoly(drawing.Magenta, 3)
	}

	// don't forget to save the canvas to a file
	err := canvas.SaveImage("poly.png")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
