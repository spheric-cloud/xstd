// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package hashmap

import (
	"hash/maphash"
	"iter"
	"slices"
)

type entry[K, V any] struct {
	key   K
	value V
}

type HashMap[K, V any] struct {
	hash    func(K) uint64
	equal   func(K, K) bool
	len     int
	entries map[uint64][]*entry[K, V]
}

func New[K, V any](hash func(K) uint64, equal func(k1, k2 K) bool) *HashMap[K, V] {
	return &HashMap[K, V]{
		hash:    hash,
		equal:   equal,
		entries: make(map[uint64][]*entry[K, V]),
	}
}

var seed = maphash.MakeSeed()

func NewComparable[K comparable, V any]() *HashMap[K, V] {
	return New[K, V](func(k K) uint64 { return maphash.Comparable(seed, k) }, func(k1 K, k2 K) bool { return k1 == k2 })
}

func (h *HashMap[K, V]) findEntryIndex(entries []*entry[K, V], key K) int {
	return slices.IndexFunc(entries, func(e *entry[K, V]) bool { return h.equal(e.key, key) })
}

func (h *HashMap[K, V]) findEntry(entries []*entry[K, V], key K) *entry[K, V] {
	idx := h.findEntryIndex(entries, key)
	if idx == -1 {
		return nil
	}
	return entries[idx]
}

func (h *HashMap[K, V]) Put(key K, value V) {
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

func (h *HashMap[K, V]) Get(key K) (V, bool) {
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

func (h *HashMap[K, V]) Delete(key K) bool {
	hash := h.hash(key)
	entries, ok := h.entries[hash]
	if !ok {
		return false
	}

	idx := h.findEntryIndex(entries, key)
	if idx < 0 {
		return false
	}

	h.entries[hash] = slices.Delete(entries, idx, idx+1)
	h.len--
	return true
}

func (h *HashMap[K, V]) Len() int {
	return h.len
}

func (h *HashMap[K, V]) Clear() {
	clear(h.entries)
}

func (h *HashMap[K, V]) All() iter.Seq2[K, V] {
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

func (h *HashMap[K, V]) Keys() iter.Seq[K] {
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

func (h *HashMap[K, V]) Values() iter.Seq[V] {
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
