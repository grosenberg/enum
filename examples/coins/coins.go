// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package enums

import (
	"fmt"

	"github.com/grosenberg/enum"
)

// The Coins rich enum set
var (
	PENNY       = &Coin{enum.Enum{Desc: "penny"}, 1, true}
	NICKEL      = &Coin{enum.Enum{Desc: "nickel"}, 5, true}
	DIME        = &Coin{enum.Enum{Desc: "dime"}, 10, true}
	QUARTER     = &Coin{enum.Enum{Desc: "quarter"}, 25, true}
	HALF_DOLLAR = &Coin{enum.Enum{Desc: "half-dollar"}, 50, false}
	DOLLAR      = &Coin{enum.Enum{Desc: "dollar"}, 100, false}
)

type Coin struct {
	enum.Enum         // must be first field
	value     float64 // custom member fields
	common    bool
}

// Compile-time implementation prover
var _ enum.E = &Coin{}

// auto-initialization
func init() {
	err := enum.Define(PENNY, NICKEL, DIME, QUARTER, HALF_DOLLAR, DOLLAR)
	if err != nil {
		fmt.Println(err)
	}
}

//////////////////////////////////////////////////////////////
// Rich enum functions

func (c *Coin) Value() float64 {
	return c.value
}

func (c *Coin) CommonCurrency() bool {
	return c.common
}

//////////////////////////////////////////////////////////////
// Helpers for generic methods

func (c *Coin) CoinValues() []Coin {
	vs := c.Values()
	coins := make([]Coin, len(vs))
	for i, v := range vs {
		coins[i] = *(v.(*Coin))
	}
	return coins
}
