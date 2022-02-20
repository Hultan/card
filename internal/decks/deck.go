package decks

import (
	"errors"
	"math/rand"
	"time"
)

type Decks struct {
	main  *Pile
	piles []*Pile
	trash *Pile
}

func NewDecks(numDecks int) *Decks {
	decks := &Decks{}
	decks.main = decks.createPileInternal("Main", true)
	decks.trash = decks.createPileInternal("Trash", true)

	setupSuitNames()
	setupRankNames()

	for i := 0; i < numDecks; i++ {
		for j := 0; j < NumberOfCardsInDeck; j++ {
			c := &Card{
				deck: i,
				card: j,
			}
			decks.main.cards = append(decks.main.cards, c)
		}
	}

	return decks
}

func (d *Decks) SetRandomSeed() {
	d.SetSeed(time.Now().UnixNano())
}

func (d *Decks) SetSeed(seed int64) {
	rand.Seed(seed)
}

func (d *Decks) CreatePile(name string) *Pile {
	return d.createPileInternal(name, false)
}

func (d *Decks) Main() *Pile {
	return d.main
}

func (d *Decks) Trash() *Pile {
	return d.trash
}

func (d *Decks) Count() int {
	return len(d.piles)
}

func (d *Decks) Pile(i int) (*Pile, error) {
	if i < 0 || i >= len(d.piles) {
		return nil, errors.New("invalid pile index")
	}
	return d.piles[i], nil
}

func (d *Decks) PileByName(name string) (*Pile, error) {
	for i, pile := range d.piles {
		if pile.Name() == name {
			return d.piles[i], nil
		}
	}
	return nil, errors.New("invalid pile name")
}

func (d *Decks) Collect() {
	for i := 0; i < len(d.piles); i++ {
		d.piles[i].Collect()
	}
	d.trash.Collect()
}

func (d *Decks) createPileInternal(name string, isInternalPile bool) *Pile {
	// Create a new Pile
	p := &Pile{deck: d, name: name}
	if !isInternalPile {
		// Add the pile to the decks piles slice
		d.piles = append(d.piles, p)
	}
	return p
}
