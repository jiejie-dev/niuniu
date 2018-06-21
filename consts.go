package niuniu

const (
	CardColorSpade = 4
	CardColorHeart = 3
	CardColorPlum  = 2
	CardColorBlock = 1

	CardValueA  = 1
	CardValue2  = 2
	CardValue3  = 3
	CardValue4  = 4
	CardValue5  = 5
	CardValue6  = 6
	CardValue7  = 7
	CardValue8  = 8
	CardValue9  = 9
	CardValue10 = 10
	CardValueJ  = 11
	CardValueQ  = 12
	CardValueK  = 13

	CardsCount         = 52
	PerPlayerCardCount = 5

	NotNiu      = 0
	Niu1      = 1
	Niu2      = 2
	Niu3      = 3
	Niu4      = 4
	Niu5      = 5
	Niu6      = 6
	Niu7      = 7
	Niu8      = 8
	Niu9      = 9
	NiuNiu    = 10
	SilverNiu = 11
	GoldNiu   = 12
	Bomb      = 13
	SmallNiu  = 14
)

var CardsData []CardInfo

var TransColor = map[int]string{
	4: "黑桃",
	3: "红桃",
	2: "梅花",
	1: "方块",
}

var TransValue = map[int]string{
	1:  "A",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "10",
	11: "J",
	12: "Q",
	13: "K",
}

var TransNiu = map[int]string{
	0:  "没牛",
	1:  "牛一",
	2:  "牛二",
	3:  "牛三",
	4:  "牛四",
	5:  "牛五",
	6:  "牛六",
	7:  "牛七",
	8:  "牛八",
	9:  "牛九",
	10: "牛牛",
	11: "银牛",
	12: "金牛",
	13: "炸弹",
	14: "五小牛",
}

func Init52Cards() {
	CardsData = make([]CardInfo, 0)
	for color := 1; color < 5; color++ {
		for val := 1; val < 14; val++ {
			CardsData = append(CardsData, CardInfo{
				CardColor: color,
				CardValue: val,
				CardCount: getCountByValue(val),
			})
		}
	}
}

func init() {
	Init52Cards()
}

type CardInfo struct {
	CardValue int
	CardColor int
	CardCount int
}

type Player struct {
	Name  string
	Cards []CardInfo
}
