package main

import (
	"hnjournal/urls"
)

func main() {
	links := urls.GetLinks()
	for _, el := range links {
		el.Log()
	}
}
