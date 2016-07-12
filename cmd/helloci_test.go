package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	if foo() != "good job" {
		t.Fatal("foo should output good job")
	}
}
