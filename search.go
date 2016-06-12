package main

// Searchresult represents the result of a local search.
type Searchresult struct {
	name      string
	fullpath  string
	thumbnail string
}

// Bookmarks repersents a bookmarked location of a gnome desktop, e.g.
// Videos, Documents, Pictures, Downloads, ...
type Bookmarks []string

var bookmarks = Bookmarks{"", "/Documents", "/.Downloads", "/Music", "Pictures", "/Videos"}

// SearchresultNames takes a slice of Searchresults and
// returns of slice of its names.
func SearchresultNames(sr []Searchresult) []string {
	results := make([]string, len(sr))
	for i := range sr {
		results[i] = sr[i].name
	}
	return results
}
func main() {

}
