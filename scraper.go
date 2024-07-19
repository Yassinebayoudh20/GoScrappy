package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Item struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	ImageUrl    string `json:"imgurl"`
	Reviews     string `json:"reviews"`
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("webscraper.io"),
	)

	var items []Item

	c.OnHTML("div.product-wrapper.card-body", func(h *colly.HTMLElement) {
		item := getItemsFromWeb(h)
		items = append(items, item)
	})

	c.OnHTML("a[rel=next]", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://webscraper.io/test-sites/e-commerce/static/computers/laptops")

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("products.json", content, 0664)
}

func getItemsFromWeb(h *colly.HTMLElement) Item {
	Name := h.ChildText("h4 > a.title")
	Price := h.ChildText("h4.price")
	Description := h.ChildText("p.description")
	ImageUrl := h.ChildAttr("img", "src")
	Reviews := h.ChildText("p.review-count")

	item := Item{
		Name:        Name,
		Price:       Price,
		Description: Description,
		ImageUrl:    ImageUrl,
		Reviews:     Reviews,
	}

	return item
}
