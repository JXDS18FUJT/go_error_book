package wqmtcontrol

import (
	"fmt"
	"strings"
	"time"

	"347613781qq.com/goInit1/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func Brand(c *gin.Context) {
	co := colly.NewCollector()
	co.Limit(&colly.LimitRule{
		DomainGlob:  "*httpsbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})
	// model.TruncateComBrand(model.ComBrand{})
	co.OnHTML(".resp-tabs-container .resp-tab-content:nth-child(1) #CardSelectTr tbody", func(el *colly.HTMLElement) {
		//通用的烙印
		// var comBrand = model.ComBrand{}
		el.ForEach("tr", func(i int, ele *colly.HTMLElement) {

			if i > 0 {
				var comBrand = model.ComBrand{}
				comBrand.Access = ele.Attr("data-param2")
				ele.ForEach("td", func(j int, elem *colly.HTMLElement) {
					fmt.Printf("->%d\n%s", j, elem.Text)
					switch j {
					case 0:
						comBrand.ImgUrl = elem.ChildAttr("img", "src")
					case 1:
						comBrand.Name = elem.Text
					case 2:
						comBrand.LevelName = strings.TrimSpace(elem.Text)
						if comBrand.LevelName == "狂" {
							comBrand.Level = 3

						}
						if comBrand.LevelName == "危" {
							comBrand.Level = 2

						}
						if comBrand.LevelName == "普" {
							comBrand.Level = 1

						}
					case 3:
						comBrand.Suit1 = elem.Text
					case 4:
						comBrand.Suit2 = elem.Text
					case 5:
						comBrand.Suit3 = elem.Text
					case 6:
						comBrand.Effect = elem.Text
					case 7:
						comBrand.Desc = elem.Text
					default:

					}
				})
				model.CreateComBrand(&comBrand)
			}

		})

		// fmt.Printf("Link found: %q -> %s\n", el.Text, "s")

	})
	co.OnHTML(".resp-tabs-container .resp-tab-content:nth-child(2) #CardSelectTr tbody", func(el *colly.HTMLElement) {
		//专属烙印
		el.ForEach("tr", func(i int, ele *colly.HTMLElement) {
			if i > 0 {
				var exBrand = model.ExBrand{}
				ele.ForEach("td", func(j int, elem *colly.HTMLElement) {
					fmt.Printf("->%d\n%s", j, elem.Text)
					switch j {
					case 0:
						exBrand.Name = strings.TrimSpace(elem.Text)
						exBrand.ImgUrl = elem.ChildAttr("img", "src")
					case 1:
						exBrand.Character = strings.TrimSpace(elem.Text)
					case 2:
						exBrand.Effect = strings.TrimSpace(elem.Text)
					default:

					}

				})
				model.CreateExBrand(&exBrand)
			}

		})

	})
	model.TruncateExBrand()
	model.TruncateComBrand()
	co.Visit("https://wiki.biligame.com/wqmt/%E7%83%99%E5%8D%B0%E5%9B%BE%E9%89%B4")
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "成功",
	})
}
