// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package iters

import (
	"errors"
	"iter"
	"testing"

	"spheric.cloud/xstd/slices"
)

var errTest = errors.New("test error")

func TestTryAll(t *testing.T) {
	s := []int{2, 4, 6}
	seq := MapLift(slices.Values(s), func(v int) (int, error) {
		return v, nil
	})
	ok, err := TryAll(seq, func(v int) bool { return v%2 == 0 })
	if err != nil {
		t.Errorf("TryAll() error = %v, want nil", err)
	}
	if !ok {
		t.Error("TryAll() = false, want true")
	}

	s = []int{2, 4, 5}
	seq = MapLift(slices.Values(s), func(v int) (int, error) {
		return v, nil
	})
	ok, err = TryAll(seq, func(v int) bool { return v%2 == 0 })
	if err != nil {
		t.Errorf("TryAll() error = %v, want nil", err)
	}
	if ok {
		t.Error("TryAll() = true, want false")
	}

	s = []int{2, 4, 6}
	seq = MapLift(slices.Values(s), func(v int) (int, error) {
		if v == 4 {
			return v, errTest
		}
		return v, nil
	})
	ok, err = TryAll(seq, func(v int) bool { return v%2 == 0 })
	if err != errTest {
		t.Errorf("TryAll() error = %v, want %v", err, errTest)
	}
	if ok {
		t.Error("TryAll() = true, want false")
	}
}

func TestTryReduce(t *testing.T) {
	s := []int{1, 2, 3}
	seq := MapLift(slices.Values(s), func(v int) (int, error) {
		return v, nil
	})
	sum, err := TryReduce(0, seq, func(sum, v int) int { return sum + v })
	if err != nil {
		t.Errorf("TryReduce() error = %v, want nil", err)
	}
	if sum != 6 {
		t.Errorf("TryReduce() sum = %d, want 6", sum)
	}

	s = []int{1, 2, 3}
	seq = MapLift(slices.Values(s), func(v int) (int, error) {
		if v == 3 {
			return v, errTest
		}
		return v, nil
	})
	sum, err = TryReduce(0, seq, func(sum, v int) int { return sum + v })
	if err != errTest {
		t.Errorf("TryReduce() error = %v, want %v", err, errTest)
	}
	if sum != 3 {
		t.Errorf("TryReduce() sum = %d, want 3", sum)
	}
}

func TestTryTransform(t *testing.T) {
	var (
		err1 error
		err2 = errors.New("err2")
		err3 = errors.New("err3")
		err4 error
		err5 error
	)

	seq := Concat2(
		Singleton2(1, err1),
		Singleton2(2, err2),
		Singleton2(3, err3),
		Singleton2(4, err4),
		Singleton2(5, err5),
	)

	var (
		vs   []int
		errs []error
	)
	for v, err := range TryTransform(seq, func(seq iter.Seq[int]) iter.Seq[int] { return Map(Slice(seq, 0, 2), func(v int) int { return v * 3 }) }) {
		vs = append(vs, v)
		errs = append(errs, err)
	}

	expectedVs := []int{3, 0, 0, 12, 0}
	if !slices.Equal(expectedVs, vs) {
		t.Errorf("TryTransform() got vs %v, want %v", vs, expectedVs)
	}

	expectedErrs := []error{err1, err2, err3, err4, err5}
	if !slices.Equal(expectedErrs, errs) {
		t.Errorf("TryTransform() got errs %v, want %v", errs, expectedErrs)
	}
}

func BenchmarkTryFlattenVsTryTransformFlatten(b *testing.B) {
	testCases := []struct {
		name      string
		flattener func(seq iter.Seq2[iter.Seq[int], error]) iter.Seq2[int, error]
	}{
		{
			name:      "Flatten",
			flattener: TryFlatten[int],
		},
		{
			name: "TryTransform Flatten",
			flattener: func(seq iter.Seq2[iter.Seq[int], error]) iter.Seq2[int, error] {
				return TryTransform(seq, Flatten)
			},
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			cat := 1000
			seq := LiftSuccess(Map(Range(1, cat), func(i int) iter.Seq[int] { return Range(i*cat, (i+1)*cat) }))

			for v, err := range tc.flattener(seq) {
				_, _ = v, err
			}
		})
	}

}

func TestTryAny(t *testing.T) {
	s := []int{1, 3, 4}
	seq := LiftSuccess(slices.Values(s))
	ok, err := TryAny(seq, func(v int) bool { return v%2 == 0 })
	if err != nil {
		t.Errorf("TryAny() error = %v, want nil", err)
	}
	if !ok {
		t.Error("TryAny() = false, want true")
	}
}

func TestTryContains(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	ok, err := TryContains(seq, 2)
	if err != nil {
		t.Errorf("TryContains() error = %v, want nil", err)
	}
	if !ok {
		t.Error("TryContains() = false, want true")
	}
}

func TestTryFind(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	v, ok, err := TryFind(seq, func(v int) bool { return v == 2 })
	if err != nil {
		t.Errorf("TryFind() error = %v, want nil", err)
	}
	if !ok || v != 2 {
		t.Errorf("TryFind() = %d, %v, want 2, true", v, ok)
	}
}

func TestTryForEach(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	var sum int
	err := TryForEach(seq, func(v int) { sum += v })
	if err != nil {
		t.Errorf("TryForEach() error = %v, want nil", err)
	}
	if sum != 6 {
		t.Errorf("TryForEach() sum = %d, want 6", sum)
	}
}

func TestTryDrain(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	err := TryDrain(seq)
	if err != nil {
		t.Errorf("TryDrain() error = %v, want nil", err)
	}
}

func TestTryMap(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	seq2 := TryMap(seq, func(v int) int { return v * 2 })
	res, err := slices.TryCollect(seq2)
	if err != nil {
		t.Errorf("TryMap() error = %v, want nil", err)
	}
	if !slices.Equal(res, []int{2, 4, 6}) {
		t.Errorf("TryMap() = %v, want %v", res, []int{2, 4, 6})
	}
}

func TestTryFilter(t *testing.T) {
	s := []int{1, 2, 3, 4}
	seq := LiftSuccess(slices.Values(s))
	seq2 := TryFilter(seq, func(v int) bool { return v%2 == 0 })
	res, err := slices.TryCollect(seq2)
	if err != nil {
		t.Errorf("TryFilter() error = %v, want nil", err)
	}
	if !slices.Equal(res, []int{2, 4}) {
		t.Errorf("TryFilter() = %v, want %v", res, []int{2, 4})
	}
}

func TestTryTap(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	var sum int
	seq2 := TryTap(seq, func(v int) { sum += v })
	_, err := slices.TryCollect(seq2)
	if err != nil {
		t.Errorf("TryTap() error = %v, want nil", err)
	}
	if sum != 6 {
		t.Errorf("TryTap() sum = %d, want 6", sum)
	}
}

func TestTrySum(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	sum, err := TrySum[int](seq)
	if err != nil {
		t.Errorf("TrySum() error = %v, want nil", err)
	}
	if sum != 6 {
		t.Errorf("TrySum() = %d, want 6", sum)
	}
}

func TestTryCount(t *testing.T) {
	s := []int{1, 2, 3, 4}
	seq := LiftSuccess(slices.Values(s))
	count, err := TryCount[int](seq, func(v int) bool { return v%2 == 0 })
	if err != nil {
		t.Errorf("TryCount() error = %v, want nil", err)
	}
	if count != 2 {
		t.Errorf("TryCount() = %d, want 2", count)
	}
}

func TestTryLen(t *testing.T) {
	s := []int{1, 2, 3, 4}
	seq := LiftSuccess(slices.Values(s))
	length, err := TryLen[int](seq)
	if err != nil {
		t.Errorf("TryLen() error = %v, want nil", err)
	}
	if length != 4 {
		t.Errorf("TryLen() = %d, want 4", length)
	}
}

func TestTryIndex(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	v, ok, err := TryIndex[int](seq, 1)
	if err != nil {
		t.Errorf("TryIndex() error = %v, want nil", err)
	}
	if !ok || v != 2 {
		t.Errorf("TryIndex() = %d, %v, want 2, true", v, ok)
	}
}

func TestSplitError(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	seq2, err := SplitError(seq)
	res := slices.Collect(seq2)
	if *err != nil {
		t.Errorf("SplitError() error = %v, want nil", *err)
	}
	if !slices.Equal(res, []int{1, 2, 3}) {
		t.Errorf("SplitError() = %v, want %v", res, []int{1, 2, 3})
	}

	seq = Concat2(
		Singleton2[int, error](1, nil),
		Singleton2[int, error](2, errTest),
		Singleton2[int, error](3, nil),
	)
	seq2, err = SplitError(seq)
	res = slices.Collect(seq2)
	if *err != errTest {
		t.Errorf("SplitError() error = %v, want %v", *err, errTest)
	}
	if !slices.Equal(res, []int{1}) {
		t.Errorf("SplitError() = %v, want %v", res, []int{1})
	}
}

func TestTryAllErr(t *testing.T) {
	s := []int{2, 4, 6}
	seq := LiftSuccess(slices.Values(s))
	ok, err := TryAllErr(seq, func(v int) (bool, error) { return v%2 == 0, nil })
	if err != nil {
		t.Errorf("TryAllErr() error = %v, want nil", err)
	}
	if !ok {
		t.Error("TryAllErr() = false, want true")
	}
}

func TestTryAnyErr(t *testing.T) {
	s := []int{1, 3, 4}
	seq := LiftSuccess(slices.Values(s))
	ok, err := TryAnyErr(seq, func(v int) (bool, error) { return v%2 == 0, nil })
	if err != nil {
		t.Errorf("TryAnyErr() error = %v, want nil", err)
	}
	if !ok {
		t.Error("TryAnyErr() = false, want true")
	}
}

func TestTryFindErr(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	v, ok, err := TryFindErr(seq, func(v int) (bool, error) { return v == 2, nil })
	if err != nil {
		t.Errorf("TryFindErr() error = %v, want nil", err)
	}
	if !ok || v != 2 {
		t.Errorf("TryFindErr() = %d, %v, want 2, true", v, ok)
	}
}

func TestTryForEachErr(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	var sum int
	err := TryForEachErr(seq, func(v int) error {
		sum += v
		return nil
	})
	if err != nil {
		t.Errorf("TryForEachErr() error = %v, want nil", err)
	}
	if sum != 6 {
		t.Errorf("TryForEachErr() sum = %d, want 6", sum)
	}
}

func TestTryMapErr(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	seq2 := TryMapErr(seq, func(v int) (int, error) { return v * 2, nil })
	res, err := slices.TryCollect(seq2)
	if err != nil {
		t.Errorf("TryMapErr() error = %v, want nil", err)
	}
	if !slices.Equal(res, []int{2, 4, 6}) {
		t.Errorf("TryMapErr() = %v, want %v", res, []int{2, 4, 6})
	}
}

func TestTryFlatMap(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	seq2 := TryFlatMap(seq, func(v int) iter.Seq[int] { return Range(0, v) })
	res, err := slices.TryCollect(seq2)
	if err != nil {
		t.Errorf("TryFlatMap() error = %v, want nil", err)
	}
	if !slices.Equal(res, []int{0, 0, 1, 0, 1, 2}) {
		t.Errorf("TryFlatMap() = %v, want %v", res, []int{0, 0, 1, 0, 1, 2})
	}
}

func TestTryMax(t *testing.T) {
	s := []int{1, 5, 3}
	seq := LiftSuccess(slices.Values(s))
	max, err := TryMax(seq)
	if err != nil {
		t.Errorf("TryMax() error = %v, want nil", err)
	}
	if max != 5 {
		t.Errorf("TryMax() = %d, want 5", max)
	}
}

func TestTryMin(t *testing.T) {
	s := []int{5, 1, 3}
	seq := LiftSuccess(slices.Values(s))
	min, err := TryMin(seq)
	if err != nil {
		t.Errorf("TryMin() error = %v, want nil", err)
	}
	if min != 1 {
		t.Errorf("TryMin() = %d, want 1", min)
	}
}

func TestTryReduceErr(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	sum, err := TryReduceErr(0, seq, func(sum, v int) (int, error) { return sum + v, nil })
	if err != nil {
		t.Errorf("TryReduceErr() error = %v, want nil", err)
	}
	if sum != 6 {
		t.Errorf("TryReduceErr() = %d, want 6", sum)
	}
}

func TestTryIndexValue(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	v, err := TryIndexValue[int](seq, 1)
	if err != nil {
		t.Errorf("TryIndexValue() error = %v, want nil", err)
	}
	if v != 2 {
		t.Errorf("TryIndexValue() = %d, want 2", v)
	}
}

func TestTryTransformErr(t *testing.T) {
	seq := LiftSuccess(slices.Values([]int{1, 2, 3}))
	seq2 := TryTransformErr(seq, func(seq iter.Seq[int]) iter.Seq2[int, error] {
		return TryMap(LiftSuccess(seq), func(v int) int { return v * 2 })
	})
	res, err := slices.TryCollect(seq2)
	if err != nil {
		t.Errorf("TryTransformErr() error = %v, want nil", err)
	}
	if !slices.Equal(res, []int{2, 4, 6}) {
		t.Errorf("TryTransformErr() = %v, want %v", res, []int{2, 4, 6})
	}
}

func TestLiftSuccess(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	res, err := slices.TryCollect(seq)
	if err != nil {
		t.Errorf("LiftSuccess() error = %v, want nil", err)
	}
	if !slices.Equal(res, s) {
		t.Errorf("LiftSuccess() = %v, want %v", res, s)
	}
}

func TestLiftFailure(t *testing.T) {
	s := []error{errTest, nil}
	seq := LiftFailure[int](slices.Values(s))
	res, err := slices.TryCollect(seq)
	if err != errTest {
		t.Errorf("LiftFailure() error = %v, want %v", err, errTest)
	}
	if len(res) != 0 {
		t.Errorf("LiftFailure() = %v, want empty", res)
	}
}

func TestTryMapOKErr(t *testing.T) {
	s := []int{1, 2, 3, 4}
	seq := LiftSuccess(slices.Values(s))
	seq2 := TryMapOKErr(seq, func(v int) (int, error, bool) {
		if v%2 != 0 {
			return 0, nil, false
		}
		return v * 2, nil, true
	})
	res, err := slices.TryCollect(seq2)
	if err != nil {
		t.Errorf("TryMapOKErr() error = %v, want nil", err)
	}
	if !slices.Equal(res, []int{4, 8}) {
		t.Errorf("TryMapOKErr() = %v, want %v", res, []int{4, 8})
	}
}

func TestTryFlatMapErr(t *testing.T) {
	s := []int{1, 2, 3}
	seq := LiftSuccess(slices.Values(s))
	seq2 := TryFlatMapErr(seq, func(v int) (iter.Seq[int], error) { return Range(0, v), nil })
	res, err := slices.TryCollect(seq2)
	if err != nil {
		t.Errorf("TryFlatMapErr() error = %v, want nil", err)
	}
	if !slices.Equal(res, []int{0, 0, 1, 0, 1, 2}) {
		t.Errorf("TryFlatMapErr() = %v, want %v", res, []int{0, 0, 1, 0, 1, 2})
	}
}

func TestTryFilterErr(t *testing.T) {
	s := []int{1, 2, 3, 4}
	seq := LiftSuccess(slices.Values(s))
	seq2 := TryFilterErr(seq, func(v int) (bool, error) { return v%2 == 0, nil })
	res, err := slices.TryCollect(seq2)
	if err != nil {
		t.Errorf("TryFilterErr() error = %v, want nil", err)
	}
	if !slices.Equal(res, []int{2, 4}) {
		t.Errorf("TryFilterErr() = %v, want %v", res, []int{2, 4})
	}
}

func TestTryMaxFunc(t *testing.T) {
	s := []string{"a", "bb", "c"}
	seq := LiftSuccess(slices.Values(s))
	max, err := TryMaxFunc(seq, func(a, b string) int { return len(a) - len(b) })
	if err != nil {
		t.Errorf("TryMaxFunc() error = %v, want nil", err)
	}
	if max != "bb" {
		t.Errorf("TryMaxFunc() = %s, want bb", max)
	}
}

func TestTryMinFunc(t *testing.T) {
	s := []string{"aa", "b", "cc"}
	seq := LiftSuccess(slices.Values(s))
	min, err := TryMinFunc(seq, func(a, b string) int { return len(a) - len(b) })
	if err != nil {
		t.Errorf("TryMinFunc() error = %v, want nil", err)
	}
	if min != "b" {
		t.Errorf("TryMinFunc() = %s, want b", min)
	}
}
