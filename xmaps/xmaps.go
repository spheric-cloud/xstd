// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package xmaps

import (
	"maps"
	"slices"

	"spheric.cloud/xstd/container/set"
)

// Inverse returns a new map with the keys and values swapped.
// If multiple keys map to the same value, only one will be present in the result.
func Inverse[Map ~map[K]V, K, V comparable](m Map) map[V]K {
	res := make(map[V]K, len(m))
	for k, v := range m {
		res[v] = k
	}
	return res
}

// KeysByValue returns a new map that groups keys by their values.
func KeysByValue[Map ~map[K]V, K, V comparable](m Map) map[V][]K {
	res := make(map[V][]K)
	for k, v := range m {
		res[v] = append(res[v], k)
	}
	return res
}

// GetAny returns an arbitrary key-value pair from the map and whether the map is non-empty.
func GetAny[Map ~map[K]V, K comparable, V any](m Map) (K, V, bool) {
	for k, v := range m {
		return k, v, true
	}
	var (
		zeroK K
		zeroV V
	)
	return zeroK, zeroV, false
}

// GetAnyValue returns an arbitrary key-value pair from the map, ignoring the boolean.
func GetAnyValue[Map ~map[K]V, K comparable, V any](m Map) (K, V) {
	k, v, _ := GetAny(m)
	return k, v
}

// Pop removes and returns an arbitrary key-value pair from the map.
// It also returns a boolean indicating whether a pair was popped.
// If the map is empty, it returns the zero values for the key and value and false.
func Pop[Map ~map[K]V, K comparable, V any](m Map) (K, V, bool) {
	k, v, ok := GetAny(m)
	if ok {
		delete(m, k)
	}
	return k, v, ok
}

// PopValue removes and returns an arbitrary key-value pair from the map.
// It's a convenience wrapper around Pop that ignores the boolean return value.
func PopValue[Map ~map[K]V, K comparable, V any](m Map) (K, V) {
	k, v, _ := Pop(m)
	return k, v
}

// KeySlice returns a slice of all keys in the map.
func KeySlice[Map ~map[K]V, K comparable, V any](m Map) []K {
	res := make([]K, 0, len(m))
	return slices.AppendSeq(res, maps.Keys(m))
}

// ValueSlice returns a slice of all values in the map.
func ValueSlice[Map ~map[K]V, K comparable, V any](m Map) []V {
	res := make([]V, 0, len(m))
	return slices.AppendSeq(res, maps.Values(m))
}

// KeySet returns a set of all keys in the map.
func KeySet[Map ~map[K]V, K comparable, V any](m Map) set.Set[K] {
	res := make(set.Set[K], len(m))
	for k := range m {
		res.Insert(k)
	}
	return res
}

// ValueSet returns a set of all values in the map.
func ValueSet[Map ~map[K]V, K, V comparable](m Map) set.Set[V] {
	res := make(set.Set[V], len(m))
	for _, v := range m {
		res.Insert(v)
	}
	return res
}
