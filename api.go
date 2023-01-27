package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"os"
)

type Theater struct {
	Title       string
	URL         string
	Description string
	URLpath     string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	allArticles := make([]Theater, 0)
	c.OnHTML("article.tile", func(e *colly.HTMLElement) {
		e.ForEach("div._articleContent_1pzwm_26", func(i int, h *colly.HTMLElement) {
			article := Theater{}
			article.Title = e.ChildText("a._titleLinkContainer_1pzwm_44")
			article.URL = e.ChildAttr("a", "href")
			url, _ := url.Parse(article.URL)
			article.Description = e.ChildText("p._p_1vat8_1")
			articles := Theater{
				Title:       article.Title,
				URL:         article.URL,
				Description: article.Description,
				URLpath:     url.Path,
			}
			allArticles = append(allArticles, articles)

		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		js, err := json.MarshalIndent(allArticles, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("posts.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
			c.Visit("https://www.timeout.com/barcelona/culture")
		}

	})
	c.Visit("https://www.timeout.com/barcelona/culture")
}
