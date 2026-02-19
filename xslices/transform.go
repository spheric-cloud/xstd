// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package xslices

// AppendFlattened appends all elements from a slice of slices to res.
func AppendFlattened[S ~[]V, V any](res S, ss []S) S {
	for _, s := range ss {
		res = append(res, s...)
	}
	return res
}

// Flatten flattens a slice of slices into a single slice.
func Flatten[S ~[]V, V any](ss []S, opts ...InitSliceOption) S {
	res := initSliceFromOpts[S](opts)
	return AppendFlattened(res, ss)
}

// AppendMapped appends the results of calling f on each element of s to res.
func AppendMapped[SIn ~[]VIn, SOut ~[]VOut, VIn, VOut any](res SOut, s SIn, f func(VIn) VOut) []VOut {
	for _, v := range s {
		res = append(res, f(v))
	}
	return res
}

// Map returns a new slice that is the results of calling f on each value from s.
func Map[S ~[]VIn, VIn, VOut any](s S, f func(VIn) VOut, opts ...InitSliceOption) []VOut {
	res := initSlice[[]VOut](toInitSliceOptions(opts).minCapHint(len(s)))
	return AppendMapped(res, s, f)
}

// AppendMappedRef appends the results of calling f on a reference to each element of s to res.
func AppendMappedRef[SIn ~[]VIn, SOut ~[]VOut, VIn, VOut any](res SOut, s SIn, f func(*VIn) VOut) []VOut {
	for _, v := range s {
		res = append(res, f(&v))
	}
	return res
}

// MapRef returns a new slice containing the results of calling f on a reference to each element of s.
func MapRef[S ~[]VIn, VIn, VOut any](s S, f func(*VIn) VOut, opts ...InitSliceOption) []VOut {
	res := initSlice[[]VOut](toInitSliceOptions(opts).minCapHint(len(s)))
	return AppendMappedRef(res, s, f)
}

// MapRefDeref returns a new slice containing the dereferenced results of calling f on a reference to each element of s.
func MapRefDeref[S ~[]VIn, VIn, VOut any](s S, f func(*VIn) *VOut) []VOut {
	res := make([]VOut, 0, len(s))
	for _, v := range s {
		res = append(res, *f(&v))
	}
	return res
}

// MapErr returns a new slice containing the results of calling f on each element of s, stopping at the first error.
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

// MapRefErr returns a new slice containing the results of calling f on a reference to each element,
// stopping at the first error.
func MapRefErr[S ~[]VIn, VIn, VOut any](s S, f func(*VIn) (VOut, error)) ([]VOut, error) {
	var res []VOut
	for _, vIn := range s {
		vOut, err := f(&vIn)
		if err != nil {
			return res, err
		}

		res = append(res, vOut)
	}
	return res, nil
}

// MapRefDerefErr returns a new slice containing the dereferenced results of calling f on a reference
// to each element, stopping at the first error.
func MapRefDerefErr[S ~[]VIn, VIn, VOut any](s S, f func(*VIn) (*VOut, error)) ([]VOut, error) {
	var res []VOut
	for _, vIn := range s {
		vOut, err := f(&vIn)
		if err != nil {
			return res, err
		}

		res = append(res, *vOut)
	}
	return res, nil
}

// MapOKErr returns a new slice containing the results of calling f on each element,
// skipping values for which f returns false, and stopping at the first error.
func MapOKErr[S ~[]VIn, VIn, VOut any](s S, f func(VIn) (VOut, bool, error)) ([]VOut, error) {
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

// MapRefOKErr returns a new slice containing the results of calling f on a reference to each element,
// skipping values for which f returns false, and stopping at the first error.
func MapRefOKErr[S ~[]VIn, VIn, VOut any](s S, f func(*VIn) (VOut, bool, error)) ([]VOut, error) {
	var res []VOut
	for _, vIn := range s {
		vOut, ok, err := f(&vIn)
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

// MapRefDerefOKErr returns a new slice containing the dereferenced results of calling f on a reference
// to each element, skipping values for which f returns false, and stopping at the first error.
func MapRefDerefOKErr[S ~[]VIn, VIn, VOut any](s S, f func(*VIn) (*VOut, bool, error)) ([]VOut, error) {
	var res []VOut
	for _, vIn := range s {
		vOut, ok, err := f(&vIn)
		if err != nil {
			return res, err
		}
		if !ok {
			continue
		}

		res = append(res, *vOut)
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
