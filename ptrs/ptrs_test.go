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

func TestIsNil(t *testing.T) {
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
			if got := IsNil(tt.args.v); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
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
