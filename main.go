package main

import (
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {

}

func drawGui() {
	MainWindow{
		Title: "Murphy's tech law",
		Size:  Size{Width: 300, Height: 300},
		Children: []Widget{
			Label{Text: getTodaysLaw()},
		},
		OnMouseUp: ,
	}.Run()
}
