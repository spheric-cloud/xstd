// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package iters

import (
	"cmp"
	"iter"

	"spheric.cloud/xstd/constraints"
)

// TryTransform is a helper to implement iterators that transform a sequence of values that can fail.
func TryTransform[VIn, VOut any](seq iter.Seq2[VIn, error], f func(iter.Seq[VIn]) iter.Seq[VOut]) iter.Seq2[VOut, error] {
	return func(yield func(VOut, error) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()

		var (
			err error
			ok  bool
		)
		for vOut := range f(func(yieldInner func(VIn) bool) {
			for {
				var vIn VIn
				vIn, err, ok = next()
				if !ok {
					return
				}
				if err != nil {
					var zero VOut
					ok = yield(zero, err)
					if !ok {
						return
					}
					continue
				}
				if !yieldInner(vIn) {
					return
				}
			}
		}) {
			if !yield(vOut, nil) {
				return
			}
		}
		if !ok {
			return
		}
		for {
			_, err, ok = next()
			if !ok {
				return
			}
			var zero VOut
			if !yield(zero, err) {
				return
			}
		}
	}
}

// TryTransformErr is a helper to implement iterators that transform a sequence of values that can fail.
func TryTransformErr[VIn, VOut any](seq iter.Seq2[VIn, error], f func(iter.Seq[VIn]) iter.Seq2[VOut, error]) iter.Seq2[VOut, error] {
	return func(yield func(VOut, error) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()

		var (
			err error
			ok  bool
		)
		for vOut, err := range f(func(yieldInner func(VIn) bool) {
			for {
				var vIn VIn
				vIn, err, ok = next()
				if !ok {
					return
				}
				if err != nil {
					var zero VOut
					ok = yield(zero, err)
					if !ok {
						return
					}
					continue
				}
				if !yieldInner(vIn) {
					return
				}
			}
		}) {
			if !yield(vOut, err) {
				return
			}
		}
		if !ok {
			return
		}
		for {
			_, err, ok = next()
			if !ok {
				return
			}
			var zero VOut
			if !yield(zero, err) {
				return
			}
		}
	}
}

// LiftSuccess lifts a sequence of values to a sequence of values that can fail, with no errors.
func LiftSuccess[V any](seq iter.Seq[V]) iter.Seq2[V, error] {
	return LiftZeroValues[error](seq)
}

// LiftFailure lifts a sequence of errors to a sequence of values that can fail, with zero values.
func LiftFailure[V any](seq iter.Seq[error]) iter.Seq2[V, error] {
	return LiftZeroKeys[V](seq)
}

// TryAll returns true if f returns true for all values in seq.
// If seq contains an error, it returns false and the error.
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

// TryAllErr returns true if f returns true for all values in seq.
// If seq contains an error, or f returns an error, it returns false and the error.
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

// TryAny returns true if f returns true for any value in seq.
// If seq contains an error, it returns false and the error.
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

// TryAnyErr returns true if f returns true for any value in seq.
// If seq contains an error, or f returns an error, it returns false and the error.
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

// TryContains returns true if seq contains needle.
// If seq contains an error, it returns false and the error.
func TryContains[V comparable](seq iter.Seq2[V, error], needle V) (bool, error) {
	return TryAny(seq, func(v V) bool { return v == needle })
}

// TryFind returns the first value in seq for which f returns true.
// If seq contains an error, it returns the zero value, false, and the error.
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

// TryFindErr returns the first value in seq for which f returns true.
// If seq contains an error, or f returns an error, it returns the zero value, false, and the error.
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

// TryForEach calls f for each value in seq.
// If seq contains an error, it returns the error.
func TryForEach[V any](seq iter.Seq2[V, error], f func(V)) error {
	for v, err := range seq {
		if err != nil {
			return err
		}
		f(v)
	}
	return nil
}

// TryForEachErr calls f for each value in seq.
// If seq contains an error, or f returns an error, it returns the error.
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

// TryDrain consumes all values in seq.
// If seq contains an error, it returns the error.
func TryDrain[V any](seq iter.Seq2[V, error]) error {
	for _, err := range seq {
		if err != nil {
			return err
		}
	}
	return nil
}

// TryMapOKErr returns a new iterator that yields the results of calling f on each value from seq.
// If f returns false, the value is skipped.
// If seq contains an error, or f returns an error, it yields the error.
func TryMapOKErr[VIn, VOut any](seq iter.Seq2[VIn, error], f func(VIn) (VOut, error, bool)) iter.Seq2[VOut, error] {
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

// TryFlatten flattens a sequence of sequences into a single sequence.
// If seq contains an error, it yields the error.
func TryFlatten[V any](seq iter.Seq2[iter.Seq[V], error]) iter.Seq2[V, error] {
	return func(yield func(V, error) bool) {
		for seq, err := range seq {
			if err != nil {
				var zero V
				if !yield(zero, err) {
					return
				}
				continue
			}

			for v := range seq {
				if !yield(v, nil) {
					return
				}
			}
		}
	}
}

// TryMap returns a new iterator that yields the results of calling f on each value from seq.
// If seq contains an error, it yields the error.
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

// TryMapErr returns a new iterator that yields the results of calling f on each value from seq.
// If seq contains an error, or f returns an error, it yields the error.
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

// TryFlatMap returns a new iterator that yields the results of calling f on each value from seq and flattening the result.
// If seq contains an error, it yields the error.
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

// TryFlatMapErr returns a new iterator that yields the results of calling f on each value from seq and flattening the result.
// If seq contains an error, or f returns an error, it yields the error.
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

// TryFlatMapLiftErr returns a new iterator that yields the results of calling f on each value from seq and flattening the result.
// This function lifts a sequence of values to a sequence of key-value pairs.
// If seq contains an error, or f returns an error, it yields the error.
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

// TryFlatMapErrLiftErr returns a new iterator that yields the results of calling f on each value from seq and flattening the result.
// This function lifts a sequence of values to a sequence of key-value pairs.
// If seq contains an error, or f returns an error, it yields the error.
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

// TryFilter returns a new iterator that yields only the values from seq for which f returns true.
// If seq contains an error, it yields the error.
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

// TryFilterErr returns a new iterator that yields only the values from seq for which f returns true.
// If seq contains an error, or f returns an error, it yields the error.
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

// TryTap calls f for each value in seq, and yields the value.
// If seq contains an error, it yields the error.
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

// TryMax returns the maximum value in seq.
// It panics if seq is empty.
// If seq contains an error, it returns the zero value and the error.
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

// TryMaxFunc returns the maximum value in seq, using the given comparison function.
// It panics if seq is empty.
// If seq contains an error, it returns the zero value and the error.
func TryMaxFunc[V any](seq iter.Seq2[V, error], compare func(V, V) int) (V, error) {
	var (
		best V
		ok   bool
	)
	for v, err := range seq {
		if err != nil {
			var zero V
			return zero, err
		}
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

// TryMin returns the minimum value in seq.
// It panics if seq is empty.
// If seq contains an error, it returns the zero value and the error.
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

// TryMinFunc returns the minimum value in seq, using the given comparison function.
// It panics if seq is empty.
// If seq contains an error, it returns the zero value and the error.
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

// TrySum returns the sum of all values in seq.
// If seq contains an error, it returns the sum of the values seen so far and the error.
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

// TryReduce reduces seq to a single value by calling f with the current sum and the next value.
// If seq contains an error, it returns the current sum and the error.
func TryReduce[Sum, V any](sum Sum, seq iter.Seq2[V, error], f func(Sum, V) Sum) (Sum, error) {
	for v, err := range seq {
		if err != nil {
			return sum, err
		}

		sum = f(sum, v)
	}
	return sum, nil
}

// TryReduceErr reduces seq to a single value by calling f with the current sum and the next value.
// If seq contains an error, or f returns an error, it returns the current sum and the error.
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

// TryCount returns the number of values in seq for which f returns true.
// If seq contains an error, it returns the count of the values seen so far and the error.
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

// TryLen returns the number of values in seq.
// If seq contains an error, it returns the length of the values seen so far and the error.
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

// TryIndex returns the value at idx in seq.
// If seq contains an error, it returns the zero value, false, and the error.
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

// TryIndexValue returns the value at idx in seq, or the zero value if idx is out of bounds.
// If seq contains an error, it returns the zero value and the error.
func TryIndexValue[Int constraints.Integer, V any](seq iter.Seq2[V, error], idx Int) (V, error) {
	v, _, err := TryIndex[Int, V](seq, idx)
	return v, err
}

// SplitError splits a sequence of values that can fail into a sequence of values and an error.
// The error is the first error encountered in the sequence.
func SplitError[V any](seq iter.Seq2[V, error]) (iter.Seq[V], *error) {
	var res error

	return func(yield func(V) bool) {
		for v, err := range seq {
			if err != nil {
				res = err
				return
			}

			if !yield(v) {
				return
			}
		}
	}, &res
}
