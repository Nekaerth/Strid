package gui

import "github.com/nequilich/gocto"
import "github.com/gopherjs/gopherjs/js"

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
	
	backgroundCanvas.Set("width", 400)
	backgroundCanvas.Set("height", 300)
	
	gridCanvas.Set("width", 400)
	gridCanvas.Set("height", 300)
	
	menuCanvas.Set("width", 400)
	menuCanvas.Set("height", 300)
}
