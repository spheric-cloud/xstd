// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package iters

import (
	"iter"
	"slices"
	"strconv"
	"testing"
)

func TestTap(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	var sum int
	got := slices.Collect(Tap(seq, func(v int) { sum += v }))
	if !slices.Equal(got, s) {
		t.Errorf("Tap() = %v, want %v", got, s)
	}
	if sum != 6 {
		t.Errorf("Tap() sum = %d, want 6", sum)
	}
}

func TestFlatten(t *testing.T) {
	s := [][]int{{1, 2}, {3, 4}}
	seq := Map(slices.Values(s), func(v []int) iter.Seq[int] { return slices.Values(v) })
	got := slices.Collect(Flatten(seq))
	want := []int{1, 2, 3, 4}
	if !slices.Equal(got, want) {
		t.Errorf("Flatten() = %v, want %v", got, want)
	}
}

func TestConcat(t *testing.T) {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	seq1 := slices.Values(s1)
	seq2 := slices.Values(s2)
	got := slices.Collect(Concat(seq1, seq2))
	want := []int{1, 2, 3, 4}
	if !slices.Equal(got, want) {
		t.Errorf("Concat() = %v, want %v", got, want)
	}
}

func TestMap(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	f := func(v int) string { return strconv.Itoa(v) }
	got := slices.Collect(Map(seq, f))
	want := []string{"1", "2", "3"}
	if !slices.Equal(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}

func TestFlatMap(t *testing.T) {
	s := []int{1, 2}
	seq := slices.Values(s)
	f := func(v int) iter.Seq[int] { return Range(0, v) }
	got := slices.Collect(FlatMap(seq, f))
	want := []int{0, 0, 1}
	if !slices.Equal(got, want) {
		t.Errorf("FlatMap() = %v, want %v", got, want)
	}
}

func TestValues(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	got := slices.Collect(Values(seq))
	want := []string{"a", "b"}
	if !slices.Equal(got, want) {
		t.Errorf("Values() = %v, want %v", got, want)
	}
}

func TestKeys(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	got := slices.Collect(Keys(seq))
	want := []int{1, 2}
	if !slices.Equal(got, want) {
		t.Errorf("Keys() = %v, want %v", got, want)
	}
}

func TestSwap(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	got := collect2(Swap(seq))
	want := []struct {
		K string
		V int
	}{{"a", 1}, {"b", 2}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("Swap() = %v, want %v", got, want)
	}
}

func TestJoin(t *testing.T) {
	s1 := []int{1, 2}
	s2 := []string{"a", "b"}
	seq1 := slices.Values(s1)
	seq2 := slices.Values(s2)
	got := collect2(Join(seq1, seq2))
	want := []struct {
		K int
		V string
	}{{1, "a"}, {2, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("Join() = %v, want %v", got, want)
	}
}

func TestEnumerate(t *testing.T) {
	s := []string{"a", "b", "c"}
	seq := slices.Values(s)
	got := collect2(Enumerate[int](seq))
	want := []struct {
		K int
		V string
	}{{0, "a"}, {1, "b"}, {2, "c"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V || got[2].K != want[2].K || got[2].V != want[2].V {
		t.Errorf("Enumerate() = %v, want %v", got, want)
	}
}

func TestRange(t *testing.T) {
	got := slices.Collect(Range(0, 5))
	want := []int{0, 1, 2, 3, 4}
	if !slices.Equal(got, want) {
		t.Errorf("Range() = %v, want %v", got, want)
	}
}

func TestRangeStep(t *testing.T) {
	got := slices.Collect(RangeStep(0, 5, 2))
	want := []int{0, 2, 4}
	if !slices.Equal(got, want) {
		t.Errorf("RangeStep() = %v, want %v", got, want)
	}
	got = slices.Collect(RangeStep(5, 0, -2))
	want = []int{5, 3, 1}
	if !slices.Equal(got, want) {
		t.Errorf("RangeStep() = %v, want %v", got, want)
	}
}

func TestSlice(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	seq := slices.Values(s)
	got := slices.Collect(Slice[int](seq, 2, 4))
	want := []int{2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("Slice() = %v, want %v", got, want)
	}
}

func TestTap2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	var sumK int
	var sumV string
	got := collect2(Tap2(seq, func(k int, v string) {
		sumK += k
		sumV += v
	}))
	want := []struct {
		K int
		V string
	}{{1, "a"}, {2, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("Tap2() = %v, want %v", got, want)
	}
	if sumK != 3 || sumV != "ab" {
		t.Errorf("Tap2() sumK = %d, sumV = %s, want 3, ab", sumK, sumV)
	}
}

func TestTapKey(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	var sumK int
	collect2(TapKey(seq, func(k int) {
		sumK += k
	}))
	if sumK != 3 {
		t.Errorf("TapKey() sumK = %d, want 3", sumK)
	}
}

func TestTapValue(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	var sumV string
	collect2(TapValue(seq, func(v string) {
		sumV += v
	}))
	if sumV != "ab" {
		t.Errorf("TapValue() sumV = %s, want ab", sumV)
	}
}

func TestFlatten2(t *testing.T) {
	seq := Concat(
		Singleton(Singleton2(1, "a")),
		Singleton(Singleton2(2, "b")),
	)
	got := collect2(Flatten2(seq))
	want := []struct {
		K int
		V string
	}{{1, "a"}, {2, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("Flatten2() = %v, want %v", got, want)
	}
}

func TestConcat2(t *testing.T) {
	seq1 := Singleton2(1, "a")
	seq2 := Singleton2(2, "b")
	got := collect2(Concat2(seq1, seq2))
	want := []struct {
		K int
		V string
	}{{1, "a"}, {2, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("Concat2() = %v, want %v", got, want)
	}
}

func TestMap2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	got := collect2(Map2(seq, func(k int, v string) (int, string) {
		return k * 2, v + v
	}))
	want := []struct {
		K int
		V string
	}{{2, "aa"}, {4, "bb"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("Map2() = %v, want %v", got, want)
	}
}

func TestMapKeys(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	got := collect2(MapKeys(seq, func(k int) int {
		return k * 2
	}))
	want := []struct {
		K int
		V string
	}{{2, "a"}, {4, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("MapKeys() = %v, want %v", got, want)
	}
}

func TestMapValues(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	got := collect2(MapValues(seq, func(v string) string {
		return v + v
	}))
	want := []struct {
		K int
		V string
	}{{1, "aa"}, {2, "bb"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("MapValues() = %v, want %v", got, want)
	}
}

func TestFlatMap2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, 1),
		Singleton2(3, 3),
	)
	got := collect2(FlatMap2(seq, func(k, v int) iter.Seq2[int, int] {
		return Join(Range(k, k+2), Range(v, v+2))
	}))
	want := []KV[int, int]{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	if !slices.Equal(got, want) {
		t.Errorf("FlatMap2() = %v, want %v", got, want)
	}
}

func TestEmpty(t *testing.T) {
	got := slices.Collect(Empty[int]())
	if len(got) != 0 {
		t.Errorf("Empty() should be empty, got %v", got)
	}
}

func TestEmpty2(t *testing.T) {
	got := collect2(Empty2[int, string]())
	if len(got) != 0 {
		t.Errorf("Empty2() should be empty, got %v", got)
	}
}

func TestSingleton(t *testing.T) {
	got := slices.Collect(Singleton(1))
	want := []int{1}
	if !slices.Equal(got, want) {
		t.Errorf("Singleton() = %v, want %v", got, want)
	}
}

func TestSingleton2(t *testing.T) {
	got := collect2(Singleton2(1, "a"))
	want := []struct {
		K int
		V string
	}{{1, "a"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V {
		t.Errorf("Singleton2() = %v, want %v", got, want)
	}
}

func TestSlice2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	got := collect2(Slice2[int](seq, 1, 3))
	want := []struct {
		K int
		V string
	}{{2, "b"}, {3, "c"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("Slice2() = %v, want %v", got, want)
	}
}

func TestFlatMapKeys(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	got := collect2(FlatMapKeys(seq, func(k int) iter.Seq[int] {
		return Range(0, k)
	}))
	want := []struct {
		K int
		V string
	}{{0, "a"}, {0, "b"}, {1, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("FlatMapKeys() = %v, want %v", got, want)
	}
}

func TestFlatMapValues(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	got := collect2(FlatMapValues(seq, func(v string) iter.Seq[string] {
		return Singleton(v + v)
	}))
	want := []struct {
		K int
		V string
	}{{1, "aa"}, {2, "bb"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("FlatMapValues() = %v, want %v", got, want)
	}
}

func TestFlatMapLift(t *testing.T) {
	s := []int{1, 2}
	seq := slices.Values(s)
	got := collect2(FlatMapLift(seq, func(v int) iter.Seq2[int, int] {
		return Singleton2(v, v*2)
	}))
	want := []struct {
		K int
		V int
	}{{1, 2}, {2, 4}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("FlatMapLift() = %v, want %v", got, want)
	}
}

func TestFlatMapLower(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	got := slices.Collect(FlatMapLower(seq, func(k int, v string) iter.Seq[string] {
		return Singleton(v + strconv.Itoa(k))
	}))
	want := []string{"a1", "b2"}
	if !slices.Equal(got, want) {
		t.Errorf("FlatMapLower() = %v, want %v", got, want)
	}
}

func TestMapLift(t *testing.T) {
	s := []int{1, 2}
	seq := slices.Values(s)
	got := collect2(MapLift(seq, func(v int) (int, int) {
		return v, v * 2
	}))
	want := []struct {
		K int
		V int
	}{{1, 2}, {2, 4}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("MapLift() = %v, want %v", got, want)
	}
}

func TestLiftZeroValues(t *testing.T) {
	s := []int{1, 2}
	seq := slices.Values(s)
	got := collect2(LiftZeroValues[string](seq))
	want := []struct {
		K int
		V string
	}{{1, ""}, {2, ""}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("LiftZeroValues() = %v, want %v", got, want)
	}
}

func TestLiftZeroKeys(t *testing.T) {
	s := []string{"a", "b"}
	seq := slices.Values(s)
	got := collect2(LiftZeroKeys[int](seq))
	want := []struct {
		K int
		V string
	}{{0, "a"}, {0, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("LiftZeroKeys() = %v, want %v", got, want)
	}
}

func TestMapLower(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	got := slices.Collect(MapLower(seq, func(k int, v string) string {
		return v + strconv.Itoa(k)
	}))
	want := []string{"a1", "b2"}
	if !slices.Equal(got, want) {
		t.Errorf("MapLower() = %v, want %v", got, want)
	}
}

func TestSplit(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	kSeq, vSeq, stop := Split(seq)
	defer stop()
	keys := slices.Collect(kSeq)
	values := slices.Collect(vSeq)
	wantKeys := []int{1, 2}
	wantValues := []string{"a", "b"}
	if !slices.Equal(keys, wantKeys) {
		t.Errorf("Split() keys = %v, want %v", keys, wantKeys)
	}
	if !slices.Equal(values, wantValues) {
		t.Errorf("Split() values = %v, want %v", values, wantValues)
	}
}
