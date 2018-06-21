package niuniu

func RandCardList() (p1, p2 Player) {
	cards := Shuffle(CardsData)
	return Player{
		Name:  "p1",
		Cards: cards[:5],
	}, Player{
		Name:  "p2",
		Cards: cards[5:10],
	}
}
