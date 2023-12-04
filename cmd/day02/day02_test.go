package main

import "testing"

func TestDay02First(t *testing.T) {
	ans := process("input2_test.txt")
	if ans != 8 {
		t.Fatalf("Returned: %d. Expected: %d", ans, 8)
	}
}

func TestDay02Second(t *testing.T) {
	ans := process_part2("input2_test.txt")
	if ans != 2286 {
		t.Fatalf("Returned: %d. Expected: %d", ans, 8)
	}
}
