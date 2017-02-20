package gui

import(
	"math"
	"math/rand"
	"github.com/nequilich/gocto"
	"github.com/gopherjs/gopherjs/js"
)

type Body struct {
    canvasBody *js.Object
    canvasMap map[string]*js.Object
	contextMap map[string]*js.Object
}

func (this Body) resizeAllCanvas() {
	for _, canvas := range this.canvasMap {
		canvas.Set("width", gocto.GetWindowInnerWidth())
		canvas.Set("height", gocto.GetWindowInnerHeight())
	}
}

func (this Body) refreshUi() {
	drawGrid(this.contextMap["gridContext"])
}

func SetUpInterface() {
	var body Body = Body{gocto.GetElementById("canvasBody"), make(map[string]*js.Object), make(map[string]*js.Object)}

	backgroundCanvas := gocto.CreateElement("canvas")
	gridCanvas := gocto.CreateElement("canvas")
	menuCanvas := gocto.CreateElement("canvas")
	
	body.canvasMap["backgroundCanvas"] = backgroundCanvas
	body.canvasMap["gridCanvas"] = gridCanvas
	body.canvasMap["menuCanvas"] = menuCanvas
	
	gocto.AppendChild(body.canvasBody, backgroundCanvas)
	gocto.AppendChild(body.canvasBody, gridCanvas)
	gocto.AppendChild(body.canvasBody, menuCanvas)
	
	body.contextMap["backgroundContext"] = gocto.GetContext2d(backgroundCanvas)
	body.contextMap["gridContext"] = gocto.GetContext2d(gridCanvas)
	body.contextMap["menuContext"] = gocto.GetContext2d(menuCanvas)
	
	gocto.AddEventListener("resize", func() {
		body.resizeAllCanvas()
		body.refreshUi()
	})
	
	body.resizeAllCanvas()
	body.refreshUi()
}

func drawGrid(context *js.Object) {
	var polygonSides float64 = 6
	var polygonRadius float64 = 32
	
	widthInPolygons := 10
	heightInPolygons := 10
	
	context.Set("strokeStyle", "black")
	context.Set("lineWidth", 2)
	
	for y := 0; y < heightInPolygons; y++ {
		for x := 0; x < widthInPolygons; x++ {
			var xAdjust float64 = polygonRadius * math.Cos(math.Pi / 6)
			var yAdjust float64 = polygonRadius * math.Sin(math.Pi / 6) * 3
			
			var xPixel float64 = float64(x) * xAdjust * 2 + (float64(y) * xAdjust)
			var yPixel float64 = float64(y) * yAdjust
			
			drawPolygon(context, xPixel, yPixel, polygonSides, polygonRadius)
		}
	}
}

func drawPolygon(context *js.Object, xPixel float64, yPixel float64, sides float64, radius float64) {
	context.Set("fillStyle", getRandomColour())
	context.Call("beginPath")
	context.Call("moveTo",
		xPixel + radius * math.Cos(math.Pi / sides),
		yPixel + radius * math.Sin(math.Pi / sides))
	for i := 0; float64(i) < sides + 1; i++ {
		context.Call("lineTo",
			xPixel + radius * math.Cos(float64(i) * 2 * math.Pi / sides + math.Pi / sides),
			yPixel + radius * math.Sin(float64(i) * 2 * math.Pi / sides + math.Pi / sides))
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