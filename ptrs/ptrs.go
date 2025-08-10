// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

// Package ptrs provides utility functions for working with pointers.
package ptrs

import (
	"iter"
	"reflect"

	"spheric.cloud/xstd/slices"
)

// To returns a pointer to the given value.
func To[V any](v V) *V {
	return &v
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

// IsNil returns true if the given value is a nil pointer.
// It returns false for non-pointer types.
func IsNil(v any) bool {
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
