package main

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
	"os/user"
	"path"
	"strings"

	"github.com/codeskyblue/go-sh"
)

// Searchresult represents the result of a local search.
type Searchresult struct {
	name      string
	fullpath  string
	thumbnail string
}

// Bookmarks repersents a bookmarked location of a gnome desktop, e.g.
// Videos, Documents, Pictures, Downloads, ...
type Bookmarks []string

var bookmarks = Bookmarks{"", "/Documents", "/.Downloads",
	"/Music", "Pictures", "/Videos"}

// SearchresultNames takes a slice of Searchresults and
// returns of slice of its names.
func SearchresultNames(sr []*Searchresult) []string {
	results := make([]string, len(sr))
	for i := range sr {
		results[i] = sr[i].name
	}
	return results
}

// locateCommand takes a queryvalue ant returns an exec.Command
// with the 'locate' command, ignoring case, limited to 20 queries,
// matching only the 'base name', searching for the query value.
func locateCommand(query string) *exec.Cmd {
	return exec.Command("locate", "-l", "20", "-b", "-i", query)
}
func findCommandBookmarks(loc, value string) (*exec.Cmd, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	if loc == "" {
		return exec.Command("find", usr.HomeDir+loc, "-maxdepth", "1",
			"-iname", "*"+value+"*"), nil
	}
	return exec.Command("find", usr.HomeDir+loc, "-maxdepth", "2",
		"-iname", "*"+value+"*"), nil
}

func findCommandBinries(loc, value string) *exec.Cmd {
	return exec.Command("find", loc, "-maxdepth", "2", "-iname", "*"+value+"*")
}

func scanner(out []byte) *bufio.Scanner {
	return bufio.NewScanner(bytes.NewReader(out))
}

func locateOutput(query string) []*Searchresult {
	var res []*Searchresult
	cmd := locateCommand(query)
	output, err := sh.Command(cmd.Path, cmd.Args[1:]).Output()
	if err != nil {
		log.Fatal("locate command fails with error: ", err)
	}
	scanner := scanner(output)
	for scanner.Scan() {
		res = append(res, newSearchresult(scanner.Text()))
	}
	return res
}

func commandOutput(query, location string, cmd *exec.Cmd) []*Searchresult {
	var res []*Searchresult
	output, err := sh.Command(cmd.Path, cmd.Args[1:]).Output()
	if err != nil {
		log.Fatal("search command failed with error: ", err)
	}
	scanner := scanner(output)
	for scanner.Scan() {
		res = append(res, newSearchresult(scanner.Text()))
	}
	return res
}

func findBinariesOutput(query, location string) []*Searchresult {
	var res []*Searchresult
	cmd := findCommandBinries(location, query)
	output, err := sh.Command(cmd.Path, cmd.Args[1:]).Output()
}

func getMimeType(file string) (string, error) {
	mime, err := sh.Command("file", "--mime-type", "--b", file).Output()
	if err != nil {
		return "", err
	}
	return string(mime), nil
}
func newSearchresult(result string) *Searchresult {
	var sr Searchresult
	res := strings.Trim(result, " ")
	sr.name = path.Base(res)
	sr.fullpath = res
	return &sr
}
