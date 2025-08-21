// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package slices

// Flatten flattens a slice of slices into a single slice.
func Flatten[S ~[]V, V any](ss []S) S {
	var res S
	for _, s := range ss {
		res = append(res, s...)
	}
	return res
}

// Map returns a new slice that is the results of calling f on each value from s.
func Map[S ~[]VIn, VIn, VOut any](s S, f func(VIn) VOut) []VOut {
	res := make([]VOut, 0, len(s))
	for _, v := range s {
		res = append(res, f(v))
	}
	return res
}

func MapErr[S ~[]VIn, VIn, VOut any](s S, f func(VIn) (VOut, error)) ([]VOut, error) {
	var res []VOut
	for _, vIn := range s {
		vOut, err := f(vIn)
		if err != nil {
			return res, err
		}

		res = append(res, vOut)
	}
	return res, nil
}

func MapOkErr[S ~[]VIn, VIn, VOut any](s S, f func(VIn) (VOut, bool, error)) ([]VOut, error) {
	var res []VOut
	for _, vIn := range s {
		vOut, ok, err := f(vIn)
		if err != nil {
			return res, err
		}
		if !ok {
			continue
		}

		res = append(res, vOut)
	}
	return res, nil
}

// FlatMap returns a new slice that is the results of calling f on each value from s and flattening the result.
func FlatMap[S ~[]VIn, VIn, VOut any](s S, f func(VIn) []VOut) []VOut {
	var res []VOut
	for _, v := range s {
		res = append(res, f(v)...)
	}
	return res
}
