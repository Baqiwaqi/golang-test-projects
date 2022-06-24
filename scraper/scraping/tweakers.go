package scraping

import "github.com/gocolly/colly"

var (
	news []TweakersNews
)

// TweakersNews Struct

type TweakersNews struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	Lead  string `json:"lead"`
}

func ScrapeWebForNews(c *colly.Collector) []TweakersNews {
	return scrapeTweakers(c)
}

func scrapeTweakers(c *colly.Collector) []TweakersNews {
	c.OnHTML("div[class=newsContentBlock]", func(h *colly.HTMLElement) {
		title := h.ChildText("h1")
		link := h.ChildAttr("a[href]", "href")
		lead := h.ChildText("p[class=lead]")
		news = append(news, TweakersNews{title, link, lead})
	})
	c.Visit("https://tweakers.net/nieuws/list/")
	return news
	// writeJson("tweakers.json", news)
	// 	fmt.Println(h.ChildText("H1"))
	// 	fmt.Println(h.ChildAttr("a[href]", "href"))
	// 	fmt.Println(h.ChildText("p[class=lead]"))
	// 	fmt.Println("")
}
