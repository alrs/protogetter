package main

import (
	"testing"
)

func TestSaneVersion(t *testing.T) {
	cases := map[string]bool{
		"0.0.0":      true,
		"200.100.33": true,
		"12.12":      false,
		"":           false,
		"9.9.a9":     false,
	}

	template := "got:%t expected:%t"
	for ver, exp := range cases {
		got := saneVersion(ver)
		if got != exp {
			t.Fatalf(template, got, exp)
		}
		t.Logf(template, got, exp)
	}
}
