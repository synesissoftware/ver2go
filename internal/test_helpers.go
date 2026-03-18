package internal

import "testing"

func EqualInteger[E int | uint | uint16 | uint64, A int | uint | uint16 | uint64](t *testing.T, expected E, actual A) {
	t.Helper()

	if int64(expected) != int64(actual) {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func EqualString(t *testing.T, expected, actual string) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

