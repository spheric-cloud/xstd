// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package sets

import (
	"errors"
	"slices"
	"testing"

	"spheric.cloud/xstd/iters"
	"spheric.cloud/xstd/set"
)

func TestHasAll(t *testing.T) {
	s := set.New("a", "b", "c")
	if !HasAll(s, slices.Values([]string{"a", "b"})) {
		t.Error("expected set to have all values")
	}
	if HasAll(s, slices.Values([]string{"a", "d"})) {
		t.Error("expected set to not have all values")
	}
}

func TestHasAny(t *testing.T) {
	s := set.New("a", "b", "c")
	if !HasAny(s, slices.Values([]string{"a", "d"})) {
		t.Error("expected set to have any value")
	}
	if HasAny(s, slices.Values([]string{"d", "e"})) {
		t.Error("expected set to not have any value")
	}
}

func TestClone(t *testing.T) {
	s1 := set.New("a", "b", "c")
	s2 := Clone(s1)
	if !Equal(s1, s2) {
		t.Error("expected sets to be equal")
	}
	s2.Insert("d")
	if Equal(s1, s2) {
		t.Error("expected sets to not be equal")
	}
}

func TestDifference(t *testing.T) {
	s1 := set.New("a", "b", "c")
	s2 := set.New("b", "c", "d")
	s3 := Difference(s1, s2)
	if !s3.Has("a") || s3.Has("b") || s3.Has("c") || s3.Has("d") {
		t.Error("invalid difference")
	}
}

func TestUnion(t *testing.T) {
	s1 := set.New("a", "b", "c")
	s2 := set.New("b", "c", "d")
	s3 := Union(s1, s2)
	if !s3.Has("a") || !s3.Has("b") || !s3.Has("c") || !s3.Has("d") {
		t.Error("invalid union")
	}
}

func TestEqual(t *testing.T) {
	s1 := set.New("a", "b", "c")
	s2 := set.New("a", "b", "c")
	s3 := set.New("a", "b", "d")
	if !Equal(s1, s2) {
		t.Error("expected sets to be equal")
	}
	if Equal(s1, s3) {
		t.Error("expected sets to not be equal")
	}
}

func TestValues(t *testing.T) {
	s := set.New("a", "b", "c")
	vals := make([]string, 0, 3)
	for v := range Values(s) {
		vals = append(vals, v)
	}
	slices.Sort(vals)
	if !slices.Equal(vals, []string{"a", "b", "c"}) {
		t.Errorf("unexpected values: %v", vals)
	}
}

func TestPop(t *testing.T) {
	s := set.New("a", "b", "c")
	for i := 0; i < 3; i++ {
		v, ok := Pop(s)
		if !ok {
			t.Error("expected ok")
		}
		if s.Has(v) {
			t.Errorf("value %q should have been removed", v)
		}
	}
	if _, ok := Pop(s); ok {
		t.Error("expected not ok")
	}
}

func TestInsert(t *testing.T) {
	s := set.New[string]()
	Insert(s, slices.Values([]string{"a", "b", "c"}))
	if !s.Has("a") || !s.Has("b") || !s.Has("c") {
		t.Error("invalid insertion")
	}
}

func TestTryInsert(t *testing.T) {
	s := set.New[string]()
	err := TryInsert(s, iters.LiftSingletonValue(slices.Values([]string{"a", "b", "c"}), (error)(nil)))
	if err != nil {
		t.Fatal(err)
	}
	if !s.Has("a") || !s.Has("b") || !s.Has("c") {
		t.Error("invalid insertion")
	}

	s = set.New[string]()
	err = TryInsert(s, iters.LiftSingletonValue(slices.Values([]string{"a", "b", "c"}), errors.New("some error")))
	if err == nil {
		t.Fatal("expected an error")
	}
	if s.Len() != 0 {
		t.Error("expected set to be empty")
	}
}

func TestCollect(t *testing.T) {
	s := Collect(slices.Values([]string{"a", "b", "c"}))
	if !s.Has("a") || !s.Has("b") || !s.Has("c") {
		t.Error("invalid collection")
	}
}

func TestTryCollect(t *testing.T) {
	s, err := TryCollect(iters.LiftSuccess(slices.Values([]string{"a", "b", "c"})))
	if err != nil {
		t.Fatal(err)
	}
	if !s.Has("a") || !s.Has("b") || !s.Has("c") {
		t.Error("invalid collection")
	}

	_, err = TryCollect(iters.LiftFailure[string](iters.Singleton(errors.New("some error"))))
	if err == nil {
		t.Fatal("expected an error")
	}
}
