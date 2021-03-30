package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"regexp"
	"strings"
)

type trail struct {
	Name   string `json:"trail"`
	Status string `json:"status"`
	Date   string `json:"last_updated"`
	Link   string `json:"trail_info"`
}

func NewTrail(name string, status string, date string, link string) trail {

	// Some of the trails have a weird character thrown in as they are sub-trails, but we just treat them as normal
	re, err := regexp.Compile(`[^A-Za-z0-9 - ( )]`)

	if err != nil {
		log.Fatal(err)
	}

	name = re.ReplaceAllString(name, " ")
	status = strings.ToLower(status)

	t := trail{
		Date:   date,
		Name:   strings.TrimSpace(name),
		Link:   link,
		Status: strings.TrimSpace(status),
	}
	return t
}

func parseColly(e *colly.HTMLElement) []string{

	var results []string

	// Maybe colly has some native support for what I'm doing, but I couldn't find it
	e.ForEach("td", func(_ int, td *colly.HTMLElement){
		results = append(results, td.Text)
	})


	page := e.ChildAttr("a", "href")
	link := fmt.Sprintf("%s%s", "https://www.trianglemtb.com/", page)

	results = append(results, link)
	return results
}

func FetchTrailStatus() []trail {

	trails []trail
	c := colly.NewCollector()

	// On every a element which has tr attribute call callback
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

		elements := parseColly(e)
		currentTrail := NewTrail(elements[0], elements[1], elements[2], elements[3])

		trails = append(trails, currentTrail)
	})

	c.Visit("https://www.trianglemtb.com/mobiletrailstatus.php")

	return trails
}
