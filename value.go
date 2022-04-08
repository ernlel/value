package value

import (
	"time"
)

// Value type
type Value struct {
	value interface{}
}

// New value
func New(i interface{}) Value {
	return Value{value: i}
}

// IsNil ...
func (v Value) IsNil() bool {
	return v.value == nil
}

// IsEmpty ...
func (v Value) IsEmpty() bool {
	return v.String() == ""
}

// Interface ...
func (v Value) Interface(defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 && v.value == nil {
		return defaultValue[0]
	}
	return v.value
}

// String ...
func (v Value) String(defaultValue ...string) string {
	return ToString(v.value, defaultValue...)
}

// Bool ...
func (v Value) Bool(boolTrue ...string) bool {
	return ToBool(v.value, boolTrue...)
}

// Float64 ...
func (v Value) Float64(defaultValue ...float64) float64 {
	return ToFloat64(v.value, defaultValue...)
}

// Float32 ...
func (v Value) Float32(defaultValue ...float32) float32 {
	return ToFloat32(v.value, defaultValue...)
}

// Int64 ...
func (v Value) Int64(defaultValue ...int64) int64 {
	return ToInt64(v.value, defaultValue...)
}

// Int32 ...
func (v Value) Int32(defaultValue ...int32) int32 {
	return ToInt32(v.value, defaultValue...)
}

// Int16 ...
func (v Value) Int16(defaultValue ...int16) int16 {
	return ToInt16(v.value, defaultValue...)
}

// Int8 ...
func (v Value) Int8(defaultValue ...int8) int8 {
	return ToInt8(v.value, defaultValue...)
}

// Int ...
func (v Value) Int(defaultValue ...int) int {
	return ToInt(v.value, defaultValue...)
}

// Uint64 ...
func (v Value) Uint64(defaultValue ...uint64) uint64 {
	return ToUint64(v.value, defaultValue...)
}

// Uint32 ...
func (v Value) Uint32(defaultValue ...uint32) uint32 {
	return ToUint32(v.value, defaultValue...)
}

// Uint16 ...
func (v Value) Uint16(defaultValue ...uint16) uint16 {
	return ToUint16(v.value, defaultValue...)
}

// Uint8 ...
func (v Value) Uint8(defaultValue ...uint8) uint8 {
	return ToUint8(v.value, defaultValue...)
}

// Uint ...
func (v Value) Uint(defaultValue ...uint) uint {
	return ToUint(v.value, defaultValue...)
}

// Time ...
func (v Value) Time(timeFormat ...string) time.Time {
	return ToTime(v.value, timeFormat...)
}

// TimeString. timeFormat[0] : format of time string, timeFormat[1] : if interface string optionly provide specific format
func (v Value) TimeString(timeFormat ...string) string {
	return ToTimeString(v.value, timeFormat...)
}

// StringSlice ...
func (v Value) StringSlice(seperator ...string) []string {
	return ToStringSlice(v.value, seperator...)
}

// Map ...
func (v Value) Map() map[string]interface{} {
	return ToMap(v.value)
}

// MapSlice ...
func (v Value) MapSlice() []map[string]interface{} {
	return ToMapSlice(v.value)
}

// MapGet ...
func (v Value) MapGet(path string) Value {
	return ToMapGet(v.value, path)
}

// Slice ...
func (v Value) Slice(seperator ...string) []interface{} {
	return ToSlice(v.value, seperator...)
}

// IntSlice ...
func (v Value) IntSlice(seperator ...string) []int {
	return ToIntSlice(v.value, seperator...)
}

// BoolSlice ...
// boolTrueAndSeperator 1: boolTrue, 2:seperator, empty string "" to skip parameter
func (v Value) BoolSlice(boolTrueAndSeperator ...string) []bool {
	return ToBoolSlice(v.value, boolTrueAndSeperator...)
}

// TimeSlice ...
// timeFormatAndSeperator 1: timeFormat, 2:seperator, empty string "" to skip parameter
func (v Value) TimeSlice(timeFormatAndSeperator ...string) []time.Time {
	return ToTimeSlice(v.value, timeFormatAndSeperator...)
}

// ValueMap ...
func (v Value) ValueMap() map[string]Value {
	return ToValueMap(v.value)
}

// ValueSlice ...
func (v Value) ValueSlice(seperator ...string) []Value {
	return ToValueSlice(v.value, seperator...)
}
