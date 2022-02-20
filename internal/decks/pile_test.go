package decks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPile_Name(t *testing.T) {
	d := NewDecks(1)
	h1 := d.CreatePile("Hand1")
	assert.Equalf(t, "Hand1", h1.Name(), "invalid pile name : hand1")
}

func TestPile_DealToPile(t *testing.T) {
	d := NewDecks(1)
	h1 := d.CreatePile("Hand1")
	d.Main().DealToPile(5, h1)
	assert.Equalf(t, 5, h1.Count(), "invalid card count : hand1")
	assert.Equalf(t, 47, d.Main().Count(), "invalid card count : main")
}

func TestPile_Collect(t *testing.T) {
	d := NewDecks(1)
	h1 := d.CreatePile("Hand1")
	d.Main().DealToPile(5, h1)
	h2 := d.CreatePile("Hand2")
	assert.NotEqual(t, h1, h2, "h1 and h2 should be different piles")
	d.Main().DealToPile(5, h2)
	h1.Collect()
	assert.Equalf(t, 0, h1.Count(), "invalid card count : hand1")
	assert.Equalf(t, 5, h2.Count(), "invalid card count : hand2")
	assert.Equalf(t, 47, d.Main().Count(), "invalid card count : main")
	h2.Collect()
	assert.Equalf(t, 0, h1.Count(), "invalid card count : hand1")
	assert.Equalf(t, 0, h2.Count(), "invalid card count : hand2")
	assert.Equalf(t, 52, d.Main().Count(), "invalid card count : main")
}

func TestPile_Trash(t *testing.T) {
	d := NewDecks(1)
	h1 := d.CreatePile("Hand1")
	d.Main().DealToPile(5, h1)
	h2 := d.CreatePile("Hand2")
	assert.NotEqual(t, h1, h2, "h1 and h2 should be different piles")
	d.Main().DealToPile(5, h2)
	h1.Trash()
	assert.Equalf(t, 0, h1.Count(), "invalid card count : hand1")
	assert.Equalf(t, 5, h2.Count(), "invalid card count : hand2")
	assert.Equalf(t, 42, d.Main().Count(), "invalid card count : main")
	assert.Equalf(t, 5, d.Trash().Count(), "invalid card count : trash")
	h2.Trash()
	assert.Equalf(t, 0, h1.Count(), "invalid card count : hand1")
	assert.Equalf(t, 0, h2.Count(), "invalid card count : hand2")
	assert.Equalf(t, 42, d.Main().Count(), "invalid card count : main")
	assert.Equalf(t, 10, d.Trash().Count(), "invalid card count : trash")
}

func TestPile_Shuffle(t *testing.T) {
	d := NewDecks(1)
	h1 := d.Main().hash()
	d.Main().Shuffle()
	h2 := d.Main().hash()
	assert.NotEqual(t, h2, h1, "should be different (in most cases)")
}

func TestPile_Card(t *testing.T) {
	d := NewDecks(1)
	h1 := d.CreatePile("Hand1")
	d.Main().DealToPile(5, h1)
	c := h1.Card(0)
	fmt.Println(c.Suit())
	fmt.Println(c.Rank())
	assert.NotNil(t, c, "card should not be nil")
	assert.Equal(t, SuitHeart, c.Suit(), "card color should be heart")
	assert.Equal(t, RankTwo, c.Rank(), "card rank should be two")
}
