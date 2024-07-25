package main

import (
	"reflect"
	"testing"
)

func Test_censorAds(t *testing.T) {
	type args struct {
		ads    []Ad
		censor map[string]string
	}
	tests := []struct {
		name string
		args args
		want []Ad
	}{
		{
			name: "Regular test",
			args: args{
				ads: []Ad{
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
				},
				censor: map[string]string{
					"велосипед merida": "телефон Apple",
					"ваз":              "ВАЗ",
					"бмв":              "BMW",
					"macbook pro":      "Macbook Pro",
				},
			},
			want: []Ad{
				{
					Title:       "Куплю телефон Apple",
					Description: "Куплю телефон Apple в хорошем состоянии",
				},
				{
					Title:       "Продам ВАЗ 2101",
					Description: "Продам ВАЗ 2101 в хорошем состоянии",
				},
				{
					Title:       "Продам BMW",
					Description: "Продам BMW в хорошем состоянии",
				},
				{
					Title:       "Продам Macbook Pro",
					Description: "Продам Macbook Pro в хорошем состоянии",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := censorAds(tt.args.ads, tt.args.censor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("censorAds() = %v, want %v", got, tt.want)
			}
		})
	}
}
