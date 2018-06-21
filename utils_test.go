package niuniu

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	results := Shuffle(CardsData)
	t.Log(len(results))
	for i := 0; i < len(results); i++ {
		t.Logf("Value: %d  Count: %d  Color: %d", results[i].CardValue, results[i].CardCount, results[i].CardColor)
	}
	assert.Equal(t, len(results), 52)
}

func TestTrans(t *testing.T) {
	r := Trans(1, 1)
	assert.Equal(t, "方块 A", r)
}

func TestCompareCards(t *testing.T) {
	a := CardInfo{
		CardValue: CardValue9,
		CardCount: CardValue9,
		CardColor: CardColorSpade,
	}
	b := CardInfo{
		CardValue: CardValue6,
		CardCount: CardValue6,
		CardColor: CardColorSpade,
	}
	flag := CompareCards(a, b)
	assert.True(t, flag)
	if flag {
		t.Log(flag)
	}
	// TODO: more sample
}

func TestIsSmallNiu(t *testing.T) {
	cards := []CardInfo{
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValue2,
			CardCount: CardValue2,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValue3,
			CardCount: CardValue3,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
	}
	flag := IsSmallNiu(cards)
	assert.True(t, flag)
}

func TestIsBomb(t *testing.T) {
	cards := []CardInfo{
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
	}
	flag := IsBomb(cards)
	assert.True(t, flag)
}

func TestIsGoldNiu(t *testing.T) {
	cards := []CardInfo{
		{
			CardValue: CardValueJ,
			CardCount: CardValueJ,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueJ,
			CardCount: CardValueJ,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueJ,
			CardCount: CardValueJ,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueQ,
			CardCount: CardValueQ,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueK,
			CardCount: CardValueK,
			CardColor: CardColorSpade,
		},
	}
	flag := IsGoldNiu(cards)
	assert.True(t, flag)
}

func TestIsSilverNiu(t *testing.T) {
	cards := []CardInfo{
		{
			CardValue: CardValue10,
			CardCount: CardValue10,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueJ,
			CardCount: CardValueJ,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueJ,
			CardCount: CardValueJ,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueJ,
			CardCount: CardValueJ,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueJ,
			CardCount: CardValueJ,
			CardColor: CardColorSpade,
		},
	}
	flag := IsSilverNiu(cards)
	assert.True(t, flag)
}

func TestSortCards(t *testing.T) {
	cards := []CardInfo{
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValue2,
			CardCount: CardValue2,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValue3,
			CardCount: CardValue3,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
	}
	cards = SortCards(cards)
	m := []int{
		1, 1, 1, 2, 3,
	}
	for i := 0; i < len(cards); i++ {
		assert.Equal(t, m[i], cards[i].CardValue)
	}
}

func TestBankerIsWin(t *testing.T) {
	cards := []CardInfo{
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValue2,
			CardCount: CardValue2,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValue3,
			CardCount: CardValue3,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
	}
	cards2 := []CardInfo{
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValue2,
			CardCount: CardValue2,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
	}
	flag := BankerIsWin(cards2, cards)
	assert.True(t, flag)
}

func TestGetNiu(t *testing.T) {
	cards := []CardInfo{
		{
			CardValue: CardValue10,
			CardCount: CardValue10,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValue10,
			CardCount: CardValue10,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValue10,
			CardCount: CardValue10,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
	}
	niu := GetNiu(cards)
	assert.Equal(t, 2, niu)
}

func TestGetNiuType(t *testing.T) {
	cards := []CardInfo{
		{
			CardValue: CardValue10,
			CardCount: CardValue10,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValue10,
			CardCount: CardValue10,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValue10,
			CardCount: CardValue10,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
		{
			CardValue: CardValueA,
			CardCount: CardValueA,
			CardColor: CardColorSpade,
		},
	}
	niu := GetNiuType(cards)
	assert.Equal(t, 2, niu)
}
