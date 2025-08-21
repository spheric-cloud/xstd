// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package slices

import (
	"iter"
	"math/rand/v2"
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

// Collect2 collects the elements of a sequence of pairs into two new slices,
// one for keys and one for values.
func Collect2[K, V any](seq iter.Seq2[K, V]) ([]K, []V) {
	var (
		kSlice []K
		vSlice []V
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
func TryCollect[V any](seq iter.Seq2[V, error]) ([]V, error) {
	var res []V
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

// PtrValues returns a sequence of pointers to the elements of a slice.
func PtrValues[Slice ~[]V, V any](s Slice) iter.Seq[*V] {
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
