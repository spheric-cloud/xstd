// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package gen

import (
	"fmt"
	"strings"
	"testing"
)

func TestCast(t *testing.T) {
	var v any = "hello"
	s := Cast[string](v)
	if s != "hello" {
		t.Errorf("Cast failed, got %s, want hello", s)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Cast should have panicked")
		}
	}()
	Cast[int](v)
}

func TestCastOK(t *testing.T) {
	var v any = "hello"
	s, ok := CastOK[string](v)
	if !ok || s != "hello" {
		t.Errorf("CastOK failed, got %s, %t, want hello, true", s, ok)
	}

	i, ok := CastOK[int](v)
	if ok || i != 0 {
		t.Errorf("CastOK failed, got %d, %t, want 0, false", i, ok)
	}
}

func TestIsA(t *testing.T) {
	var v any = "hello"
	if !IsA[string](v) {
		t.Error("IsA failed, should be string")
	}
	if IsA[int](v) {
		t.Error("IsA failed, should not be int")
	}
}

func TestZero(t *testing.T) {
	s := Zero[string]()
	if s != "" {
		t.Errorf("Zero[string] failed, got %s, want empty string", s)
	}

	i := Zero[int]()
	if i != 0 {
		t.Errorf("Zero[int] failed, got %d, want 0", i)
	}
}

func TestIsZero(t *testing.T) {
	if !IsZero(0) {
		t.Error("IsZero(0) failed, should be true")
	}
	if IsZero(1) {
		t.Error("IsZero(1) failed, should be false")
	}
	if !IsZero("") {
		t.Error("IsZero(\"\") failed, should be true")
	}
	if IsZero("a") {
		t.Error("IsZero(\"a\") failed, should be false")
	}
	var p *int
	if !IsZero(p) {
		t.Error("IsZero(nil) failed, should be true")
	}
	p = new(int)
	if IsZero(p) {
		t.Error("IsZero(new(int)) failed, should be false")
	}
}

func TestNew(t *testing.T) {
	p := New[int]()
	if p == nil {
		t.Fatal("New[int] returned nil")
	}
	if *p != 0 {
		t.Errorf("New[int] returned %d, want 0", *p)
	}
}

func TestTODO(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("TODO should have panicked")
		}
		msg := fmt.Sprint(r)
		expected := "TODO: provide a value of type int"
		if !strings.HasPrefix(msg, expected) {
			t.Errorf("TODO panic message got %q, want prefix %q", msg, expected)
		}
	}()
	TODO[int]()
}

func TestTODO_WithMessage(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("TODO should have panicked")
		}
		msg := fmt.Sprint(r)
		expected := "TODO: my custom message"
		if !strings.HasPrefix(msg, expected) {
			t.Errorf("TODO panic message got %q, want prefix %q", msg, expected)
		}
	}()
	TODO[int]("my custom message")
}

func TestStub(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Stub should have panicked")
		}
		msg := fmt.Sprint(r)
		expected := "Stub was called - this should not happen: "
		if !strings.HasPrefix(msg, expected) {
			t.Errorf("Stub panic message got %q, want prefix %q", msg, expected)
		}
	}()
	Stub[int]()
}
