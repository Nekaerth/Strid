package gui

import (
	"math"
	"math/rand"

	"github.com/gopherjs/gopherjs/js"
	"github.com/nequilich/gocto"
)

type CanvasBody struct {
	body       *gocto.Body
	canvasMap  map[string]*gocto.Canvas
	contextMap map[string]*js.Object
}

func (this CanvasBody) resizeAllCanvas() {
	for _, canvas := range this.canvasMap {
		canvas.Set("width", gocto.GetWindowInnerWidth())
		canvas.Set("height", gocto.GetWindowInnerHeight())
	}
}

func (this CanvasBody) refreshUi() {
	drawGrid(this.contextMap["gridContext"])
	drawMenu(this.contextMap["menuContext"])
}

func SetUpInterface() {
	var canvasBody CanvasBody = CanvasBody{&gocto.Body{gocto.GetElementById("canvasBody")}, make(map[string]*gocto.Canvas), make(map[string]*js.Object)}

	backgroundCanvas := &gocto.Canvas{gocto.CreateElement("canvas")}
	gridCanvas := &gocto.Canvas{gocto.CreateElement("canvas")}
	menuCanvas := &gocto.Canvas{gocto.CreateElement("canvas")}

	canvasBody.canvasMap["backgroundCanvas"] = backgroundCanvas
	canvasBody.canvasMap["gridCanvas"] = gridCanvas
	canvasBody.canvasMap["menuCanvas"] = menuCanvas

	canvasBody.body.AppendChild(backgroundCanvas)
	canvasBody.body.AppendChild(gridCanvas)
	canvasBody.body.AppendChild(menuCanvas)

	canvasBody.contextMap["backgroundContext"] = backgroundCanvas.GetContext2d()
	canvasBody.contextMap["gridContext"] = gridCanvas.GetContext2d()
	canvasBody.contextMap["menuContext"] = menuCanvas.GetContext2d()

	gocto.AddEventListener("resize", func() {
		canvasBody.resizeAllCanvas()
		canvasBody.refreshUi()
	})

	canvasBody.resizeAllCanvas()
	canvasBody.refreshUi()
}

func drawGrid(context *js.Object) {
	var polygonSides float64 = 6
	var polygonRadius float64 = 32

	var xAdjust float64 = polygonRadius * math.Cos(math.Pi/polygonSides)
	var yAdjust float64 = polygonRadius * math.Sin(math.Pi/polygonSides) * (polygonSides / 2)

	heightInPolygons := (gocto.GetWindowInnerHeight() / int(polygonRadius*1.5)) + 1
	widthInPolygons := (gocto.GetWindowInnerWidth() / (int(xAdjust) * 2)) + 1

	context.Set("strokeStyle", "black")
	context.Set("lineWidth", 2)

	xStart := -(heightInPolygons / 2)

	for y := 0; y < heightInPolygons; y++ {
		for x := xStart; x < widthInPolygons; x++ {
			var xPixel float64 = float64(x)*xAdjust*2 + (float64(y) * xAdjust)
			var yPixel float64 = float64(y) * yAdjust

			drawPolygon(context, xPixel, yPixel, polygonSides, polygonRadius)
		}
	}
}

func drawPolygon(context *js.Object, xPixel float64, yPixel float64, sides float64, radius float64) {
	context.Set("fillStyle", getRandomColour())
	context.Call("beginPath")
	context.Call("moveTo",
		xPixel+radius*math.Cos(math.Pi/sides),
		yPixel+radius*math.Sin(math.Pi/sides))
	for i := 0; float64(i) < sides+1; i++ {
		context.Call("lineTo",
			xPixel+radius*math.Cos(float64(i)*2*math.Pi/sides+math.Pi/sides),
			yPixel+radius*math.Sin(float64(i)*2*math.Pi/sides+math.Pi/sides))
	}
	context.Call("fill")
	context.Call("stroke")
	context.Call("closePath")
}

func getRandomColour() string {
	n := rand.Intn(3)
	if n == 1 {
		return "green"
	} else if n == 2 {
		return "blue"
	} else {
		return "yellow"
	}
}

func drawMenu(context *js.Object) {
	gocto.DrawImage(context, "gui/menu-elements/menu_test_L.png", 0, gocto.GetWindowInnerHeight()-320)
	gocto.DrawImage(context, "gui/menu-elements/menu_test_R.png", gocto.GetWindowInnerWidth()-500, gocto.GetWindowInnerHeight()-125)
}
