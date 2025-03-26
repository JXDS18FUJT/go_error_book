package wqmtcontrol

import (
	"strings"
	"time"

	"347613781qq.com/goInit1/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func Achieve(c *gin.Context) {

	co := colly.NewCollector()
	co.Limit(&colly.LimitRule{
		DomainGlob:  "*httpsbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})
	co.OnHTML("#CardSelectTr", func(el *colly.HTMLElement) {
		el.ForEach("tr", func(i int, ele *colly.HTMLElement) {
			if i == 0 {
				return
			}
			var achieve = &model.Achieve{}
			ele.ForEach("td", func(j int, elem *colly.HTMLElement) {
				switch j {
				case 0:
					achieve.ImgUrl = elem.ChildAttr("img", "src")
				case 1:
					achieve.Name = strings.TrimSpace(elem.Text)
				case 2:
					achieve.LevelName = strings.TrimSpace(elem.Text)
					if achieve.LevelName == "狂" {
						achieve.Level = 3

					} else if achieve.LevelName == "危" {
						achieve.Level = 2

					} else if achieve.LevelName == "普" {
						achieve.Level = 1

					}
				case 3:
					achieve.Access = elem.Text
				case 4:
					achieve.UnlockDesc = elem.Text
				default:

				}

			})
			model.CreateAchieve(achieve)
		})

	})
	model.TruncateAchieve()
	co.Visit("https://wiki.biligame.com/wqmt/%E7%A7%B0%E5%8F%B7")
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "成功",
	})

}
