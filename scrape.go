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
	Notes string `json:"notes"`
	Date   string `json:"last_updated"`
	Link   string `json:"trail_info"`
}

func NewTrail(name string, status string, date string, link string, notes string) trail {

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
		Notes: notes,
		Link:   link,
		Status: strings.TrimSpace(status),
	}
	return t
}


func parseColly(e *colly.HTMLElement) (details []string){
	var note string

	// Maybe colly has some native support for what I'm doing, but I couldn't find it
	e.ForEach("td", func(_ int, td *colly.HTMLElement){
		nAttr := td.Attr("onclick")
		if len(nAttr) > 0 {
			note = nAttr[strings.Index(nAttr, "'") + 1 : strings.Index(nAttr, ",") - 1]
		}
		details = append(details, td.Text)
	})


	page := e.ChildAttr("a", "href")
	link := fmt.Sprintf("%s%s", "https://www.trianglemtb.com/", page)

	details = append(details, link)
	details = append(details, note)
	return details
}

func FetchTrailStatus() []trail {


	notes := make(map[string]string)
	var trails []trail

	c := colly.NewCollector()

	// It's not pretty - but we need to build up the map of all notes before
	// we build out each trail object
	c.OnHTML("tr", func(e *colly.HTMLElement){
		if len(e.Attr("id")) > 0 {
			notes[e.Attr("id")] = e.Text
		}
	})

	// On every a element which has tr attribute call callback
	c.OnHTML("tr", func(e *colly.HTMLElement) {

		// tr elements with an id have already been handled
		if len(e.Attr("id")) > 0 {
			return
		}

		// Rules out the header and footer classes on the page
		if e.ChildAttr("td","class") != "trail" {
			return
		}

		elements := parseColly(e)
		currentTrail := NewTrail(elements[0], elements[1], elements[2], elements[3], notes[elements[4]])
		trails = append(trails, currentTrail)
	})

	c.Visit("https://www.trianglemtb.com/mobiletrailstatus.php")

	return trails
}
