// Copyright 2026 Axel Christ and Spheric contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package xcmp

import (
	"iter"
	"slices"
)

// ZeroPredicate returns a function that checks if the given value is zero.
// It only allocates the zero value once.
func ZeroPredicate[V comparable]() func(v V) bool {
	var zero V
	return func(v V) bool {
		return v == zero
	}
}

// IsZero checks if a value is the zero value for its type.
func IsZero[V comparable](v V) bool {
	var zero V
	return v == zero
}

// AnyZero checks if any of the given values is the zero value for its type.
func AnyZero[V comparable](vs ...V) bool {
	return AnyZeroSeq(slices.Values(vs))
}

// AnyZeroSeq checks if any of the given values in the sequence is the zero value for its type.
func AnyZeroSeq[V comparable](vs iter.Seq[V]) bool {
	var zero V
	for v := range vs {
		if v == zero {
			return true
		}
	}
	return false
}

// AllZero checks if all the given values are the zero value for their type.
func AllZero[V comparable](vs ...V) bool {
	return AllZeroSeq(slices.Values(vs))
}

// AllZeroSeq checks if all the given values are the zero value for their type.
func AllZeroSeq[V comparable](vs iter.Seq[V]) bool {
	var zero V
	for v := range vs {
		if v != zero {
			return false
		}
	}
	return true
}

// OrSeq returns the first non-zero value of the given seq.
// It returns zero if the sequence is empty or if all values are zero.
func OrSeq[V comparable](vs iter.Seq[V]) V {
	var zero V
	for v := range vs {
		if v != zero {
			return v
		}
	}
	return zero
}

// OrFunc is a function-driven alternative of cmp.Or.
func OrFunc[V comparable](fs ...func() V) V {
	return OrFuncSeq(slices.Values(fs))
}

// OrFuncSeq returns the first non-zero value from evaluating the functions in the given sequence.
// It returns zero if the sequence is empty or if all values are zero.
func OrFuncSeq[V comparable](fs iter.Seq[func() V]) V {
	var zero V
	for f := range fs {
		if v := f(); v != zero {
			return v
		}
	}
	return zero
}
