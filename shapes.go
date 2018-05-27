package drawing

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
