package decks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const hash = "H2H3H4H5H6H7H8H9HTHJHQHKHAS2S3S4S5S6S7S8S9STSJSQSKSAD2D3D4D5D6D7D8D9DTDJDQDKDAC2C3C4C5C6C7C8C9CTCJCQCKCA"

func TestCard_AllCards(t *testing.T) {
	d := NewDecks(1)
	h := d.Main().hash()
	assert.Equal(t, hash, h, "hash not correct")
}

func TestCard_Rank(t *testing.T) {
	d := NewDecks(1)
	c := d.Main().Card(13)
	assert.Equal(t, RankTwo, c.Rank(), "wrong rank")
	assert.Equal(t, "2", c.RankName(), "wrong rank name")
}

func TestCard_Suit(t *testing.T) {
	d := NewDecks(1)
	c := d.Main().Card(14)
	assert.Equal(t, SuitSpade, c.Suit(), "wrong suit")
	assert.Equal(t, "Spades", c.SuitName(), "wrong suit name")
}
