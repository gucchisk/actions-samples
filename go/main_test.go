package main

import(
	"testing"
)

func TestHello(t *testing.T) {
	str := hello()
	if str != "hello" {
		t.Errorf("string not equal hello")
	}
}
