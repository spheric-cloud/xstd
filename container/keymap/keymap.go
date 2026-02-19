// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package hashmap

import (
	"iter"
)

type entry[K, V any] struct {
	key   K
	value V
}

type KeyMap[IK comparable, K, V any] struct {
	internalKey func(K) IK
	entries     map[IK]*entry[K, V]
}

func New[IK comparable, K, V any](internalKey func(K) IK) *KeyMap[IK, K, V] {
	return &KeyMap[IK, K, V]{
		internalKey: internalKey,
		entries:     make(map[IK]*entry[K, V]),
	}
}

func (h *KeyMap[IK, K, V]) Set(key K, value V) {
	internalKey := h.internalKey(key)
	e := h.entries[internalKey]
	if e == nil {
		h.entries[internalKey] = &entry[K, V]{key, value}
	} else {
		e.value = value
	}
}

func (h *KeyMap[IK, K, V]) Get(key K) (V, bool) {
	e, ok := h.entries[h.internalKey(key)]
	if !ok {
		var zero V
		return zero, false
	}
	return e.value, true
}

func (h *KeyMap[IK, K, V]) Value(key K) V {
	v, _ := h.Get(key)
	return v
}

func (h *KeyMap[IK, K, V]) Delete(key K) {
	delete(h.entries, h.internalKey(key))
}

func (h *KeyMap[IK, K, V]) Len() int {
	return len(h.entries)
}

func (h *KeyMap[IK, K, V]) Clear() {
	clear(h.entries)
}

func (h *KeyMap[IK, K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, e := range h.entries {
			if !yield(e.key, e.value) {
				return
			}
		}
	}
}

func (h *KeyMap[IK, K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for _, e := range h.entries {
			if !yield(e.key) {
				return
			}
		}
	}
}

func (h *KeyMap[IK, K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, e := range h.entries {
			if !yield(e.value) {
				return
			}
		}
	}
}
