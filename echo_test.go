package echo_test

import (
	"bytes"
	"testing"

	"github.com/theantichris/echo"
)

func TestEcho(t *testing.T) {
	t.Run("it should write out the commandline-arguments", func(t *testing.T) {
		args := []string{"echo", "hello", "world"}
		output := bytes.Buffer{}

		echo.Echo(args, &output)
		got := output.String()
		want := "hello world\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
