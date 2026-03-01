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

package wait

import (
	"context"
	"time"
)

// Poll runs the given function at the start and again every interval until it either returns
// success (true and no error) or an error.
func Poll(ctx context.Context, interval time.Duration, f func(context.Context) (bool, error)) error {
	if ok, err := f(ctx); err != nil || ok {
		return err
	}

	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			if ok, err := f(ctx); err != nil || ok {
				return err
			}
		}
	}
}

// PollValue runs the given function at the start and again every interval until it either returns
// success (value, true and no error) or an error.
func PollValue[V any](ctx context.Context, interval time.Duration, f func(ctx context.Context) (V, bool, error)) (V, error) {
	if v, ok, err := f(ctx); err != nil || ok {
		return v, err
	}

	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			var zero V
			return zero, ctx.Err()
		case <-t.C:
			if v, ok, err := f(ctx); err != nil || ok {
				return v, err
			}
		}
	}
}

// Loop runs the given function immediately and at every interval until the context ends.
func Loop(ctx context.Context, interval time.Duration, f func(context.Context)) {
	f(ctx)

	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			f(ctx)
		}
	}
}

// Until runs the given function immediately and at every interval until the context ends or
// the function returns an error.
// It returns the error of the function or of the context.
func Until(ctx context.Context, interval time.Duration, f func(context.Context) error) error {
	if err := f(ctx); err != nil {
		return err
	}

	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			if err := f(ctx); err != nil {
				return err
			}
		}
	}
}
