// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package xstrings

import (
	"iter"
	"math/rand/v2"
	"slices"
	"strings"
)

// JoinSeq joins the elements of the given sequence with the given separator.
func JoinSeq(elems iter.Seq[string], sep string) string {
	var sb strings.Builder
	WriteJoiningSeq(&sb, elems, sep)
	return sb.String()
}

// WriteJoining appends the elements of elems to the Builder, separated by sep.
func WriteJoining(sb *strings.Builder, elems []string, sep string) {
	WriteJoiningSeq(sb, slices.Values(elems), sep)
}

// WriteJoiningSeq appends the elements of elems to the Builder, separated by sep.
func WriteJoiningSeq(sb *strings.Builder, elems iter.Seq[string], sep string) {
	var needSep bool
	for elem := range elems {
		if needSep {
			sb.WriteString(sep)
		} else {
			needSep = true
		}
		sb.WriteString(elem)
	}
}

// Random returns a random string of the given length.
func Random(n int, charset []rune) string {
	if n < 0 {
		panic("strings.Random: negative n")
	}
	if n > 0 && len(charset) == 0 {
		panic("strings.Random: non-zero n and empty charset")
	}
	var sb strings.Builder
	sb.Grow(n)
	for range n {
		r := charset[rand.IntN(len(charset))]
		sb.WriteRune(r)
	}
	return sb.String()
}
