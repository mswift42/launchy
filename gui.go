package main

import (
	"container/list"
	"fmt"
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

func addRemoveWindow() *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("unable to setup window: ", err)
	}
	grid := setupGrid(gtk.ORIENTATION_VERTICAL)
	win.Add(grid)
	sw := setupScrolledWindow(nil, nil)
	grid.Attach(sw, 0, 0, 2, 1)
	sw.SetHExpand(true)
	sw.SetVExpand(true)
	labelsGrid := setupGrid(gtk.ORIENTATION_VERTICAL)
	labelsGrid.SetHExpand(true)
	return win
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
	grid := setupGrid(gtk.ORIENTATION_VERTICAL)
	win.Add(grid)
	entry := setupTextentry()
	grid.Add(entry)
	sw := setupScrolledWindow(nil, nil)
	sw.SetHExpand(true)
	sw.SetVExpand(true)
	grid.Add(sw)
	labelsGrid := setupGrid(gtk.ORIENTATION_VERTICAL)
	labelsGrid.SetHExpand(true)
	sw.Add(labelsGrid)
	var labellist = list.New()
	entry.Connect("activate", func() {
		text, err := entry.GetText()
		if err != nil {
			fmt.Println(err)
		}
		resultnames := SearchresultNames(locateOutput(text))
		for _, j := range resultnames {
			label := setupLabel(j)
			label.SetHExpand(true)
			labellist.PushBack(label)
			labelsGrid.Add(label)
			labelsGrid.ShowAll()
		}
		labelsGrid.ShowAll()

	})
	return win
}
