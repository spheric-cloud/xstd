// SPDX-FileCopyrightText: 2026 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package xslices

type lengthOption int

func (l lengthOption) initSliceOption() {}

func (l lengthOption) ApplyToMakeMap(o MakeMapOptions) {
	o.makeMapOptions().len = (*int)(&l)
}

func (l lengthOption) ApplyToInitSlice(o InitSliceOptions) {
	o.initSliceOptions().lengthAndCapacity = &lenAndCap{cap: int(l), len: int(l)}
}

func WithLen(length int) interface{ initSliceOption() } {
	return lengthOption(length)
}

func WithLenAndCap(length, capacity int) interface{ initSliceOption() } {
	return &lenAndCap{length, capacity}
}

type InitSliceOption interface {
	ApplyToInitSlice(o InitSliceOptions)
}

type InitSliceOptions interface {
	initSliceOptions() *initSliceOptions
}

type MakeMapOption interface {
	ApplyToMakeMap(o MakeMapOptions)
}

type MakeMapOptions interface {
	makeMapOptions() *makeMapOptions
}

type makeMapOptions struct {
	len *int
}

func (o *makeMapOptions) ApplyOptions(opts []MakeMapOption) *makeMapOptions {
	for _, opt := range opts {
		opt.ApplyToMakeMap(o)
	}
	return o
}

func (o *makeMapOptions) makeMapOptions() *makeMapOptions {
	return o
}

func toMakeMapOptions(opts []MakeMapOption) *makeMapOptions {
	return (&makeMapOptions{}).ApplyOptions(opts)
}

func makeMapFromOptions[M ~map[K]V, K comparable, V any](opts []MakeMapOption) M {
	o := toMakeMapOptions(opts)
	return makeMap[M](o)
}

func makeMap[M ~map[K]V, K comparable, V any](o *makeMapOptions) M {
	if o.len == nil {
		return make(M)
	}
	return make(M, *o.len)
}

type lenAndCap struct {
	len int
	cap int
}

func (l *lenAndCap) initSliceOption() {}

type initSliceOptions struct {
	lengthAndCapacity *lenAndCap
}

func (o *initSliceOptions) minCapHint(cap int) *initSliceOptions {
	if o.lengthAndCapacity == nil {
		o.lengthAndCapacity = &lenAndCap{cap: cap}
		return o
	}

	o.lengthAndCapacity.cap = max(o.lengthAndCapacity.cap, cap+o.lengthAndCapacity.len)
	return o
}

func (o *initSliceOptions) initSliceOptions() *initSliceOptions {
	return o
}

func (o *initSliceOptions) ApplyOptions(opts []InitSliceOption) *initSliceOptions {
	for _, opt := range opts {
		opt.ApplyToInitSlice(o)
	}
	return o
}

func toInitSliceOptions(opts []InitSliceOption) *initSliceOptions {
	return (&initSliceOptions{}).ApplyOptions(opts)
}

func initSlice[S ~[]V, V any](o *initSliceOptions) S {
	if o.lengthAndCapacity == nil {
		var zero S
		return zero
	}
	return make(S, o.lengthAndCapacity.len, o.lengthAndCapacity.cap)
}

func initSliceFromOpts[S ~[]V, V any](opts []InitSliceOption) S {
	o := toInitSliceOptions(opts)
	return initSlice[S](o)
}
