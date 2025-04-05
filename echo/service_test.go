package echo_test

import (
	"testing"

	"github.com/OzBank/oz-echo-core/echo"
)

func TestEchoService(t *testing.T) {
	svc := echo.NewEchoService()

	phrase := "Hello, World!"
	response := svc.Echo(phrase)
	if response != phrase {
		t.Errorf("Expected response '%s', got '%s'", phrase, response)
	}
}
