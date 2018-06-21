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
