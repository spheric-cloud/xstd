// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package xstrings

import (
	"slices"
	"strings"
	"testing"

	"spheric.cloud/xstd/iters"
)

func TestJoinSeq(t *testing.T) {
	actual := JoinSeq(slices.Values([]string{"a", "b", "c"}), ",")
	want := "a,b,c"
	if actual != want {
		t.Errorf("JoinSeq(...) = %s; want %s", actual, want)
	}
}

func TestWriteJoining(t *testing.T) {
	var sb strings.Builder
	WriteJoining(&sb, []string{"a", "b", "c"}, ",")
	if sb.String() != "a,b,c" {
		t.Errorf("WriteJoining failed, got %s, want a,b,c", sb.String())
	}
}

func TestWriteJoiningSeq(t *testing.T) {
	var sb strings.Builder
	WriteJoiningSeq(&sb, slices.Values([]string{"a", "b", "c"}), ",")
	if sb.String() != "a,b,c" {
		t.Errorf("WriteJoiningSeq failed, got %s, want a,b,c", sb.String())
	}

	// Test with empty seq
	sb.Reset()
	WriteJoiningSeq(&sb, iters.Empty[string](), ",")
	if sb.String() != "" {
		t.Errorf("WriteJoiningSeq with empty seq failed, got %s, want empty string", sb.String())
	}

	// Test with one element
	sb.Reset()
	WriteJoiningSeq(&sb, slices.Values([]string{"a"}), ",")
	if sb.String() != "a" {
		t.Errorf("WriteJoiningSeq with one element failed, got %s, want a", sb.String())
	}
}

func TestRandom(t *testing.T) {
	charset := []rune("abc")
	s := Random(10, charset)
	if len(s) != 10 {
		t.Errorf("Random length = %d, want 10", len(s))
	}
	for _, r := range s {
		if r != 'a' && r != 'b' && r != 'c' {
			t.Errorf("Random contains unexpected rune %c", r)
		}
	}

	// Test zero-length
	s = Random(0, charset)
	if s != "" {
		t.Errorf("Random(0) = %q, want empty string", s)
	}
}

func TestRandom_NegativePanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Random with negative n should panic")
		}
	}()
	Random(-1, []rune("abc"))
}

func TestRandom_EmptyCharsetPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Random with empty charset and n > 0 should panic")
		}
	}()
	Random(5, nil)
}
