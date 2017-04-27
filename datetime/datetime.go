package datetime

import "time"
import "fmt"

const (
	// DateFormat yyyy-MM-dd
	DateFormat = "2006-01-02"

	// DateTimeFormat yyyy-MM-dd HH:mm:ss
	DateTimeFormat = "2006-01-02 15:04:05"
)

// Timestamp 解析json成毫秒
type Timestamp time.Time

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprint(time.Time(t).UnixNano() / 1000000)), nil
}

// ParseDate yyyy-MM-dd
func ParseDate(val string) time.Time {
	t, err := time.Parse(DateFormat, val)
	if err != nil {
		return time.Time{}
	}

	return t
}

// ParseDateTime yyyy-MM-dd HH:mm:ss
func ParseDateTime(val string) time.Time {
	t, err := time.Parse(DateTimeFormat, val)
	if err != nil {
		return time.Time{}
	}

	return t
}
