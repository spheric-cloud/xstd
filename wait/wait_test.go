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
	"errors"
	"testing"
	"time"
)

func TestPoll(t *testing.T) {
	c := time.After(100 * time.Millisecond)
	err := Poll(t.Context(), 50*time.Millisecond, func(ctx context.Context) (bool, error) {
		select {
		case <-c:
			return true, nil
		default:
			return false, nil
		}
	})
	if err != nil {
		t.Fatalf("Poll: expected no error, got %v", err)
	}
}

func TestPoll_Timeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 50*time.Millisecond)
	defer cancel()

	err := Poll(ctx, 50*time.Millisecond, func(ctx context.Context) (bool, error) {
		return false, nil
	})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Poll: expected context.DeadlineExceeded got %v", err)
	}
}

func TestPoll_Error(t *testing.T) {
	customErr := errors.New("custom error")
	err := Poll(t.Context(), 50*time.Millisecond, func(ctx context.Context) (bool, error) {
		return false, customErr
	})
	if !errors.Is(err, customErr) {
		t.Errorf("Poll: expected custom error, got %v", err)
	}
}
