// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package xslices

import (
	"context"
	"errors"
	"slices"
	"strconv"
	"testing"

	"spheric.cloud/xstd/iters"
)

func TestCollect(t *testing.T) {
	got := Collect(slices.Values([]int{1, 2, 3}))
	want := []int{1, 2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("Collect() = %v, want %v", got, want)
	}
}

func TestUnique(t *testing.T) {
	got := Unique([]int{1, 2, 2, 3, 1})
	want := []int{1, 2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("Unique() = %v, want %v", got, want)
	}
}

func TestUniqueByKey(t *testing.T) {
	type item struct {
		id   int
		name string
	}
	items := []item{{1, "a"}, {2, "b"}, {1, "c"}}
	got := UniqueByKey(items, func(i item) int { return i.id })
	if len(got) != 2 || got[0].name != "a" || got[1].name != "b" {
		t.Errorf("UniqueByKey() = %v", got)
	}
}

func TestFilter(t *testing.T) {
	got := Filter([]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 })
	want := []int{2, 4}
	if !slices.Equal(got, want) {
		t.Errorf("Filter() = %v, want %v", got, want)
	}
}

func TestFilterNonNil(t *testing.T) {
	a, b := 1, 2
	got := FilterNonNil([]*int{&a, nil, &b})
	if len(got) != 2 || *got[0] != 1 || *got[1] != 2 {
		t.Error("FilterNonNil failed")
	}
}

func TestFilterNonNilDeref(t *testing.T) {
	a, b := 1, 2
	got := FilterNonNilDeref([]*int{&a, nil, &b})
	want := []int{1, 2}
	if !slices.Equal(got, want) {
		t.Errorf("FilterNonNilDeref() = %v, want %v", got, want)
	}
}

func TestMapOK(t *testing.T) {
	got := MapOK([]int{1, 2, 3}, func(v int) (string, bool) {
		if v%2 == 0 {
			return strconv.Itoa(v), true
		}
		return "", false
	})
	want := []string{"2"}
	if !slices.Equal(got, want) {
		t.Errorf("MapOK() = %v, want %v", got, want)
	}
}

func TestChunks(t *testing.T) {
	got := Chunks([]int{1, 2, 3, 4, 5}, 2)
	if len(got) != 3 {
		t.Fatalf("Chunks() length = %d, want 3", len(got))
	}
	if !slices.Equal(got[0], []int{1, 2}) {
		t.Errorf("Chunks()[0] = %v, want [1 2]", got[0])
	}
	if !slices.Equal(got[1], []int{3, 4}) {
		t.Errorf("Chunks()[1] = %v, want [3 4]", got[1])
	}
	if !slices.Equal(got[2], []int{5}) {
		t.Errorf("Chunks()[2] = %v, want [5]", got[2])
	}
}

func TestFlatten(t *testing.T) {
	got := Flatten([][]int{{1, 2}, {3, 4}, {5}})
	want := []int{1, 2, 3, 4, 5}
	if !slices.Equal(got, want) {
		t.Errorf("Flatten() = %v, want %v", got, want)
	}
}

func TestMap(t *testing.T) {
	got := Map([]int{1, 2, 3}, func(v int) string { return strconv.Itoa(v) })
	want := []string{"1", "2", "3"}
	if !slices.Equal(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}

func TestMapErr(t *testing.T) {
	got, err := MapErr([]int{1, 2, 3}, func(v int) (string, error) { return strconv.Itoa(v), nil })
	if err != nil {
		t.Fatalf("MapErr() error = %v", err)
	}
	want := []string{"1", "2", "3"}
	if !slices.Equal(got, want) {
		t.Errorf("MapErr() = %v, want %v", got, want)
	}

	// With error
	_, err = MapErr([]int{1, 2}, func(v int) (string, error) {
		if v == 2 {
			return "", errors.New("fail")
		}
		return strconv.Itoa(v), nil
	})
	if err == nil {
		t.Error("MapErr should return error")
	}
}

func TestFlatMap(t *testing.T) {
	got := FlatMap([]int{1, 2, 3}, func(v int) []int { return []int{v, v * 10} })
	want := []int{1, 10, 2, 20, 3, 30}
	if !slices.Equal(got, want) {
		t.Errorf("FlatMap() = %v, want %v", got, want)
	}
}

func TestIndexValues(t *testing.T) {
	s := []string{"a", "b", "c"}
	var keys []int
	var vals []string
	for i, v := range IndexValues(s) {
		keys = append(keys, i)
		vals = append(vals, v)
	}
	if !slices.Equal(keys, []int{0, 1, 2}) {
		t.Errorf("IndexValues keys = %v", keys)
	}
	if !slices.Equal(vals, []string{"a", "b", "c"}) {
		t.Errorf("IndexValues vals = %v", vals)
	}
}

func TestAppendRecvChan(t *testing.T) {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	got := AppendRecvChan([]int{0}, c)
	want := []int{0, 1, 2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("AppendRecvChan() = %v, want %v", got, want)
	}
}

func TestCollectRecvChan(t *testing.T) {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	got := CollectRecvChan(c)
	want := []int{1, 2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("CollectRecvChan() = %v, want %v", got, want)
	}
}

func TestAppendPollChan(t *testing.T) {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	got, err := AppendPollChan(context.Background(), []int{0}, c)
	if err != nil {
		t.Fatalf("AppendPollChan() error = %v", err)
	}
	want := []int{0, 1, 2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("AppendPollChan() = %v, want %v", got, want)
	}
}

func TestCollectPollChan(t *testing.T) {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	got, err := CollectPollChan(context.Background(), c)
	if err != nil {
		t.Fatalf("CollectPollChan() error = %v", err)
	}
	want := []int{1, 2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("CollectPollChan() = %v, want %v", got, want)
	}
}

func TestAppendPollChanNoError(t *testing.T) {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	got := AppendPollChanNoError(context.Background(), []int{0}, c)
	want := []int{0, 1, 2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("AppendPollChanNoError() = %v, want %v", got, want)
	}
}

func TestCollectPollChanNoError(t *testing.T) {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	got := CollectPollChanNoError(context.Background(), c)
	want := []int{1, 2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("CollectPollChanNoError() = %v, want %v", got, want)
	}
}

// Ensure the iters import is not flagged
var _ = iters.Of[int]
