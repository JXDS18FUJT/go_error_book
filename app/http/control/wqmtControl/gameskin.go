package wqmtcontrol

import (
	"strings"
	"time"

	"347613781qq.com/goInit1/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func Gameskin(c *gin.Context) {
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
			var gameskin = &model.Gameskin{}
			ele.ForEach("td", func(j int, elem *colly.HTMLElement) {
				switch j {
				case 0:
					gameskin.Character = elem.Text
				case 1:
					gameskin.Name = elem.Text
				case 2:
					gameskin.ImgUrl = elem.ChildAttr(".showOnImgSource", "src")
				case 3:
					gameskin.LevelName = strings.TrimSpace(elem.Text)
					if gameskin.LevelName == "三星" {
						gameskin.Level = 3

					}
					if gameskin.LevelName == "二星" {
						gameskin.Level = 2

					}
					if gameskin.LevelName == "一星" {
						gameskin.Level = 1

					}
				case 4:
					gameskin.Series = elem.Text
				case 5:
					gameskin.OnlineTime = elem.Text
				case 6:
					gameskin.Access = elem.Text
				default:
				}
				//fmt.Printf("->%s", gameskin.ImgUrl)

			})
			model.CreateGameskin(gameskin)

		})

	})
	co.Visit("https://wiki.biligame.com/wqmt/%E7%9A%AE%E8%82%A4%E5%9B%BE%E9%89%B4")
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "成功",
	})
}
