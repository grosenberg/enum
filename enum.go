// Copyright © 2015 Gerald Rosenberg.
// Use of this source code is governed by a BSD-style
// license that can be found in the License.md file.
//
package enum

//////////////////////////////////////////////////////////////
// Standard Enum functions

// Get the ordinal value of this enum instance
func (e *Enum) Ordinal() int {
	return e.Ord
}

// Get the descriptive name of this enum instance
func (e *Enum) Name() string {
	return e.Desc
}

// Get the set of enums that this enum belongs to
func (e *Enum) Values() []E {
	for _, v := range enums {
		if v.eId == e.Id {
			return v.eSet
		}
	}
	return nil
}

// Get the type of the set of enums that this enum belongs to.
// By design, this is the name of the implementing enum structure.
func (e *Enum) Type() string {
	for _, v := range enums {
		if v.eId == e.Id {
			return v.eType
		}
	}
	return "<Unknown enum type>" // should never occur
}

// Determines if these two enums are comparable (belong to the same enum set)
func (e *Enum) Comparable(v E) bool {
	if e.Id == toEnum(v).Id {
		return true
	}
	return false
}

// Determines if these two enums are equal.
func (e *Enum) Equals(v E) bool {
	if e.Comparable(v) && e.Ord == toEnum(v).Ord {
		return true
	}
	return false
}

// Find the enum in the receivers' enum set with the given description
func (e *Enum) ValueOf(desc string) E {
	for _, v := range e.Values() {
		if v.Name() == desc {
			return v
		}
	}
	return nil
}
