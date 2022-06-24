package main

import (
	"baqiwaqi/scraper/middleware"
	"baqiwaqi/scraper/scraping"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func main() {
	colly := colly.NewCollector(colly.AllowedDomains("tweakers.net"))
	router := gin.Default()
	router.Use(middleware.ValidateApiKey())
	router.GET("/tweakers", func(c *gin.Context) {
		news := scraping.ScrapeWebForNews(colly)
		c.IndentedJSON(http.StatusOK, gin.H{
			"tweakers": news,
		})
	})

	router.Run()
}
