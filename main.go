package main

import (
	"hnjournal/urls"
)

func main() {
	m := urls.NewConnection("mongodb://localhost:27017", "journal", "hackernews")

	links := urls.GetLinks()
	for _, el := range links {
		m.Insert(*el)
	}

	m.Close()
}
