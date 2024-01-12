package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"golang-webscraper/handlers"
	"os"
)

func main() {
	var pokemonProducts []handlers.PokemonProduct
	pagesToScrape := []string{}

	pageToScrape := "https://scrapeme.live/shop/page/1/"

	pagesDiscovered := []string{pageToScrape}

	i := 1
	limit := 5

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnHTML("a.page-numbers", handlers.HandlePagination(&pagesToScrape, &pagesDiscovered))
	c.OnHTML("li.product", handlers.HandleProduct(&pokemonProducts))

	c.OnScraped(handlers.HandleScraped(c, &pagesToScrape, &pageToScrape, &i, limit))

	err := c.Visit(pageToScrape)
	handlers.ErrorVisit(err)

	file, err := os.Create("products.csv")
	handlers.ErrorCreateFile(err)
	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}

	writer.Write(headers)
	for _, pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.Url,
			pokemonProduct.Image,
			pokemonProduct.Name,
			pokemonProduct.Price,
		}
		err := writer.Write(record)
		handlers.ErrorWriteFile(err)
	}
	defer writer.Flush()
}
