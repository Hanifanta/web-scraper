package handlers

import (
	"github.com/gocolly/colly"
)

type PokemonProduct struct {
	Url, Image, Name, Price string
}

func HandleProduct(pokemonProducts *[]PokemonProduct) colly.HTMLCallback {
	return func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}
		pokemonProduct.Url = e.ChildAttr("a", "href")
		pokemonProduct.Image = e.ChildAttr("img", "src")
		pokemonProduct.Name = e.ChildText("h2")
		pokemonProduct.Price = e.ChildText(".price")
		*pokemonProducts = append(*pokemonProducts, pokemonProduct)
	}
}
