// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package iters

import (
	"slices"
	"testing"
)

func TestFromNext(t *testing.T) {
	s := []int{1, 2, 3}
	i := 0
	next := func() (int, bool) {
		if i >= len(s) {
			return 0, false
		}
		v := s[i]
		i++
		return v, true
	}
	got := slices.Collect(FromNext(next))
	if !slices.Equal(got, s) {
		t.Errorf("FromNext() = %v, want %v", got, s)
	}
}

func TestFromNext2(t *testing.T) {
	type pair struct {
		K int
		V string
	}
	s := []pair{{K: 1, V: "a"}, {K: 2, V: "b"}, {K: 3, V: "c"}}
	i := 0
	next := func() (int, string, bool) {
		if i >= len(s) {
			return 0, "", false
		}
		v := s[i]
		i++
		return v.K, v.V, true
	}

	got := collect2(FromNext2(next))
	if len(got) != len(s) || got[0].K != s[0].K || got[0].V != s[0].V || got[1].K != s[1].K || got[1].V != s[1].V || got[2].K != s[2].K || got[2].V != s[2].V {
		t.Errorf("FromNext2() = %v, want %v", got, s)
	}
}

func TestRepeat(t *testing.T) {
	got := slices.Collect(Repeat(1, 3))
	want := []int{1, 1, 1}
	if !slices.Equal(got, want) {
		t.Errorf("Repeat() = %v, want %v", got, want)
	}
}

func TestRepeat2(t *testing.T) {
	got := collect2(Repeat2(1, "a", 3))
	want := []struct {
		K int
		V string
	}{{1, "a"}, {1, "a"}, {1, "a"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V || got[2].K != want[2].K || got[2].V != want[2].V {
		t.Errorf("Repeat2() = %v, want %v", got, want)
	}
}

func TestCycle(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	got := slices.Collect(Take(Cycle(seq), 5))
	want := []int{1, 2, 3, 1, 2}
	if !slices.Equal(got, want) {
		t.Errorf("Cycle() = %v, want %v", got, want)
	}
}

func TestCycle2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	got := collect2(Take2(Cycle2(seq), 3))
	want := []struct {
		K int
		V string
	}{{1, "a"}, {2, "b"}, {1, "a"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V || got[2].K != want[2].K || got[2].V != want[2].V {
		t.Errorf("Cycle2() = %v, want %v", got, want)
	}
}
