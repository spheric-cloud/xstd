// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package xslices

import "slices"

// AppendUnique appends only unique values from s to res.
func AppendUnique[S ~[]V, V comparable](res, s S) S {
	var (
		seen = make(map[V]struct{})
	)
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}

		seen[v] = struct{}{}
		res = append(res, v)
	}
	return res
}

// Unique returns a new slice with all duplicate elements removed.
func Unique[S ~[]V, V comparable](s S, opts ...InitSliceOption) S {
	res := initSliceFromOpts[S](opts)
	return AppendUnique(res, s)
}

// AppendUniqueByKey appends values from s to res, using the given key function for uniqueness checks.
func AppendUniqueByKey[S ~[]V, V any, K comparable](res, s S, f func(V) K) S {
	var (
		seen = make(map[K]struct{})
	)
	for _, v := range s {
		key := f(v)
		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		res = append(res, v)
	}
	return res
}

// UniqueByKey returns a new slice that contains only the unique values from s,
// using the given function to generate the key for uniqueness checks.
func UniqueByKey[S ~[]V, V any, Key comparable](s S, f func(V) Key, opts ...InitSliceOption) S {
	res := initSliceFromOpts[S](opts)
	return AppendUniqueByKey(res, s, f)
}

// AppendFiltered appends values from s to res for which f returns true.
func AppendFiltered[S ~[]V, V any](res, s S, f func(V) bool) S {
	for _, v := range s {
		if !f(v) {
			continue
		}

		res = append(res, v)
	}
	return res
}

// Filter returns a new slice that contains only the values from s for which f returns true.
func Filter[S ~[]V, V any](s S, f func(V) bool, opts ...InitSliceOption) S {
	res := initSliceFromOpts[S](opts)
	return AppendFiltered(res, s, f)
}

// AppendFilteredNonNil appends non-nil pointer values from s to res.
func AppendFilteredNonNil[S ~[]*V, V any](res, s S) S {
	for _, v := range s {
		if v == nil {
			continue
		}

		res = append(res, v)
	}
	return res
}

// FilterNonNil returns a new slice that contains only the non-nil values from s.
func FilterNonNil[S ~[]*V, V any](s S, opts ...InitSliceOption) S {
	res := initSliceFromOpts[S, *V](opts)
	return AppendFilteredNonNil(res, s)
}

// AppendFilteredNonNilDeref returns a new slice that returns only the non-nil values from s, dereferenced.
func AppendFilteredNonNilDeref[S ~[]V, PS ~[]*V, V any](res S, s PS) S {
	for _, v := range s {
		if v == nil {
			continue
		}

		res = append(res, *v)
	}
	return res
}

// FilterNonNilDeref returns a new slice that returns only the non-nil values from s, dereferenced.
func FilterNonNilDeref[S ~[]*V, V any](s S, opts ...InitSliceOption) []V {
	res := initSliceFromOpts[[]V](opts)
	return AppendFilteredNonNilDeref(res, s)
}

// AppendMappedOK appends transformed values from s to res, skipping values for which f returns false.
func AppendMappedOK[SOut ~[]VOut, SIn ~[]VIn, VIn, VOut any](res SOut, s SIn, f func(VIn) (VOut, bool)) SOut {
	for _, vIn := range s {
		vOut, ok := f(vIn)
		if !ok {
			continue
		}

		res = append(res, vOut)
	}
	return res
}

// MapOK returns a new slice that contains the results of calling f on each value from seq.
// If f returns false, the value is skipped.
func MapOK[S ~[]VIn, VIn, VOut any](s S, f func(VIn) (VOut, bool), opts ...InitSliceOption) []VOut {
	res := initSliceFromOpts[[]VOut](opts)
	return AppendMappedOK(res, s, f)
}

func divCeil(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	div := a / b
	rem := a % b
	if rem != 0 {
		// only add 1 if signs match
		if (a > 0 && b > 0) || (a < 0 && b < 0) {
			div++
		}
	}
	return div
}

// AppendChunked appends chunked slices of s (each of size n) to res.
func AppendChunked[S ~[]V, V any](res []S, s S, n int) []S {
	if n <= 0 {
		panic("iters.AppendChunked: n must be > 0")
	}
	for chunk := range slices.Chunk(s, n) {
		res = append(res, slices.Clone(chunk))
	}
	return res
}

// Chunks returns chunk slices of size n.
// The last chunk may be smaller than n.
func Chunks[S ~[]V, V any](s S, n int, opts ...InitSliceOption) []S {
	res := initSlice[[]S](toInitSliceOptions(opts).minCapHint(divCeil(len(s), n)))
	return AppendChunked(res, s, n)
}
