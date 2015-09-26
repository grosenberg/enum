// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package example

import (
	"testing"

	"github.com/grosenberg/enum/examples/cards"
	"github.com/grosenberg/enum/examples/coins"
)

func TestName(t *testing.T) {
	n := cards.SPADE.Name()
	v := "spade"
	if n != v {
		t.Errorf("Expected %v but got %v.\n", v, n)
	}
}

func TestOrdinal(t *testing.T) {
	n := cards.DIAMOND.Ordinal()
	v := 2
	if n != v {
		t.Errorf("Expected %v but got %v.\n", v, n)
	}
}

func TestType(t *testing.T) {
	n := cards.DIAMOND.Type()
	v := "Card"
	if n != v {
		t.Errorf("Expected %v but got %v.\n", v, n)
	}
}

func TestCmp(t *testing.T) {
	n := cards.DIAMOND
	v := cards.CLUB
	if !n.Comparable(v) {
		t.Errorf("Should be comparable: %v and %v.\n", v, n)
	}
}

func TestCmp1(t *testing.T) {
	n := cards.SPADE
	v := coins.PENNY
	if n.Comparable(v) {
		t.Errorf("Should not be comparable: %v and %v.\n", v, n)
	}
}

func TestEquals(t *testing.T) {
	n := cards.DIAMOND
	v := cards.CLUB
	if n.Equals(v) {
		t.Errorf("Should not be equal: %v and %v.\n", v, n)
	}
}

func TestEquals(t *testing.T) {
	n := cards.SPADE // ordinal 1 of cards
	v := cards.PENNY // ordinal 1 of coins
	if n.Equals(v) {
		t.Errorf("Should not be equal: %v and %v.\n", v, n)
	}
}

func TestEquals1(t *testing.T) {
	n := coins.DIME
	v := coins.DIME
	if !n.Equals(v) {
		t.Errorf("Should be equal: %v and %v.\n", v, n)
	}
}
