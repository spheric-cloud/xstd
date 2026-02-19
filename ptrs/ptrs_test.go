// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package ptrs

import (
	"slices"
	"testing"
)

func TestTo(t *testing.T) {
	v := 1
	p := To(v)
	if *p != v {
		t.Errorf("To() = %v, want %v", *p, v)
	}
}

func TestDeref(t *testing.T) {
	v := 1
	p := &v
	if Deref(p) != v {
		t.Errorf("Deref() = %v, want %v", Deref(p), v)
	}
}

func TestDerefOr(t *testing.T) {
	v := 1
	p := &v
	if DerefOr(p, 2) != 1 {
		t.Errorf("DerefOr() = %v, want %v", DerefOr(p, 2), 1)
	}
	if DerefOr(nil, 2) != 2 {
		t.Errorf("DerefOr() = %v, want %v", DerefOr(nil, 2), 2)
	}
}

func TestDerefOrElse(t *testing.T) {
	v := 1
	p := &v
	if DerefOrElse(p, func() int { return 2 }) != 1 {
		t.Errorf("DerefOrElse() = %v, want %v", DerefOrElse(p, func() int { return 2 }), 1)
	}
	if DerefOrElse(nil, func() int { return 2 }) != 2 {
		t.Errorf("DerefOrElse() = %v, want %v", DerefOrElse(nil, func() int { return 2 }), 2)
	}
}

func TestDerefOrZero(t *testing.T) {
	v := 1
	p := &v
	if DerefOrZero(p) != 1 {
		t.Errorf("DerefOrZero() = %v, want %v", DerefOrZero(p), 1)
	}
	if DerefOrZero[int](nil) != 0 {
		t.Errorf("DerefOrZero() = %v, want %v", DerefOrZero[int](nil), 0)
	}
}

func TestIsNilAny(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nilable value",
			args: args{v: ([]int)(nil)},
			want: true,
		},
		{
			name: "non-nil value",
			args: args{v: []int{}},
			want: false,
		},
		{
			name: "non-nilable value",
			args: args{v: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNilAny(tt.args.v); got != tt.want {
				t.Errorf("IsNilAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoalesceSeq(t *testing.T) {
	v1 := 1
	v2 := 2
	if *CoalesceSeq(slices.Values([]*int{nil, &v1, &v2})) != 1 {
		t.Error("CoalesceSeq returned wrong value")
	}
	if CoalesceSeq(slices.Values([]*int{nil, nil})) != nil {
		t.Error("CoalesceSeq returned non-nil")
	}
}

func TestCoalesce(t *testing.T) {
	v1 := 1
	v2 := 2
	if *Coalesce(nil, &v1, &v2) != 1 {
		t.Error("Coalesce returned wrong value")
	}
	if Coalesce((*int)(nil), nil) != nil {
		t.Error("Coalesce returned non-nil")
	}
}

func TestEqual(t *testing.T) {
	v1 := 1
	v2 := 2
	if !Equal(&v1, &v1) {
		t.Error("expected pointers to be equal")
	}
	if Equal(&v1, &v2) {
		t.Error("expected pointers to be not equal")
	}
	if !Equal[int](nil, nil) {
		t.Error("expected nil pointers to be equal")
	}
	if Equal[int](&v1, nil) {
		t.Error("expected pointers to be not equal")
	}
}

func TestEqualFunc(t *testing.T) {
	v1 := 1
	v2 := 2
	eq := func(a, b int) bool { return a == b }
	if !EqualFunc(&v1, &v1, eq) {
		t.Error("expected pointers to be equal")
	}
	if EqualFunc(&v1, &v2, eq) {
		t.Error("expected pointers to be not equal")
	}
	if !EqualFunc(nil, (*int)(nil), eq) {
		t.Error("expected nil pointers to be equal")
	}
	if EqualFunc(&v1, nil, eq) {
		t.Error("expected pointers to be not equal")
	}
}

func TestMap(t *testing.T) {
	v := 42
	result := Map(&v, func(n int) string { return "val" })
	if result == nil || *result != "val" {
		t.Error("Map on non-nil pointer failed")
	}
	result = Map((*int)(nil), func(n int) string { return "val" })
	if result != nil {
		t.Error("Map on nil pointer should return nil")
	}
}

func TestMapRef(t *testing.T) {
	v := 42
	result := MapRef(&v, func(n *int) string { return "val" })
	if result == nil || *result != "val" {
		t.Error("MapRef on non-nil pointer failed")
	}
	result = MapRef((*int)(nil), func(n *int) string { return "val" })
	if result != nil {
		t.Error("MapRef on nil pointer should return nil")
	}
}

func TestFlatMap(t *testing.T) {
	v := 42
	s := "hello"
	result := FlatMap(&v, func(n *int) *string { return &s })
	if result == nil || *result != "hello" {
		t.Error("FlatMap on non-nil pointer failed")
	}
	result = FlatMap((*int)(nil), func(n *int) *string { return &s })
	if result != nil {
		t.Error("FlatMap on nil pointer should return nil")
	}
}

func TestFlatten(t *testing.T) {
	v := 42
	p := &v
	result := Flatten(&p)
	if result == nil || *result != 42 {
		t.Error("Flatten on non-nil double pointer failed")
	}
	result = Flatten((**int)(nil))
	if result != nil {
		t.Error("Flatten on nil double pointer should return nil")
	}
}

func TestIsNil(t *testing.T) {
	v := 42
	if IsNil(&v) {
		t.Error("IsNil on non-nil pointer should return false")
	}
	if !IsNil[int](nil) {
		t.Error("IsNil on nil pointer should return true")
	}
}
