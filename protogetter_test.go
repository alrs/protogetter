package main

import (
	"testing"
)

func TestAssembleURL(t *testing.T) {
	exp := "https://github.com/protocolbuffers/protobuf/releases/download/v3.10.1/protobuf-all-3.10.1.tar.gz"
	got := assembleURL("3.10.1")
	template := "got:%s\nexpected:%s"
	if got != exp {
		t.Fatalf(template, got, exp)
	}
	t.Logf(template, got, exp)
}

func TestAssembleFilename(t *testing.T) {
	exp := "protobuf-all-10.10.10rc1.tar.gz"
	template := "got:%s expected:%s"
	got := assembleFilename("10.10.10rc1")
	if got != exp {
		t.Fatalf(template, got, exp)
	}
	t.Logf(template, got, exp)
}

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
