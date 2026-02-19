// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package keyset

import (
	"iter"
	"maps"
)

// KeySet is a generic set data structure.
type KeySet[IK comparable, V any] struct {
	internalKey func(V) IK
	values      map[IK]V
}

// New creates a new set from the given values.
func New[IK comparable, V any](internalKey func(V) IK, vs ...V) *KeySet[IK, V] {
	s := &KeySet[IK, V]{
		internalKey: internalKey,
		values:      make(map[IK]V),
	}
	s.Insert(vs...)
	return s
}

// Insert adds the given values to the set.
func (s *KeySet[IK, V]) Insert(vs ...V) *KeySet[IK, V] {
	for _, v := range vs {
		s.values[s.internalKey(v)] = v
	}
	return s
}

// Delete removes the given values from the set.
func (s *KeySet[IK, V]) Delete(vs ...V) *KeySet[IK, V] {
	for _, v := range vs {
		delete(s.values, s.internalKey(v))
	}
	return s
}

// Values returns an iterator over all values in the set.
func (s *KeySet[IK, V]) Values() iter.Seq[V] {
	return maps.Values(s.values)
}

// Has returns true if the set contains the given value.
func (s *KeySet[IK, V]) Has(v V) bool {
	_, ok := s.values[s.internalKey(v)]
	return ok
}

// Len returns the length of the set.
func (s *KeySet[IK, V]) Len() int {
	return len(s.values)
}
