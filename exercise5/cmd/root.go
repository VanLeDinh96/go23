package cmd

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func Execute() {
	c := colly.NewCollector()

	c.OnHTML(".gs-c-promo-heading__title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Println(title)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("https://www.bbc.com/news")
	if err != nil {
		log.Fatal(err)
	}
}