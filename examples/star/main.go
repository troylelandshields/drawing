package main

import (
	"fmt"

	"github.com/troylelandshields/drawing"
)

func main() {
	// create a new canvas on which to draw
	canvas := drawing.NewCanvas(1000, 1000)

	// draw the first line
	canvas.DrawLine(drawing.Magenta, 18, 7)

	// draw the other 4 lines
	for i := 0; i < 4; i++ {
		canvas.DrawLine(drawing.Magenta, 144, 7)
	}

	// don't forget to save the canvas to a file
	err := canvas.SaveImage("star.png")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
