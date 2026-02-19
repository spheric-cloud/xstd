// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package hashset

import (
	"iter"

	"spheric.cloud/xstd/container/list"
)

// HashSet is a generic set data structure.
type HashSet[H comparable, V any] struct {
	hash    func(V) H
	equal   func(V, V) bool
	entries map[H]*list.List[V]
	len     int
}

// New creates a new set from the given values.
func New[H comparable, V any](hash func(V) H, equal func(V, V) bool) *HashSet[H, V] {
	s := &HashSet[H, V]{
		hash:    hash,
		equal:   equal,
		entries: make(map[H]*list.List[V]),
	}
	return s
}

func (s *HashSet[H, V]) findEntry(entries *list.List[V], value V) *list.Element[V] {
	for e := range entries.Elems() {
		if s.equal(value, e.Value) {
			return e
		}
	}
	return nil
}

// Insert adds the given values to the set.
func (s *HashSet[H, V]) Insert(vs ...V) *HashSet[H, V] {
	for _, v := range vs {
		hash := s.hash(v)
		entries := s.entries[hash]
		if entries == nil {
			entries = list.New[V]()
			s.entries[hash] = entries
			entries.PushBack(v)
			s.len++
		} else if entry := s.findEntry(entries, v); entry == nil {
			entries.PushBack(v)
			s.len++
		}
	}
	return s
}

// Delete removes the given values from the set.
func (s *HashSet[H, V]) Delete(vs ...V) *HashSet[H, V] {
	for _, v := range vs {
		hash := s.hash(v)
		entries, ok := s.entries[hash]
		if !ok {
			continue
		}

		entry := s.findEntry(entries, v)
		if entry == nil {
			continue
		}

		s.len--
		entries.Remove(entry)
		if entries.Len() == 0 {
			delete(s.entries, hash)
		}
	}
	return s
}

// Has returns true if the set contains the given value.
func (s *HashSet[H, V]) Has(v V) bool {
	hash := s.hash(v)
	entries, ok := s.entries[hash]
	if !ok {
		return false
	}

	entry := s.findEntry(entries, v)
	return entry != nil
}

// Values returns an iterator over all values in the set.
func (s *HashSet[H, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, entry := range s.entries {
			for v := range entry.Values() {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Len returns the length of the set.
func (s *HashSet[H, V]) Len() int {
	return s.len
}
