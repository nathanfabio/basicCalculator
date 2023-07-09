package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	test := sum(2, 2)

	expected := 4

	if test != expected {
		t.Error("Expected:", expected, "Actual:", test)
	}
}

func TestShouldSumIncorrect(t *testing.T) {
	test := sum(2, 1)

	expected := 5

	if test != expected {
		t.Error("Expected:", expected, "Actual:", test)
	}
}

func TestShouldSubtractCorrect(t *testing.T) {
	test := subtract(10, 5)

	expected := 5

	if test != expected {
		t.Error("Expected:", expected, "Actual", test)
	}
}

func TestShouldSubtractIncorrect(t *testing.T) {
	test := subtract(10, 2)

	expected := 2

	if test != expected {
		t.Error("Expected:", expected, "Actual", test)
	}
}

func TestShouldMultiplyCorrect(t *testing.T) {
	test := multiply(3, 2)

	expected := 6

	if test != expected {
		t.Error("Expected:", expected, "Actual", test)
	}
}

func TestShouldMultiplyIncorrect(t *testing.T) {
	test := multiply(3, 3)

	expected := 10

	if test != expected {
		t.Error("Expected:", expected, "Actual", test)
	}
}


func TestShouldDivideCorrect(t *testing.T) {
	test := divide(30, 2)

	expected := float64(15)

	if test != expected {
		t.Error("Expected:", expected, "Actual", test)
	}
}

func TestShouldDivideIncorrect(t *testing.T) {
	test := divide(30, 2)

	expected := float64(10)

	if test != expected {
		t.Error("Expected:", expected, "Actual", test)
	}
}