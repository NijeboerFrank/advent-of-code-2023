package main

import "testing"

func TestMain(t *testing.T) {
	if output := day01First("testinput1.txt"); output != 142 {
		t.Errorf("%d is not equal to %d", output, 142)
	}
	if output := day01Second("testinput1_1.txt"); output != 281 {
		t.Errorf("%d is not equal to %d", output, 281)
	}
}

func TestConverter(t *testing.T) {
	if number, err := stringToDigit("one", 0); err != nil || number != 1 {
		t.Errorf("%d should be 1", number)
	}
	if number, err := stringToDigit("two", 0); err != nil || number != 2 {
		t.Errorf("%d should be 2", number)
	}
	if number, err := stringToDigit("three", 0); err != nil || number != 3 {
		t.Errorf("%d should be 3", number)
	}
	if number, err := stringToDigit("four", 0); err != nil || number != 4 {
		t.Errorf("%d should be 4", number)
	}
	if number, err := stringToDigit("five", 0); err != nil || number != 5 {
		t.Errorf("%d should be 5", number)
	}
	if number, err := stringToDigit("six", 0); err != nil || number != 6 {
		t.Errorf("%d should be 6", number)
	}
	if number, err := stringToDigit("seven", 0); err != nil || number != 7 {
		t.Errorf("%d should be 7", number)
	}
	if number, err := stringToDigit("eight", 0); err != nil || number != 8 {
		t.Errorf("%d should be 8", number)
	}
	if number, err := stringToDigit("nine", 0); err != nil || number != 9 {
		t.Errorf("%d should be 9", number)
	}
}
