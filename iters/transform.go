package iters

import (
	"fmt"
	"iter"
	"slices"
	"xstd/constraints"
)

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

func TapKey[K, V any](seq iter.Seq2[K, V], f func(K)) iter.Seq2[K, V] {
	return Tap2(seq, func(k K, v V) { f(k) })
}

func TapValue[K, V any](seq iter.Seq2[K, V], f func(V)) iter.Seq2[K, V] {
	return Tap2(seq, func(k K, v V) { f(v) })
}

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

func Concat[V any](seq ...iter.Seq[V]) iter.Seq[V] {
	return Flatten(slices.Values(seq))
}

func Concat2[K, V any](seq ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return Flatten2(slices.Values(seq))
}

func Map[VIn, VOut any](seq iter.Seq[VIn], f func(VIn) VOut) iter.Seq[VOut] {
	return func(yield func(VOut) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func Map2[KIn, VIn, KOut, VOut any](seq iter.Seq2[KIn, VIn], f func(KIn, VIn) (KOut, VOut)) iter.Seq2[KOut, VOut] {
	return func(yield func(KOut, VOut) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

func MapKeys[KIn, V, KOut any](seq iter.Seq2[KIn, V], f func(KIn) KOut) iter.Seq2[KOut, V] {
	return Map2(seq, func(k KIn, v V) (KOut, V) { return f(k), v })
}

func MapValues[K, VIn, VOut any](seq iter.Seq2[K, VIn], f func(VIn) VOut) iter.Seq2[K, VOut] {
	return Map2(seq, func(k K, v VIn) (K, VOut) { return k, f(v) })
}

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

func MapLift[VIn, KOut, VOut any](seq iter.Seq[VIn], f func(VIn) (KOut, VOut)) iter.Seq2[KOut, VOut] {
	return func(yield func(KOut, VOut) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func MapLower[KIn, VIn, VOut any](seq iter.Seq2[KIn, VIn], f func(KIn, VIn) VOut) iter.Seq[VOut] {
	return func(yield func(VOut) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

func Values[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				return
			}
		}
	}
}

func Keys[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range seq {
			if !yield(k) {
				return
			}
		}
	}
}

func Swap[K, V any](seq iter.Seq2[K, V]) iter.Seq2[V, K] {
	return Map2(seq, func(k K, v V) (V, K) { return v, k })
}

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

func Split[K, V any](seq iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V], func()) {
	var (
		next, stop   = iter.Pull2(seq)
		kDone, vDone bool
		kBuf         []K
		vBuf         []V
	)

	kSeq := func(yield func(K) bool) {
		if kDone {
			return
		}

		k, v, ok := next()
		if !ok {
			kDone = true
			vDone = true
			return
		}

		vBuf = append(vBuf, v)
		if !yield(k) {
			kDone = true
			if vDone {
				stop()
			}
			return
		}
	}

	vSeq := func(yield func(V) bool) {
		if vDone {
			return
		}

		k, v, ok := next()
		if !ok {
			kDone = true
			vDone = true
			return
		}

		kBuf = append(kBuf, k)
		if !yield(v) {
			vDone = true
			if kDone {
				stop()
			}
			return
		}
	}

	return kSeq, vSeq, stop
}

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

func Empty[V any]() iter.Seq[V] {
	return func(yield func(V) bool) {}
}

func Empty2[K, V any]() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {}
}

func Singleton[V any](v V) iter.Seq[V] {
	return func(yield func(V) bool) {
		yield(v)
	}
}

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
		for i := start; i < end; i += step {
			if !yield(i) {
				return
			}
		}
	}
}

func Range[Int constraints.Integer](start, end Int) iter.Seq[Int] {
	return rangeInternal("Range", start, end, 1)
}

func RangeStep[Int constraints.Integer](start, end, step Int) iter.Seq[Int] {
	return rangeInternal("RangeStep", start, end, step)
}

func Slice[Int constraints.Integer, V any](seq iter.Seq[V], start, end Int) iter.Seq[V] {
	if start > end {
		panic("iters.Slice: start > end")
	}
	return func(yield func(V) bool) {
		var i Int
		for v := range seq {
			if i >= end {
				return
			}

			if i < start {
				i++
				continue
			}

			if !yield(v) {
				return
			}
			i++
		}
	}
}

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
