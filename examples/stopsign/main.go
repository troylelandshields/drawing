package main

import (
	"fmt"

	"github.com/troylelandshields/drawing"
)

func main() {
	// create a new canvas on which to draw
	canvas := drawing.NewCanvas(1000, 1500)

	// centering picture
	canvas.Move(drawing.Straight, 5)
	canvas.Move(drawing.Left, 4)
	canvas.Turn(drawing.Right)

	canvas.DrawStopSign(drawing.Red, drawing.Cyan)

	// don't forget to save the canvas to a file
	err := canvas.SaveImage("stopsign.png")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
