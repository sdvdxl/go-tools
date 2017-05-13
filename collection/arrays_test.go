package collection

import "testing"

func TestUniqueString(t *testing.T) {
	slice := []string{"1", "1"}
	if len(UniqueString(slice)) != 1 {
		t.Fail()
	}
}
