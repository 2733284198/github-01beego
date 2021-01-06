package main

import (
	"github.com/gocolly/colly/v2"
	//"github.com/prometheus/common/log"
	"log"
	//"github.com/prometheus/common/log"
)

func main() {
	// create a new collector
	c := colly.NewCollector()

	// authenticate
	//err := c.Post("http://www.baidu.com/login", map[string]string{"username": "admin", "password": "admin"})
	err := c.Post("http://www.baidu.com", map[string]string{"username": "admin", "password": "admin"})
	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	// start scraping
	c.Visit("https://www.baidu.com/")
}
