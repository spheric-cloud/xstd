// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package strings

import (
	"slices"
	"strings"
	"testing"

	"spheric.cloud/xstd/iters"
)

func TestWriteJoining(t *testing.T) {
	var sb strings.Builder
	WriteJoining(&sb, []string{"a", "b", "c"}, ",")
	if sb.String() != "a,b,c" {
		t.Errorf("WriteJoining failed, got %s, want a,b,c", sb.String())
	}
}

func TestWriteJoiningSeq(t *testing.T) {
	var sb strings.Builder
	WriteJoiningSeq(&sb, slices.Values([]string{"a", "b", "c"}), ",")
	if sb.String() != "a,b,c" {
		t.Errorf("WriteJoiningSeq failed, got %s, want a,b,c", sb.String())
	}

	// Test with empty seq
	sb.Reset()
	WriteJoiningSeq(&sb, iters.Empty[string](), ",")
	if sb.String() != "" {
		t.Errorf("WriteJoiningSeq with empty seq failed, got %s, want empty string", sb.String())
	}

	// Test with one element
	sb.Reset()
	WriteJoiningSeq(&sb, slices.Values([]string{"a"}), ",")
	if sb.String() != "a" {
		t.Errorf("WriteJoiningSeq with one element failed, got %s, want a", sb.String())
	}
}
