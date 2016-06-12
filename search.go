package main

import (
	"os/exec"
	"os/user"
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
func SearchresultNames(sr []Searchresult) []string {
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
	return exec.Command(loc, "-maxdepth", "2", "-iname", "*"+value+"*")
}
func main() {

}
