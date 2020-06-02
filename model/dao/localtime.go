package dao

import (
	"database/sql/driver"
	"fmt"
	"time"

	. "github.com/dingdayu/dnsx/internal/consts"
)

// LocalDate 本地日期
type LocalTime struct {
	time.Time
}

// MarshalJSON LocalDate 序列号
func (t LocalTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Format(DateTimeFormat))), nil
}

// Value LocalDate 转 time
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan Gorm 扫描时的数据赋值
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
