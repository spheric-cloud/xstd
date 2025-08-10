// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package iters

import (
	"fmt"
	"iter"
	"slices"

	"spheric.cloud/xstd/constraints"
)

// Tap calls f for each value in seq, and yields the value.
func Tap[V any](seq iter.Seq[V], f func(V)) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			f(v)
			if !yield(v) {
				return
			}
		}
	}
}

// Tap2 calls f for each key-value pair in seq, and yields the key-value pair.
func Tap2[K, V any](seq iter.Seq2[K, V], f func(K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			f(k, v)
			if !yield(k, v) {
				return
			}
		}
	}
}

// TapKey calls f for each key in seq, and yields the key-value pair.
func TapKey[K, V any](seq iter.Seq2[K, V], f func(K)) iter.Seq2[K, V] {
	return Tap2(seq, func(k K, v V) { f(k) })
}

// TapValue calls f for each value in seq, and yields the key-value pair.
func TapValue[K, V any](seq iter.Seq2[K, V], f func(V)) iter.Seq2[K, V] {
	return Tap2(seq, func(k K, v V) { f(v) })
}

// Flatten flattens a sequence of sequences into a single sequence.
func Flatten[V any](seq iter.Seq[iter.Seq[V]]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for seq := range seq {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Flatten2 flattens a sequence of sequences of key-value pairs into a single sequence of key-value pairs.
func Flatten2[K, V any](seq iter.Seq[iter.Seq2[K, V]]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for seq := range seq {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Concat concatenates multiple sequences into a single sequence.
func Concat[V any](seq ...iter.Seq[V]) iter.Seq[V] {
	return Flatten(slices.Values(seq))
}

// Concat2 concatenates multiple sequences of key-value pairs into a single sequence of key-value pairs.
func Concat2[K, V any](seq ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return Flatten2(slices.Values(seq))
}

// Map returns a new iterator that yields the results of calling f on each value from seq.
func Map[VIn, VOut any](seq iter.Seq[VIn], f func(VIn) VOut) iter.Seq[VOut] {
	return func(yield func(VOut) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// Map2 returns a new iterator that yields the results of calling f on each key-value pair from seq.
func Map2[KIn, VIn, KOut, VOut any](seq iter.Seq2[KIn, VIn], f func(KIn, VIn) (KOut, VOut)) iter.Seq2[KOut, VOut] {
	return func(yield func(KOut, VOut) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

// MapKeys returns a new iterator that yields the results of calling f on each key from seq.
func MapKeys[KIn, V, KOut any](seq iter.Seq2[KIn, V], f func(KIn) KOut) iter.Seq2[KOut, V] {
	return Map2(seq, func(k KIn, v V) (KOut, V) { return f(k), v })
}

// MapValues returns a new iterator that yields the results of calling f on each value from seq.
func MapValues[K, VIn, VOut any](seq iter.Seq2[K, VIn], f func(VIn) VOut) iter.Seq2[K, VOut] {
	return Map2(seq, func(k K, v VIn) (K, VOut) { return k, f(v) })
}

// FlatMap returns a new iterator that yields the results of calling f on each value from seq and flattening the result.
func FlatMap[VIn, VOut any](seq iter.Seq[VIn], f func(VIn) iter.Seq[VOut]) iter.Seq[VOut] {
	return func(yield func(VOut) bool) {
		for vIn := range seq {
			for vOut := range f(vIn) {
				if !yield(vOut) {
					return
				}
			}
		}
	}
}

// FlatMap2 returns a new iterator that yields the results of calling f on each key-value pair from seq and flattening the result.
func FlatMap2[KIn, VIn, KOut, VOut any](seq iter.Seq2[KIn, VIn], f func(KIn, VIn) iter.Seq2[KOut, VOut]) iter.Seq2[KOut, VOut] {
	return func(yield func(KOut, VOut) bool) {
		for kIn, vIn := range seq {
			for kOut, vOut := range f(kIn, vIn) {
				if !yield(kOut, vOut) {
					return
				}
			}
		}
	}
}

// FlatMapKeys returns a new iterator that yields the results of calling f on each key from seq and flattening the result.
func FlatMapKeys[KIn, V, KOut any](seq iter.Seq2[KIn, V], f func(KIn) iter.Seq[KOut]) iter.Seq2[KOut, V] {
	return FlatMap2(seq, func(kIn KIn, v V) iter.Seq2[KOut, V] {
		return func(yield func(KOut, V) bool) {
			for kOut := range f(kIn) {
				if !yield(kOut, v) {
					return
				}
			}
		}
	})
}

// FlatMapValues returns a new iterator that yields the results of calling f on each value from seq and flattening the result.
func FlatMapValues[K, VIn, VOut any](seq iter.Seq2[K, VIn], f func(VIn) iter.Seq[VOut]) iter.Seq2[K, VOut] {
	return FlatMap2(seq, func(k K, v VIn) iter.Seq2[K, VOut] {
		return func(yield func(K, VOut) bool) {
			for vOut := range f(v) {
				if !yield(k, vOut) {
					return
				}
			}
		}
	})
}

// FlatMapLift returns a new iterator that yields the results of calling f on each value from seq and flattening the result.
// This function lifts a sequence of values to a sequence of key-value pairs.
func FlatMapLift[VIn, KOut, VOut any](seq iter.Seq[VIn], f func(VIn) iter.Seq2[KOut, VOut]) iter.Seq2[KOut, VOut] {
	return func(yield func(KOut, VOut) bool) {
		for vIn := range seq {
			for kOut, vOut := range f(vIn) {
				if !yield(kOut, vOut) {
					return
				}
			}
		}
	}
}

// FlatMapLower returns a new iterator that yields the results of calling f on each key-value pair from seq and flattening the result.
// This function lowers a sequence of key-value pairs to a sequence of values.
func FlatMapLower[KIn, VIn, VOut any](seq iter.Seq2[KIn, VIn], f func(KIn, VIn) iter.Seq[VOut]) iter.Seq[VOut] {
	return func(yield func(VOut) bool) {
		for kIn, vIn := range seq {
			for vOut := range f(kIn, vIn) {
				if !yield(vOut) {
					return
				}
			}
		}
	}
}

// MapLift returns a new iterator that yields the results of calling f on each value from seq.
// This function lifts a sequence of values to a sequence of key-value pairs.
func MapLift[VIn, KOut, VOut any](seq iter.Seq[VIn], f func(VIn) (KOut, VOut)) iter.Seq2[KOut, VOut] {
	return func(yield func(KOut, VOut) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// LiftZeroValues lifts a sequence of keys to a sequence of key-value pairs, with zero values.
func LiftZeroValues[V, K any](seq iter.Seq[K]) iter.Seq2[K, V] {
	return MapLift(seq, func(k K) (K, V) {
		var zero V
		return k, zero
	})
}

// LiftZeroKeys lifts a sequence of values to a sequence of key-value pairs, with zero keys.
func LiftZeroKeys[K, V any](seq iter.Seq[V]) iter.Seq2[K, V] {
	return MapLift(seq, func(v V) (K, V) {
		var zero K
		return zero, v
	})
}

func LiftSingletonKey[K, V any](seq iter.Seq[V], k K) iter.Seq2[K, V] {
	return MapLift(seq, func(v V) (K, V) {
		return k, v
	})
}

func LiftSingletonValue[K, V any](seq iter.Seq[K], v V) iter.Seq2[K, V] {
	return MapLift(seq, func(k K) (K, V) {
		return k, v
	})
}

// MapLower returns a new iterator that yields the results of calling f on each key-value pair from seq.
// This function lowers a sequence of key-value pairs to a sequence of values.
func MapLower[KIn, VIn, VOut any](seq iter.Seq2[KIn, VIn], f func(KIn, VIn) VOut) iter.Seq[VOut] {
	return func(yield func(VOut) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

// Values returns a new iterator that yields only the values from seq.
func Values[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				return
			}
		}
	}
}

// Keys returns a new iterator that yields only the keys from seq.
func Keys[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range seq {
			if !yield(k) {
				return
			}
		}
	}
}

// Swap returns a new iterator that swaps the key and value of each pair in seq.
func Swap[K, V any](seq iter.Seq2[K, V]) iter.Seq2[V, K] {
	return Map2(seq, func(k K, v V) (V, K) { return v, k })
}

// Join returns a new iterator that yields pairs of values from two sequences.
func Join[K, V any](kSeq iter.Seq[K], vSeq iter.Seq[V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		nextV, stop := iter.Pull(vSeq)
		defer stop()

		for k := range kSeq {
			v, ok := nextV()
			if !ok {
				return
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

// Split splits a sequence of key-value pairs into two sequences, one of keys and one of values.
func Split[K, V any](seq iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V], func()) {
	var (
		next, stop   = iter.Pull2(seq)
		kDone, vDone bool
		kBuf         []K
		vBuf         []V
	)

	kSeq := func(yield func(K) bool) {
		for {
			for len(kBuf) > 0 {
				k := kBuf[0]
				kBuf = kBuf[1:]
				if !yield(k) {
					kBuf = nil
					kDone = true
					if vDone {
						stop()
					}
					return
				}
			}
			if kDone {
				return
			}

			k, v, ok := next()
			if !ok {
				kDone = true
				vDone = true
				return
			}

			if !vDone {
				vBuf = append(vBuf, v)
			}
			if !yield(k) {
				kDone = true
				if vDone {
					stop()
				}
				return
			}
		}
	}

	vSeq := func(yield func(V) bool) {
		for {
			for len(vBuf) > 0 {
				v := vBuf[0]
				vBuf = vBuf[1:]
				if !yield(v) {
					vBuf = nil
					vDone = true
					if kDone {
						stop()
					}
					return
				}
			}
			if vDone {
				return
			}

			k, v, ok := next()
			if !ok {
				kDone = true
				vDone = true
				return
			}

			if !kDone {
				kBuf = append(kBuf, k)
			}
			if !yield(v) {
				vDone = true
				if kDone {
					stop()
				}
				return
			}
		}
	}

	return kSeq, vSeq, stop
}

// Enumerate returns a new iterator that yields the index and value of each value in seq.
func Enumerate[Int constraints.Integer, V any](s iter.Seq[V]) iter.Seq2[Int, V] {
	return func(yield func(Int, V) bool) {
		var i Int
		for v := range s {
			if !yield(i, v) {
				return
			}

			i++
		}
	}
}

// Empty returns an empty iterator.
func Empty[V any]() iter.Seq[V] {
	return func(yield func(V) bool) {}
}

// Empty2 returns an empty iterator of key-value pairs.
func Empty2[K, V any]() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {}
}

// Singleton returns an iterator that yields a single value.
func Singleton[V any](v V) iter.Seq[V] {
	return func(yield func(V) bool) {
		yield(v)
	}
}

// Singleton2 returns an iterator that yields a single key-value pair.
func Singleton2[K, V any](k K, v V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		yield(k, v)
	}
}

func rangeInternal[Int constraints.Integer](name string, start, end, step Int) iter.Seq[Int] {
	if step == 0 {
		panic(fmt.Sprintf("iters.%s step cannot be zero", name))
	}
	if start < end && step < 0 || start > end && step > 0 {
		panic(fmt.Sprintf("iters.%s %d to %d step %d is not a valid range", name, start, end, step))
	}
	return func(yield func(Int) bool) {
		if step > 0 {
			for i := start; i < end; i += step {
				if !yield(i) {
					return
				}
			}
		} else {
			for i := start; i > end; i += step {
				if !yield(i) {
					return
				}
			}
		}
	}
}

// Range returns a new iterator that yields values from start to end (exclusive).
func Range[Int constraints.Integer](start, end Int) iter.Seq[Int] {
	return rangeInternal("Range", start, end, 1)
}

// RangeStep returns a new iterator that yields values from start to end (exclusive) with the given step.
func RangeStep[Int constraints.Integer](start, end, step Int) iter.Seq[Int] {
	return rangeInternal("RangeStep", start, end, step)
}

// Slice returns a new iterator that yields a slice of seq from start to end (exclusive).
func Slice[Int constraints.Integer, V any](seq iter.Seq[V], start, end Int) iter.Seq[V] {
	if start > end {
		panic("iters.Slice: start > end")
	}
	if start == end {
		return Empty[V]()
	}
	return func(yield func(V) bool) {
		var i Int
		for v := range seq {
			if i < start {
				i++
				continue
			}

			if !yield(v) {
				return
			}
			i++

			if i >= end {
				return
			}
		}
	}
}

// Slice2 returns a new iterator that yields a slice of seq from start to end (exclusive).
func Slice2[Int constraints.Integer, K, V any](seq iter.Seq2[K, V], start, end Int) iter.Seq2[K, V] {
	if start > end {
		panic("iters.Slice: start > end")
	}
	return func(yield func(K, V) bool) {
		var i Int
		for k, v := range seq {
			if i >= end {
				return
			}

			if i < start {
				i++
				continue
			}

			if !yield(k, v) {
				return
			}
			i++
		}
	}
}
