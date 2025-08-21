// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package strings

import (
	"iter"
	"math/rand/v2"
	"slices"
	"strings"
)

// WriteJoining appends the elements of elems to the Builder, separated by sep.
func WriteJoining(sb *strings.Builder, elems []string, sep string) {
	WriteJoiningSeq(sb, slices.Values(elems), sep)
}

// WriteJoiningSeq appends the elements of elems to the Builder, separated by sep.
func WriteJoiningSeq(sb *strings.Builder, elems iter.Seq[string], sep string) {
	var next bool
	for elem := range elems {
		if next {
			sb.WriteString(sep)
		} else {
			next = true
		}
		sb.WriteString(elem)
	}
}

// Random returns a random string of the given length.
func Random(n int, charset []rune) string {
	if n < 0 {
		panic("strings.Random: negative n")
	}
	if len(charset) == 0 {
		panic("strings.Random: empty charset")
	}
	var sb strings.Builder
	sb.Grow(n)
	for i := 0; i < n; i++ {
		r := charset[rand.IntN(len(charset))]
		sb.WriteRune(r)
	}
	return sb.String()
}
