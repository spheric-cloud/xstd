// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package slices

import (
	"errors"
	"slices"
	"testing"

	"spheric.cloud/xstd/iters"
)

func TestAppendSeq2(t *testing.T) {
	kSlice := []string{"a"}
	vSlice := []int{1}
	seq := func(yield func(string, int) bool) {
		yield("b", 2)
		yield("c", 3)
	}
	kSlice, vSlice = AppendSeq2(kSlice, vSlice, seq)
	if !slices.Equal(kSlice, []string{"a", "b", "c"}) {
		t.Errorf("AppendSeq2 keys got %v, want %v", kSlice, []string{"a", "b", "c"})
	}
	if !slices.Equal(vSlice, []int{1, 2, 3}) {
		t.Errorf("AppendSeq2 values got %v, want %v", vSlice, []int{1, 2, 3})
	}
}

func TestCollect2(t *testing.T) {
	seq := func(yield func(string, int) bool) {
		yield("a", 1)
		yield("b", 2)
	}
	kSlice, vSlice := Collect2(seq)
	if !slices.Equal(kSlice, []string{"a", "b"}) {
		t.Errorf("Collect2 keys got %v, want %v", kSlice, []string{"a", "b"})
	}
	if !slices.Equal(vSlice, []int{1, 2}) {
		t.Errorf("Collect2 values got %v, want %v", vSlice, []int{1, 2})
	}
}

func TestTryAppendSeq(t *testing.T) {
	s := []int{1}
	seq := func(yield func(int, error) bool) {
		yield(2, nil)
		yield(3, nil)
	}
	s, err := TryAppendSeq(s, seq)
	if err != nil {
		t.Errorf("TryAppendSeq failed with error: %v", err)
	}
	if !slices.Equal(s, []int{1, 2, 3}) {
		t.Errorf("TryAppendSeq got %v, want %v", s, []int{1, 2, 3})
	}

	// Test with error
	s = []int{1}
	seqErr := func(yield func(int, error) bool) {
		yield(2, nil)
		yield(0, errors.New("test error"))
	}
	s, err = TryAppendSeq(s, seqErr)
	if err == nil {
		t.Error("TryAppendSeq should have returned an error")
	}
	if !slices.Equal(s, []int{1, 2}) {
		t.Errorf("TryAppendSeq with error got %v, want %v", s, []int{1, 2})
	}
}

func TestTryCollect(t *testing.T) {
	seq := func(yield func(int, error) bool) {
		yield(1, nil)
		yield(2, nil)
	}
	s, err := TryCollect(seq)
	if err != nil {
		t.Errorf("TryCollect failed with error: %v", err)
	}
	if !slices.Equal(s, []int{1, 2}) {
		t.Errorf("TryCollect got %v, want %v", s, []int{1, 2})
	}

	// Test with error
	seqErr := func(yield func(int, error) bool) {
		yield(1, nil)
		yield(0, errors.New("test error"))
	}
	s, err = TryCollect(seqErr)
	if err == nil {
		t.Error("TryCollect should have returned an error")
	}
	if !slices.Equal(s, []int{1}) {
		t.Errorf("TryCollect with error got %v, want %v", s, []int{1})
	}
}

func TestCopySeq(t *testing.T) {
	dst := make([]int, 2)
	src := slices.Values([]int{1, 2, 3})
	n := CopySeq(dst, src)
	if n != 2 {
		t.Errorf("CopySeq returned %d, want 2", n)
	}
	if !slices.Equal(dst, []int{1, 2}) {
		t.Errorf("CopySeq got %v, want %v", dst, []int{1, 2})
	}

	// Test with smaller source
	dst = make([]int, 3)
	src = slices.Values([]int{1, 2})
	n = CopySeq(dst, src)
	if n != 2 {
		t.Errorf("CopySeq with smaller source returned %d, want 2", n)
	}
	if !slices.Equal(dst, []int{1, 2, 0}) {
		t.Errorf("CopySeq with smaller source got %v, want %v", dst, []int{1, 2, 0})
	}
}

func TestTryCopySeq(t *testing.T) {
	dst := make([]int, 2)
	seq := iters.LiftSuccess(iters.Of(1, 2, 3))
	n, err := TryCopySeq(dst, seq)
	if err != nil {
		t.Errorf("TryCopySeq failed with error: %v", err)
	}
	if n != 2 {
		t.Errorf("TryCopySeq returned %d, want 2", n)
	}
	if !slices.Equal(dst, []int{1, 2}) {
		t.Errorf("TryCopySeq got %v, want %v", dst, []int{1, 2})
	}

	// Test with error
	dst = make([]int, 3)
	seqErr := iters.Concat2(iters.Singleton2[int, error](1, nil), iters.Singleton2[int, error](0, errors.New("test error")))
	n, err = TryCopySeq(dst, seqErr)
	if err == nil {
		t.Error("TryCopySeq should have returned an error")
	}
	if n != 1 {
		t.Errorf("TryCopySeq with error returned %d, want 1", n)
	}
	if !slices.Equal(dst, []int{1, 0, 0}) {
		t.Errorf("TryCopySeq with error got %v, want %v", dst, []int{1, 0, 0})
	}
}

func TestPtrValues(t *testing.T) {
	s := []int{1, 2, 3}
	seq := PtrValues(s)
	var ptrs []*int
	for p := range seq {
		ptrs = append(ptrs, p)
	}
	if len(ptrs) != 3 {
		t.Fatalf("PtrValues returned %d pointers, want 3", len(ptrs))
	}
	for i, p := range ptrs {
		if *p != s[i] {
			t.Errorf("PtrValues pointer %d has value %d, want %d", i, *p, s[i])
		}
	}

	// Modify slice through pointers
	for _, p := range ptrs {
		*p = *p + 1
	}
	if !slices.Equal(s, []int{2, 3, 4}) {
		t.Errorf("Modifying through pointers failed, got %v, want %v", s, []int{2, 3, 4})
	}
}
