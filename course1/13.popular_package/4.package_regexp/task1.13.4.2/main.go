package main

import (
	"fmt"
	"regexp"
)

type Ad struct {
	Title       string
	Description string
}

func main() {
	ads := []Ad{
		{
			Title:       "Куплю велосипед MeRiDa",
			Description: "Куплю велосипед meriDA в хорошем состоянии",
		},
		{
			Title:       "Продам ВаЗ 2101",
			Description: "Продам ваз 2101 в хорошем состоянии",
		},
		{
			Title:       "Продам БМВ",
			Description: "Продам бМв в хорошем состоянии",
		},
		{
			Title:       "Продам macBook pro",
			Description: "Продам macBook PRO в хорошем состоянии",
		},
	}

	cansor := map[string]string{
		"велосипед merida": "телефон Apple",
		"ваз":              "ВАЗ",
		"бмв":              "BMW",
		"macbook pro":      "Macbook Pro",
	}

	fmt.Println(censorAds(ads, cansor))
}

func censorAds(ads []Ad, censor map[string]string) []Ad {
	for illegal, legal := range censor {
		re := regexp.MustCompile(fmt.Sprintf(`(?i)%s`, illegal))

		for i := range ads {
			ads[i].Title = re.ReplaceAllString(ads[i].Title, legal)
			ads[i].Description = re.ReplaceAllString(ads[i].Description, legal)
		}
	}

	return ads
}
