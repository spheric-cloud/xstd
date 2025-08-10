// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package maps

import (
	"testing"
)

func TestPop(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	k, v, ok := Pop(m)
	if !ok {
		t.Error("Pop failed to pop a value")
	}
	if len(m) != 1 {
		t.Errorf("Pop did not remove the element, got %v, want 1", len(m))
	}
	if (k == "a" && v != 1) || (k == "b" && v != 2) {
		t.Errorf("Pop returned an incorrect value for key %s: got %d", k, v)
	}
	if _, exists := m[k]; exists {
		t.Errorf("Pop did not remove the element with key %s from the map", k)
	}

	// Test empty map
	mEmpty := map[string]int{}
	_, _, ok = Pop(mEmpty)
	if ok {
		t.Error("Pop returned ok for an empty map")
	}
}

func TestPopValue(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	k, v := PopValue(m)
	if len(m) != 1 {
		t.Errorf("PopValue did not remove the element, got %v, want 1", len(m))
	}
	if (k == "a" && v != 1) || (k == "b" && v != 2) {
		t.Errorf("PopValue returned an incorrect value for key %s: got %d", k, v)
	}
	if _, exists := m[k]; exists {
		t.Errorf("PopValue did not remove the element with key %s from the map", k)
	}
}

func TestSingle(t *testing.T) {
	// Test with single element
	mSingle := map[string]int{"a": 1}
	k, v, n := Single(mSingle)
	if n != 1 {
		t.Errorf("Single on a single-element map returned n=%d, want 1", n)
	}
	if k != "a" || v != 1 {
		t.Errorf("Single on a single-element map returned %s, %d, want a, 1", k, v)
	}

	// Test with multiple elements
	mMulti := map[string]int{"a": 1, "b": 2}
	_, _, n = Single(mMulti)
	if n != 2 {
		t.Errorf("Single on a multi-element map returned n=%d, want 2", n)
	}

	// Test with empty map
	mEmpty := map[string]int{}
	_, _, n = Single(mEmpty)
	if n != 0 {
		t.Errorf("Single on an empty map returned n=%d, want 0", n)
	}
}

func TestSingleValue(t *testing.T) {
	// Test with single element
	mSingle := map[string]int{"a": 1}
	k, v := SingleValue(mSingle)
	if k != "a" || v != 1 {
		t.Errorf("SingleValue on a single-element map returned %s, %d, want a, 1", k, v)
	}
}
