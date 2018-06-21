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
