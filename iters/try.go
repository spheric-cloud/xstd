package iters

import (
	"cmp"
	"iter"
	"xstd/constraints"
)

func TryAll[V any](seq iter.Seq2[V, error], f func(V) bool) (bool, error) {
	for v, err := range seq {
		if err != nil {
			return false, err
		}
		if !f(v) {
			return false, nil
		}
	}
	return true, nil
}

func TryAllErr[V any](seq iter.Seq2[V, error], f func(V) (bool, error)) (bool, error) {
	for v, err := range seq {
		if err != nil {
			return false, err
		}
		ok, err := f(v)
		if err != nil || !ok {
			return false, err
		}
	}
	return true, nil
}

func TryAny[V any](seq iter.Seq2[V, error], f func(V) bool) (bool, error) {
	for v, err := range seq {
		if err != nil {
			return false, err
		}
		if f(v) {
			return true, nil
		}
	}
	return false, nil
}

func TryAnyErr[V any](seq iter.Seq2[V, error], f func(V) (bool, error)) (bool, error) {
	for v, err := range seq {
		if err != nil {
			return false, err
		}
		if ok, err := f(v); err != nil || ok {
			return ok, err
		}
	}
	return false, nil
}

func TryContains[V comparable](seq iter.Seq2[V, error], needle V) (bool, error) {
	return TryAny(seq, func(v V) bool { return v == needle })
}

func TryFind[V any](seq iter.Seq2[V, error], f func(V) bool) (V, bool, error) {
	for v, err := range seq {
		if err != nil {
			var zero V
			return zero, false, err
		}
		if f(v) {
			return v, true, nil
		}
	}
	var zero V
	return zero, false, nil
}

func TryFindErr[V any](seq iter.Seq2[V, error], f func(V) (bool, error)) (V, bool, error) {
	for v, err := range seq {
		if err != nil {
			var zero V
			return zero, false, err
		}
		if ok, err := f(v); err != nil || ok {
			return v, ok, err
		}
	}
	var zero V
	return zero, false, nil
}

func TryForEach[V any](seq iter.Seq2[V, error], f func(V)) error {
	for v, err := range seq {
		if err != nil {
			return err
		}
		f(v)
	}
	return nil
}

func TryForEachErr[V any](seq iter.Seq2[V, error], f func(V) error) error {
	for v, err := range seq {
		if err != nil {
			return err
		}
		if err := f(v); err != nil {
			return err
		}
	}
	return nil
}

func TryDrain[V any](seq iter.Seq2[V, error]) error {
	for _, err := range seq {
		if err != nil {
			return err
		}
	}
	return nil
}

func TryCollectErr[VIn, VOut any](seq iter.Seq2[VIn, error], f func(VIn) (VOut, error, bool)) iter.Seq2[VOut, error] {
	return func(yield func(VOut, error) bool) {
		for vIn, err := range seq {
			if err != nil {
				var vOutZero VOut
				if !yield(vOutZero, err) {
					return
				}
				continue
			}

			vOut, err, ok := f(vIn)
			if !ok {
				continue
			}
			if !yield(vOut, err) {
				return
			}
		}
	}
}

func TryMap[VIn, VOut any](seq iter.Seq2[VIn, error], f func(VIn) VOut) iter.Seq2[VOut, error] {
	return func(yield func(VOut, error) bool) {
		for vIn, err := range seq {
			if err != nil {
				var vOutZero VOut
				if !yield(vOutZero, err) {
					return
				}
				continue
			}

			if !yield(f(vIn), nil) {
				return
			}
		}
	}
}

func TryMapErr[VIn, VOut any](seq iter.Seq2[VIn, error], f func(VIn) (VOut, error)) iter.Seq2[VOut, error] {
	return func(yield func(VOut, error) bool) {
		for vIn, err := range seq {
			if err != nil {
				var vOutZero VOut
				if !yield(vOutZero, err) {
					return
				}
				continue
			}

			if !yield(f(vIn)) {
				return
			}
		}
	}
}

func TryFlatMap[VIn, VOut any](seq iter.Seq2[VIn, error], f func(VIn) iter.Seq[VOut]) iter.Seq2[VOut, error] {
	return func(yield func(VOut, error) bool) {
		for vIn, err := range seq {
			if err != nil {
				var vOutZero VOut
				if !yield(vOutZero, err) {
					return
				}
				continue
			}

			for vOut := range f(vIn) {
				if !yield(vOut, nil) {
					return
				}
			}
		}
	}
}

func TryFlatMapErr[VIn, VOut any](seq iter.Seq2[VIn, error], f func(VIn) (iter.Seq[VOut], error)) iter.Seq2[VOut, error] {
	return func(yield func(VOut, error) bool) {
		for vIn, err := range seq {
			if err != nil {
				var vOutZero VOut
				if !yield(vOutZero, err) {
					return
				}
				continue
			}

			vOuts, err := f(vIn)
			if err != nil {
				var vOutZero VOut
				if !yield(vOutZero, err) {
					return
				}
				continue
			}

			for vOut := range vOuts {
				if !yield(vOut, nil) {
					return
				}
			}
		}
	}
}

func TryFlatMapLiftErr[VIn, VOut any](seq iter.Seq2[VIn, error], f func(VIn) iter.Seq2[VOut, error]) iter.Seq2[VOut, error] {
	return func(yield func(VOut, error) bool) {
		for vIn, err := range seq {
			if err != nil {
				var vOutZero VOut
				if !yield(vOutZero, err) {
					return
				}
				continue
			}

			for vOut, err := range f(vIn) {
				if !yield(vOut, err) {
					return
				}
			}
		}
	}
}

func TryFlatMapErrLiftErr[VIn, VOut any](seq iter.Seq2[VIn, error], f func(VIn) (iter.Seq2[VOut, error], error)) iter.Seq2[VOut, error] {
	return func(yield func(VOut, error) bool) {
		for vIn, err := range seq {
			if err != nil {
				var vOutZero VOut
				if !yield(vOutZero, err) {
					return
				}
				continue
			}

			vOuts, err := f(vIn)
			if err != nil {
				var vOutZero VOut
				if !yield(vOutZero, err) {
					return
				}
				continue
			}

			for vOut, err := range vOuts {
				if !yield(vOut, err) {
					return
				}
			}
		}
	}
}

func TryFilter[V any](seq iter.Seq2[V, error], f func(V) bool) iter.Seq2[V, error] {
	return func(yield func(V, error) bool) {
		for v, err := range seq {
			if err != nil {
				if !yield(v, err) {
					return
				}
				continue
			}

			if !f(v) {
				continue
			}
			if !yield(v, nil) {
				return
			}
		}
	}
}

func TryFilterErr[V any](seq iter.Seq2[V, error], f func(V) (bool, error)) iter.Seq2[V, error] {
	return func(yield func(V, error) bool) {
		for v, err := range seq {
			if err != nil {
				if !yield(v, err) {
					return
				}
				continue
			}

			ok, err := f(v)
			if err != nil {
				var vZero V
				if !yield(vZero, err) {
					return
				}
				continue
			}
			if !ok {
				continue
			}
			if !yield(v, nil) {
				return
			}
		}
	}
}

func TryTap[V any](seq iter.Seq2[V, error], f func(V)) iter.Seq2[V, error] {
	return func(yield func(V, error) bool) {
		for v, err := range seq {
			if err == nil {
				f(v)
			}
			if !yield(v, err) {
				return
			}
		}
	}
}

func TryMax[V cmp.Ordered](seq iter.Seq2[V, error]) (V, error) {
	var (
		best V
		ok   bool
	)
	for v, err := range seq {
		if err != nil {
			var zero V
			return zero, err
		}

		if !ok || v > best {
			best = v
			ok = true
		}
	}
	if !ok {
		panic("iters.TryMax: empty seq")
	}
	return best, nil
}

func TryMaxFunc[V any](seq iter.Seq2[V, error], compare func(V, V) int) (V, error) {
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
		panic("iters.TryMaxFunc: empty seq")
	}
	return best, nil
}

func TryMin[V cmp.Ordered](seq iter.Seq2[V, error]) (V, error) {
	var (
		best V
		ok   bool
	)
	for v, err := range seq {
		if err != nil {
			var zero V
			return zero, err
		}

		if !ok || v < best {
			best = v
			ok = true
		}
	}
	if !ok {
		panic("iters.TryMin: empty seq")
	}
	return best, nil
}

func TryMinFunc[V any](seq iter.Seq2[V, error], compare func(V, V) int) (V, error) {
	var (
		best V
		ok   bool
	)
	for v, err := range seq {
		if err != nil {
			var zero V
			return zero, err
		}

		if !ok || compare(v, best) < 0 {
			best = v
			ok = true
		}
	}
	if !ok {
		panic("iters.TryMinFunc: empty seq")
	}
	return best, nil
}

func TrySum[V cmp.Ordered](seq iter.Seq2[V, error]) (V, error) {
	var sum V
	for v, err := range seq {
		if err != nil {
			return sum, err
		}
		sum += v
	}
	return sum, nil
}

func TryReduce[Sum, V any](sum Sum, seq iter.Seq2[V, error], f func(Sum, V) Sum) (Sum, error) {
	for v, err := range seq {
		if err != nil {
			return sum, err
		}

		sum = f(sum, v)
	}
	return sum, nil
}

func TryReduceErr[Sum, V any](sum Sum, seq iter.Seq2[V, error], f func(Sum, V) (Sum, error)) (Sum, error) {
	for v, err := range seq {
		if err != nil {
			return sum, err
		}

		sum, err = f(sum, v)
		if err != nil {
			return sum, err
		}
	}
	return sum, nil
}

func TryCount[Int constraints.Integer, V any](seq iter.Seq2[V, error], f func(V) bool) (Int, error) {
	var ct Int
	for v, err := range seq {
		if err != nil {
			return ct, err
		}

		if f(v) {
			ct++
		}
	}
	return ct, nil
}

func TryLen[Int constraints.Integer, V any](seq iter.Seq2[V, error]) (Int, error) {
	var ct Int
	for _, err := range seq {
		if err != nil {
			return ct, err
		}
		ct++
	}
	return ct, nil
}

func TryIndex[Int constraints.Integer, V any](seq iter.Seq2[V, error], idx Int) (V, bool, error) {
	var ct Int
	for v, err := range seq {
		if err != nil {
			var zero V
			return zero, false, err
		}

		if ct == idx {
			return v, true, nil
		}
		ct++
	}
	var zero V
	return zero, false, nil
}

func TryIndexValue[Int constraints.Integer, V any](seq iter.Seq2[V, error], idx Int) (V, error) {
	v, _, err := TryIndex[Int, V](seq, idx)
	return v, err
}
