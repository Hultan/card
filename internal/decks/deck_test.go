package decks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecks_NewDecks(t *testing.T) {
	d := NewDecks(1)
	assert.NotNil(t, d, "deck is nil")
	assert.Equalf(t, 52, d.Main().Count(), "new deck main pile is not full")
	assert.Equalf(t, 0, d.Trash().Count(), "new deck trash pile is not empty")
	assert.Equalf(t, 0, d.Count(), "new deck piles slice is not empty")
}

func TestDecks_CreatePile(t *testing.T) {
	d := NewDecks(1)
	p := d.CreatePile("Hand1")
	assert.NotNil(t, p, "pile is nil")
	assert.Equalf(t, 1, d.Count(), "new pile is not added to deck slice of piles")

	pByIndex, err := d.Pile(0)
	assert.Nil(t, err, err)
	assert.Equalf(t, p, pByIndex, "pile by index failed")

	pByName, err := d.PileByName("Hand1")
	assert.Nil(t, err, err)
	assert.Equalf(t, p, pByName, "pile by name failed")

	assert.Equalf(t, "Hand1", p.Name(), "new pile has the wrong name")
	assert.Equalf(t, 0, p.Count(), "new pile is not empty")
}

func TestDecks_Pile(t *testing.T) {
	d := NewDecks(1)
	p := d.CreatePile("Hand1")
	pByIndex, err := d.Pile(0)
	assert.Nil(t, err, err)
	assert.Equalf(t, p, pByIndex, "pile by index failed")
	assert.Equalf(t, "Hand1", p.Name(), "new pile has the wrong name")
	assert.Equalf(t, 0, p.Count(), "new pile is not empty")
}

func TestDecks_PileByName(t *testing.T) {
	d := NewDecks(1)
	p := d.CreatePile("Hand1")
	pByName, err := d.PileByName("Hand1")
	assert.Nil(t, err, err)
	assert.Equalf(t, p, pByName, "pile by name failed")
	assert.Equalf(t, "Hand1", p.Name(), "new pile has the wrong name")
	assert.Equalf(t, 0, p.Count(), "new pile is not empty")
}

func TestDecks_Collect(t *testing.T) {
	d := NewDecks(1)
	p := d.CreatePile("Hand1")
	err := d.Main().DealToPile(5, p)
	assert.Nil(t, err, "failed to deal cards")
	assert.Equalf(t, 5, p.Count(), "invalid card count : hand1")
	assert.Equalf(t, 47, d.Main().Count(), "invalid card count : main")
	d.Collect()
	assert.Equalf(t, 0, p.Count(), "invalid card count : hand1")
	assert.Equalf(t, 52, d.Main().Count(), "invalid card count : main")
}

func TestDecks_Seed(t *testing.T) {
	d := NewDecks(1)
	d.SetSeed(1)
	d.Main().Shuffle()
	c := d.Main().Card(0)

	d2 := NewDecks(1)
	d2.SetRandomSeed()
	d2.Main().Shuffle()
	c2 := d2.Main().Card(0)

	assert.NotEqual(t, c, c2, "cards should not be equal (in most cases)")
}
