// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package iters

import (
	"cmp"
	"iter"

	"spheric.cloud/xstd/constraints"
	"spheric.cloud/xstd/gen"
)

// All returns true if f returns true for all values in seq.
func All[V any](seq iter.Seq[V], f func(V) bool) bool {
	for v := range seq {
		if !f(v) {
			return false
		}
	}
	return true
}

// All2 returns true if f returns true for all key-value pairs in seq.
func All2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) bool {
	for k, v := range seq {
		if !f(k, v) {
			return false
		}
	}
	return true
}

// AllKeys returns true if f returns true for all keys in seq.
func AllKeys[K, V any](seq iter.Seq2[K, V], f func(K) bool) bool {
	return All2(seq, func(k K, v V) bool { return f(k) })
}

// AllValues returns true if f returns true for all values in seq.
func AllValues[K, V any](seq iter.Seq2[K, V], f func(V) bool) bool {
	return All2(seq, func(k K, v V) bool { return f(v) })
}

// Any returns true if f returns true for any value in seq.
func Any[V any](seq iter.Seq[V], f func(V) bool) bool {
	for v := range seq {
		if f(v) {
			return true
		}
	}
	return false
}

// Any2 returns true if f returns true for any key-value pair in seq.
func Any2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) bool {
	for k, v := range seq {
		if f(k, v) {
			return true
		}
	}
	return false
}

// Contains returns true if seq contains needle.
func Contains[V comparable](seq iter.Seq[V], needle V) bool {
	return Any(seq, func(v V) bool { return v == needle })
}

// Contains2 returns true if seq contains a key-value pair equal to needleK and needleV.
func Contains2[K, V comparable](seq iter.Seq2[K, V], needleK K, needleV V) bool {
	return Any2(seq, func(k K, v V) bool { return k == needleK && v == needleV })
}

// ContainsKey returns true if seq contains a key equal to needle.
func ContainsKey[K comparable, V any](seq iter.Seq2[K, V], needle K) bool {
	return Any2(seq, func(k K, v V) bool { return k == needle })
}

// ContainsValue returns true if seq contains a value equal to needle.
func ContainsValue[K any, V comparable](seq iter.Seq2[K, V], needle V) bool {
	return Any2(seq, func(k K, v V) bool { return v == needle })
}

// Find returns the first value in seq for which f returns true.
func Find[V any](seq iter.Seq[V], f func(V) bool) (V, bool) {
	for v := range seq {
		if f(v) {
			return v, true
		}
	}
	var zero V
	return zero, false
}

// Find2 returns the first key-value pair in seq for which f returns true.
func Find2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) (K, V, bool) {
	for k, v := range seq {
		if f(k, v) {
			return k, v, true
		}
	}
	var (
		zeroK K
		zeroV V
	)
	return zeroK, zeroV, false
}

// ForEach calls f for each value in seq.
func ForEach[V any](seq iter.Seq[V], f func(V)) {
	for v := range seq {
		f(v)
	}
}

// ForEach2 calls f for each key-value pair in seq.
func ForEach2[K, V any](seq iter.Seq2[K, V], f func(K, V)) {
	for k, v := range seq {
		f(k, v)
	}
}

// Drain consumes all values in seq.
func Drain[V any](seq iter.Seq[V]) {
	for v := range seq {
		_ = v
	}
}

// Drain2 consumes all key-value pairs in seq.
func Drain2[K, V any](seq iter.Seq2[K, V]) {
	for k, v := range seq {
		_, _ = k, v
	}
}

// Reduce reduces seq to a single value by calling f with the current sum and the next value.
func Reduce[Sum, V any](sum Sum, seq iter.Seq[V], f func(Sum, V) Sum) Sum {
	for v := range seq {
		sum = f(sum, v)
	}
	return sum
}

// Reduce2 reduces seq to a single value by calling f with the current sum and the next key-value pair.
func Reduce2[Sum, K, V any](sum Sum, seq iter.Seq2[K, V], f func(Sum, K, V) Sum) Sum {
	for k, v := range seq {
		sum = f(sum, k, v)
	}
	return sum
}

// Count returns the number of values in seq for which f returns true.
func Count[Int constraints.Integer, V any](seq iter.Seq[V], f func(V) bool) Int {
	var ct Int
	for v := range seq {
		if f(v) {
			ct++
		}
	}
	return ct
}

// Count2 returns the number of key-value pairs in seq for which f returns true.
func Count2[Int constraints.Integer, K, V any](seq iter.Seq2[K, V], f func(K, V) bool) Int {
	var ct Int
	for k, v := range seq {
		if f(k, v) {
			ct++
		}
	}
	return ct
}

// Len returns the number of values in seq.
func Len[Int constraints.Integer, V any](seq iter.Seq[V]) Int {
	var ct Int
	for range seq {
		ct++
	}
	return ct
}

// Len2 returns the number of key-value pairs in seq.
func Len2[Int constraints.Integer, K, V any](seq iter.Seq2[K, V]) Int {
	var ct Int
	for range seq {
		ct++
	}
	return ct
}

// First returns the first value in seq.
func First[V any](seq iter.Seq[V]) (V, bool) {
	for v := range seq {
		return v, true
	}
	var zero V
	return zero, false
}

// First2 returns the first key-value pair in seq.
func First2[K, V any](seq iter.Seq2[K, V]) (K, V, bool) {
	for k, v := range seq {
		return k, v, true
	}
	var (
		zeroK K
		zeroV V
	)
	return zeroK, zeroV, false
}

// FirstValue returns the first value in seq, or the zero value if seq is empty.
func FirstValue[V any](seq iter.Seq[V]) V {
	v, _ := First(seq)
	return v
}

// First2Value returns the first key-value pair in seq, or the zero values if seq is empty.
func First2Value[K, V any](seq iter.Seq2[K, V]) (K, V) {
	k, v, _ := First2(seq)
	return k, v
}

// Index returns the value at idx in seq.
func Index[Int constraints.Integer, V any](seq iter.Seq[V], idx Int) (V, bool) {
	var ct Int
	for v := range seq {
		if ct == idx {
			return v, true
		}
		ct++
	}
	var zero V
	return zero, false
}

// Index2 returns the key-value pair at idx in seq.
func Index2[Int constraints.Integer, K, V any](seq iter.Seq2[K, V], idx Int) (K, V, bool) {
	var ct Int
	for k, v := range seq {
		if ct == idx {
			return k, v, true
		}
		ct++
	}
	var (
		zeroK K
		zeroV V
	)
	return zeroK, zeroV, false
}

// IndexValue returns the value at idx in seq, or the zero value if idx is out of bounds.
func IndexValue[Int constraints.Integer, V any](seq iter.Seq[V], idx Int) V {
	v, _ := Index(seq, idx)
	return v
}

// Index2Value returns the key-value pair at idx in seq, or the zero values if idx is out of bounds.
func Index2Value[Int constraints.Integer, K, V any](seq iter.Seq2[K, V], idx Int) (K, V) {
	k, v, _ := Index2(seq, idx)
	return k, v
}

// Max returns the maximum value in seq.
// It panics if seq is empty.
func Max[V cmp.Ordered](seq iter.Seq[V]) V {
	var (
		best V
		ok   bool
	)
	for v := range seq {
		if !ok || v > best {
			best = v
			ok = true
		}
	}
	if !ok {
		panic("iters.Max: empty seq")
	}
	return best
}

// MaxFunc returns the maximum value in seq, using the given comparison function.
// It panics if seq is empty.
func MaxFunc[V any](seq iter.Seq[V], compare func(V, V) int) V {
	var (
		best V
		ok   bool
	)
	for v := range seq {
		if !ok || compare(v, best) > 0 {
			best = v
			ok = true
		}
	}
	if !ok {
		panic("iters.MaxFunc: empty seq")
	}
	return best
}

// Min returns the minimum value in seq.
// It panics if seq is empty.
func Min[V cmp.Ordered](seq iter.Seq[V]) V {
	var (
		best V
		ok   bool
	)
	for v := range seq {
		if !ok || v < best {
			best = v
			ok = true
		}
	}
	if !ok {
		panic("iters.Min: empty seq")
	}
	return best
}

// MinFunc returns the minimum value in seq, using the given comparison function.
// It panics if seq is empty.
func MinFunc[V any](seq iter.Seq[V], compare func(V, V) int) V {
	var (
		best V
		ok   bool
	)
	for v := range seq {
		if !ok || compare(v, best) < 0 {
			best = v
			ok = true
		}
	}
	if !ok {
		panic("iters.MinFunc: empty seq")
	}
	return best
}

// Sum returns the sum of all values in seq.
func Sum[V cmp.Ordered](seq iter.Seq[V]) V {
	var sum V
	for v := range seq {
		sum += v
	}
	return sum
}

// Equal returns true if seq1 and seq2 are equal.
func Equal[V comparable](seq1, seq2 iter.Seq[V]) bool {
	next1, stop := iter.Pull(seq1)
	defer stop()
	next2, stop := iter.Pull(seq2)
	defer stop()

	for {
		v1, ok1 := next1()
		v2, ok2 := next2()
		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if v1 != v2 {
			return false
		}
	}
}

// EqualFunc returns true if seq1 and seq2 are equal, using the given comparison function.
func EqualFunc[V1, V2 any](seq1 iter.Seq[V1], seq2 iter.Seq[V2], f func(V1, V2) bool) bool {
	next1, stop := iter.Pull(seq1)
	defer stop()
	next2, stop := iter.Pull(seq2)
	defer stop()

	for {
		v1, ok1 := next1()
		v2, ok2 := next2()
		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if !f(v1, v2) {
			return false
		}
	}
}

// Equal2 returns true if seq1 and seq2 are equal.
func Equal2[K, V comparable](seq1, seq2 iter.Seq2[K, V]) bool {
	next1, stop := iter.Pull2(seq1)
	defer stop()
	next2, stop := iter.Pull2(seq2)
	defer stop()

	for {
		k1, v1, ok1 := next1()
		k2, v2, ok2 := next2()
		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if k1 != k2 || v1 != v2 {
			return false
		}
	}
}

// EqualFunc2 returns true if seq1 and seq2 are equal, using the given comparison function.
func EqualFunc2[K1, V1, K2, V2 any](seq1 iter.Seq2[K1, V1], seq2 iter.Seq2[K2, V2], f func(K1, V1, K2, V2) bool) bool {
	next1, stop := iter.Pull2(seq1)
	defer stop()
	next2, stop := iter.Pull2(seq2)
	defer stop()

	for {
		k1, v1, ok1 := next1()
		k2, v2, ok2 := next2()
		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if !f(k1, v1, k2, v2) {
			return false
		}
	}
}

// Single returns the single value in seq.
// It returns the zero value and 0 if seq is empty.
// It returns the zero value and 2 if seq has more than one value.
func Single[V any](seq iter.Seq[V]) (V, int) {
	var (
		ct    int
		value V
	)
	for v := range seq {
		if ct == 0 {
			value = v
		} else {
			return gen.Zero[V](), 2
		}
		ct++
	}
	return value, ct
}

// SingleValue returns the single value in seq.
// It returns the zero value if seq is empty or has more than one value.
func SingleValue[V any](seq iter.Seq[V]) V {
	v, _ := Single(seq)
	return v
}

// Single2 returns the single key-value pair in seq.
// It returns the zero values and 0 if seq is empty.
// It returns the zero values and 2 if seq has more than one value.
func Single2[K, V any](seq iter.Seq2[K, V]) (K, V, int) {
	var (
		ct    int
		key   K
		value V
	)
	for k, v := range seq {
		if ct == 0 {
			key = k
			value = v
		} else {
			return gen.Zero[K](), gen.Zero[V](), 2
		}
		ct++
	}
	return key, value, ct
}

// Single2Value returns the single key-value pair in seq.
// It returns the zero values if seq is empty or has more than one value.
func Single2Value[K, V any](seq iter.Seq2[K, V]) (K, V) {
	k, v, _ := Single2(seq)
	return k, v
}
