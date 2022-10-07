package main

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg := hello(name)
	if !want.MatchString(msg) {
		t.Fatalf(`hello("Gladys") = %q, want match for %#q, nil`, msg, want)
	}
}
