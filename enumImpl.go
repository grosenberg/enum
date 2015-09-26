// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package enum

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sync"
)

var (
	mu    sync.Mutex // creation lock
	initd bool       // initialization flag
	enums []enumSet  // root collection of all defined enum sets
)

var (
	errAlreadyDefined = errors.New("Enum set already defined.")
)

// An enumSet is a defined set of comparable Enums
type enumSet struct {
	eId   string // unique identifier of this enumSet
	eSet  []E    // list of the Enums that makeup this set
	eType string // internal assigned type of this set (name of the enum strct)
}

// A single Enum instance
type Enum struct {
	Id   string // unique identifier of the set this Enum belongs to
	Ord  int    // ordinal value of this Enum
	Desc string // descriptive string for this Enum (name of the enum)
}

//////////////////////////////////////////////////////////////
// Helper methods

// Create an enum set hash value
func encode(defs ...E) string {
	hash := md5.New()
	for _, v := range defs {
		io.WriteString(hash, toEnum(v).Desc)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Get the actual embedded anonymous enum structure instance - never fails to amaze
func toEnum(e E) *Enum {
	return reflect.ValueOf(e).Elem().Field(0).Addr().Interface().(*Enum)
}
