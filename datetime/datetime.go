package datetime

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

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

// UnmarshalJSON 将时间戳字符串转换为时间
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	timestamp, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}

	if timestamp < 0 {
		return errors.New("timestamp can not be negative")
	}
	fmt.Println("timestamp", timestamp)

	sec := timestamp / 1000
	nsec := timestamp % 1000 * 1000 * 1000
	// nt := Timestamp(time.Unix(sec, nsec).Local())
	var tZero time.Time

	p := (unsafe.Pointer(t))
	tsec := (*int64)(p)
	*tsec = sec - tZero.Unix()

	secPtr := uintptr(p)
	nsecPtr := secPtr + uintptr(unsafe.Sizeof(int64(0)))
	tnsec := (*int32)(unsafe.Pointer(nsecPtr))
	*tnsec = int32(nsec)
	// locPtr := nsecPtr + uintptr(unsafe.Sizeof(int32(0)))
	// tloc := (*time.Location)(unsafe.Pointer(locPtr))
	// tloc = time.Local

	return nil
}

// NewTimestamp return Local Time
func NewTimestamp() Timestamp {
	return Timestamp(time.Now().Local())
}

func (t Timestamp) String() string {
	return time.Time(t).Local().String()
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
