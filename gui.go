package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func setup_textentry() *gtk.Entry {
	te, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create new entry: ", err)
	}
	return te
}

func setup_box(orientation gtk.Orientation) *gtk.Box {
	box, err := gtk.BoxNew(orientation, 2)
	if err != nil {
		log.Fatal("unable to create new box: ", err)
	}
	return box
}

func setup_grid(orientation gtk.Orientation) *gtk.Grid {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to setup grid: ", err)
	}
	grid.SetOrientation(orientation)
	return grid
}
