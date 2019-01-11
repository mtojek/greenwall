package main_test

import (
	"github.com/surullabs/lint/errcheck"
	"github.com/surullabs/lint/gostaticcheck"
	"github.com/surullabs/lint/golint"
	"github.com/surullabs/lint/govet"
	"github.com/surullabs/lint/gofmt"
	"testing"

	"github.com/surullabs/lint"
)

func TestLint(t *testing.T) {
	var lints = lint.Group{
		gofmt.Check{},         // Verify that all files are properly formatted
		govet.Shadow,          // go vet
		golint.Check{},        // golint
		gostaticcheck.Check{}, // honnef.co/go/staticcheck
		errcheck.Check{},
	}

	if err := lints.Check("./..."); err != nil {
		t.Fatalf("Lint failures: %v", err)
	}
}
