// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package xmaps

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

func TestInverse(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	inv := Inverse(m)
	if inv[1] != "a" || inv[2] != "b" {
		t.Errorf("Inverse got %v", inv)
	}
}

func TestKeysByValue(t *testing.T) {
	m := map[string]int{"a": 1, "b": 1, "c": 2}
	kbv := KeysByValue(m)
	if len(kbv[1]) != 2 {
		t.Errorf("KeysByValue[1] len = %d, want 2", len(kbv[1]))
	}
	if len(kbv[2]) != 1 || kbv[2][0] != "c" {
		t.Errorf("KeysByValue[2] = %v, want [c]", kbv[2])
	}
}

func TestGetAny(t *testing.T) {
	m := map[string]int{"a": 1}
	k, v, ok := GetAny(m)
	if !ok || k != "a" || v != 1 {
		t.Errorf("GetAny got %s, %d, %t, want a, 1, true", k, v, ok)
	}

	mEmpty := map[string]int{}
	_, _, ok = GetAny(mEmpty)
	if ok {
		t.Error("GetAny on empty map should return false")
	}
}

func TestGetAnyValue(t *testing.T) {
	m := map[string]int{"a": 1}
	k, v := GetAnyValue(m)
	if k != "a" || v != 1 {
		t.Errorf("GetAnyValue got %s, %d, want a, 1", k, v)
	}
}
