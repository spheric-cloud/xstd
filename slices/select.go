// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package slices

// Unique returns a new slice with all duplicate elements removed.
func Unique[S ~[]V, V comparable](s S) S {
	var (
		res  S
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

// UniqueFunc returns a new slice that contains only the unique values from s,
// using the given function to generate the key for uniqueness checks.
func UniqueFunc[S ~[]V, V any, Key comparable](s S, f func(V) Key) S {
	var (
		res  S
		seen = make(map[Key]struct{})
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

// Filter returns a new slice that contains only the values from s for which f returns true.
func Filter[S ~[]V, V any](s S, f func(V) bool) S {
	var res S
	for _, v := range s {
		if !f(v) {
			continue
		}

		res = append(res, v)
	}
	return res
}

// MapOK returns a new slice that contains the results of calling f on each value from seq.
// If f returns false, the value is skipped.
func MapOK[S ~[]VIn, VIn, VOut any](s S, f func(VIn) (VOut, bool)) []VOut {
	var res []VOut
	for _, vIn := range s {
		vOut, ok := f(vIn)
		if !ok {
			continue
		}

		res = append(res, vOut)
	}
	return res
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

// Chunked returns chunk slices of size n.
// The last chunk may be smaller than n.
func Chunked[S ~[]V, V any](s S, n int) []S {
	if n <= 0 {
		panic("iters.Chunk: n must be > 0")
	}

	var (
		chunks = make([]S, divCeil(len(s), n))
		i      int
	)
	for chunk := range Chunk(s, n) {
		copy(chunks[i], chunk)
		i += n
	}
	return chunks
}
