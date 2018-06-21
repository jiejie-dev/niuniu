package niuniu

import (
	"math/rand"
	"time"
	"fmt"
)

func getCountByValue(value int) int {
	if value > 10 {
		return 10
	}
	return value
}

func Shuffle(vals []CardInfo) []CardInfo {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < len(vals); i++ {
		n := len(vals)
		randIndex := r.Intn(n - i)
		vals[n-1-i], vals[randIndex] = vals[randIndex], vals[n-1-i]
	}
	return vals
}

func Trans(color, value int) string {
	return fmt.Sprintf("%s %s", TransColor[color], TransValue[value])
}

func CompareCards(a, b CardInfo) bool {
	if a.CardValue < b.CardValue {
		return true
	}
	if a.CardValue < b.CardValue {
		return false
	}
	return a.CardColor > b.CardColor
}

func IsSmallNiu(cards []CardInfo) bool {
	sum := 0
	for _, item := range (cards) {
		sum += item.CardCount
	}
	if sum <= 10 {
		return true
	}
	return false
}

func IsBomb(cards []CardInfo) bool {
	if cards[0].CardCount == cards[3].CardCount {
		return true
	} else if cards[1].CardCount == cards[4].CardCount {
		return true
	}
	return false
}

func IsGoldNiu(cards []CardInfo) bool {
	return cards[1].CardValue > 10
}

func IsSilverNiu(cards []CardInfo) bool {
	return cards[1].CardValue > 10 && cards[0].CardValue == 10
}

func GetNiu(cards []CardInfo) int {
	lave := 0
	for i := 1; i < len(cards); i++ {
		lave = lave + cards[i].CardValue
	}
	lave = lave % 10
	for i := 0; i < len(cards)-1; i++ {
		for j := i + 1; j < len(cards); j++ {
			if ((cards[i].CardCount+cards[j].CardCount)%10 == lave) {
				if lave == 0 {
					return 10
				}
				return lave
			}
		}
	}
	return 0
}

func GetNiuType(cards []CardInfo) int {
	// TODO: sort cards
	if IsSmallNiu(cards) {
		return SmallNiu
	}
	if IsBomb(cards) {
		return Bomb
	}
	if IsGoldNiu(cards) {
		return GoldNiu
	}
	if IsSilverNiu(cards) {
		return SilverNiu
	}
	return GetNiu(cards)
}

func SortCards(cards []CardInfo) []CardInfo {
	length := len(cards)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			temp := cards[i]
			if cards[i].CardCount < cards[j].CardCount {
				cards[i] = cards[j]
				cards[j] = temp
			}
		}
	}
	return cards
}
