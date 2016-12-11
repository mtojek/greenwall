package main_test

import (
	"testing"

	"github.com/surullabs/lint"
)

func TestLint(t *testing.T) {
	if err := lint.Default.Check("./..."); err != nil {
		t.Fatalf("Lint failures: %v", err)
	}
}
