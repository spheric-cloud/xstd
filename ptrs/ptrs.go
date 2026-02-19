// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

// Package ptrs provides utility functions for working with pointers.
package ptrs

import (
	"iter"
	"reflect"
	"slices"
)

// To returns a pointer to the given value.
func To[V any](v V) *V {
	return new(v)
}

// Deref dereferences the given pointer.
// It panics if the pointer is nil.
func Deref[V any](v *V) V {
	return *v
}

// DerefOr returns the dereferenced pointer value or the default value if the pointer is nil.
func DerefOr[V any](v *V, defaultValue V) V {
	if v != nil {
		return *v
	}
	return defaultValue
}

// DerefOrElse returns the dereferenced pointer value or the result of calling orElse if the pointer is nil.
func DerefOrElse[V any](v *V, orElse func() V) V {
	if v != nil {
		return *v
	}
	return orElse()
}

// DerefOrZero returns the dereferenced pointer value or the zero value if the pointer is nil.
func DerefOrZero[V any](v *V) V {
	if v != nil {
		return *v
	}
	var zero V
	return zero
}

// Map calls the given function with the dereferenced value of v if it's non-nil.
func Map[VIn, VOut any](v *VIn, f func(VIn) VOut) *VOut {
	if v == nil {
		return nil
	}
	res := f(*v)
	return &res
}

// MapRef calls the given function with the value of v if it's non-nil.
func MapRef[VIn, VOut any](v *VIn, f func(*VIn) VOut) *VOut {
	if v == nil {
		return nil
	}
	return new(f(v))
}

// FlatMap calls the given function with the value of v if it's non-nil, dereferencing the result.
func FlatMap[VIn, VOut any](v *VIn, f func(*VIn) *VOut) *VOut {
	if v == nil {
		return nil
	}
	return f(v)
}

// Flatten dereferences a double pointer. Returns nil if the outer pointer is nil.
func Flatten[V any](v **V) *V {
	if v == nil {
		return nil
	}
	return *v
}

// IsNilAny returns true if the given value is a nil pointer.
// It returns false for non-pointer types.
func IsNilAny(v any) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map,
		reflect.Pointer, reflect.UnsafePointer,
		reflect.Interface, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

// IsNil returns true if the given pointer is nil.
func IsNil[V any](v *V) bool {
	return v == nil
}

// CoalesceSeq returns the first non-nil pointer in the given sequence.
func CoalesceSeq[V any](seq iter.Seq[*V]) *V {
	for v := range seq {
		if v != nil {
			return v
		}
	}
	return nil
}

// Coalesce returns the first non-nil pointer in the given slice.
func Coalesce[V any](vs ...*V) *V {
	return CoalesceSeq(slices.Values(vs))
}

// Equal returns true if the two pointers are equal.
// Two pointers are considered equal if they are both nil, or if they both point to equal values.
func Equal[V comparable](a, b *V) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

// EqualFunc returns true if the two pointers are equal using the given equality function.
// Two pointers are considered equal if they are both nil, or if they both point to equal values.
func EqualFunc[V1, V2 any](a *V1, b *V2, eq func(V1, V2) bool) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return eq(*a, *b)
}
