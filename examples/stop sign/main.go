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

	// drawing octagon
	for i := 0; i < 8; i++ {
		canvas.DrawLine(drawing.Red, 45, 3)
	}

	// drawing pole
	canvas.Move(drawing.Right, 4)
	canvas.Move(drawing.Right, 5)
	canvas.Turn(drawing.Left)

	canvas.DrawRectangle(drawing.Cyan, 10, 1)

	canvas.DrawLine(drawing.Green, drawing.Straight, 3)
	// don't forget to save the canvas to a file
	err := canvas.SaveImage("stopsign.png")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
