//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepequal-gen. DO NOT EDIT.

package labels

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ByKey) DeepEqual(other *ByKey) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for i, inElement := range *in {
			if !inElement.DeepEqual(&(*other)[i]) {
				return false
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Lexer) DeepEqual(other *Lexer) bool {
	if other == nil {
		return false
	}

	if in.s != other.s {
		return false
	}
	if in.pos != other.pos {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Parser) DeepEqual(other *Parser) bool {
	if other == nil {
		return false
	}

	if (in.l == nil) != (other.l == nil) {
		return false
	} else if in.l != nil {
		if !in.l.DeepEqual(other.l) {
			return false
		}
	}

	if ((in.scannedItems != nil) && (other.scannedItems != nil)) || ((in.scannedItems == nil) != (other.scannedItems == nil)) {
		in, other := &in.scannedItems, &other.scannedItems
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if in.position != other.position {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Requirement) DeepEqual(other *Requirement) bool {
	if other == nil {
		return false
	}

	if in.key != other.key {
		return false
	}
	if in.operator != other.operator {
		return false
	}
	if ((in.strValues != nil) && (other.strValues != nil)) || ((in.strValues == nil) != (other.strValues == nil)) {
		in, other := &in.strValues, &other.strValues
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Requirements) DeepEqual(other *Requirements) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for i, inElement := range *in {
			if !inElement.DeepEqual(&(*other)[i]) {
				return false
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ScannedItem) DeepEqual(other *ScannedItem) bool {
	if other == nil {
		return false
	}

	if in.tok != other.tok {
		return false
	}
	if in.literal != other.literal {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Set) DeepEqual(other *Set) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for key, inValue := range *in {
			if otherValue, present := (*other)[key]; !present {
				return false
			} else {
				if inValue != otherValue {
					return false
				}
			}
		}
	}

	return true
}
