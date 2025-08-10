// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package constraints

// Channel is a constraint that permits any channel type.
type Channel[T any] interface {
	~chan T | ~<-chan T | ~chan<- T
}

// Send is a constraint that permits send-only or bidirectional channels.
type Send[T any] interface {
	~chan T | ~chan<- T
}

// Receive is a constraint that permits receive-only or bidirectional channels.
type Receive[T any] interface {
	~chan T | ~<-chan T
}
