package iters

import "iter"

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

func FilterKeys[K, V any](seq iter.Seq2[K, V], f func(K) bool) iter.Seq2[K, V] {
	return Filter2(seq, func(k K, v V) bool { return f(k) })
}

func FilterValues[K, V any](seq iter.Seq2[K, V], f func(V) bool) iter.Seq2[K, V] {
	return Filter2(seq, func(k K, v V) bool { return f(v) })
}

func Collect[VIn, VOut any](seq iter.Seq[VIn], f func(VIn) (VOut, bool)) iter.Seq[VOut] {
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

func Collect2[KIn, VIn, KOut, VOut any](seq iter.Seq2[KIn, VIn], f func(KIn, VIn) (KOut, VOut, bool)) iter.Seq2[KOut, VOut] {
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

func CollectKeys[KIn, V, KOut any](seq iter.Seq2[KIn, V], f func(KIn) (KOut, bool)) iter.Seq2[KOut, V] {
	return Collect2(seq, func(kIn KIn, v V) (KOut, V, bool) {
		kOut, ok := f(kIn)
		return kOut, v, ok
	})
}

func CollectValues[K, VIn, VOut any](seq iter.Seq2[K, VIn], f func(VIn) (VOut, bool)) iter.Seq2[K, VOut] {
	return Collect2(seq, func(k K, vIn VIn) (K, VOut, bool) {
		vOut, ok := f(vIn)
		return k, vOut, ok
	})
}

func CollectLift[VIn, KOut, VOut any](seq iter.Seq[VIn], f func(VIn) (KOut, VOut, bool)) iter.Seq2[KOut, VOut] {
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

func CollectLower[KIn, VIn, VOut any](seq iter.Seq2[KIn, VIn], f func(KIn, VIn) (VOut, bool)) iter.Seq[VOut] {
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

func DropWhileKeys[K, V any](seq iter.Seq2[K, V], f func(K) bool) iter.Seq2[K, V] {
	return DropWhile2(seq, func(k K, v V) bool { return f(k) })
}

func DropWhileValues[K, V any](seq iter.Seq2[K, V], f func(V) bool) iter.Seq2[K, V] {
	return DropWhile2(seq, func(k K, v V) bool { return f(v) })
}

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

func TakeWhileKeys[K, V any](seq iter.Seq2[K, V], f func(K) bool) iter.Seq2[K, V] {
	return TakeWhile2(seq, func(k K, v V) bool { return f(k) })
}

func TakeWhileValues[K, V any](seq iter.Seq2[K, V], f func(V) bool) iter.Seq2[K, V] {
	return TakeWhile2(seq, func(k K, v V) bool { return f(v) })
}

func Chunk[V any](seq iter.Seq[V], n int) iter.Seq[iter.Seq[V]] {
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
				for i := 0; i < n; i++ {
					v, ok := next()
					if !ok {
						done = true
						return
					}

					if shouldYield {
						if !yield(v) {
							shouldYield = false
						}
					}
				}
			}) || done {
				return
			}
		}
	}
}

func Chunk2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq[iter.Seq2[K, V]] {
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
				for i := 0; i < n; i++ {
					k, v, ok := next()
					if !ok {
						done = true
						return
					}

					if shouldYield {
						if !yield(k, v) {
							shouldYield = false
						}
					}
				}
			}) || done {
				return
			}
		}
	}
}
