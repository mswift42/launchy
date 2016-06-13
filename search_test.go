package main

import (
	"os/user"
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
	assert.Equal(l1.Args, []string{"locate", "-l", "20", "-b", "-i", "go"})
	assert.Equal(l1.Env, []string(nil))
}

func TestFindCommandBookmarks(t *testing.T) {
	assert := assert.New(t)
	usr, err := user.Current()
	assert.Equal(err, nil)
	f1, err := findCommandBookmarks("/Documents", "go")
	assert.Equal(err, nil)
	assert.Equal(f1.Path, "/usr/bin/find")
	assert.Equal(f1.Env, []string(nil))
	assert.Equal(f1.Args[1:], []string{usr.HomeDir + "/Documents", "-maxdepth", "2",
		"-iname", "*go*"})
}

func TestFindCommandBinaries(t *testing.T) {
	assert := assert.New(t)
	f1 := findCommandBinries("/opt", "go")
	assert.Equal(f1.Path, "/usr/bin/find")
	assert.Equal(f1.Args[1:], []string{"/opt", "-maxdepth", "2", "-iname", "*go*"})
	assert.Equal(f1.Env, []string(nil))
}
