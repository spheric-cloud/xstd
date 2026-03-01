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
	"slices"
	"testing"
	"time"
)

const eps = 50 * time.Millisecond

func TestNewValue(t *testing.T) {
	v, resolve := NewValue[int]()

	select {
	case <-v.Done():
		t.Fatalf("NewValue: already resolved")
	default:
	}

	if actual := v.Value(); actual != 0 {
		t.Fatalf("NewValue (unresolved): got %d, want 0", actual)
	}

	want := 2
	resolve(want)

	select {
	case <-v.Done():
	case <-time.After(50 * time.Millisecond):
		t.Fatalf("NewValue: not done")
	}

	if actual := v.Value(); actual != want {
		t.Fatalf("NewValue: got %d, want %d", actual, want)
	}

	resolve(3)

	if actual := v.Value(); actual != want {
		t.Fatalf("NewValue: got %d, want %d", actual, want)
	}
}

func TestNever(t *testing.T) {
	v := Never[int]()
	if v.Done() != nil {
		t.Fatalf("Never().Done(): got %v, want nil", v.Done())
	}

	if actual := v.Value(); actual != 0 {
		t.Fatalf("Never(): got %d, want 0", actual)
	}
}

func TestConstant(t *testing.T) {
	v := Const[int](1)

	select {
	case <-v.Done():
	case <-time.After(50 * time.Millisecond):
		t.Fatalf("Constant: not done")
	}

	if actual := v.Value(); actual != 1 {
		t.Fatalf("Constant: got %d, want 1", actual)
	}
}

func TestDelay(t *testing.T) {
	duration := 50 * time.Millisecond
	v := Delay(1, duration)

	select {
	case <-v.Done():
	case <-time.After(duration + eps):
		t.Fatalf("Delay: not done after duration %s", duration)
	}

	if actual := v.Value(); actual != 1 {
		t.Fatalf("Delay: got %d, want 1", actual)
	}
}

func TestAfter(t *testing.T) {
	v, resolve := NewValue[int]()

	resolved := make(chan struct{})
	After(v, func() {
		close(resolved)
	})

	select {
	case <-resolved:
		t.Fatalf("After(): called although not resolved")
	case <-time.After(50 * time.Millisecond):
	}

	resolve(2)

	select {
	case <-resolved:
	case <-time.After(50 * time.Millisecond):
		t.Fatalf("After(): not called")
	}
}

func TestAfter_Stop(t *testing.T) {
	v, resolve := NewValue[int]()

	resolved := make(chan struct{})
	stop := After(v, func() {
		close(resolved)
	})

	select {
	case <-resolved:
		t.Fatalf("After(): called although not resolved")
	case <-time.After(50 * time.Millisecond):
	}

	if stopped := stop(); !stopped {
		t.Fatalf("After().Stop(): not stopped")
	}

	resolve(1)

	select {
	case <-resolved:
		t.Fatalf("After(): called although stopped")
	case <-time.After(50 * time.Millisecond):
	}
}

func TestMap(t *testing.T) {
	v := Map(Const(2), func(v int) int { return v * v })

	select {
	case <-v.Done():
	case <-time.After(50 * time.Millisecond):
		t.Fatalf("Map(): not done")
	}

	if actual := v.Value(); actual != 4 {
		t.Fatalf("Map(): got %d, want 4", actual)
	}
}

func TestRace(t *testing.T) {
	v := Race(Delay(1, 50*time.Millisecond), Delay(2, 100*time.Millisecond), Delay(3, 100*time.Millisecond))

	select {
	case <-v.Done():
	case <-time.After(50*time.Millisecond + eps):
		t.Fatalf("Race(): not done")
	}

	if actual := v.Value(); actual != 1 {
		t.Fatalf("Race(): got %d, want 1", actual)
	}
}

func TestTryAll(t *testing.T) {
	v := TryAll(
		LiftSuccess(Delay(1, 50*time.Millisecond)),
		LiftSuccess(Delay(2, 50*time.Millisecond)),
		LiftSuccess(Const(3)),
	)

	select {
	case <-v.Done():
	case <-time.After(50*time.Millisecond + eps):
		t.Fatalf("TryAll(): not done")
	}

	vs, err := v.Values()
	if err != nil {
		t.Fatalf("TryAll(): got %v, want nil", err)
	}

	slices.Sort(vs)
	want := []int{1, 2, 3}
	if !slices.Equal(vs, want) {
		t.Fatalf("TryAll(): got %v, want %v", vs, want)
	}
}
