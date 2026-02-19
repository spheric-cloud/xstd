// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package hashmap

import (
	"iter"
	"slices"
)

type entry[K, V any] struct {
	key   K
	value V
}

type HashMap[H comparable, K, V any] struct {
	hash    func(K) H
	equal   func(K, K) bool
	len     int
	entries map[H][]*entry[K, V]
}

// New creates a new HashMap with the given hash and equality functions.
func New[H comparable, K, V any](hash func(K) H, equal func(k1, k2 K) bool) *HashMap[H, K, V] {
	return &HashMap[H, K, V]{
		hash:    hash,
		equal:   equal,
		entries: make(map[H][]*entry[K, V]),
	}
}

func (h *HashMap[H, K, V]) findEntryIndex(entries []*entry[K, V], key K) int {
	return slices.IndexFunc(entries, func(e *entry[K, V]) bool { return h.equal(e.key, key) })
}

func (h *HashMap[H, K, V]) findEntry(entries []*entry[K, V], key K) *entry[K, V] {
	idx := h.findEntryIndex(entries, key)
	if idx == -1 {
		return nil
	}
	return entries[idx]
}

// Set inserts or updates the value for the given key.
func (h *HashMap[H, K, V]) Set(key K, value V) {
	hash := h.hash(key)
	entries := h.entries[hash]
	e := h.findEntry(entries, key)
	if e == nil {
		h.entries[hash] = append(entries, &entry[K, V]{key, value})
		h.len++
	} else {
		e.value = value
	}
}

// Get returns the value for the given key and whether it was found.
func (h *HashMap[H, K, V]) Get(key K) (V, bool) {
	hash := h.hash(key)
	entries, ok := h.entries[hash]
	if !ok {
		var zero V
		return zero, false
	}

	e := h.findEntry(entries, key)
	if e == nil {
		var zero V
		return zero, false
	}
	return e.value, true
}

// Value returns the value for the given key, or the zero value if not found.
func (h *HashMap[H, K, V]) Value(key K) V {
	v, _ := h.Get(key)
	return v
}

// Delete removes the entry for the given key.
func (h *HashMap[H, K, V]) Delete(key K) {
	hash := h.hash(key)
	entries, ok := h.entries[hash]
	if !ok {
		return
	}

	idx := h.findEntryIndex(entries, key)
	if idx < 0 {
		return
	}

	h.entries[hash] = slices.Delete(entries, idx, idx+1)
	h.len--
}

// Len returns the number of entries in the map.
func (h *HashMap[H, K, V]) Len() int {
	return h.len
}

// Clear removes all entries from the map.
func (h *HashMap[H, K, V]) Clear() {
	clear(h.entries)
}

// All returns an iterator over all key-value pairs in the map.
func (h *HashMap[H, K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, es := range h.entries {
			for _, e := range es {
				if !yield(e.key, e.value) {
					return
				}
			}
		}
	}
}

// Keys returns an iterator over all keys in the map.
func (h *HashMap[H, K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for _, es := range h.entries {
			for _, e := range es {
				if !yield(e.key) {
					return
				}
			}
		}
	}
}

// Values returns an iterator over all values in the map.
func (h *HashMap[H, K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, es := range h.entries {
			for _, e := range es {
				if !yield(e.value) {
					return
				}
			}
		}
	}
}
