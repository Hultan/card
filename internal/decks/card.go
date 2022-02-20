package decks

import (
	"fmt"
)

const NumberOfCardsInDeck = 52
const NumberOfColors = 4
const NumberOfRanks = 13

var suitNames map[int]string
var rankNames map[int]string

type Suit int

const (
	SuitHeart Suit = iota
	SuitSpade
	SuitDiamond
	SuitClubs
)

type Rank int

const (
	RankTwo Rank = iota
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
	RankTen
	RankJack
	RankQueen
	RankKing
	RankAce
)

type Card struct {
	deck     int
	card     int
	isFaceUp bool
}

func (c *Card) FaceUp() bool {
	return c.isFaceUp
}

func (c *Card) Suit() Suit {
	return Suit(c.card / 13)
}

func (c *Card) Rank() Rank {
	return Rank(c.card % 13)
}

func (c *Card) SuitName() string {
	return suitNames[int(c.Suit())]
}

func (c *Card) RankName() string {
	return rankNames[int(c.Rank())]
}

func (c *Card) String() string {
	return fmt.Sprintf("%v of %v", rankNames[int(c.Rank())], suitNames[int(c.Suit())])
}

func (c *Card) hash() string {
	return string(c.SuitName()[0]) + string(c.RankName()[0])
}

func setupSuitNames() {
	suitNames = make(map[int]string, NumberOfColors)
	suitNames[0] = "Hearts"
	suitNames[1] = "Spades"
	suitNames[2] = "Diamonds"
	suitNames[3] = "Clubs"
}

func setupRankNames() {
	rankNames = make(map[int]string, NumberOfRanks)
	rankNames[0] = "2"
	rankNames[1] = "3"
	rankNames[2] = "4"
	rankNames[3] = "5"
	rankNames[4] = "6"
	rankNames[5] = "7"
	rankNames[6] = "8"
	rankNames[7] = "9"
	rankNames[8] = "Ten"
	rankNames[9] = "Jack"
	rankNames[10] = "Queen"
	rankNames[11] = "King"
	rankNames[12] = "Ace"
}
