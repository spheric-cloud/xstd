// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package iters

import "iter"

type KV[K, V any] struct {
	K K
	V V
}

func collect2[K, V any](seq iter.Seq2[K, V]) []KV[K, V] {
	var s []KV[K, V]
	for k, v := range seq {
		s = append(s, KV[K, V]{k, v})
	}
	return s
}
