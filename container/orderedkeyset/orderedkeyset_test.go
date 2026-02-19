// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package orderedkeyset

import (
	"testing"
)

func identity[V any](v V) V { return v }

func TestSet(t *testing.T) {
	s := New[string](identity[string])
	if s.Has("a") {
		t.Error("expected set to not have 'a'")
	}

	s.Insert("a")
	if !s.Has("a") {
		t.Error("expected set to have 'a'")
	}

	s.Insert("b", "c")
	if !s.Has("b") || !s.Has("c") {
		t.Error("expected set to have 'b' and 'c'")
	}

	s.Delete("a")
	if s.Has("a") {
		t.Error("expected set to not have 'a' after deletion")
	}

	s.Delete("b", "c")
	if s.Has("b") || s.Has("c") {
		t.Error("expected set to not have 'b' or 'c' after deletion")
	}
}

func TestNew(t *testing.T) {
	s := New(identity[string], "a", "b", "c")
	if !s.Has("a") || !s.Has("b") || !s.Has("c") {
		t.Error("expected set to have 'a', 'b', and 'c'")
	}
	if s.Len() != 3 {
		t.Errorf("expected set to have length 3, got %d", s.Len())
	}
}
