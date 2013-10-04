package underground_test

import (
  "testing"
  "strings"
  . "github.com/futoase/underground"
)

func TestUnderground(t *testing.T) {
  if !strings.EqualFold(WelcomeToUnderground(), "Welcome to underground...\n") {
    t.Errorf("Oh! this is heaven!")
  }
}

