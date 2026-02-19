// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package xslices

import (
	"iter"
	"math/rand/v2"
	"slices"
)

// Random returns a random element from the slice using `math/rand/v2`.
func Random[Slice ~[]V, V any](slice Slice) V {
	idx := rand.IntN(len(slice))
	return slice[idx]
}

// AppendSeq2 appends the elements of a sequence of pairs to two slices,
// one for keys and one for values.
func AppendSeq2[KSlice ~[]K, VSlice ~[]V, K, V any](kSlice KSlice, vSlice VSlice, seq iter.Seq2[K, V]) (KSlice, VSlice) {
	for k, v := range seq {
		kSlice = append(kSlice, k)
		vSlice = append(vSlice, v)
	}
	return kSlice, vSlice
}

// Collect collects the elements of a sequence into a new slice.
func Collect[V any](seq iter.Seq[V], opts ...InitSliceOption) []V {
	slice := initSliceFromOpts[[]V](opts)
	return slices.AppendSeq(slice, seq)
}

// Collect2 collects the elements of a sequence of pairs into two new slices,
// one for keys and one for values.
func Collect2[K, V any](seq iter.Seq2[K, V], opts ...InitSliceOption) ([]K, []V) {
	o := toInitSliceOptions(opts)
	var (
		kSlice = initSlice[[]K](o)
		vSlice = initSlice[[]V](o)
	)
	return AppendSeq2(kSlice, vSlice, seq)
}

// TryAppendSeq appends the elements of a sequence to a slice, stopping at the first error.
func TryAppendSeq[Slice ~[]V, V any](s Slice, seq iter.Seq2[V, error]) (Slice, error) {
	for v, err := range seq {
		if err != nil {
			return s, err
		}
		s = append(s, v)
	}
	return s, nil
}

// TryCollect collects the elements of a sequence into a new slice, stopping at the first error.
func TryCollect[V any](seq iter.Seq2[V, error], opts ...InitSliceOption) ([]V, error) {
	res := initSliceFromOpts[[]V](opts)
	return TryAppendSeq(res, seq)
}

// CopySeq copies elements from a sequence to a slice, returning the number of elements copied.
func CopySeq[Slice ~[]V, V any](dst Slice, src iter.Seq[V]) int {
	if len(dst) == 0 {
		return 0
	}
	var i int
	for v := range src {
		dst[i] = v
		i++
		if i >= len(dst) {
			break
		}
	}
	return i
}

// TryCopySeq copies elements from a sequence to a slice, stopping at the first error.
// It returns the number of elements copied and the error.
func TryCopySeq[Slice ~[]V, V any](dst Slice, src iter.Seq2[V, error]) (int, error) {
	if len(dst) == 0 {
		return 0, nil
	}
	var i int
	for v, err := range src {
		if err != nil {
			return i, err
		}

		dst[i] = v
		i++
		if i >= len(dst) {
			break
		}
	}
	return i, nil
}

// RefValues returns a sequence of pointers to the elements of a slice.
func RefValues[Slice ~[]V, V any](s Slice) iter.Seq[*V] {
	return func(yield func(*V) bool) {
		for i := 0; i < len(s); i++ {
			if !yield(&s[i]) {
				return
			}
		}
	}
}

// IndexValues returns a sequence of indexes alongside their values.
func IndexValues[Slice ~[]V, V any](s Slice) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

// EachSetMap applies f to each element in the slice and sets the result as a key-value pair in the map.
func EachSetMap[M ~map[KOut]VOut, S ~[]VIn, VIn any, KOut comparable, VOut any](m M, s S, f func(VIn) (KOut, VOut)) {
	for _, vIn := range s {
		k, vOut := f(vIn)
		m[k] = vOut
	}
}

// ToMap converts a slice to a map by applying f to each element to produce key-value pairs.
func ToMap[Slice ~[]VIn, VIn any, KOut comparable, VOut any](s Slice, f func(VIn) (KOut, VOut), opts ...MakeMapOption) map[KOut]VOut {
	res := makeMapFromOptions[map[KOut]VOut](opts)
	EachSetMap(res, s, f)
	return res
}

// EachSetMapWithKey applies f to each element to produce a key, mapping each key to the original element.
func EachSetMapWithKey[M ~map[K]V, S ~[]V, K comparable, V any](m M, s S, f func(V) K) {
	for _, v := range s {
		k := f(v)
		m[k] = v
	}
}

// ToMapWithKey converts a slice to a map using f to derive the key for each element.
func ToMapWithKey[Slice ~[]V, K comparable, V any](s Slice, f func(V) K, opts ...MakeMapOption) map[K]V {
	res := makeMapFromOptions[map[K]V](opts)
	EachSetMapWithKey(res, s, f)
	return res
}

// EachSetMapWithValue applies f to each element to produce a value, mapping each element to the computed value.
func EachSetMapWithValue[M ~map[K]V, Slice ~[]K, K comparable, V any](m M, s Slice, f func(K) V) {
	for _, k := range s {
		v := f(k)
		m[k] = v
	}
}

// ToMapWithValue converts a slice to a map using f to derive the value for each element.
func ToMapWithValue[Slice ~[]K, K comparable, V any](s Slice, f func(K) V, opts ...MakeMapOption) map[K]V {
	res := makeMapFromOptions[map[K]V](opts)
	EachSetMapWithValue(res, s, f)
	return res
}
