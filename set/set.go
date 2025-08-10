// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package set

// Set is a generic set data structure.
type Set[V comparable] map[V]struct{}

// New creates a new set from the given values.
func New[V comparable](vs ...V) Set[V] {
	s := make(Set[V])
	s.Insert(vs...)
	return s
}

// Insert adds the given values to the set.
func (s Set[V]) Insert(vs ...V) Set[V] {
	for _, v := range vs {
		s[v] = struct{}{}
	}
	return s
}

// Delete removes the given values from the set.
func (s Set[V]) Delete(vs ...V) Set[V] {
	for _, v := range vs {
		delete(s, v)
	}
	return s
}

// Has returns true if the set contains the given value.
func (s Set[V]) Has(v V) bool {
	_, ok := s[v]
	return ok
}

// Len returns the length of the set.
func (s Set[V]) Len() int {
	return len(s)
}
