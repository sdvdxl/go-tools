package collection

import (
	"testing"
	"fmt"
)

func TestNewSet(t *testing.T) {
	set:=NewSet(2,"a", "a", "b")
	if len(set.values)==0 {
		t.Fatal("should has values")
	}
}

func TestContains(t *testing.T) {
	set := NewSet(10)
	set.Add("/user/me.html")
	set.Add("/user/m2.html")
	fmt.Println(set.Contains("/user/me.html"))
}