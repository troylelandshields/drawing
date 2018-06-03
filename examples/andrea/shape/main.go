package main

import (
	"fmt"

	"github.com/troylelandshields/drawing"
)

func main() {

	canvas := drawing.NewCanvas(1000, 1000)

	for i := 0; i < 5; i++ {
		canvas.DrawTriangle(drawing.Green, 6)
		canvas.Move(-65, 3)
	}

	for i := 0; i < 4; i++ {
		canvas.Move(90, 3)
		canvas.DrawTriangle(drawing.Magenta, 4)
	}

	for i := 0; i < 18; i++ {
		canvas.DrawRectangle(drawing.Cyan, 2, 2)
		canvas.Move(25, 1)
	}

	err := canvas.SaveImage("shape.png")
	if err != nil {
		fmt.Println("Error:", err)
	}

}
