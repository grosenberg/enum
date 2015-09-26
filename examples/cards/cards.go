// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package enums

import (
	"fmt"

	"github.com/grosenberg/enum"
)

// The Cards rich enum set
var (
	SPADE   = &Card{enum.Enum{Desc: "spade"}, 1, false}
	HEART   = &Card{enum.Enum{Desc: "heart"}, 2, false}
	DIAMOND = &Card{enum.Enum{Desc: "diamond"}, 3, true}
	CLUB    = &Card{enum.Enum{Desc: "club"}, 4, true}
)

type Card struct {
	enum.Enum     // must be first field
	weight    int // custom member fields
	joker     bool
}

// Compile-time implementation prover
var _ enum.E = &Card{}

// auto-initialization
func init() {
	err := enum.Define(SPADE, HEART, DIAMOND, CLUB)
	if err != nil {
		fmt.Println(err)
	}
}

//////////////////////////////////////////////////////////////
// Rich enum functions

func (c *Card) Suit() string {
	return c.Name()
}

func (c *Card) Weight() int {
	return c.weight
}

func (c *Card) JokerWild() bool {
	return c.joker
}

//////////////////////////////////////////////////////////////
// Helpers for generic methods

func (c *Card) CardValues() []Card {
	vs := c.Values()
	cards := make([]Card, len(vs))
	for i, v := range vs {
		cards[i] = *(v.(*Card))
	}
	return cards
}
