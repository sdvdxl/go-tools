package number

import (
	"testing"
)

func TestDefaultInt(t *testing.T) {
	if DefaultInt("11", 1) != 11 {
		t.Fail()
	}
}
