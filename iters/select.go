// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package iters

import (
	"iter"

	"spheric.cloud/xstd/constraints"
)

// Unique returns a new iterator that yields only the unique values from seq.
func Unique[V comparable](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		seen := make(map[V]struct{})
		for v := range seq {
			if _, ok := seen[v]; ok {
				continue
			}

			seen[v] = struct{}{}
			if !yield(v) {
				return
			}
		}
	}
}

// UniqueFunc returns a new iterator that yields only the unique values from seq,
// using the given function to generate the key for uniqueness checks.
func UniqueFunc[V any, Key comparable](seq iter.Seq[V], f func(V) Key) iter.Seq[V] {
	return func(yield func(V) bool) {
		seen := make(map[Key]struct{})
		for v := range seq {
			key := f(v)
			if _, ok := seen[key]; ok {
				continue
			}

			seen[key] = struct{}{}
			if !yield(v) {
				return
			}
		}
	}
}

// Unique2 returns a new iterator that yields only the unique key-value pairs from seq.
func Unique2[K, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	type key struct {
		K K
		V V
	}
	return func(yield func(K, V) bool) {
		seen := make(map[key]struct{})
		for k, v := range seq {
			if _, ok := seen[key{k, v}]; ok {
				continue
			}

			seen[key{k, v}] = struct{}{}
			if !yield(k, v) {
				return
			}
		}
	}
}

// UniqueFunc2 returns a new iterator that yields only the unique key-value pairs from seq,
// using the given function to generate the key for uniqueness checks.
func UniqueFunc2[K, V any, Key comparable](seq iter.Seq2[K, V], f func(K, V) Key) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seen := make(map[Key]struct{})
		for k, v := range seq {
			key := f(k, v)
			if _, ok := seen[key]; ok {
				continue
			}

			seen[key] = struct{}{}
			if !yield(k, v) {
				return
			}
		}
	}
}

// Filter returns a new iterator that yields only the values from seq for which f returns true.
func Filter[V any](seq iter.Seq[V], f func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !f(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}

// Filter2 returns a new iterator that yields only the key-value pairs from seq for which f returns true.
func Filter2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !f(k, v) {
				continue
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

// FilterKeys returns a new iterator that yields only the key-value pairs from seq for which f returns true for the key.
func FilterKeys[K, V any](seq iter.Seq2[K, V], f func(K) bool) iter.Seq2[K, V] {
	return Filter2(seq, func(k K, v V) bool { return f(k) })
}

// FilterValues returns a new iterator that yields only the key-value pairs from seq for which f returns true for the value.
func FilterValues[K, V any](seq iter.Seq2[K, V], f func(V) bool) iter.Seq2[K, V] {
	return Filter2(seq, func(k K, v V) bool { return f(v) })
}

// MapOK returns a new iterator that yields the results of calling f on each value from seq.
// If f returns false, the value is skipped.
func MapOK[VIn, VOut any](seq iter.Seq[VIn], f func(VIn) (VOut, bool)) iter.Seq[VOut] {
	return func(yield func(VOut) bool) {
		for vIn := range seq {
			vOut, ok := f(vIn)
			if !ok {
				continue
			}

			if !yield(vOut) {
				return
			}
		}
	}
}

// MapOK2 returns a new iterator that yields the results of calling f on each key-value pair from seq.
// If f returns false, the value is skipped.
func MapOK2[KIn, VIn, KOut, VOut any](seq iter.Seq2[KIn, VIn], f func(KIn, VIn) (KOut, VOut, bool)) iter.Seq2[KOut, VOut] {
	return func(yield func(KOut, VOut) bool) {
		for kIn, vIn := range seq {
			kOut, vOut, ok := f(kIn, vIn)
			if !ok {
				continue
			}

			if !yield(kOut, vOut) {
				return
			}
		}
	}
}

// MapOKKeys returns a new iterator that yields the results of calling f on each key from seq.
// If f returns false, the value is skipped.
func MapOKKeys[KIn, V, KOut any](seq iter.Seq2[KIn, V], f func(KIn) (KOut, bool)) iter.Seq2[KOut, V] {
	return MapOK2(seq, func(kIn KIn, v V) (KOut, V, bool) {
		kOut, ok := f(kIn)
		return kOut, v, ok
	})
}

// MapOKValues returns a new iterator that yields the results of calling f on each value from seq.
// If f returns false, the value is skipped.
func MapOKValues[K, VIn, VOut any](seq iter.Seq2[K, VIn], f func(VIn) (VOut, bool)) iter.Seq2[K, VOut] {
	return MapOK2(seq, func(k K, vIn VIn) (K, VOut, bool) {
		vOut, ok := f(vIn)
		return k, vOut, ok
	})
}

// MapOKLift returns a new iterator that yields the results of calling f on each value from seq.
// This function lifts a sequence of values to a sequence of key-value pairs.
// If f returns false, the value is skipped.
func MapOKLift[VIn, KOut, VOut any](seq iter.Seq[VIn], f func(VIn) (KOut, VOut, bool)) iter.Seq2[KOut, VOut] {
	return func(yield func(KOut, VOut) bool) {
		for vIn := range seq {
			kOut, vOut, ok := f(vIn)
			if !ok {
				continue
			}

			if !yield(kOut, vOut) {
				return
			}
		}
	}
}

// MapOKLower returns a new iterator that yields the results of calling f on each key-value pair from seq.
// This function lowers a sequence of key-value pairs to a sequence of values.
// If f returns false, the value is skipped.
func MapOKLower[KIn, VIn, VOut any](seq iter.Seq2[KIn, VIn], f func(KIn, VIn) (VOut, bool)) iter.Seq[VOut] {
	return func(yield func(VOut) bool) {
		for kIn, vIn := range seq {
			vOut, ok := f(kIn, vIn)
			if !ok {
				continue
			}

			if !yield(vOut) {
				return
			}
		}
	}
}

// Drop returns a new iterator that drops the first n values from seq.
func Drop[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	if n < 0 {
		panic("iters.Drop: negative n")
	}
	return func(yield func(V) bool) {
		var dropped int
		for v := range seq {
			if dropped < n {
				dropped++
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}

// Drop2 returns a new iterator that drops the first n key-value pairs from seq.
func Drop2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	if n < 0 {
		panic("iters.Drop2: negative n")
	}
	return func(yield func(K, V) bool) {
		var dropped int
		for k, v := range seq {
			if dropped < n {
				dropped++
				continue
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

// DropWhile returns a new iterator that drops values from seq while f returns true.
func DropWhile[V any](seq iter.Seq[V], f func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		drop := true
		for v := range seq {
			if drop {
				if f(v) {
					continue
				}

				drop = false
			}
			if !yield(v) {
				return
			}
		}
	}
}

// DropWhile2 returns a new iterator that drops key-value pairs from seq while f returns true.
func DropWhile2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		drop := true
		for k, v := range seq {
			if drop {
				if f(k, v) {
					continue
				}

				drop = false
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

// DropWhileKeys returns a new iterator that drops key-value pairs from seq while f returns true for the key.
func DropWhileKeys[K, V any](seq iter.Seq2[K, V], f func(K) bool) iter.Seq2[K, V] {
	return DropWhile2(seq, func(k K, v V) bool { return f(k) })
}

// DropWhileValues returns a new iterator that drops key-value pairs from seq while f returns true for the value.
func DropWhileValues[K, V any](seq iter.Seq2[K, V], f func(V) bool) iter.Seq2[K, V] {
	return DropWhile2(seq, func(k K, v V) bool { return f(v) })
}

// Take returns a new iterator that takes the first n values from seq.
func Take[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	if n < 0 {
		panic("iters.Take: negative n")
	}
	return func(yield func(V) bool) {
		if n == 0 {
			return
		}

		var taken int
		for v := range seq {
			if !yield(v) {
				return
			}

			taken++
			if taken == n {
				return
			}
		}
	}
}

// Take2 returns a new iterator that takes the first n key-value pairs from seq.
func Take2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if n == 0 {
			return
		}

		var taken int
		for k, v := range seq {
			if !yield(k, v) {
				return
			}

			taken++
			if taken == n {
				return
			}
		}
	}
}

// TakeWhile returns a new iterator that takes values from seq while f returns true.
func TakeWhile[V any](seq iter.Seq[V], f func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !f(v) {
				return
			}

			if !yield(v) {
				return
			}
		}
	}
}

// TakeWhile2 returns a new iterator that takes key-value pairs from seq while f returns true.
func TakeWhile2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !f(k, v) {
				return
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

// TakeWhileKeys returns a new iterator that takes key-value pairs from seq while f returns true for the key.
func TakeWhileKeys[K, V any](seq iter.Seq2[K, V], f func(K) bool) iter.Seq2[K, V] {
	return TakeWhile2(seq, func(k K, v V) bool { return f(k) })
}

// TakeWhileValues returns a new iterator that takes key-value pairs from seq while f returns true for the value.
func TakeWhileValues[K, V any](seq iter.Seq2[K, V], f func(V) bool) iter.Seq2[K, V] {
	return TakeWhile2(seq, func(k K, v V) bool { return f(v) })
}

// Chunk returns a new iterator that yields chunks of size n.
// The last chunk may be smaller than n.
// The chunks must be consumed in order, otherwise there is no guarantee of how ordering will be.
func Chunk[Int constraints.Integer, V any](seq iter.Seq[V], n Int) iter.Seq[iter.Seq[V]] {
	if n <= 0 {
		panic("iters.Chunk: n must be > 0")
	}

	return func(yield func(iter.Seq[V]) bool) {
		next, stop := iter.Pull(seq)
		defer stop()

		var done bool
		for {
			if !yield(func(yield func(V) bool) {
				shouldYield := true
				for i := Int(0); i < n; i++ {
					v, ok := next()
					if !ok {
						done = true
						return
					}

					if shouldYield {
						shouldYield = yield(v)
					}
				}
			}) || done {
				return
			}
		}
	}
}

// Chunk2 returns a new iterator that yields chunks of size n.
// The last chunk may be smaller than n.
func Chunk2[Int constraints.Integer, K, V any](seq iter.Seq2[K, V], n Int) iter.Seq[iter.Seq2[K, V]] {
	if n <= 0 {
		panic("iters.Chunk: n must be > 0")
	}

	return func(yield func(iter.Seq2[K, V]) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()

		var done bool
		for {
			if !yield(func(yield func(K, V) bool) {
				shouldYield := true
				for i := Int(0); i < n; i++ {
					k, v, ok := next()
					if !ok {
						done = true
						return
					}

					if shouldYield {
						shouldYield = yield(k, v)
					}
				}
			}) || done {
				return
			}
		}
	}
}
