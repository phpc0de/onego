package fetcher

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"log"
)

func NewFetcher() *colly.Collector {
	c := colly.NewCollector()
  
	rule := &colly.LimitRule{
		RandomDelay: 10 * time.Second,
		Parallelism: 5, //并发数为5
	}
	_ = c.Limit(rule)
	extensions.Referer(c)
	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	return c
}
