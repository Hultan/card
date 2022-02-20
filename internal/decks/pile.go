package decks

import (
	"errors"
	"fmt"
	"math/rand"
)

type Pile struct {
	deck  *Decks
	name  string
	cards []*Card
}

func (p *Pile) Card(i int) *Card {
	return p.cards[i]
}

func (p *Pile) Name() string {
	return p.name
}

func (p *Pile) Count() int {
	return len(p.cards)
}

func (p *Pile) Collect() {
	p.moveToPile(p.deck.main)
}

func (p *Pile) Trash() {
	p.moveToPile(p.deck.trash)
}

func (p *Pile) DealToPile(cards int, toPile *Pile) error {
	if p.deck != toPile.deck {
		return errors.New("cannot deal to a pile of another deck")
	}
	if len(p.cards) < cards {
		return errors.New("not enough cards in pile")
	}

	toPile.cards = append(toPile.cards, p.cards[:cards]...)
	p.cards = p.cards[cards:]
	return nil
}

func (p *Pile) Shuffle() {
	rand.Shuffle(len(p.cards), func(i, j int) { p.cards[i], p.cards[j] = p.cards[j], p.cards[i] })
}

func (p *Pile) moveToPile(toPile *Pile) {
	toPile.cards = append(toPile.cards, p.cards...)
	p.cards = nil
}

func (p *Pile) hash() string {
	h := ""
	for _, c := range p.cards {
		h += c.hash()
	}
	return h
}

func (p *Pile) printPile() {
	for _, c := range p.cards {
		fmt.Println(c)
	}
}
