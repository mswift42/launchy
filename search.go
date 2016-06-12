package main

import "os/exec"

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
func main() {

}
