package drawing

import (
	"math/rand"
	"time"
)

//DrawSickAsTriangles draws a bunch of sick-as triangles and stuff
func (c *Canvas) DrawSickAsTriangles() {
	for i := 0; i < 5; i++ {
		c.DrawTriangle(Green, 6)
		c.Move(-65, 3)
	}
	for i := 0; i < 4; i++ {
		c.Move(90, 3)
		c.DrawTriangle(Magenta, 4)
	}
	for i := 0; i < 18; i++ {
		c.DrawRectangle(Cyan, 2, 2)
		c.Move(25, 1)
	}
}

// DrawTriangle draws an equilateral triangle from the bottom-left corner, ending in the same position that it started in.
func (c *Canvas) DrawTriangle(color string, length int) {
	c.DrawLine(color, 30, length)
	c.DrawLine(color, 120, length)
	c.DrawLine(color, 120, length)
	c.Turn(Right)
}

// DrawRectangle draws a rectangle with the given parameters from the top-right corner and ends where it begins.
func (c *Canvas) DrawRectangle(color string, length int, width int) {
	c.DrawLine(color, Right, length)
	c.DrawLine(color, Right, width)
	c.DrawLine(color, Right, length)
	c.DrawLine(color, Right, width)
}

// Move moves the pen without drawing; it turns the given number of degrees and then moves the given number of spaces.
func (c *Canvas) Move(numDegrees int, numSpaces int) {
	c.PenUp()
	c.Turn(numDegrees)
	c.Forward(numSpaces)
}

// DrawLine moves the pen to draw a line; it turns the given number of degrees and then moves the given number of spaces.
func (c *Canvas) DrawLine(color string, numDegrees int, numSpaces int) {
	c.PenDown(color)
	c.Turn(numDegrees)
	c.Forward(numSpaces)
}

//DrawPoly added by Holly
func (c *Canvas) DrawPoly(color string, length int) {
	c.DrawLine(color, 90, 3)

	for i := 0; i < 5; i++ {
		c.DrawLine(Magenta, 60, length)
	}

	c.Turn(Right)

}

// DrawRandom will draw a random number of lines with random colors and random lengths
func (c *Canvas) DrawRandom() {

	rand.Seed(time.Now().Unix())

	numLines := rand.Int31n(12)

	for i := 0; i < int(numLines); i++ {

		var color string
		rColor := rand.Int31n(5)
		switch rColor {
		case 0:
			color = Blue
		case 1:
			color = Red
		case 2:
			color = Cyan
		case 3:
			color = Magenta
		case 4:
			color = Green
		}

		numDegrees := rand.Int31n(360)
		numSpaces := rand.Int31n(5)

		c.DrawLine(color, int(numDegrees), int(numSpaces))
	}
}

// DrawStopSign allows you to draw a stop sign with your choice of sign and pole color--avoid messing with the sizing on this one though, it's a bitch
func (c *Canvas) DrawStopSign(signColor string, poleColor string) {
	// drawing octagon
	for i := 0; i < 8; i++ {
		c.DrawLine(signColor, 45, 3)
	}
	// drawing pole
	c.Move(Right, 4)
	c.Move(Right, 5)
	c.Turn(Left)

	c.DrawRectangle(poleColor, 10, 1)

	// face gopher forwards again
	c.Turn(Left)
}
