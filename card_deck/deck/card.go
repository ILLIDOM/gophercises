//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

type Suit uint8 // e.g. heart, diamond, club, spade

const (
	Spade   Suit = iota //first value is 0, then increment by 1
	Diamond             // value is 1
	Club                // value is 2
	Heart
	Joker // value is 4
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_   Rank = iota // skip first value of iota (0)
	Ace             // value 1
	Two             // value 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King // value 13
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New() []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	return cards
}
