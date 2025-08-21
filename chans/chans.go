// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package chans

import (
	"context"
	"iter"

	"spheric.cloud/xstd/constraints"
)

func Offer[C constraints.Send[V], V any](ctx context.Context, c C, v V) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case c <- v:
		return nil
	}
}

func Poll[C constraints.Receive[V], V any](ctx context.Context, c C) (V, error) {
	select {
	case <-ctx.Done():
		var zero V
		return zero, ctx.Err()
	case v := <-c:
		return v, nil
	}
}

func RecvSeq[C constraints.Receive[V], V any](c C) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}

func SendSeq[Int constraints.Integer, C constraints.Send[V], V any](c C, seq iter.Seq[V]) Int {
	var n Int
	for v := range seq {
		c <- v
		n++
	}
	return n
}

func OfferSeq[Int constraints.Integer, C constraints.Send[V], V any](ctx context.Context, c C, seq iter.Seq[V]) (Int, error) {
	var n Int
	for v := range seq {
		select {
		case <-ctx.Done():
			return n, ctx.Err()
		case c <- v:
			n++
		}
	}
	return n, nil
}

func PollSeq[C constraints.Receive[V], V any](ctx context.Context, c C) iter.Seq2[V, error] {
	return func(yield func(V, error) bool) {
		for {
			select {
			case <-ctx.Done():
				var zero V
				yield(zero, ctx.Err())
				return
			case v := <-c:
				if !yield(v, nil) {
					return
				}
			}
		}
	}
}
