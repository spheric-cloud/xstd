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
	"sync"
	"sync/atomic"

	"spheric.cloud/xstd/container/set"
)

var (
	closedChan = func() chan struct{} {
		ch := make(chan struct{})
		close(ch)
		return ch
	}()
	nilChan <-chan struct{}
)

type Value[V any] interface {
	Value() V
	Done() <-chan struct{}
}

type Value2[K, V any] interface {
	Values() (K, V)
	Done() <-chan struct{}
}

type callback interface {
	callback()
}

type callbackRegistry interface {
	registerCallback(cb callback)
	deregisterCallback(cb callback)
}

type value[V any] struct {
	done  atomic.Value
	value atomic.Value

	mu       sync.Mutex
	children set.Set[callback]
}

func (v *value[V]) registerCallback(cb callback) {
	v.mu.Lock()
	defer v.mu.Unlock()

	if val := v.value.Load(); val != nil {
		cb.callback()
		return
	}

	if v.children == nil {
		v.children = set.New[callback]()
	}
	v.children.Insert(cb)
}

func (v *value[V]) deregisterCallback(cb callback) {
	v.mu.Lock()
	defer v.mu.Unlock()
	delete(v.children, cb)
}

func (v *value[V]) resolve(val V) {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.value.Store(val)

	done, _ := v.done.Load().(chan struct{})
	if done == nil {
		v.done.Store(closedChan)
	} else {
		close(done)
	}

	for child := range v.children {
		child.callback()
	}
	v.children = nil
}

func (v *value[V]) Value() V {
	if val, ok := v.value.Load().(V); ok {
		<-v.Done()
		return val
	}
	var zero V
	return zero
}

func (v *value[V]) Done() <-chan struct{} {
	done, _ := v.done.Load().(chan struct{})
	if done != nil {
		return done
	}

	v.mu.Lock()
	defer v.mu.Unlock()
	done, _ = v.done.Load().(chan struct{})
	if done == nil {
		done = make(chan struct{})
		v.done.Store(done)
	}
	return done
}

func NewValue[V any]() (v Value[V], resolve func(V)) {
	val := new(value[V])

	var once sync.Once
	resolve = func(v V) {
		once.Do(func() {
			val.resolve(v)
		})
	}
	return val, resolve
}

func Compute[V any](f func() V) Value[V] {
	val := new(value[V])
	go func() {
		val.resolve(f())
	}()
	return val
}

func ComputeContext[V any](ctx context.Context, f func(ctx context.Context) V) Value[V] {
	return Compute(func() V { return f(ctx) })
}

type tuple[K, V any] struct {
	K K
	V V
}

type value2[K, V any] struct {
	done   atomic.Value
	values atomic.Value

	mu       sync.Mutex
	children set.Set[callback]
}

func (v *value2[K, V]) registerCallback(cb callback) {
	v.mu.Lock()
	defer v.mu.Unlock()

	if val := v.values.Load(); val != nil {
		cb.callback()
		return
	}

	if v.children == nil {
		v.children = set.New[callback]()
	}
	v.children.Insert(cb)
}

func (v *value2[K, V]) deregisterCallback(cb callback) {
	v.mu.Lock()
	defer v.mu.Unlock()
	delete(v.children, cb)
}

func (v *value2[K, V]) resolve(key K, value V) {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.values.Store(tuple[K, V]{key, value})

	done, _ := v.done.Load().(chan struct{})
	if done == nil {
		v.done.Store(closedChan)
	} else {
		close(done)
	}

	for child := range v.children {
		child.callback()
	}
	v.children = nil
}

func (v *value2[K, V]) Values() (K, V) {
	if val, ok := v.values.Load().(tuple[K, V]); ok {
		<-v.Done()
		return val.K, val.V
	}
	var (
		zeroK K
		zeroV V
	)
	return zeroK, zeroV
}

func (v *value2[K, V]) Done() <-chan struct{} {
	done, _ := v.done.Load().(chan struct{})
	if done != nil {
		return done
	}

	v.mu.Lock()
	defer v.mu.Unlock()
	done, _ = v.done.Load().(chan struct{})
	if done == nil {
		done = make(chan struct{})
		v.done.Store(done)
	}
	return done
}

func NewValue2[K, V any]() (v Value2[K, V], resolve func(K, V)) {
	val := new(value2[K, V])

	var once sync.Once
	resolve = func(k K, v V) {
		once.Do(func() {
			val.resolve(k, v)
		})
	}
	return val, resolve
}

func Compute2[K, V any](f func() (K, V)) Value2[K, V] {
	val := new(value2[K, V])
	go func() {
		val.resolve(f())
	}()
	return val
}

func Compute2Context[K, V any](ctx context.Context, f func(ctx context.Context) (K, V)) Value2[K, V] {
	return Compute2(func() (K, V) { return f(ctx) })
}

func ComputeContextErr[V any](ctx context.Context, f func(ctx context.Context) (V, error)) Value2[V, error] {
	return Compute2Context(ctx, f)
}

type constant[V any] struct {
	value V
}

func (c *constant[V]) Value() V {
	return c.value
}

func (c *constant[V]) Done() <-chan struct{} {
	return closedChan
}

func Const[V any](v V) Value[V] {
	return &constant[V]{
		value: v,
	}
}

type constant2[K, V any] struct {
	key   K
	value V
}

func (c *constant2[K, V]) Values() (K, V) {
	return c.key, c.value
}

func (c *constant2[K, V]) Done() <-chan struct{} {
	return closedChan
}

func Const2[K, V any](k K, v V) Value2[K, V] {
	return &constant2[K, V]{
		key:   k,
		value: v,
	}
}

type never[V any] struct{}

func (never[V]) Value() V {
	var zero V
	return zero
}

func (never[V]) Done() <-chan struct{} {
	return nilChan
}

func Never[V any]() Value[V] {
	return never[V]{}
}

type never2[K, V any] struct{}

func (never2[K, V]) Values() (K, V) {
	var (
		zeroK K
		zeroV V
	)
	return zeroK, zeroV
}

func (never2[K, V]) Done() <-chan struct{} {
	return nilChan
}

func Never2[K, V any]() Value2[K, V] {
	return never2[K, V]{}
}

type cancellableCallbackFunc struct {
	mu   sync.Mutex
	done atomic.Value
	once sync.Once
	f    func()
}

func (f *cancellableCallbackFunc) Cancel() {
	f.mu.Lock()
	defer f.mu.Unlock()

	done, _ := f.done.Load().(chan struct{})
	if done != nil {
		close(done)
	} else {
		f.done.Store(closedChan)
	}
}

func (f *cancellableCallbackFunc) Done() <-chan struct{} {
	done, _ := f.done.Load().(chan struct{})
	if done != nil {
		return done
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	done, _ = f.done.Load().(chan struct{})
	if done != nil {
		return done
	}

	done = make(chan struct{})
	f.done.Store(done)
	return done
}

func (f *cancellableCallbackFunc) callback() {
	f.once.Do(func() {
		go f.f()
	})
}

type doner interface {
	Done() <-chan struct{}
}

type canceller interface {
	Cancel()
}

func registerCallback(v doner, cb callback) {
	done := v.Done()
	if done == nil {
		// Value will never complete.
		return
	}

	select {
	case <-done:
		cb.callback()
		return
	default:
	}

	if cbr, ok := v.(callbackRegistry); ok {
		cbr.registerCallback(cb)
		return
	}

	go func() {
		if cbDone, ok := cb.(doner); ok {
			select {
			case <-done:
				cb.callback()
			case <-cbDone.Done():
			}
			return
		}

		<-done
		cb.callback()
	}()
}

func removeCallback(v doner, cb callback) {
	if cbr, ok := v.(callbackRegistry); ok {
		cbr.deregisterCallback(cb)
		return
	}

	if c, ok := v.(canceller); ok {
		c.Cancel()
	}
}

func After[V any](v Value[V], f func()) (stop func() bool) {
	cb := &cancellableCallbackFunc{f: f}

	registerCallback(v, cb)

	return func() bool {
		var stopped bool
		cb.once.Do(func() {
			stopped = true
		})

		if stopped {
			removeCallback(v, cb)
		}

		return stopped
	}
}

func After2[K, V any](v Value2[K, V], f func()) (stop func() bool) {
	cb := &cancellableCallbackFunc{f: f}

	registerCallback(v, cb)

	return func() bool {
		var stopped bool
		cb.once.Do(func() {
			stopped = true
		})

		if stopped {
			removeCallback(v, cb)
		}

		return stopped
	}
}

type plainCallbackFunc struct {
	f func()
}

func (f *plainCallbackFunc) callback() {
	go f.f()
}

func callbackFunc(f func()) callback {
	return &plainCallbackFunc{f: f}
}
