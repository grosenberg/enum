// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package enum

import "reflect"

// Generic Enum instance - standard set of enum functons
type E interface {
	Values() []E
	Name() string
	Ordinal() int
	Type() string
	Comparable(E) bool
	Equals(E) bool
	ValueOf(string) E
}

// Define a new set of Enums
func Define(defs ...E) error {
	mu.Lock()
	defer mu.Unlock()

	// initialize root
	if !initd {
		initd = true
		enums = make([]enumSet, 0)
	}

	// check for uniqueness
	newId := encode(defs...)
	for _, v := range enums {
		if v.eId == newId {
			return errAlreadyDefined
		}
	}

	// create a set container for these new enums
	set := enumSet{
		eId:   newId,
		eSet:  make([]E, 0),
		eType: reflect.ValueOf(defs[0]).Elem().Type().Name(),
	}
	enums = append(enums, set)

	// update the enum instances and add enums to their set
	for i, v := range defs {
		e := toEnum(v)
		e.Id = newId // unique hash
		e.Ord = i    // ordinal value
		set.eSet = append(set.eSet, v)
	}
	return nil
}
