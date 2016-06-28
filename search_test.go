package main

import (
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

var SampleSearchResults = []string{
	"/Documents/GoBook.pdf",
	"/Documents/Dive_into_Python3.pdf",
}

func TestSearchresultNames(t *testing.T) {
	assert := assert.New(t)
	s1 := newSearchresult(SampleSearchResults[0])
	assert.Equal(s1.name, "GoBook.pdf")
	assert.Equal(s1.fullpath, SampleSearchResults[0])
	s2 := newSearchresult(SampleSearchResults[1])
	assert.Equal(s2.name, "Dive_into_Python3.pdf")
	assert.Equal(s2.fullpath, SampleSearchResults[1])
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

func TestGetMimeType(t *testing.T) {
	assert := assert.New(t)
	m1, err := getMimeType("/home/martin/Documents/ModernC.pdf")
	assert.Equal(err, nil)
	assert.Equal(m1, "application/pdf\n")
}

func TestNewSearchresult(t *testing.T) {
	assert := assert.New(t)
	l1 := "/home/martin/Documents/ModernC.pdf"
	l2 := "/home/martin/Documents/Go in Action.pdf"
	q := newSearchresult(l1)
	assert.Equal(q.name, "ModernC.pdf")
	assert.Equal(q.fullpath, l1)
	q2 := newSearchresult(l2)
	assert.Equal(q2.name, "Go in Action.pdf")
	assert.Equal(q2.fullpath, l2)
}

func TestScanner(t *testing.T) {
	assert := assert.New(t)
	b := []byte("some text")
	scanner := scanner(b)
	for scanner.Scan() {
		assert.Equal(scanner.Text(), "some text")
	}
}
