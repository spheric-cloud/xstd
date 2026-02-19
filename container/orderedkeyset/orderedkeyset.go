// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package orderedkeyset

import (
	"iter"

	"spheric.cloud/xstd/container/list"
)

type entry[IK comparable, V any] struct {
	node  *list.Element[IK]
	value V
}

// KeySet is a generic set data structure.
type KeySet[IK comparable, V any] struct {
	internalKey func(V) IK
	entries     map[IK]*entry[IK, V]
	order       *list.List[IK]
}

// New creates a new set from the given values.
func New[IK comparable, V any](internalKey func(V) IK, vs ...V) *KeySet[IK, V] {
	s := &KeySet[IK, V]{
		internalKey: internalKey,
		entries:     make(map[IK]*entry[IK, V]),
		order:       list.New[IK](),
	}
	s.Insert(vs...)
	return s
}

func (s *KeySet[IK, V]) insert(v V) {
	internalKey := s.internalKey(v)
	if e, ok := s.entries[internalKey]; ok {
		e.value = v
		return
	}

	e := &entry[IK, V]{
		value: v,
	}
	e.node = s.order.PushBack(internalKey)
	s.entries[internalKey] = e
}

// Insert adds the given values to the set.
func (s *KeySet[IK, V]) Insert(vs ...V) *KeySet[IK, V] {
	for _, v := range vs {
		s.insert(v)
	}
	return s
}

func (s *KeySet[IK, V]) delete(v V) {
	internalKey := s.internalKey(v)
	e, ok := s.entries[internalKey]
	if !ok {
		return
	}

	s.order.Remove(e.node)
	delete(s.entries, internalKey)
}

// Delete removes the given values from the set.
func (s *KeySet[IK, V]) Delete(vs ...V) *KeySet[IK, V] {
	for _, v := range vs {
		s.delete(v)
	}
	return s
}

// Values returns an iterator over all values in the set in insertion order.
func (s *KeySet[IK, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for ik := range s.order.Values() {
			e := s.entries[ik]
			if !yield(e.value) {
				return
			}
		}
	}
}

// Has returns true if the set contains the given value.
func (s *KeySet[IK, V]) Has(v V) bool {
	internalKey := s.internalKey(v)
	_, ok := s.entries[internalKey]
	return ok
}

// Len returns the length of the set.
func (s *KeySet[IK, V]) Len() int {
	return len(s.entries)
}
