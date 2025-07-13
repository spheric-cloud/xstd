package iters

import "iter"

func FromNext[V any](next func() (V, bool)) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v, ok := next(); ok; v, ok = next() {
			if !yield(v) {
				return
			}
		}
	}
}

func FromNext2[K, V any](next func() (K, V, bool)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v, ok := next(); ok; k, v, ok = next() {
			if !yield(k, v) {
				return
			}
		}
	}
}

func Repeat[V any](v V, n int) iter.Seq[V] {
	if n < 0 {
		panic("iters.Repeat: negative n")
	}
	return func(yield func(V) bool) {
		for i := 0; i < n; i++ {
			if !yield(v) {
				return
			}
		}
	}
}

func Repeat2[K, V any](k K, v V, n int) iter.Seq2[K, V] {
	if n < 0 {
		panic("iters.Repeat2: negative n")
	}
	return func(yield func(K, V) bool) {
		for i := 0; i < n; i++ {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Cycle cycles over the given sequence by repeatedly calling it.
//
// Cycle will not work with non-reusable iterators, as it does not cache.
func Cycle[V any](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Cycle2 cycles over the given sequence by repeatedly calling it.
//
// Cycle2 will not work with non-reusable iterators, as it does not cache.
func Cycle2[K, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
