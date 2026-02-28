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

package xcmp

import (
	"slices"
	"testing"
)

func TestZeroPredicate(t *testing.T) {
	p := ZeroPredicate[string]()

	if s := ""; !p(s) {
		t.Errorf("ZeroPredicate(%q): got false, want true", s)
	}

	if s := "foo"; p(s) {
		t.Errorf("ZeroPredicate(%q): got true, want false", s)
	}
}

func TestIsZero(t *testing.T) {
	if s := ""; !IsZero(s) {
		t.Errorf("IsZero(%q): got false, want true", s)
	}

	if s := "foo"; IsZero(s) {
		t.Errorf("IsZero(%q): got true, want false", s)
	}
}

func TestAnyZeroSeq(t *testing.T) {
	for _, tc := range []struct {
		name string
		vs   []string
		want bool
	}{
		{
			name: "zero in middle",
			vs:   []string{"foo", "", "baz"},
			want: true,
		},
		{
			name: "none zero",
			vs:   []string{"foo", "bar", "baz"},
			want: false,
		},
		{
			name: "all zero",
			vs:   []string{"", "", ""},
			want: true,
		},
		{
			name: "empty sequence",
			vs:   nil,
			want: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res := AnyZeroSeq(slices.Values(tc.vs))
			if res != tc.want {
				t.Errorf("AnyZeroSeq(%v): got %v, want %v", tc.vs, res, tc.want)
			}
		})
	}
}
