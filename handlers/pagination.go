package handlers

import "github.com/gocolly/colly"

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func HandlePagination(pagesToScrape, pagesDiscovered *[]string) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {
		newPaginationLink := e.Attr("href")
		if !contains(*pagesToScrape, newPaginationLink) {
			if !contains(*pagesDiscovered, newPaginationLink) {
				*pagesToScrape = append(*pagesToScrape, newPaginationLink)
			}
			*pagesDiscovered = append(*pagesDiscovered, newPaginationLink)
		}
	}
}
