package gui

import "github.com/nekaerth/strid/game"
import "github.com/nequilich/gocto"
import "github.com/gopherjs/gopherjs/js"

type body struct {
    *js.Object
    backgroundCanvas js.c
}

var body = gocto.GetElementById("canvasBody")

var backgroundCanvas = gocto.CreateElement("canvas")
var gridCanvas = gocto.CreateElement("canvas")
var menuCanvas = gocto.CreateElement("canvas")

gocto.AddChildren(body, backgroundCanvas, gridCanvas, menuCanvas)

var backgroundContext = gocto.GetContext2D(backgroundCanvas)
var gridContext = gocto.GetContext2D(gridCanvas)
var menuContext = gocto.GetContext2D(menuCanvas)

func SetUpInterface(state *game.GameState) {

}
