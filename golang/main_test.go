package main

import "testing"

func TestAdd(t *testing.T) {
	var a = 1
	var b = 1

	var expected = 2
	var actual = Add(a, b)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestSubtract(t *testing.T) {
	var a = 4
	var b = 2

	var expected = 2
	var actual = Subtract(a, b)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestMultiply(t *testing.T) {
	var a = 2
	var b = 2

	var expected = 4
	var actual = Multiply(a, b)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
