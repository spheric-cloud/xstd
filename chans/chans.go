// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package chans

import (
	"context"
	"iter"

	"golang.org/x/exp/constraints"
	"spheric.cloud/xstd/xconstraints"
)

// Offer attempts to send a value on the channel, respecting context cancellation.
// It returns the context error if the context is cancelled before the send completes.
func Offer[C xconstraints.Send[V], V any](ctx context.Context, c C, v V) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case c <- v:
		return nil
	}
}

// Poll attempts to receive a value from the channel, respecting context cancellation.
// It returns the context error if the context is cancelled before a value is received.
func Poll[C xconstraints.Receive[V], V any](ctx context.Context, c C) (V, error) {
	select {
	case <-ctx.Done():
		var zero V
		return zero, ctx.Err()
	case v := <-c:
		return v, nil
	}
}

// RecvSeq returns an iterator that yields all values received from the channel until it is closed.
func RecvSeq[C xconstraints.Receive[V], V any](c C) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}

// SendSeq sends all values from the iterator sequence to the channel.
// It returns the number of values sent.
func SendSeq[Int constraints.Integer, C xconstraints.Send[V], V any](c C, seq iter.Seq[V]) Int {
	var n Int
	for v := range seq {
		c <- v
		n++
	}
	return n
}

// OfferSeq sends all values from the iterator sequence to the channel, respecting context cancellation.
// It returns the number of values sent and the context error if the context is cancelled.
func OfferSeq[Int constraints.Integer, C xconstraints.Send[V], V any](ctx context.Context, c C, seq iter.Seq[V]) (Int, error) {
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

// OfferSeqNoError sends all values from the iterator sequence to the channel, respecting context cancellation.
// It returns the number of values sent, discarding any context error.
func OfferSeqNoError[Int constraints.Integer, C xconstraints.Send[V], V any](ctx context.Context, c C, seq iter.Seq[V]) Int {
	var n Int
	for v := range seq {
		select {
		case <-ctx.Done():
			return n
		case c <- v:
			n++
		}
	}
	return n
}

// PollSeqNoError returns an iterator that receives values from the channel, respecting context cancellation.
// It stops when the context is cancelled, discarding the error.
func PollSeqNoError[C xconstraints.Receive[V], V any](ctx context.Context, c C) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-c:
				if !yield(v) {
					return
				}
			}
		}
	}
}

// PollSeq returns an iterator that receives values from the channel, respecting context cancellation.
// It yields pairs of (value, error), where error is non-nil when the context is cancelled.
func PollSeq[C xconstraints.Receive[V], V any](ctx context.Context, c C) iter.Seq2[V, error] {
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
