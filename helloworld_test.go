package gohelloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	res := HelloWorld()
	if res != "Hello World" {
		t.Errorf("HelloWorld() = %v, want %v", res, "Hello World")
	}
}
