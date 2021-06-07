package cachitoGomodTest

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := CachitoGomodTest(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
