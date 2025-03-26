package wqmtcontrol

import (
	"strings"
	"time"

	"347613781qq.com/goInit1/app/model"
	"347613781qq.com/goInit1/app/service"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func CharacterProperty(c *gin.Context) {
	co := colly.NewCollector()
	co.Limit(&colly.LimitRule{
		DomainGlob:  "*httpsbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})
	co.OnHTML("#CardSelectTr .divsort", func(el *colly.HTMLElement) {
		var character = model.CharacterProperty{}
		el.ForEach(".mtd", func(i int, ele *colly.HTMLElement) {
			// if i == 0 {
			// 	name := strings.Replace(ele.Text, "新", "", 1)
			// 	name = strings.Replace(name, "概率提升", "", 1)
			// 	name = strings.TrimSpace(name)
			// 	character.Name = name
			// }

			switch i {
			case 0:
				name := strings.Replace(ele.Text, "新", "", 1)
				name = strings.Replace(name, "概率提升", "", 1)
				name = strings.TrimSpace(name)
				character.Character = strings.TrimSpace(name)
				character.HeadImage = ele.ChildAttr(".img-kk", "src")
			case 1:
				character.Code = strings.TrimSpace(ele.Text)
			case 2:
				character.Career = strings.TrimSpace(ele.Text)
			case 3:
				alt := ele.ChildAttr("img", "alt")
				levelName := strings.Replace(alt, "tj", "", 1)
				levelName = strings.Replace(levelName, ".png", "", 1)
				levelName = strings.TrimSpace(levelName)
				character.LevelName = levelName
				if levelName == "狂" {
					character.Level = 3

				}
				if levelName == "危" {
					character.Level = 2

				}
				if levelName == "普" {
					character.Level = 1

				}

			case 4:
				character.Camp = strings.TrimSpace(ele.Text)
			case 5:
				character.Hp = strings.TrimSpace(ele.Text)
			case 6:
				character.Atk = strings.TrimSpace(ele.Text)
			case 7:
				character.Def = strings.TrimSpace(ele.Text)
			case 8:
				character.Rgs = strings.TrimSpace(ele.Text)
			case 9:
				character.Block = strings.TrimSpace(ele.Text)
			case 10:
				character.Tag = strings.TrimSpace(ele.Text)

			default:
			}

		})
		model.CreateCharacterProperty(&character)

	})
	//清除对应的数据库
	model.TruncateCharacterProperty()
	co.Visit("https://wiki.biligame.com/wqmt/%E7%A6%81%E9%97%AD%E8%80%85")
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "成功",
	})

}

func CharacterSkill(c *gin.Context) {

	co := colly.NewCollector()
	co.Limit(&colly.LimitRule{
		DomainGlob:  "*httpsbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})
	var skills = &model.CharacterSkill{}
	skills.Character = "卓娅"
	skills.Code = "MBCC-S-028"
	model.FirstOrCreateCharacterSkill(skills)
	service.WqmtCharacterSkill(&struct {
		Character string
		Code      string
	}{
		Character: skills.Character,
		Code:      skills.Code,
	}, &colly.LimitRule{
		DomainGlob:  "*httpsbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})
	// error handling...

	// co.OnHTML(".col-sm-6 .main-line-wrap .resp-tabs .resp-tabs-container .resp-tab-content td", func(el *colly.HTMLElement) {
	// 	// var html, _ = el.DOM.Html()

	// })

	// co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(1) tr", func(el *colly.HTMLElement) {
	// 	//爬取普攻名字
	// 	fmt.Printf("%s----(%v)\n", el.Text, el.Index)
	// 	switch el.Index {
	// 	case 0:
	// 		skills.BasicAttackName = strings.TrimSpace(el.Text)

	// 	case 1:
	// 		skills.BasicAttackDesc = strings.TrimSpace(el.Text)
	// 		model.UpdatesCharacterSkill(skills)
	// 	default:

	// 	}

	// })

	//co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(1) tr td>div", func(el *colly.HTMLElement) {
	//普攻爬取lv2-lv10
	//fmt.Printf("%s----(%v)\n", el.Text, el.Index)
	// var lvDesc = ""
	// switch el.Index {
	// case 10, 11, 12, 13, 14, 15, 16, 17:
	// 	skills.BasicAttackLvDesc = skills.BasicAttackLvDesc + "LV" + el.Text + ","
	// case 18:
	// 	skills.BasicAttackLvDesc = skills.BasicAttackLvDesc + "LV" + el.Text
	// 	model.UpdatesCharacterSkill(skills)
	// default:

	//}

	//})

	// co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(2) tr", func(el *colly.HTMLElement) {
	// 	//爬取必杀名字
	// 	fmt.Printf("%s----(%v)\n", el.Text, el.Index)

	// 	switch el.Index {
	// 	case 0:
	// 		skills.ChargeSkillName = strings.TrimSpace(el.Text)

	// 	case 1:
	// 		skills.ChargeSkillTag = strings.TrimSpace(el.Text)
	// 	case 2:
	// 		skills.ChargeSkillDesc = strings.TrimSpace(el.Text)
	// 		model.UpdatesCharacterSkill(skills)
	// 	default:

	// 	}

	// })

	// co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(2) tr td>div", func(el *colly.HTMLElement) {
	// 	//必杀爬取lv2-lv10
	// 	fmt.Printf("%s----(%v)\n", el.Text, el.Index)
	// 	// var lvDesc = ""
	// 	switch el.Index {
	// 	case 0, 1, 2, 3, 4, 5, 6, 7:
	// 		skills.ChargeSkillLvDesc = skills.ChargeSkillLvDesc + "LV" + el.Text + ","
	// 	case 8:
	// 		skills.ChargeSkillLvDesc = skills.ChargeSkillLvDesc + "LV" + el.Text
	// 		model.UpdatesCharacterSkill(skills)
	// 	default:

	// 	}

	// })

	// co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(3) tr", func(el *colly.HTMLElement) {
	// 	//爬取被动1名字
	// 	fmt.Printf("%s----(%v)\n", el.Text, el.Index)

	// 	switch el.Index {
	// 	case 0:
	// 		skills.PassiveSkill1Name = strings.TrimSpace(el.Text)

	// 	case 1:
	// 		skills.PassiveSkill1Desc = strings.TrimSpace(el.Text)
	// 		model.UpdatesCharacterSkill(skills)
	// 	default:

	// 	}

	// })

	// co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(3) tr td>div", func(el *colly.HTMLElement) {
	// 	爬取被动1lv2-lv10
	// 	fmt.Printf("%s----(%v)\n", el.Text, el.Index)
	// 	// var lvDesc = ""
	// 	switch el.Index {
	// 	case 0, 1, 2, 3, 4, 5, 6, 7:
	// 		skills.PassiveSkill1LvDesc = skills.PassiveSkill1LvDesc + "LV" + el.Text + ","
	// 	case 8:
	// 		skills.PassiveSkill1LvDesc = skills.PassiveSkill1LvDesc + "LV" + el.Text
	// 		model.UpdatesCharacterSkill(skills)
	// 	default:

	// 	}

	// })

	// co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(3) tr", func(el *colly.HTMLElement) {
	// 	//爬取被动2名字
	// 	fmt.Printf("%s----(%v)\n", el.Text, el.Index)

	// 	switch el.Index {
	// 	case 0:
	// 		skills.PassiveSkill2Name = strings.TrimSpace(el.Text)

	// 	case 1:
	// 		skills.PassiveSkill2Desc = strings.TrimSpace(el.Text)
	// 		model.UpdatesCharacterSkill(skills)
	// 	default:

	// 	}

	// })

	// co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(3) tr td>div", func(el *colly.HTMLElement) {
	// 	//爬取被动2lv2-lv10
	// 	fmt.Printf("%s----(%v)\n", el.Text, el.Index)
	// 	// var lvDesc = ""
	// 	switch el.Index {
	// 	case 0, 1, 2, 3, 4, 5, 6, 7:
	// 		skills.PassiveSkill2LvDesc = skills.PassiveSkill2LvDesc + "LV" + el.Text + ","
	// 	case 8:
	// 		skills.PassiveSkill2LvDesc = skills.PassiveSkill2LvDesc + "LV" + el.Text
	// 		model.UpdatesCharacterSkill(skills)
	// 	default:

	// 	}

	// })

	// co.Visit("https://wiki.biligame.com/wqmt/%E7%A6%81%E9%97%AD%E8%80%85:" + skills.Character)
	// co.OnError(func(r *colly.Response, err error) {
	// 	fmt.Printf("发生错误,指定网页不存在")

	// })
	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "成功",
	})
}
