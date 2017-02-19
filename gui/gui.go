package gui

import(
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
	context.Set("fillStyle", "green")
	context.Call("fillRect", 0, 0 , gocto.GetWindowInnerWidth()/2, gocto.GetWindowInnerHeight()/2)
}