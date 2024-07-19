package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Config struct {
	URL                string            `json:"url"`
	ContainerSelector  string            `json:"container-selector"`
	PaginationSelector string            `json:"pagination-selector"`
	AllowPagination    bool              `json:"allow-pagination"`
	Selectors          map[string]string `json:"selectors"`
}

func main() {

	configFile := flag.String("config", "config.json", "./")
	flag.Parse()

	file, err := os.ReadFile(*configFile)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return
	}

	c := colly.NewCollector()

	var items []map[string]string

	c.OnHTML(config.ContainerSelector, func(h *colly.HTMLElement) {
		item := make(map[string]string)
		for key, selector := range config.Selectors {
			if key == "imgurl" {
				item[key] = h.ChildAttr(selector, "src")
			} else {
				item[key] = h.ChildText(selector)
			}
		}
		items = append(items, item)
	})

	if config.AllowPagination {
		c.OnHTML(config.PaginationSelector, func(h *colly.HTMLElement) {
			nextPage := h.Request.AbsoluteURL(h.Attr("href"))
			c.Visit(nextPage)
		})
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	c.Visit(config.URL)

	content, err := json.Marshal(items)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile("products.json", content, 0664)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
}
