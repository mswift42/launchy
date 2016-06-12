package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var SampleSearchResults = []Searchresult{
	Searchresult{"gobook", "/Documents/GoBook.pdf", "pdf.png"},
	Searchresult{"Dive_into_Python3", "/Documents/Dive_into_Python3.pdf", "pdf.png"},
}

func TestSearchresultNames(t *testing.T) {
	assert := assert.New(t)
	s1 := SearchresultNames(SampleSearchResults)
	assert.Equal(s1[0], "gobook")
}

func TestLocateCommand(t *testing.T) {
	assert := assert.New(t)
	l1 := locateCommand("go")
	assert.Equal(l1.Path, "/usr/bin/locate")
}
