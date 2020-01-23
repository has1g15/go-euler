package ApplesAndOranges

import "testing"

func TestApplesAndOranges(t *testing.T) {

	appleResult, orangeResult := countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
	if appleResult != 1 {
		t.Error("Test Failed: {} expected, received {}", 1, 1, appleResult)
	}
	if orangeResult != 1 {
		t.Error("Test Failed: {} expected, received {}", 1, 1, orangeResult)
	}
}
