package collections

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	set:=NewSet(2,"a", "a", "b")
	if len(set.values)==0 {
		t.Fatal("should has values")
	}
}