package myutil

import "testing"

func TestAdd(t *testing.T) {
	result := Add(4, 3)
	expected := 7
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(6, 9)
	expected := -3
	if result != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d.", expected, result)
	}
}

func TestConcatenate(t *testing.T) {
	result := Concatenate("Learning", "Go")
	expected := "Learning Go"
	if result != expected {
		t.Errorf("Result was incorrect, got: %s, want: %s.", expected, result)
	}
}

func TestGreet(t *testing.T) {
	result := Greet("Ayesha")
	expected := "Hello Ayesha!"
	if result != expected {
		t.Errorf("Result was incorrect, got: %s, want: %s.", expected, result)
	}
}
