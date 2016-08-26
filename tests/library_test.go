package library

import "testing"

func TestReturnThree(t *testing.T) {
	value := ReturnThree()
	if value != 3 {
		t.Fatalf("ReturnThree did not return 3 instead it returned %d", value)
	}
}
