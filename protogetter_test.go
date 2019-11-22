package main

import (
	"testing"
)

func TestProtoFilter(t *testing.T) {
	cases := map[string]bool{
		"proto/src/google/deeper/file.proto": true,
		"":                                   false,
		"proto/src":                          false,
		"proto/src/google/file.proto":        true,
	}
	template := "got:%t expected:%t"
	for dir, exp := range cases {
		got := protoFilter(dir)
		if got != exp {
			t.Fatalf(template, got, exp)
		}
		t.Logf(template, got, exp)
	}
}

func TestStripPath(t *testing.T) {
	cases := map[string]string{
		"one/two/three/four": "three/four",
		"eins/zwei/drei":     "drei",
	}

	template := "got:%q expected:%q"
	for path, exp := range cases {
		got, err := stripPath(path)
		if err != nil {
			t.Fatal(err)
		}
		if got != exp {
			t.Fatalf(template, got, exp)
		}
		t.Logf(template, got, exp)
	}
}

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
