package iters

import (
	"cmp"
	"iter"
	"xstd/constraints"
)

func All[V any](seq iter.Seq[V], f func(V) bool) bool {
	for v := range seq {
		if !f(v) {
			return false
		}
	}
	return true
}

func All2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) bool {
	for k, v := range seq {
		if !f(k, v) {
			return false
		}
	}
	return true
}

func AllKeys[K, V any](seq iter.Seq2[K, V], f func(K) bool) bool {
	return All2(seq, func(k K, v V) bool { return f(k) })
}

func AllValues[K, V any](seq iter.Seq2[K, V], f func(V) bool) bool {
	return All2(seq, func(k K, v V) bool { return f(v) })
}

func Any[V any](seq iter.Seq[V], f func(V) bool) bool {
	for v := range seq {
		if f(v) {
			return true
		}
	}
	return false
}

func Any2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) bool {
	for k, v := range seq {
		if f(k, v) {
			return true
		}
	}
	return false
}

func Contains[V comparable](seq iter.Seq[V], needle V) bool {
	return Any(seq, func(v V) bool { return v == needle })
}

func Contains2[K, V comparable](seq iter.Seq2[K, V], needleK K, needleV V) bool {
	return Any2(seq, func(k K, v V) bool { return k == needleK && v == needleV })
}

func ContainsKey[K comparable, V any](seq iter.Seq2[K, V], needle K) bool {
	return Any2(seq, func(k K, v V) bool { return k == needle })
}

func ContainsValue[K any, V comparable](seq iter.Seq2[K, V], needle V) bool {
	return Any2(seq, func(k K, v V) bool { return v == needle })
}

func Find[V any](seq iter.Seq[V], f func(V) bool) (V, bool) {
	for v := range seq {
		if f(v) {
			return v, true
		}
	}
	var zero V
	return zero, false
}

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

func ForEach[V any](seq iter.Seq[V], f func(V)) {
	for v := range seq {
		f(v)
	}
}

func ForEach2[K, V any](seq iter.Seq2[K, V], f func(K, V)) {
	for k, v := range seq {
		f(k, v)
	}
}

func Drain[V any](seq iter.Seq[V]) {
	for v := range seq {
		_ = v
	}
}

func Drain2[K, V any](seq iter.Seq2[K, V]) {
	for k, v := range seq {
		_, _ = k, v
	}
}

func Reduce[Sum, V any](sum Sum, seq iter.Seq[V], f func(Sum, V) Sum) Sum {
	for v := range seq {
		sum = f(sum, v)
	}
	return sum
}

func Reduce2[Sum, K, V any](sum Sum, seq iter.Seq2[K, V], f func(Sum, K, V) Sum) Sum {
	for k, v := range seq {
		sum = f(sum, k, v)
	}
	return sum
}

func Count[Int constraints.Integer, V any](seq iter.Seq[V], f func(V) bool) Int {
	var ct Int
	for v := range seq {
		if f(v) {
			ct++
		}
	}
	return ct
}

func Count2[Int constraints.Integer, K, V any](seq iter.Seq2[K, V], f func(K, V) bool) Int {
	var ct Int
	for k, v := range seq {
		if f(k, v) {
			ct++
		}
	}
	return ct
}

func Len[Int constraints.Integer, V any](seq iter.Seq[V]) Int {
	var ct Int
	for range seq {
		ct++
	}
	return ct
}

func Len2[Int constraints.Integer, K, V any](seq iter.Seq2[K, V]) Int {
	var ct Int
	for range seq {
		ct++
	}
	return ct
}

func First[V any](seq iter.Seq[V]) (V, bool) {
	for v := range seq {
		return v, true
	}
	var zero V
	return zero, false
}

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

func FirstValue[V any](seq iter.Seq[V]) V {
	v, _ := First(seq)
	return v
}

func First2Value[K, V any](seq iter.Seq2[K, V]) (K, V) {
	k, v, _ := First2(seq)
	return k, v
}

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

func IndexValue[Int constraints.Integer, V any](seq iter.Seq[V], idx Int) V {
	v, _ := Index(seq, idx)
	return v
}

func Index2Value[Int constraints.Integer, K, V any](seq iter.Seq2[K, V], idx Int) (K, V) {
	k, v, _ := Index2(seq, idx)
	return k, v
}

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

func Sum[V cmp.Ordered](seq iter.Seq[V]) V {
	var sum V
	for v := range seq {
		sum += v
	}
	return sum
}
