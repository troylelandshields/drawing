// package Drawing provides an interface to move an artist around a canvas to draw shapes and lines. It is very similar to the functionality provided on http://goplay.space
package drawing

import (
	"math"

	"github.com/fogleman/gg"
)

const (
	// These values can be used to help when drawing lines or moving
	Right     = 90
	Left      = -90
	Straight  = 0
	Backwards = 180

	// These are the available colors to use when executing PenDown on a canvas
	Red     = "red"
	Blue    = "blue"
	Green   = "green"
	Cyan    = "cyan"
	Magenta = "magenta"

	x = math.Pi / 180
)

func rad(d float64) float64 { return d * x }

// Canvas allows an artist to draw shapes and lines on an empty canvas and to save the image to file
type Canvas struct {
	currentX         float64
	currentY         float64
	currentDirection float64
	currentColor     string

	multiplier float64
	drawingCtx *gg.Context
}

// NewCanvas can be called to create a new drawing canvas on which to create a nice image
func NewCanvas(canvasWidth int, canvasHeight int) *Canvas {
	newDrawingCtx := gg.NewContext(canvasWidth, canvasHeight)

	canvas := &Canvas{
		currentX:         float64(canvasWidth / 2),
		currentY:         float64(canvasHeight / 2),
		currentDirection: 0,
		currentColor:     "",
		multiplier:       50,
		drawingCtx:       newDrawingCtx,
	}

	return canvas
}

// SaveImage can be called to save whatever has been drawn on the canvas to an image on disk
func (c *Canvas) SaveImage(fileName string) error {
	return c.drawingCtx.SavePNG(fileName)
}

// PenUp removes the current color
func (c *Canvas) PenUp() {
	c.currentColor = ""
}

// PenDown sets the color so that a line will be drawn when the artist moves
func (c *Canvas) PenDown(color string) {
	c.currentColor = color

	switch color {
	case Red:
		c.drawingCtx.SetRGB255(255, 0, 0)
	case Green:
		c.drawingCtx.SetRGB255(0, 255, 0)
	case Blue:
		c.drawingCtx.SetRGB255(0, 0, 255)
	case Magenta:
		c.drawingCtx.SetRGB255(255, 0, 255)
	case Cyan:
		c.drawingCtx.SetRGB255(0, 255, 255)
	default:
		c.drawingCtx.SetRGB255(0, 0, 0)
	}
}

// Forward moves the artist forward the given number of spaces. If the pen is down, it will draw a line
func (c *Canvas) Forward(numSpaces int) {
	xDelta, yDelta := c.calcXYDelta(numSpaces)

	xDelta = xDelta * c.multiplier
	yDelta = yDelta * c.multiplier

	newX := c.currentX + xDelta
	newY := c.currentY + yDelta

	if c.currentColor != "" {
		c.drawingCtx.DrawLine(c.currentX, c.currentY, newX, newY)
		c.drawingCtx.Stroke()
	}

	c.currentX = newX
	c.currentY = newY

	return
}

func (c *Canvas) calcXYDelta(numSpaces int) (float64, float64) {
	if c.currentDirection == 0 {
		return 0, float64(-numSpaces)
	} else if c.currentDirection == 180 {
		return 0, float64(numSpaces)
	} else if c.currentDirection == 90 {
		return float64(numSpaces), 0
	} else if c.currentDirection == 270 {
		return float64(-numSpaces), 0
	}

	angle1 := c.currentDirection

	xNeg := 1.0
	yNeg := -1.0

	if c.currentDirection > 90 && c.currentDirection < 270 {
		angle1 = angle1 - 90
		if angle1 < 90 {
			angle1 = 90 - angle1
		}
		yNeg = 1.0
	}

	if c.currentDirection > 180 && c.currentDirection < 360 {
		if angle1 > 180 {
			angle1 = 360 - angle1
		} else {
			angle1 = angle1 - 90
		}
		xNeg = -1.0
	}

	angle2 := 90 - angle1

	xDelta := math.Sin(rad(angle1)) * float64(numSpaces) * xNeg
	yDelta := math.Sin(rad(angle2)) * float64(numSpaces) * yNeg

	return xDelta, yDelta
}

// Turn will change the current direction of the artist. Positive degrees turns the artist right, negative degrees turns the artist left
func (c *Canvas) Turn(degrees int) {
	if degrees < 0 {
		c.turnLeft(-degrees)
	} else {
		c.turnRight(degrees)
	}
}

func (c *Canvas) turnRight(degrees int) {
	c.currentDirection = c.currentDirection + float64(degrees)
	if c.currentDirection >= 360 {
		c.currentDirection = c.currentDirection - 360
	}
}

func (c *Canvas) turnLeft(degrees int) {
	rightDeg := 360 - degrees

	c.turnRight(rightDeg)
}
