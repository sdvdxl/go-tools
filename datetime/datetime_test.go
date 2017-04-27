package datetime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	timestamp := Timestamp(time.Now())
	tm, err := json.Marshal(timestamp)
	if err != nil {
		t.Fail()
	}
	fmt.Println(string(tm) + " ===")
}
