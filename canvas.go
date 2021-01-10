package canvas

import (
	"log"
	"syscall/js"
)

// Canvas The type holding the actual canvas
type Canvas struct {
	canvas  js.Value
	context js.Value
}

// Initialise the canvas and context
func Initialise(canvasID string) *Canvas {
	var canvas Canvas
	window := js.Global()
	document := window.Get("document")

	canvas.canvas = document.Call("getElementById", canvasID)
	canvas.context = canvas.canvas.Call("getContext", "2d")

	log.SetPrefix("Canvas: ")

	return &canvas
}

// SetFillStyle Sets the fill style
func (canvas *Canvas) SetFillStyle(colour string) {
	canvas.context.Set("fillStyle", colour)
	log.Println("SetFillStyle: ", colour)
}

// FillRect draws a filled rect at x, y with width and height
func (canvas *Canvas) FillRect(x, y, width, height float64) {
	canvas.context.Call("fillRect", x, y, width, height)
	log.Println("FillRect: ", x, y, width, height)
}
