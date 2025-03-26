package wqmtcontrol

import (
	"fmt"
	"strings"
	"time"

	"347613781qq.com/goInit1/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func Uppool(c *gin.Context) {
	co := colly.NewCollector()
	co.Limit(&colly.LimitRule{
		DomainGlob:  "*httpsbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// Find and visit all links
	co.OnHTML("#CardSelectTr .divsort", func(el *colly.HTMLElement) {
		var upPool = &model.UpPool{}
		upPool.KindName = el.Attr("data-param1")
		el.ForEach("td", func(i int, elem *colly.HTMLElement) {

			switch i {
			case 0:
				upPool.Name = strings.TrimSpace(elem.Text)
				// case 1:
				// 	if strings.Contains(elem.Text, "常规追踪") {
				// 		upPool.Kind = 1
				// 		upPool.KindName = "常规追踪"

				// 	} else if strings.Contains(elem.Text, "活动追踪") {
				// 		upPool.Kind = 2
				// 		upPool.KindName = "活动追踪"

				// 	} else if strings.Contains(elem.Text, "限定追踪") {
				// 		upPool.Kind = 3
				// 		upPool.KindName = "限定追踪"

				// 	} else if strings.Contains(elem.Text, "统合追踪") {
				// 		upPool.Kind = 4
				// 		upPool.KindName = "统合追踪"

				// 	} else if strings.Contains(elem.Text, "定向追踪") {
				// 		upPool.Kind = 5
				// 		upPool.KindName = "定向追踪"
				// 	} else if strings.Contains(elem.Text, "返场·限定追踪") {
				// 		upPool.Kind = 6
				// 		upPool.KindName = "返场·限定追踪"
				// 	}
				upPool.Remark = elem.Text
			case 2:
				upPool.PoolCharacterImgUrl = strings.Join(elem.ChildAttrs("img", "src"), ",")
			default:

			}

		})
		model.CreateUpPool(upPool)
		// title := ""
		// title = el.Text
		// // Print link
		// titleArr := strings.Split(title, "\n")
		// service.WqmtParseUpPool(titleArr)
		// fmt.Println(titleArr[3], len(titleArr))
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
