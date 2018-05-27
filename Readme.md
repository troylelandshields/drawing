
# About & GoDoc

[![GoDoc](https://godoc.org/github.com/troylelandshields/drawing?status.svg)](https://godoc.org/github.com/troylelandshields/drawing)

The drawing package is very much inspired by the drawing capabilities on http://goplay.space, but allows you to take the Gopher artist to your local machine so that you can save images to disk.

# Getting Started 

To start drawing an image, you must first call the `NewCanvas` function with the desired `height` and `width`.

You must call `canvas.SaveImage(fileName)` when you have finished drawing to save the canvas to an image file.

```go
canvas := drawing.NewCanvas(1000, 1000)

// ... draw on the canvas

err := canvas.SaveImage("out.png")
if err != nil {
    fmt.Println("Error saving image:", err)
}
```

# Example Drawings

There are examples in the `./examples` folder that can be useful to see more about how to use the `drawing` package.

* [Star](https://github.com/troylelandshields/drawing/blob/master/examples/star/main.go)
* [Tree](https://github.com/troylelandshields/drawing/blob/master/examples/tree/main.go)

