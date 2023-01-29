package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"os"
)

type Theater struct { // define struct
	Title       string
	URL         string
	Description string
	URLpath     string //backup to get url in case of error
}

func colly1() {
	c2 := colly.NewCollector()
	allArticles := make([]Theater, 0)
	c2.OnHTML("article._article_nod91_1", func(e *colly.HTMLElement) {
		e.ForEach("div._articleContent_nod91_27", func(i int, h *colly.HTMLElement) {
			article := Theater{}
			article.Title = e.ChildText("h3._h3_cuogz_1")
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
	c2.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c2.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	c2.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		js, err := json.MarshalIndent(allArticles, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("allPosts.txt", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}
	})
	c2.Visit("https://www.timeout.com/barcelona")
}
func colly2() {
	c := colly.NewCollector() // Create new collector
	allOthers := make([]Theater, 0)
	c.OnHTML("article.tile", func(e *colly.HTMLElement) {
		e.ForEach("div._articleContent_1pzwm_26", func(i int, h *colly.HTMLElement) {
			article := Theater{}
			article.Title = e.ChildText("a._titleLinkContainer_1pzwm_44") // Title HTML element
			article.URL = e.ChildAttr("a", "href")
			url, _ := url.Parse(article.URL) // parse for anything resembling a url
			article.Description = e.ChildText("p._p_1vat8_1")
			articles := Theater{
				Title:       article.Title,
				URL:         article.URL,
				Description: article.Description,
				URLpath:     url.Path,
			}
			allOthers = append(allOthers, articles) // Append to list for saving

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
		js, err := json.MarshalIndent(allOthers, "", " ") // convert to json
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("culturePosts.txt", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}

	})
	c.Visit("https://www.timeout.com/barcelona/culture")
}

func main() {
	colly1()
	colly2()

}
