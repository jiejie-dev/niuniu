package niuniu

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_16(t *testing.T) {
	a := 0x21
	t.Log(a/16 + 1)
	b := a ^ 0x0F
	t.Log(b)
}

func TestInit52Cards(t *testing.T) {
	Init52Cards()
	for i := 0; i < len(CardsData); i++ {
		t.Logf("Value: %d  Count: %d  Color: %d", CardsData[i].CardValue, CardsData[i].CardCount, CardsData[i].CardColor)
	}
	assert.Equal(t, 52, len(CardsData))
}
