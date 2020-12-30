package scrape

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

var (
	trails []Trail
)

type Trail struct {
	Name   string `json:"trail"`
	Link   string `json:"trail_info"`
	Date   string `json:"last_updated"`
	Status string `json:"status"`
}

func newTrail(e *colly.HTMLElement) Trail {
	var (
		name string
		status string
		date string
		elements []string
	)

	e.ForEach("td", func(_ int, td *colly.HTMLElement){
		elements = append(elements, td.Text)
	})

	name = elements[0]
	status = strings.ToLower(elements[1])
	date = elements[2]

	page := e.ChildAttr("a", "href")
	link := fmt.Sprintf("%s%s", "https://www.trianglemtb.com/", page)


	t := Trail{
		Date:   date,
		Name:   strings.TrimSpace(name),
		Link:   link,
		Status: status,
	}
	return t
}

func FetchTrails() []Trail {
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		// For trails with comments, they have an id attribute on the tr element
		// Since we don't serve the comments as a part of this API, we just skip over it :)
		if len(e.Attr("id")) > 0 {
			return
		}

		// Rules out the header and footer classes on the page
		if e.ChildAttr("td","class") != "trail" {
			return
		}

		currentTrail := newTrail(e)

		trails = append(trails, currentTrail)
	})

	c.Visit("https://www.trianglemtb.com/mobiletrailstatus.php")

	return trails
}