package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 3)

	if result != 4 {
		t.Errorf("Add (1, 3) FAILED. Expected %d, got %d\n", 4, result)
	} else {
		t.Logf("Add (1, 3) PASSED. Expected %d, got %d\n", 4, result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(5.0, 2.0)

	if result != 2.5 {
		t.Errorf("Divide (5, 2) FAILED. Expected %.1f, got %.1f\n", 2.5, result)
	} else {
		t.Logf("Add (5, 2) PASSED. Expected %.1f, got %.1f", 2.5, result)
	}

}
