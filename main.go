package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	drawGui()
}

func drawGui() {
	var mw *walk.MainWindow
	w := &MainWindow{
		Title:   "Murphy's tech law",
		MinSize: Size{Width: 100, Height: 100},
		Layout:  VBox{},
		Children: []Widget{
			Label{
				Text: getTodaysLaw(),
				Font: Font{Family: "Arial", PointSize: 10},
			},
			PushButton{
				MinSize: Size{Width: 30, Height: 30},
				MaxSize: Size{Width: 30, Height: 30},
				Text:    "Close",
				OnClicked: func() {
					mw.Close()
				}},
		},
	}
	w.Run()
}
