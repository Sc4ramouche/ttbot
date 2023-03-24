package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("sttkberlin.de"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("href") == "https://sttkberlin.de/corona/anfrage-erwachsenen-training-freizeit-2/" {
			link := e.Attr("href")
			fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://sttkberlin.de/trainingszeiten/")
}
