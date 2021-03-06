package log

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"time"
	"unicode"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// sqlRegexp To ? split
	sqlRegexp = regexp.MustCompile(`\?`)
	// numericPlaceHolderRegexp Match number
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

// isPrintable String is it possible to print
func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

// Print Gorm print
func (l *Logger) Print(values ...interface{}) {
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
		)
		fields := []zapcore.Field{
			zap.String("category", "gorm"),
			zap.String("source", values[1].(string)),
		}
		if level == "sql" {
			// duration
			fields = append(fields, zap.String("du", fmt.Sprintf("%.2fms", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0)))
			// formatted
			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				valueStr := "NULL"
				if indirectValue.IsValid() {
					// value to string
					valueStr = interfaceToFormatted(indirectValue.Interface())
				}
				formattedValues = append(formattedValues, valueStr)
			}

			// differentiate between $n placeholders or else treat like ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}
			fields = append(fields, zap.String("sql", sql), zap.Int64("affected", values[5].(int64)))
		} else if level == "log" {
			fields = append(fields, zap.Any("error", values[2]))
		}
		l.Info("SQL", fields...)
	}
}

// interfaceToFormatted Interface convert string
func interfaceToFormatted(value interface{}) string {
	// time
	if t, ok := value.(time.Time); ok {
		return fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05"))
	}

	// bytes
	if b, ok := value.([]byte); ok {
		if str := string(b); isPrintable(str) {
			return fmt.Sprintf("'%v'", str)
		}
		return "'<binary>'"
	}
	if r, ok := value.(driver.Valuer); ok {
		if value, err := r.Value(); err == nil && value != nil {
			return fmt.Sprintf("'%v'", value)
		}
		return "NULL"
	}

	// number
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
		return fmt.Sprintf("%v", value)
	}
	return fmt.Sprintf("'%v'", value)
}
