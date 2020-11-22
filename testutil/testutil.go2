package testutil

import "testing"

func eq[T comparable](a, b T) bool {
	return a == b
}

func Equal[T comparable](t *testing.T, expected, got T) {
	t.Helper()
	EqualFn(t, expected, got, eq[T])
}

func Equalf[T comparable](t *testing.T, expected, got T, f string, args ...interface{}) {
	t.Helper()
	EqualFnf(t, expected, got, eq[T], f, args...)
}

func EqualFn[T any](t *testing.T, expected, got T, cmp func(a, b T) bool) {
	t.Helper()
	EqualFnf(t, expected, got, cmp, "expected %v, got %v", expected, got)
}

func EqualFnf[T any](t *testing.T, expected, got T, cmp func(a, b T) bool, f string, args ...interface{}) {
	t.Helper()
	if !cmp(expected, got) {
		t.Fatalf(f, args...)
	}
}

func NotEqual[T comparable](t *testing.T, expected, got T) {
	t.Helper()
	NotEqualFn(t, expected, got, eq[T])
}

func NotEqualf[T comparable](t *testing.T, expected, got T, f string, args ...interface{}) {
	t.Helper()
	NotEqualFnf(t, expected, got, eq[T], f, args...)
}

func NotEqualFn[T any](t *testing.T, expected, got T, cmp func(a, b T) bool) {
	t.Helper()
	NotEqualFnf(t, expected, got, cmp, "expected not equal to %v, got %v", expected, got)
}

func NotEqualFnf[T any](t *testing.T, expected, got T, cmp func(a, b T) bool, f string, args ...interface{}) {
	t.Helper()
	if cmp(expected, got) {
		t.Fatalf(f, args...)
	}
}

func False(t *testing.T, b bool) {
	t.Helper()
	if b {
		t.Fatalf("expected false")
	}
}

func True(t *testing.T, b bool) {
	t.Helper()
	if !b {
		t.Fatalf("expected true")
	}
}
