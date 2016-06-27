package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func setupTextentry() *gtk.Entry {
	te, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create new entry: ", err)
	}
	return te
}

func setupBox(orientation gtk.Orientation) *gtk.Box {
	box, err := gtk.BoxNew(orientation, 2)
	if err != nil {
		log.Fatal("unable to create new box: ", err)
	}
	return box
}

func setupGrid(orientation gtk.Orientation) *gtk.Grid {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to setup grid: ", err)
	}
	grid.SetOrientation(orientation)
	return grid
}

func setupScrolledWindow(hadjust, vadjust *gtk.Adjustment) *gtk.ScrolledWindow {
	scrollwin, err := gtk.ScrolledWindowNew(hadjust, vadjust)
	if err != nil {
		log.Fatal("unable to setup scrolled window: ", err)
	}
	return scrollwin
}

func setupLabel(text string) *gtk.Label {
	label, err := gtk.LabelNew(text)
	if err != nil {
		log.Fatal("unable to setup label: ", err)
	}
	return label
}

func setupButtonWithLabel(text string) *gtk.Button {
	btn, err := gtk.ButtonNewWithLabel(text)
	if err != nil {
		log.Fatal("unable to setup button: ", err)
	}
	return btn
}

func launchyWindow() *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("Launchy")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	if err != nil {
		log.Fatal("couldn't create window: ", err)
	}
	box := setupBox(gtk.ORIENTATION_HORIZONTAL)
	win.Add(box)
	entry := setupTextentry()
	box.PackStart(entry, true, false, 4)
	return win
}
