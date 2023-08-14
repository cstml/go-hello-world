package greeting

import (
	"testing"
)

func TestGreeting (t *testing.T) {
  msg := "Hello World"
  if msg != Greet {
    t.Fatalf("Strings \"%s\" != \"%s\"", msg, Greet)
  }
}
