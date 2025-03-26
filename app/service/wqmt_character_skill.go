package service

import (
	"strings"

	"347613781qq.com/goInit1/app/model"
	"github.com/gocolly/colly"
)

func WqmtCharacterSkill(character *struct {
	Character string
	Code      string
}, limitRule *colly.LimitRule) {
	co := colly.NewCollector()
	co.Limit(limitRule)
	var skills = &model.CharacterSkill{}
	skills.Character = character.Character
	skills.Code = character.Code
	model.FirstOrCreateCharacterSkill(skills)

	co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(1) tr", func(el *colly.HTMLElement) {
		//爬取普攻名字
		// fmt.Printf("%s----(%v)\n", el.Text, el.Index)
		switch el.Index {
		case 0:
			skills.BasicAttackName = strings.TrimSpace(el.Text)

		case 1:
			skills.BasicAttackDesc = strings.TrimSpace(el.Text)
			model.UpdatesCharacterSkill(skills)
		default:

		}

	})

	co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(1) tr td>div", func(el *colly.HTMLElement) {
		//普攻爬取lv2 - lv10
		// fmt.Printf("%s----(%v)\n", el.Text, el.Index)

		switch el.Index {
		case 10, 11, 12, 13, 14, 15, 16, 17:
			skills.BasicAttackLvDesc = skills.BasicAttackLvDesc + "LV" + el.Text + ","
		case 18:
			skills.BasicAttackLvDesc = skills.BasicAttackLvDesc + "LV" + el.Text
			model.UpdatesCharacterSkill(skills)
		default:

		}

	})

	co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(2) tr", func(el *colly.HTMLElement) {
		//爬取必杀名字
		// fmt.Printf("%s----(%v)\n", el.Text, el.Index)

		switch el.Index {
		case 0:
			skills.ChargeSkillName = strings.TrimSpace(el.Text)

		case 1:
			skills.ChargeSkillTag = strings.TrimSpace(el.Text)
		case 2:
			skills.ChargeSkillDesc = strings.TrimSpace(el.Text)
			model.UpdatesCharacterSkill(skills)
		default:

		}

	})

	co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(2) tr td>div", func(el *colly.HTMLElement) {
		//必杀爬取lv2-lv10
		// fmt.Printf("%s----(%v)\n", el.Text, el.Index)
		//
		switch el.Index {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			skills.ChargeSkillLvDesc = skills.ChargeSkillLvDesc + "LV" + el.Text + ","
		case 8:
			skills.ChargeSkillLvDesc = skills.ChargeSkillLvDesc + "LV" + el.Text
			model.UpdatesCharacterSkill(skills)
		default:

		}

	})

	co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(3) tr", func(el *colly.HTMLElement) {
		//爬取被动1名字
		// fmt.Printf("%s----(%v)\n", el.Text, el.Index)

		switch el.Index {
		case 0:
			skills.PassiveSkill1Name = strings.TrimSpace(el.Text)

		case 1:
			skills.PassiveSkill1Desc = strings.TrimSpace(el.Text)
			model.UpdatesCharacterSkill(skills)
		default:

		}

	})

	co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(3) tr td>div", func(el *colly.HTMLElement) {
		//爬取被动1lv2 - lv10
		// fmt.Printf("%s----(%v)\n", el.Text, el.Index)
		//
		switch el.Index {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			skills.PassiveSkill1LvDesc = skills.PassiveSkill1LvDesc + "LV" + el.Text + ","
		case 8:
			skills.PassiveSkill1LvDesc = skills.PassiveSkill1LvDesc + "LV" + el.Text
			model.UpdatesCharacterSkill(skills)
		default:

		}

	})

	co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(4) tr", func(el *colly.HTMLElement) {
		//爬取被动2名字
		// fmt.Printf("%s----(%v)\n", el.Text, el.Index)

		switch el.Index {
		case 0:
			skills.PassiveSkill2Name = strings.TrimSpace(el.Text)

		case 1:
			skills.PassiveSkill2Desc = strings.TrimSpace(el.Text)
			model.UpdatesCharacterSkill(skills)
		default:

		}

	})

	co.OnHTML(".col-sm-6 h2+div>.resp-tabs>.resp-tabs-container>.resp-tab-content>.wikitable:nth-child(4) tr td>div", func(el *colly.HTMLElement) {
		//爬取被动2lv2-lv10
		// fmt.Printf("%s----(%v)\n", el.Text, el.Index)
		//
		switch el.Index {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			skills.PassiveSkill2LvDesc = skills.PassiveSkill2LvDesc + "LV" + el.Text + ","
		case 8:
			skills.PassiveSkill2LvDesc = skills.PassiveSkill2LvDesc + "LV" + el.Text
			model.UpdatesCharacterSkill(skills)
		default:

		}

	})

	co.Visit("https://wiki.biligame.com/wqmt/%E7%A6%81%E9%97%AD%E8%80%85:" + skills.Character)
	co.OnError(func(r *colly.Response, err error) {
		// fmt.Printf("发生错误,指定网页不存在")

	})

}
