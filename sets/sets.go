// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

// Package sets provides utility functions for working with sets.
package sets

import (
	"iter"
	"maps"

	"spheric.cloud/xstd/set"
)

// HasAll returns true if the set contains all values in the given sequence.
func HasAll[V comparable](s set.Set[V], seq iter.Seq[V]) bool {
	for v := range seq {
		if !s.Has(v) {
			return false
		}
	}
	return true
}

// HasAny returns true if the set contains any value in the given sequence.
func HasAny[V comparable](s set.Set[V], seq iter.Seq[V]) bool {
	for v := range seq {
		if s.Has(v) {
			return true
		}
	}
	return false
}

// Clone returns a copy of the set.
func Clone[V comparable](s set.Set[V]) set.Set[V] {
	return maps.Clone(s)
}

// Difference returns a new set with all values from s1 that are not in s2.
func Difference[V comparable](s1, s2 set.Set[V]) set.Set[V] {
	result := set.New[V]()
	for key := range s1 {
		if !s2.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// Union returns a new set with all values from s1 and s2.
func Union[V comparable](s1, s2 set.Set[V]) set.Set[V] {
	res := Clone(s1)
	for key := range s2 {
		res.Insert(key)
	}
	return res
}

// Equal returns true if the two sets are equal.
func Equal[V comparable](s1, s2 set.Set[V]) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

// Values returns a sequence of the values in the set.
func Values[V comparable](s set.Set[V]) iter.Seq[V] {
	return maps.Keys(s)
}

// Pop removes and returns an arbitrary value from the set.
// The second return value is false if the set is empty.
func Pop[V comparable](s set.Set[V]) (V, bool) {
	for v := range s {
		delete(s, v)
		return v, true
	}
	var zero V
	return zero, false
}

// Insert inserts all values from the sequence into the set.
func Insert[V comparable](s set.Set[V], seq iter.Seq[V]) {
	for v := range seq {
		s.Insert(v)
	}
}

// TryInsert inserts all values from the sequence into the set.
// It returns the first error encountered.
func TryInsert[V comparable](s set.Set[V], seq iter.Seq2[V, error]) error {
	for v, err := range seq {
		if err != nil {
			return err
		}
		s.Insert(v)
	}
	return nil
}

// Collect collects all values from the sequence into a new set.
func Collect[V comparable](seq iter.Seq[V]) set.Set[V] {
	s := set.New[V]()
	Insert(s, seq)
	return s
}

// TryCollect collects all values from the sequence into a new set.
// It returns the first error encountered.
func TryCollect[V comparable](seq iter.Seq2[V, error]) (set.Set[V], error) {
	s := set.New[V]()
	if err := TryInsert(s, seq); err != nil {
		return nil, err
	}
	return s, nil
}
