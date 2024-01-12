package handlers

import (
	"github.com/gocolly/colly"
	"log"
)

func HandleScraped(c *colly.Collector, pagesToScrape *[]string, pageToScrape *string, i *int, limit int) colly.ScrapedCallback {
	return func(response *colly.Response) {
		if len(*pagesToScrape) != 0 && *i < limit {
			*pageToScrape = (*pagesToScrape)[0]
			*pagesToScrape = (*pagesToScrape)[1:]
			*i++
			err := c.Visit(*pageToScrape)
			if err != nil {
				log.Printf("Failed to visit page: %v", err)
			}
		}
	}
}
