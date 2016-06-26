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
	box, err := gtk.BoxNew(or, 2)
	if err != nil {
		log.Fatal("unable to create new box: ", err)
	}
	return box
}
