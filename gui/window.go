package gui

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func NewWindow(width, height int, title string) *gtk.Window {
	gtk.Init(nil)
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal(err)
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetTitle(title)
	win.SetDefaultSize(width, height)
	return win
}
