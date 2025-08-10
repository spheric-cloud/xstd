// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package iters

import (
	"slices"
	"testing"
)

func TestAll(t *testing.T) {
	s := []int{2, 4, 6}
	seq := slices.Values(s)
	if !All(seq, func(v int) bool { return v%2 == 0 }) {
		t.Error("All() = false, want true")
	}
	s = []int{2, 4, 5}
	seq = slices.Values(s)
	if All(seq, func(v int) bool { return v%2 == 0 }) {
		t.Error("All() = true, want false")
	}
}

func TestAll2(t *testing.T) {
	seq := Concat2(
		Singleton2(2, "a"),
		Singleton2(4, "b"),
		Singleton2(6, "c"),
	)
	if !All2(seq, func(k int, v string) bool { return k%2 == 0 }) {
		t.Error("All2() = false, want true")
	}
	seq = Concat2(
		Singleton2(2, "a"),
		Singleton2(4, "b"),
		Singleton2(5, "c"),
	)
	if All2(seq, func(k int, v string) bool { return k%2 == 0 }) {
		t.Error("All2() = true, want false")
	}
}

func TestAny(t *testing.T) {
	s := []int{1, 3, 4}
	seq := slices.Values(s)
	if !Any(seq, func(v int) bool { return v%2 == 0 }) {
		t.Error("Any() = false, want true")
	}
	s = []int{1, 3, 5}
	seq = slices.Values(s)
	if Any(seq, func(v int) bool { return v%2 == 0 }) {
		t.Error("Any() = true, want false")
	}
}

func TestAny2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(3, "b"),
		Singleton2(4, "c"),
	)
	if !Any2(seq, func(k int, v string) bool { return k%2 == 0 }) {
		t.Error("Any2() = false, want true")
	}
	seq = Concat2(
		Singleton2(1, "a"),
		Singleton2(3, "b"),
		Singleton2(5, "c"),
	)
	if Any2(seq, func(k int, v string) bool { return k%2 == 0 }) {
		t.Error("Any2() = true, want false")
	}
}

func TestContains(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	if !Contains(seq, 2) {
		t.Error("Contains() = false, want true")
	}
	s = []int{1, 3, 5}
	seq = slices.Values(s)
	if Contains(seq, 2) {
		t.Error("Contains() = true, want false")
	}
}

func TestFind(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	v, ok := Find(seq, func(v int) bool { return v == 2 })
	if !ok || v != 2 {
		t.Errorf("Find() = %d, %v, want 2, true", v, ok)
	}
	s = []int{1, 3, 5}
	seq = slices.Values(s)
	v, ok = Find(seq, func(v int) bool { return v == 2 })
	if ok || v != 0 {
		t.Errorf("Find() = %d, %v, want 0, false", v, ok)
	}
}

func TestForEach(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	var sum int
	ForEach(seq, func(v int) { sum += v })
	if sum != 6 {
		t.Errorf("ForEach() sum = %d, want 6", sum)
	}
}

func TestReduce(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	sum := Reduce(0, seq, func(sum, v int) int { return sum + v })
	if sum != 6 {
		t.Errorf("Reduce() sum = %d, want 6", sum)
	}
}

func TestCount(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	seq := slices.Values(s)
	count := Count[int](seq, func(v int) bool { return v%2 == 0 })
	if count != 2 {
		t.Errorf("Count() = %d, want 2", count)
	}
}

func TestLen(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	seq := slices.Values(s)
	length := Len[int](seq)
	if length != 5 {
		t.Errorf("Len() = %d, want 5", length)
	}
}

func TestFirst(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	v, ok := First(seq)
	if !ok || v != 1 {
		t.Errorf("First() = %d, %v, want 1, true", v, ok)
	}
	s = []int{}
	seq = slices.Values(s)
	v, ok = First(seq)
	if ok || v != 0 {
		t.Errorf("First() = %d, %v, want 0, false", v, ok)
	}
}

func TestIndex(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	v, ok := Index[int](seq, 1)
	if !ok || v != 2 {
		t.Errorf("Index() = %d, %v, want 2, true", v, ok)
	}
	s = []int{1, 2, 3}
	seq = slices.Values(s)
	v, ok = Index[int](seq, 3)
	if ok || v != 0 {
		t.Errorf("Index() = %d, %v, want 0, false", v, ok)
	}
}

func TestMax(t *testing.T) {
	s := []int{1, 5, 3}
	seq := slices.Values(s)
	max := Max(seq)
	if max != 5 {
		t.Errorf("Max() = %d, want 5", max)
	}
}

func TestMin(t *testing.T) {
	s := []int{5, 1, 3}
	seq := slices.Values(s)
	min := Min(seq)
	if min != 1 {
		t.Errorf("Min() = %d, want 1", min)
	}
}

func TestSum(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	sum := Sum[int](seq)
	if sum != 6 {
		t.Errorf("Sum() = %d, want 6", sum)
	}
}

func TestSumFloat(t *testing.T) {
	s := []float64{1.1, 2.2, 3.3}
	seq := slices.Values(s)
	sum := Sum[float64](seq)
	if sum != 6.6 {
		t.Errorf("Sum() = %f, want 6.6", sum)
	}
}

func TestSumCmpOrdered(t *testing.T) {
	s := []string{"a", "b", "c"}
	seq := slices.Values(s)
	sum := Sum[string](seq)
	if sum != "abc" {
		t.Errorf("Sum() = %s, want abc", sum)
	}
}

func TestAllKeys(t *testing.T) {
	seq := Concat2(
		Singleton2(2, "a"),
		Singleton2(4, "b"),
		Singleton2(6, "c"),
	)
	if !AllKeys(seq, func(k int) bool { return k%2 == 0 }) {
		t.Error("AllKeys() = false, want true")
	}
}

func TestAllValues(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "c"),
		Singleton2(3, "e"),
	)
	if !AllValues(seq, func(v string) bool { return v < "f" }) {
		t.Error("AllValues() = false, want true")
	}
}

func TestContains2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	if !Contains2(seq, 2, "b") {
		t.Error("Contains2() = false, want true")
	}
}

func TestContainsKey(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	if !ContainsKey(seq, 2) {
		t.Error("ContainsKey() = false, want true")
	}
}

func TestContainsValue(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	if !ContainsValue(seq, "b") {
		t.Error("ContainsValue() = false, want true")
	}
}

func TestFind2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	k, v, ok := Find2(seq, func(k int, v string) bool { return k == 2 })
	if !ok || k != 2 || v != "b" {
		t.Errorf(`Find2() = %d, %s, %v, want 2, "b", true`, k, v, ok)
	}
}

func TestForEach2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	var keys int
	var values string
	ForEach2(seq, func(k int, v string) {
		keys += k
		values += v
	})
	if keys != 3 || values != "ab" {
		t.Errorf(`ForEach2() keys = %d, values = %s, want 3, "ab"`, keys, values)
	}
}

func TestDrain(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	Drain(seq)
}

func TestDrain2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	Drain2(seq)
}

func TestReduce2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	sum := Reduce2(0, seq, func(sum, k int, v string) int { return sum + k })
	if sum != 3 {
		t.Errorf("Reduce2() sum = %d, want 3", sum)
	}
}

func TestCount2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	count := Count2[int](seq, func(k int, v string) bool { return k%2 != 0 })
	if count != 2 {
		t.Errorf("Count2() = %d, want 2", count)
	}
}

func TestLen2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
		Singleton2(3, "c"),
	)
	length := Len2[int](seq)
	if length != 3 {
		t.Errorf("Len2() = %d, want 3", length)
	}
}

func TestFirst2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	k, v, ok := First2(seq)
	if !ok || k != 1 || v != "a" {
		t.Errorf(`First2() = %d, %s, %v, want 1, "a", true`, k, v, ok)
	}
}

func TestFirstValue(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	v := FirstValue(seq)
	if v != 1 {
		t.Errorf("FirstValue() = %d, want 1", v)
	}
}

func TestFirst2Value(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	k, v := First2Value(seq)
	if k != 1 || v != "a" {
		t.Errorf(`First2Value() = %d, %s, want 1, "a"`, k, v)
	}
}

func TestIndex2(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	k, v, ok := Index2[int](seq, 1)
	if !ok || k != 2 || v != "b" {
		t.Errorf(`Index2() = %d, %s, %v, want 2, "b", true`, k, v, ok)
	}
}

func TestIndexValue(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	v := IndexValue[int](seq, 1)
	if v != 2 {
		t.Errorf("IndexValue() = %d, want 2", v)
	}
}

func TestIndex2Value(t *testing.T) {
	seq := Concat2(
		Singleton2(1, "a"),
		Singleton2(2, "b"),
	)
	k, v := Index2Value[int](seq, 1)
	if k != 2 || v != "b" {
		t.Errorf(`Index2Value() = %d, %s, want 2, "b"`, k, v)
	}
}

func TestMaxFunc(t *testing.T) {
	s := []string{"a", "bb", "c"}
	seq := slices.Values(s)
	max := MaxFunc(seq, func(a, b string) int { return len(a) - len(b) })
	if max != "bb" {
		t.Errorf(`MaxFunc() = %s, want "bb"`, max)
	}
}

func TestMinFunc(t *testing.T) {
	s := []string{"aa", "b", "cc"}
	seq := slices.Values(s)
	min := MinFunc(seq, func(a, b string) int { return len(a) - len(b) })
	if min != "b" {
		t.Errorf(`MinFunc() = %s, want "b"`, min)
	}
}

func TestEqual(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{1, 2, 4}
	if !Equal(slices.Values(s1), slices.Values(s2)) {
		t.Error("Equal() = false, want true")
	}
	if Equal(slices.Values(s1), slices.Values(s3)) {
		t.Error("Equal() = true, want false")
	}
}

func TestEqualFunc(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []string{"a", "bb", "ccc"}
	if !EqualFunc(slices.Values(s1), slices.Values(s2), func(v1 int, v2 string) bool { return v1 == len(v2) }) {
		t.Error("EqualFunc() = false, want true")
	}
}

func TestEqual2(t *testing.T) {
	seq1 := Concat2(Singleton2(1, "a"), Singleton2(2, "b"))
	seq2 := Concat2(Singleton2(1, "a"), Singleton2(2, "b"))
	seq3 := Concat2(Singleton2(1, "a"), Singleton2(2, "c"))
	if !Equal2(seq1, seq2) {
		t.Error("Equal2() = false, want true")
	}
	if Equal2(seq1, seq3) {
		t.Error("Equal2() = true, want false")
	}
}

func TestEqualFunc2(t *testing.T) {
	seq1 := Concat2(Singleton2(1, "a"), Singleton2(2, "b"))
	seq2 := Concat2(Singleton2(1, 1), Singleton2(2, 1))
	if !EqualFunc2(seq1, seq2, func(k1 int, v1 string, k2 int, v2 int) bool { return k1 == k2 && len(v1) == v2 }) {
		t.Error("EqualFunc2() = false, want true")
	}
}

func TestSingle(t *testing.T) {
	s := []int{1}
	v, n := Single(slices.Values(s))
	if n != 1 || v != 1 {
		t.Errorf("Single() = %d, %d, want 1, 1", v, n)
	}
	s = []int{}
	v, n = Single(slices.Values(s))
	if n != 0 || v != 0 {
		t.Errorf("Single() = %d, %d, want 0, 0", v, n)
	}
	s = []int{1, 2}
	v, n = Single(slices.Values(s))
	if n != 2 || v != 0 {
		t.Errorf("Single() = %d, %d, want 0, 2", v, n)
	}
}

func TestSingleValue(t *testing.T) {
	s := []int{1}
	v := SingleValue(slices.Values(s))
	if v != 1 {
		t.Errorf("SingleValue() = %d, want 1", v)
	}
}

func TestSingle2(t *testing.T) {
	seq := Singleton2(1, "a")
	k, v, n := Single2(seq)
	if n != 1 || k != 1 || v != "a" {
		t.Errorf(`Single2() = %d, %s, %d, want 1, "a", 1`, k, v, n)
	}
	seq = Empty2[int, string]()
	k, v, n = Single2(seq)
	if n != 0 || k != 0 || v != "" {
		t.Errorf(`Single2() = %d, %s, %d, want 0, "", 0`, k, v, n)
	}
	seq = Concat2(Singleton2(1, "a"), Singleton2(2, "b"))
	k, v, n = Single2(seq)
	if n != 2 || k != 0 || v != "" {
		t.Errorf(`Single2() = %d, %s, %d, want 0, "", 2`, k, v, n)
	}
}

func TestSingle2Value(t *testing.T) {
	seq := Singleton2(1, "a")
	k, v := Single2Value(seq)
	if k != 1 || v != "a" {
		t.Errorf(`Single2Value() = %d, %s, want 1, "a"`, k, v)
	}
}
