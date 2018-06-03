package main

import (
	"fmt"

	"github.com/troylelandshields/drawing"
)

func branch(canvas *drawing.Canvas, length float64, ratio float64, angle int, numBranches int) {
	// draw this branch on the canvas
	canvas.DrawLine(drawing.Cyan, drawing.Straight, int(length))

	// if there are more branches to draw, draw them
	if numBranches > 0 {
		// draw the left branch
		canvas.Turn(-angle)
		branch(canvas, length*ratio, ratio, angle, numBranches-1)

		// draw the right branch
		canvas.Turn(angle * 2)
		branch(canvas, length*ratio, ratio, angle, numBranches-1)

		// turn back to original direction
		canvas.Turn(-angle)
	}

	// get back into original position
	canvas.Move(drawing.Backwards, int(length))
	canvas.Turn(drawing.Backwards)
}

func main() {
	// create a new canvas on which to draw
	canvas := drawing.NewCanvas(1000, 1500)

	// draw a tree using the recursive branch function
	branch(canvas, 5, 0.7, 7, 5)

	// don't forget to save the canvas to a file
	err := canvas.SaveImage("tree.png")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
