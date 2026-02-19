// SPDX-FileCopyrightText: 2026 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package xslices

import (
	"context"

	"spheric.cloud/xstd/xconstraints"
)

// AppendRecvChan appends all values received from the channel to the slice until the channel is closed.
func AppendRecvChan[S ~[]V, C xconstraints.Receive[V], V any](s S, c C) S {
	for v := range c {
		s = append(s, v)
	}
	return s
}

// CollectRecvChan collects all values received from the channel into a new slice.
func CollectRecvChan[C xconstraints.Receive[V], V any](c C, opts ...InitSliceOption) []V {
	s := initSliceFromOpts[[]V](opts)
	return AppendRecvChan(s, c)
}

// AppendPollChan appends values received from the channel to the slice, respecting context cancellation.
// It returns the updated slice and the context error if the context is cancelled.
func AppendPollChan[S ~[]V, C xconstraints.Receive[V], V any](ctx context.Context, s S, c C) (S, error) {
	for {
		select {
		case <-ctx.Done():
			return s, ctx.Err()
		case v, ok := <-c:
			if !ok {
				return s, nil
			}
			s = append(s, v)
		}
	}
}

// CollectPollChan collects values from the channel into a new slice, respecting context cancellation.
func CollectPollChan[C xconstraints.Receive[V], V any](ctx context.Context, c C, opts ...InitSliceOption) ([]V, error) {
	s := initSliceFromOpts[[]V](opts)
	return AppendPollChan(ctx, s, c)
}

// AppendPollChanNoError appends values from the channel to the slice, respecting context cancellation.
// It discards any context error.
func AppendPollChanNoError[S ~[]V, C xconstraints.Receive[V], V any](ctx context.Context, s S, c C) S {
	for {
		select {
		case <-ctx.Done():
			return s
		case v, ok := <-c:
			if !ok {
				return s
			}
			s = append(s, v)
		}
	}
}

// CollectPollChanNoError collects values from the channel into a new slice, respecting context cancellation.
// It discards any context error.
func CollectPollChanNoError[C xconstraints.Receive[V], V any](ctx context.Context, c C, opts ...InitSliceOption) []V {
	s := initSliceFromOpts[[]V](opts)
	return AppendPollChanNoError(ctx, s, c)
}
