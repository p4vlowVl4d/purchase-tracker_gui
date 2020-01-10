package main

import (
	"github.com/p4vlowVl4d/purchase-tracker_gui/gui"
	"log"
)

func main() {
	log.Println("Starting")
	win := gui.NewWindow(640, 480, "example")
	win.Show()
}
