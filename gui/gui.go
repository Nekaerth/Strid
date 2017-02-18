package gui

import(
	"github.com/nequilich/gocto"
	"github.com/gopherjs/gopherjs/js"
)

type Body struct {
    canvasBody *js.Object
    canvasMap map[string]*js.Object
}

func SetUpInterface() {
	var body Body = Body{gocto.GetElementById("canvasBody"), make(map[string]*js.Object)}

	backgroundCanvas := gocto.CreateElement("canvas")
	gridCanvas := gocto.CreateElement("canvas")
	menuCanvas := gocto.CreateElement("canvas")
	
	body.canvasMap["backgroundCanvas"] = backgroundCanvas
	body.canvasMap["gridCanvas"] = gridCanvas
	body.canvasMap["menuCanvas"] = menuCanvas
	
	gocto.AppendChild(body.canvasBody, backgroundCanvas)
	gocto.AppendChild(body.canvasBody, gridCanvas)
	gocto.AppendChild(body.canvasBody, menuCanvas)
	
	gocto.GetContext2d(backgroundCanvas)
	gocto.GetContext2d(gridCanvas)
	gocto.GetContext2d(menuCanvas)
	
	//resized := true
	gocto.AddEventListener("resize", func() { resizeAllCanvas(body) })
	
	//var draw func()
	//draw = func() {
	//	if resized {
	//		resizeAllCanvas(body)
	//		resized = false
	//	}
	//	gocto.RequestAnimationFrame(draw)
	//}
	//draw()
}

func resizeAllCanvas(body Body) {
	for _, canvas := range body.canvasMap {
		canvas.Set("width", gocto.GetWindowInnerWidth())
		canvas.Set("height", gocto.GetWindowInnerHeight())
	}
}
