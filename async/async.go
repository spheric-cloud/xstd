// Copyright 2026 Axel Christ and Spheric contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package async

import (
	"context"
	"iter"
	"slices"
	"sync"
	"sync/atomic"
	"time"
)

func Delay[V any](v V, t time.Duration) Value[V] {
	return Compute(func() V {
		<-time.After(t)
		return v
	})
}

func Delay2[K, V any](k K, v V, t time.Duration) Value2[K, V] {
	return Compute2(func() (K, V) {
		<-time.After(t)
		return k, v
	})
}

func Map[VIn, VOut any](valIn Value[VIn], f func(VIn) VOut) Value[VOut] {
	if done := valIn.Done(); done == nil {
		return Never[VOut]()
	}

	valOut, resolve := NewValue[VOut]()
	registerCallback(valIn, callbackFunc(func() {
		resolve(f(valIn.Value()))
	}))
	return valOut
}

func Map2[KIn, VIn, KOut, VOut any](valIn Value2[KIn, VIn], f func(KIn, VIn) (KOut, VOut)) Value2[KOut, VOut] {
	if done := valIn.Done(); done == nil {
		return Never2[KOut, VOut]()
	}

	valOut, resolve := NewValue2[KOut, VOut]()
	registerCallback(valIn, callbackFunc(func() {
		resolve(f(valIn.Values()))
	}))
	return valOut
}

func FlatMap[VIn, VOut any](valIn Value[VIn], f func(VIn) Value[VOut]) Value[VOut] {
	if done := valIn.Done(); done == nil {
		return Never[VOut]()
	}
	valOut, resolve := NewValue[VOut]()
	registerCallback(valIn, callbackFunc(func() {
		valInt := f(valIn.Value())
		if done := valInt.Done(); done == nil {
			return
		}

		registerCallback(valInt, callbackFunc(func() {
			resolve(valInt.Value())
		}))
	}))
	return valOut
}

func FlatMap2[KIn, VIn, KOut, VOut any](valIn Value2[KIn, VIn], f func(KIn, VIn) Value2[KOut, VOut]) Value2[KOut, VOut] {
	if done := valIn.Done(); done == nil {
		return Never2[KOut, VOut]()
	}
	valOut, resolve := NewValue2[KOut, VOut]()
	registerCallback(valIn, callbackFunc(func() {
		valInt := f(valIn.Values())
		if done := valInt.Done(); done == nil {
			return
		}

		registerCallback(valInt, callbackFunc(func() {
			resolve(valInt.Values())
		}))
	}))
	return valOut
}

func Flatten[V any](val Value[Value[V]]) Value[V] {
	if done := val.Done(); done == nil {
		return Never[V]()
	}
	valOut, resolve := NewValue[V]()
	registerCallback(val, callbackFunc(func() {
		valInt := val.Value()
		if done := valInt.Done(); done == nil {
			return
		}

		registerCallback(valInt, callbackFunc(func() {
			resolve(valInt.Value())
		}))
	}))
	return valOut
}

func Flatten2[K, V any](val Value[Value2[K, V]]) Value2[K, V] {
	if done := val.Done(); done == nil {
		return Never2[K, V]()
	}
	valOut, resolve := NewValue2[K, V]()
	registerCallback(val, callbackFunc(func() {
		valInt := val.Value()
		if done := valInt.Done(); done == nil {
			return
		}

		registerCallback(valInt, callbackFunc(func() {
			resolve(valInt.Values())
		}))
	}))
	return valOut
}

func Race[V any](vals ...Value[V]) Value[V] {
	if len(vals) == 0 {
		return Never[V]()
	}
	return RaceSeq(slices.Values(vals))
}

func RaceSeq[V any](vals iter.Seq[Value[V]]) Value[V] {
	res, resolve := NewValue[V]()

	var (
		once  sync.Once
		stops []func() bool
		stop  = func() {
			once.Do(func() {
				for _, stop := range stops {
					stop()
				}
			})
		}
	)
	for val := range vals {
		stops = append(stops, After(val, func() {
			resolve(val.Value())
			stop()
		}))
	}
	if len(stops) == 0 {
		return Never[V]()
	}

	return res
}

func Race2[K, V any](vals ...Value2[K, V]) Value2[K, V] {
	if len(vals) == 0 {
		return Never2[K, V]()
	}
	return Race2Seq(slices.Values(vals))
}

func Race2Seq[K, V any](vals iter.Seq[Value2[K, V]]) Value2[K, V] {
	res, resolve := NewValue2[K, V]()

	var (
		once  sync.Once
		stops []func() bool
		stop  = func() {
			once.Do(func() {
				for _, stop := range stops {
					stop()
				}
			})
		}
	)
	for val := range vals {
		stops = append(stops, After2(val, func() {
			resolve(val.Values())
			stop()
		}))
	}
	if len(stops) == 0 {
		return Never2[K, V]()
	}

	return res
}

func All[V any](vs ...Value[V]) Value[[]V] {
	if len(vs) == 0 {
		return Const[[]V](nil)
	}
	return AllSeq(slices.Values(vs))
}

func AllSeq[V any](vs iter.Seq[Value[V]]) Value[[]V] {
	var (
		res, resolve = NewValue[[]V]()
		mu           sync.Mutex
		values       []V
		ct           int
	)

	for v := range vs {
		ct++
		After(v, func() {
			mu.Lock()
			defer mu.Unlock()
			values = append(values, v.Value())
			if len(values) == ct {
				resolve(values)
			}
		})
	}
	return res
}

func All2Seq[K, V any](vs iter.Seq[Value2[K, V]]) Value2[[]K, []V] {
	var (
		res, resolve = NewValue2[[]K, []V]()
		mu           sync.Mutex
		keys         []K
		values       []V
		ct           int
	)

	for v := range vs {
		ct++
		After2(v, func() {
			mu.Lock()
			defer mu.Unlock()

			key, val := v.Values()
			keys = append(keys, key)
			values = append(values, val)
			if len(values) == ct {
				resolve(keys, values)
			}
		})
	}
	return res
}

func Await[V any](ctx context.Context, v Value[V]) (V, error) {
	select {
	case <-ctx.Done():
		var zero V
		return zero, ctx.Err()
	case <-v.Done():
		return v.Value(), nil
	}
}

func Await2[K, V any](ctx context.Context, v Value2[K, V]) (K, V, error) {
	select {
	case <-ctx.Done():
		var (
			zeroK K
			zeroV V
		)
		return zeroK, zeroV, ctx.Err()
	case <-v.Done():
		k, v := v.Values()
		return k, v, nil
	}
}

func TryMap[VIn, VOut any](v Value2[VIn, error], f func(VIn) VOut) Value2[VOut, error] {
	return Map2(v, func(vIn VIn, err error) (VOut, error) {
		if err != nil {
			var zero VOut
			return zero, err
		}
		return f(vIn), nil
	})
}

func TryMapErr[VIn, VOut any](v Value2[VIn, error], f func(VIn) (VOut, error)) Value2[VOut, error] {
	return Map2(v, func(vIn VIn, err error) (VOut, error) {
		if err != nil {
			var zero VOut
			return zero, err
		}
		return f(vIn)
	})
}

func TryAll[V any](vs ...Value2[V, error]) Value2[[]V, error] {
	if len(vs) == 0 {
		return Const2[[]V, error](nil, nil)
	}
	return TryAllSeq(slices.Values(vs))
}

func TryAllSeq[V any](vs iter.Seq[Value2[V, error]]) Value2[[]V, error] {
	var (
		res, resolve = NewValue2[[]V, error]()
		mu           sync.Mutex
		values       []V
		resErr       atomic.Value
		ct           int
		stops        []func() bool
	)

	for v := range vs {
		ct++
		stops = append(stops, After2(v, func() {
			if err, _ := resErr.Load().(error); err != nil {
				return
			}

			mu.Lock()
			defer mu.Unlock()

			if err, _ := resErr.Load().(error); err != nil {
				return
			}

			v, err := v.Values()
			if err != nil {
				resErr.Store(err)
				resolve(nil, err)
				for _, stop := range stops {
					stop()
				}
				return
			}

			values = append(values, v)
			if len(values) == ct {
				resolve(values, nil)
			}
		}))
	}
	return res
}

func MapLift[VIn, KOut, VOut any](valIn Value[VIn], f func(VIn) (KOut, VOut)) Value2[KOut, VOut] {
	if done := valIn.Done(); done == nil {
		return Never2[KOut, VOut]()
	}

	valOut, resolve := NewValue2[KOut, VOut]()
	registerCallback(valIn, callbackFunc(func() {
		resolve(f(valIn.Value()))
	}))
	return valOut
}

func MapLower[KIn, VIn, VOut any](valIn Value2[KIn, VIn], f func(KIn, VIn) VOut) Value[VOut] {
	if done := valIn.Done(); done == nil {
		return Never[VOut]()
	}

	valOut, resolve := NewValue[VOut]()
	registerCallback(valIn, callbackFunc(func() {
		resolve(f(valIn.Values()))
	}))
	return valOut
}

func LiftSuccess[V any](v Value[V]) Value2[V, error] {
	return MapLift(v, func(v V) (V, error) { return v, nil })
}

func LiftFailure[V any](v Value[error]) Value2[V, error] {
	return MapLift(v, func(err error) (V, error) { var zero V; return zero, nil })
}
