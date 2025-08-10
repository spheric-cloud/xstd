// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package iters

import (
	"slices"
	"strconv"
	"testing"
)

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	seq := slices.Values(s)
	f := func(v int) bool { return v%2 == 0 }
	got := slices.Collect(Filter(seq, f))
	want := []int{2, 4}
	if !slices.Equal(got, want) {
		t.Errorf("Filter() = %v, want %v", got, want)
	}
}

func TestFilter2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(k int, v string) bool { return k%2 == 0 }
	got := collect2(Filter2(seq, f))
	want := []struct {
		K int
		V string
	}{{2, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V {
		t.Errorf("Filter2() = %v, want %v", got, want)
	}
}

func TestFilterKeys(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(k int) bool { return k%2 == 0 }
	got := collect2(FilterKeys(seq, f))
	want := []struct {
		K int
		V string
	}{{2, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V {
		t.Errorf("FilterKeys() = %v, want %v", got, want)
	}
}

func TestFilterValues(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(v string) bool { return v == "b" }
	got := collect2(FilterValues(seq, f))
	want := []struct {
		K int
		V string
	}{{2, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V {
		t.Errorf("FilterValues() = %v, want %v", got, want)
	}
}

func TestCollect(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	seq := slices.Values(s)
	f := func(v int) (string, bool) {
		if v%2 == 0 {
			return strconv.Itoa(v), true
		}
		return "", false
	}
	got := slices.Collect(MapOK(seq, f))
	want := []string{"2", "4"}
	if !slices.Equal(got, want) {
		t.Errorf("MapOK() = %v, want %v", got, want)
	}
}

func TestCollect2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(k int, v string) (string, int, bool) {
		if k%2 == 0 {
			return v, k, true
		}
		return "", 0, false
	}
	got := collect2(MapOK2(seq, f))
	want := []struct {
		K string
		V int
	}{{"b", 2}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V {
		t.Errorf("MapOK2() = %v, want %v", got, want)
	}
}

func TestCollectKeys(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(k int) (string, bool) {
		if k%2 == 0 {
			return strconv.Itoa(k), true
		}
		return "", false
	}
	got := collect2(MapOKKeys(seq, f))
	want := []struct {
		K string
		V string
	}{{"2", "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V {
		t.Errorf("MapOKKeys() = %v, want %v", got, want)
	}
}

func TestCollectValues(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(v string) (string, bool) {
		if v == "b" {
			return "B", true
		}
		return "", false
	}
	got := collect2(MapOKValues(seq, f))
	want := []struct {
		K int
		V string
	}{{2, "B"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V {
		t.Errorf("MapOKValues() = %v, want %v", got, want)
	}
}

func TestCollectLift(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	seq := slices.Values(s)
	f := func(v int) (int, string, bool) {
		if v%2 == 0 {
			return v, strconv.Itoa(v), true
		}
		return 0, "", false
	}
	got := collect2(MapOKLift(seq, f))
	want := []struct {
		K int
		V string
	}{{2, "2"}, {4, "4"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("MapOKLift() = %v, want %v", got, want)
	}
}

func TestCollectLower(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(k int, v string) (string, bool) {
		if k%2 == 0 {
			return v, true
		}
		return "", false
	}
	got := slices.Collect(MapOKLower(seq, f))
	want := []string{"b"}
	if !slices.Equal(got, want) {
		t.Errorf("MapOKLower() = %v, want %v", got, want)
	}
}

func TestDrop(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	seq := slices.Values(s)
	got := slices.Collect(Drop(seq, 2))
	want := []int{3, 4, 5}
	if !slices.Equal(got, want) {
		t.Errorf("Drop() = %v, want %v", got, want)
	}
}

func TestDrop2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	got := collect2(Drop2(seq, 1))
	want := []struct {
		K int
		V string
	}{{2, "b"}, {3, "c"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("Drop2() = %v, want %v", got, want)
	}
}

func TestDropWhile(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	seq := slices.Values(s)
	f := func(v int) bool { return v < 3 }
	got := slices.Collect(DropWhile(seq, f))
	want := []int{3, 4, 5}
	if !slices.Equal(got, want) {
		t.Errorf("DropWhile() = %v, want %v", got, want)
	}
}

func TestDropWhile2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(k int, v string) bool { return k < 2 }
	got := collect2(DropWhile2(seq, f))
	want := []struct {
		K int
		V string
	}{{2, "b"}, {3, "c"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("DropWhile2() = %v, want %v", got, want)
	}
}

func TestTake(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	seq := slices.Values(s)
	got := slices.Collect(Take(seq, 2))
	want := []int{1, 2}
	if !slices.Equal(got, want) {
		t.Errorf("Take() = %v, want %v", got, want)
	}
}

func TestTake2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	got := collect2(Take2(seq, 1))
	want := []struct {
		K int
		V string
	}{{1, "a"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V {
		t.Errorf("Take2() = %v, want %v", got, want)
	}
}

func TestTakeWhile(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	seq := slices.Values(s)
	f := func(v int) bool { return v < 3 }
	got := slices.Collect(TakeWhile(seq, f))
	want := []int{1, 2}
	if !slices.Equal(got, want) {
		t.Errorf("TakeWhile() = %v, want %v", got, want)
	}
}

func TestTakeWhile2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(k int, v string) bool { return k < 2 }
	got := collect2(TakeWhile2(seq, f))
	want := []struct {
		K int
		V string
	}{{1, "a"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V {
		t.Errorf("TakeWhile2() = %v, want %v", got, want)
	}
}

func TestChunk(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	seq := slices.Values(s)

	chunks := slices.Collect(Map(Chunk(seq, 2), slices.Collect))
	if len(chunks) != 3 {
		t.Fatalf("Chunk() len = %d, want %d", len(chunks), 3)
	}
	if got, want := chunks[0], []int{1, 2}; !slices.Equal(got, want) {
		t.Errorf("Chunk()[0] = %v, want %v", got, want)
	}
	if got, want := chunks[1], []int{3, 4}; !slices.Equal(got, want) {
		t.Errorf("Chunk()[1] = %v, want %v", got, want)
	}
	if got, want := chunks[2], []int{5}; !slices.Equal(got, want) {
		t.Errorf("Chunk()[2] = %v, want %v", got, want)
	}
}

func TestChunk2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
		Singleton2(4, "d"),
		Singleton2(5, "e"),
	)
	chunks := slices.Collect(Map(Chunk2(seq, 2), collect2[int, string]))
	if len(chunks) != 3 {
		t.Fatalf("Chunk2() len = %d, want %d", len(chunks), 3)
	}

	got1 := chunks[0]
	want1 := []struct {
		K int
		V string
	}{{1, "a"}, {2, "b"}}
	if len(got1) != len(want1) || got1[0].K != want1[0].K || got1[0].V != want1[0].V || got1[1].K != want1[1].K || got1[1].V != want1[1].V {
		t.Errorf("Chunk2()[0] = %v, want %v", got1, want1)
	}

	got2 := chunks[1]
	want2 := []struct {
		K int
		V string
	}{{3, "c"}, {4, "d"}}
	if len(got2) != len(want2) || got2[0].K != want2[0].K || got2[0].V != want2[0].V || got2[1].K != want2[1].K || got2[1].V != want2[1].V {
		t.Errorf("Chunk2()[1] = %v, want %v", got2, want2)
	}

	got3 := chunks[2]
	want3 := []struct {
		K int
		V string
	}{{5, "e"}}
	if len(got3) != len(want3) || got3[0].K != want3[0].K || got3[0].V != want3[0].V {
		t.Errorf("Chunk2()[2] = %v, want %v", got3, want3)
	}
}

func TestUnique(t *testing.T) {
	s := []int{1, 2, 2, 3, 3, 3}
	seq := slices.Values(s)
	got := slices.Collect(Unique(seq))
	want := []int{1, 2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("Unique() = %v, want %v", got, want)
	}
}

func TestUniqueFunc(t *testing.T) {
	s := []string{"a", "bb", "c", "dd"}
	seq := slices.Values(s)
	got := slices.Collect(UniqueFunc(seq, func(s string) int { return len(s) }))
	want := []string{"a", "bb"}
	if !slices.Equal(got, want) {
		t.Errorf("UniqueFunc() = %v, want %v", got, want)
	}
}

func TestUnique2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(1, "a"),
	)
	got := collect2(Unique2(seq))
	want := []struct {
		K int
		V string
	}{{1, "a"}, {2, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("Unique2() = %v, want %v", got, want)
	}
}

func TestUniqueFunc2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "a"),
	)
	got := collect2(UniqueFunc2(seq, func(k int, v string) string { return v }))
	want := []struct {
		K int
		V string
	}{{1, "a"}, {2, "b"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("UniqueFunc2() = %v, want %v", got, want)
	}
}

func TestDropWhileKeys(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(k int) bool { return k < 2 }
	got := collect2(DropWhileKeys(seq, f))
	want := []struct {
		K int
		V string
	}{{2, "b"}, {3, "c"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("DropWhileKeys() = %v, want %v", got, want)
	}
}

func TestDropWhileValues(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(v string) bool { return v < "b" }
	got := collect2(DropWhileValues(seq, f))
	want := []struct {
		K int
		V string
	}{{2, "b"}, {3, "c"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V || got[1].K != want[1].K || got[1].V != want[1].V {
		t.Errorf("DropWhileValues() = %v, want %v", got, want)
	}
}

func TestTakeWhileKeys(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(k int) bool { return k < 2 }
	got := collect2(TakeWhileKeys(seq, f))
	want := []struct {
		K int
		V string
	}{{1, "a"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V {
		t.Errorf("TakeWhileKeys() = %v, want %v", got, want)
	}
}

func TestTakeWhileValues(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	f := func(v string) bool { return v < "b" }
	got := collect2(TakeWhileValues(seq, f))
	want := []struct {
		K int
		V string
	}{{1, "a"}}
	if len(got) != len(want) || got[0].K != want[0].K || got[0].V != want[0].V {
		t.Errorf("TakeWhileValues() = %v, want %v", got, want)
	}
}
