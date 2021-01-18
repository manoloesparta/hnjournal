package main

import (
	"cronjob/urls"
)

func main() {
	m := urls.NewConnection("mongodb://database:27017", "journal", "hackernews")

	links := urls.GetLinks()
	for _, el := range links {
		m.Insert(*el)
	}

	m.Close()
}
