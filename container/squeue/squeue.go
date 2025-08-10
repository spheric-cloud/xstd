// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package squeue

// SQueue is a queue that internally uses a slice, hence S(lice)Queue.
// Depending on the amount of data, the queue dynamically shrinks or grows.
type SQueue[E any] struct {
	data     []E
	start    int // Index where first element is located at
	readable int
}

// New constructs a new SQueue instance with provided parameters.
func New[E any](initialSize int) *SQueue[E] {
	return &SQueue[E]{
		data: make([]E, initialSize),
	}
}

// Len returns the number of elements in the queue.
func (r *SQueue[E]) Len() int {
	return r.readable
}

// Cap returns the capacity of the queue.
func (r *SQueue[E]) Cap() int {
	return len(r.data)
}

// migrateTo migrates the queue to a new underlying slice.
func (r *SQueue[E]) migrateTo(newData []E) {
	to := r.start + r.readable
	if to <= len(r.data) {
		copy(newData, r.data[r.start:to])
	} else {
		copied := copy(newData, r.data[r.start:])
		copy(newData[copied:], r.data[:(to%len(r.data))])
	}
	r.start = 0
	r.data = newData
}

// Dequeue dequeues an element.
func (r *SQueue[E]) Dequeue() (data E, ok bool) {
	if r.readable == 0 {
		var zero E
		return zero, false
	}
	r.readable--
	element := r.data[r.start]
	var zero E
	r.data[r.start] = zero // Zero the value to help GC
	if r.start == len(r.data)-1 {
		// Was the last element
		r.start = 0
	} else {
		r.start++
	}

	if r.readable > 0 && r.readable < len(r.data)/4 {
		// need to shrink
		r.migrateTo(make([]E, len(r.data)/2))
	}

	return element, true
}

// Enqueue enqueues an item.
func (r *SQueue[E]) Enqueue(data E) {
	if r.readable == len(r.data) {
		// need to grow
		if len(r.data) == 0 {
			r.data = make([]E, 1)
		}
		r.migrateTo(make([]E, len(r.data)*2))
	}

	if len(r.data) == 0 {
		r.data = make([]E, 1)
	}
	r.data[(r.readable+r.start)%len(r.data)] = data
	r.readable++
}
