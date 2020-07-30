package urls

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var c *colly.Collector = colly.NewCollector(
	colly.MaxDepth(0),
)

// URL represents every link in hackernews
type URL struct {
	link  string
	title string
}

func newURL(link string, title string) *URL {
	return &URL{
		link,
		title,
	}
}

// Log prints title and link into terminal
func (u *URL) Log() {
	fmt.Printf("Title: %s, Link: %s\n", u.title, u.link)
}

// GetLinks brings all the url from hackernews
func GetLinks() []*URL {
	urls := []*URL{}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		condition := !(strings.Contains(link, "ycombinator") || strings.Contains(link, "HackerNews"))

		if strings.Contains(link, "http") && condition {
			urls = append(urls, newURL(link, e.Text))
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://news.ycombinator.com/")
	c.Visit("https://news.ycombinator.com/news?p=2")
	c.Visit("https://news.ycombinator.com/news?p=3")
	return urls
}
