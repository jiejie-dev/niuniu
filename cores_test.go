package niuniu

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRandCardList(t *testing.T) {
	p1, p2 := RandCardList()
	assert.Equal(t, "p1", p1.Name)
	assert.Equal(t, "p2", p2.Name)
	assert.Equal(t, 5, len(p1.Cards))
	assert.Equal(t, 5, len(p2.Cards))
}
