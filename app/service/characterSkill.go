package service

import (
	"time"

	"github.com/gocolly/colly"
)

type GetCharacterSkillQuery struct {
	character string
	code      string
}

func GetCharacterSkill(limitRule *colly.LimitRule) {
	co := colly.NewCollector()
	co.Limit(&colly.LimitRule{
		DomainGlob:  "*httpsbin.*",
		Parallelism: 2,
		RandomDelay: 2 * time.Second,
	})

}
