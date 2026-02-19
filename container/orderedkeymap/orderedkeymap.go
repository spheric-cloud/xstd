// SPDX-FileCopyrightText: 2026 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package orderedkeymap

import (
	"iter"

	"spheric.cloud/xstd/container/list"
)

type entry[IK comparable, K, V any] struct {
	node  *list.Element[IK]
	key   K
	value V
}

type OrderedKeyMap[IK comparable, K, V any] struct {
	internalKey func(K) IK
	entries     map[IK]*entry[IK, K, V]
	order       *list.List[IK]
}

// New creates a new OrderedKeyMap with the given internal key function.
func New[IK comparable, K, V any](internalKey func(K) IK) *OrderedKeyMap[IK, K, V] {
	return &OrderedKeyMap[IK, K, V]{
		internalKey: internalKey,
		entries:     make(map[IK]*entry[IK, K, V]),
		order:       list.New[IK](),
	}
}

// Get returns the value for the given key and whether it was found.
func (m *OrderedKeyMap[IK, K, V]) Get(key K) (V, bool) {
	e, ok := m.entries[m.internalKey(key)]
	if !ok {
		var zero V
		return zero, false
	}
	return e.value, true
}

// Value returns the value for the given key, or the zero value if not found.
func (m *OrderedKeyMap[IK, K, V]) Value(key K) V {
	v, _ := m.Get(key)
	return v
}

// Set inserts or updates the value for the given key, preserving insertion order.
func (m *OrderedKeyMap[IK, K, V]) Set(key K, value V) {
	internalKey := m.internalKey(key)
	if e, ok := m.entries[internalKey]; ok {
		e.value = value
		return
	}

	e := &entry[IK, K, V]{
		key:   key,
		value: value,
	}
	e.node = m.order.PushBack(internalKey)
	m.entries[internalKey] = e
}

// Delete removes the entry for the given key.
func (m *OrderedKeyMap[IK, K, V]) Delete(key K) {
	internalKey := m.internalKey(key)
	e, ok := m.entries[internalKey]
	if !ok {
		return
	}

	m.order.Remove(e.node)
	delete(m.entries, internalKey)
}

// Len returns the number of entries in the map.
func (m *OrderedKeyMap[IK, K, V]) Len() int {
	return len(m.entries)
}

// Clear removes all entries from the map.
func (m *OrderedKeyMap[IK, K, V]) Clear() {
	m.order = list.New[IK]()
	clear(m.entries)
}

// All returns an iterator over all key-value pairs in insertion order.
func (m *OrderedKeyMap[IK, K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for ik := range m.order.Values() {
			e := m.entries[ik]
			if !yield(e.key, e.value) {
				return
			}
		}
	}
}

// Keys returns an iterator over all keys in insertion order.
func (m *OrderedKeyMap[IK, K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for ik := range m.order.Values() {
			e := m.entries[ik]
			if !yield(e.key) {
				return
			}
		}
	}
}

// Values returns an iterator over all values in insertion order.
func (m *OrderedKeyMap[IK, K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for ik := range m.order.Values() {
			v := m.entries[ik].value
			if !yield(v) {
				return
			}
		}
	}
}
