// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package maps

// Pop removes and returns an arbitrary key-value pair from the map.
// It also returns a boolean indicating whether a pair was popped.
// If the map is empty, it returns the zero values for the key and value and false.
func Pop[Map ~map[K]V, K comparable, V any](m Map) (K, V, bool) {
	for k, v := range m {
		delete(m, k)
		return k, v, true
	}
	var (
		zeroK K
		zeroV V
	)
	return zeroK, zeroV, false
}

// PopValue removes and returns an arbitrary key-value pair from the map.
// It's a convenience wrapper around Pop that ignores the boolean return value.
func PopValue[Map ~map[K]V, K comparable, V any](m Map) (K, V) {
	k, v, _ := Pop(m)
	return k, v
}

// Single returns the single key-value pair from the map.
// It also returns an integer indicating the number of pairs in the map (0, 1, or 2).
// If the map contains zero or more than one pair, it returns the zero values for the key and value.
func Single[Map ~map[K]V, K comparable, V any](m Map) (K, V, int) {
	var (
		ct    int
		key   K
		value V
	)
	for k, v := range m {
		if ct == 0 {
			key = k
			value = v
		} else {
			var (
				zeroK K
				zeroV V
			)
			return zeroK, zeroV, 2
		}
		ct++
	}
	return key, value, ct
}

// SingleValue returns the single key-value pair from the map.
// It's a convenience wrapper around Single that ignores the integer return value.
func SingleValue[Map ~map[K]V, K comparable, V any](m Map) (K, V) {
	k, v, _ := Single(m)
	return k, v
}
