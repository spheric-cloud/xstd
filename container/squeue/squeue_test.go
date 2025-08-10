// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package squeue

import (
	"testing"
)

func TestSQueue(t *testing.T) {
	q := New[int](2)

	if q.Len() != 0 {
		t.Errorf("Len() = %d, want 0", q.Len())
	}
	if q.Cap() != 2 {
		t.Errorf("Cap() = %d, want 2", q.Cap())
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if q.Len() != 2 {
		t.Errorf("Len() = %d, want 2", q.Len())
	}

	// Test grow
	q.Enqueue(3)
	if q.Len() != 3 {
		t.Errorf("Len() = %d, want 3", q.Len())
	}
	if q.Cap() != 4 {
		t.Errorf("Cap() = %d, want 4", q.Cap())
	}

	v, ok := q.Dequeue()
	if !ok || v != 1 {
		t.Errorf("Dequeue() = %d, %v, want 1, true", v, ok)
	}

	v, ok = q.Dequeue()
	if !ok || v != 2 {
		t.Errorf("Dequeue() = %d, %v, want 2, true", v, ok)
	}

	// Test shrink
	v, ok = q.Dequeue()
	if !ok || v != 3 {
		t.Errorf("Dequeue() = %d, %v, want 3, true", v, ok)
	}
	if q.Len() != 0 {
		t.Errorf("Len() = %d, want 0", q.Len())
	}
	if q.Cap() != 4 {
		t.Errorf("Cap() = %d, want 4", q.Cap())
	}

	_, ok = q.Dequeue()
	if ok {
		t.Error("Dequeue() on empty queue should return false")
	}
}

func TestSQueue_EnqueueDequeue(t *testing.T) {
	q := New[int](0)
	for i := 0; i < 100; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 100; i++ {
		v, ok := q.Dequeue()
		if !ok || v != i {
			t.Errorf("Dequeue() = %d, %v, want %d, true", v, ok, i)
		}
	}
}
