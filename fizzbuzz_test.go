package fizzbuzz

import (
	"testing"
)

func TestFizzbuzz_1_equal_1(t *testing.T) {
	expected := "1"
	actual := fizzbuzz(1)

	if expected != actual {
		t.Errorf("expected: %s, actual: %d\n", expected, actual)
	}
}

func TestFizzbuzz_3_equal_fizz(t *testing.T) {
	expected := "fizz"
	actual := fizzbuzz(3)

	if expected != actual {
		t.Errorf("expected: %s, actual: %d\n", expected, actual)
	}
}

func TestFizzbuzz_5_equal_buzz(t *testing.T) {
	expected := "buzz"
	actual := fizzbuzz(5)

	if expected != actual {
		t.Errorf("expected: %s, actual: %d\n", expected, actual)
	}
}

func TestFizzbuzz_15_equal_fizzbuzz(t *testing.T) {
	expected := "fizzbuzz"
	actual := fizzbuzz(15)

	if expected != actual {
		t.Errorf("expected: %s, actual: %d\n", expected, actual)
	}
}
