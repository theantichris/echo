package echo_test

import (
	"testing"
)

func TestEcho(t *testing.T) {
	args := []string{"hello", "world"}

	want := "hello world"
	got := Echo(args)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
