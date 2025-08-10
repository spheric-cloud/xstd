// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package gen

import (
	"fmt"
	"reflect"
	"strings"
)

// Cast performs a type assertion to the given type.
// It will panic if the type assertion fails.
func Cast[To any](in any) To {
	return any(in).(To)
}

// CastOK performs a type assertion to the given type.
// It returns the asserted value and a boolean indicating whether the assertion succeeded.
func CastOK[To any](in any) (To, bool) {
	vTo, ok := in.(To)
	return vTo, ok
}

// IsA checks if a value is of a certain type.
func IsA[V any](v any) bool {
	_, ok := v.(V)
	return ok
}

// Zero returns the zero value for a given type.
func Zero[V any]() V {
	var zero V
	return zero
}

// IsZero checks if a value is the zero value for its type.
func IsZero[V any](v V) bool {
	rv := reflect.ValueOf(v)
	return rv.IsZero()
}

// New returns a pointer to a new zero value for a given type.
func New[V any]() *V {
	return new(V)
}

// TODO is a function to create holes when stubbing out more complex mechanisms.
//
// By default, it will panic with 'TODO: provide a value of type <type>' where <type> is the type of V.
// The panic message can be altered by passing in additional args that will be printed as
// 'TODO: <args separated by space>'
func TODO[V any](args ...any) V {
	var sb strings.Builder
	sb.WriteString("TODO: ")
	if len(args) > 0 {
		_, _ = fmt.Fprintln(&sb, args...)
	} else {
		_, _ = fmt.Fprintf(&sb, "provide a value of type %T", Zero[V]())
	}
	panic(sb.String())
}

// Stub is a stub left empty by intent.
//
// For instance, to check generic types implement certain interfaces, a blank ('_') function declaration
// can be used to check the type definition.
//
// Example:
//
//	func _[K]() MyInterface[K] {
//		return Stub[MyImplementation[K]]()
//	}
func Stub[V any](args ...any) V {
	var sb strings.Builder
	sb.WriteString("Stub was called - this should not happen: ")
	if len(args) > 0 {
		_, _ = fmt.Fprintln(&sb, args...)
	} else {
		_, _ = fmt.Fprintf(&sb, "provide a value of type %T", Zero[V]())
	}
	panic(sb.String())
}
