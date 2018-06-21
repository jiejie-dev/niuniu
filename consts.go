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
