package main

import (
	"fmt"

	"github.com/troylelandshields/drawing"
)

func main() {

	canvas := drawing.NewCanvas(1000, 1000)

	canvas.DrawSickAsTriangles()

	err := canvas.SaveImage("shape.png")
	if err != nil {
		fmt.Println("Error:", err)
	}

}
