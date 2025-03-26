package apiControl

import (
	"fmt"
	"strings"
	"time"

	"347613781qq.com/goInit1/app/model"
	"347613781qq.com/goInit1/app/service"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func Colly(c *gin.Context) {
	co := colly.NewCollector()
	co.Limit(&colly.LimitRule{
		DomainGlob:  "*httpsbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// Find and visit all links
	co.OnHTML("#CardSelectTr .divsort", func(el *colly.HTMLElement) {
		title := ""
		title = el.Text
		// Print link
		titleArr := strings.Split(title, "\n")
		service.WqmtParseUpPool(titleArr)
		fmt.Println(titleArr[3])
		// Visit link found on page on a new thread
		// e.Request.Visit(link)
	})

	co.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	model.TruncateUpPool()
	co.Visit("https://wiki.biligame.com/wqmt/%E8%BF%BD%E8%B8%AA%E4%B8%80%E8%A7%88")
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "成功",
	})

}
