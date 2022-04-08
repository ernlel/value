# Value: converting any value type to another any value type

## Example:
```go
var x int = 25
valueX := value.New(x)
stringX := valueX.String()
```

## Available operations with value
```go
func (v Value) IsNil() bool 

// IsEmpty ...
func (v Value) IsEmpty() bool 

// Interface ...
func (v Value) Interface(defaultValue ...interface{}) interface{} 

// String ...
func (v Value) String(defaultValue ...string) string 

// Bool ...
func (v Value) Bool(boolTrue ...string) bool 

// Float64 ...
func (v Value) Float64(defaultValue ...float64) float64 

// Float32 ...
func (v Value) Float32(defaultValue ...float32) float32

// Int64 ...
func (v Value) Int64(defaultValue ...int64) int64 

// Int32 ...
func (v Value) Int32(defaultValue ...int32) int32 

// Int16 ...
func (v Value) Int16(defaultValue ...int16) int16

// Int8 ...
func (v Value) Int8(defaultValue ...int8) int8 

// Int ...
func (v Value) Int(defaultValue ...int) int 

// Uint64 ...
func (v Value) Uint64(defaultValue ...uint64) uint64

// Uint32 ...
func (v Value) Uint32(defaultValue ...uint32) uint32 

// Uint16 ...
func (v Value) Uint16(defaultValue ...uint16) uint16

// Uint8 ...
func (v Value) Uint8(defaultValue ...uint8) uint8 

// Uint ...
func (v Value) Uint(defaultValue ...uint) uint 

// Time ...
func (v Value) Time(timeFormat ...string) time.Time 

// TimeString. timeFormat[0] : format of time string, timeFormat[1] : if interface string optionly provide specific format
func (v Value) TimeString(timeFormat ...string) string 

// StringSlice ...
func (v Value) StringSlice(seperator ...string) []string 

// Map ...
func (v Value) Map() map[string]interface{} 

// MapSlice ...
func (v Value) MapSlice() []map[string]interface{} 

// MapGet ...
func (v Value) MapGet(path string) Value 

// Slice ...
func (v Value) Slice(seperator ...string) []interface{} 

// IntSlice ...
func (v Value) IntSlice(seperator ...string) []int 

// BoolSlice ...
// boolTrueAndSeperator 1: boolTrue, 2:seperator, empty string "" to skip parameter
func (v Value) BoolSlice(boolTrueAndSeperator ...string) []bool 

// TimeSlice ...
// timeFormatAndSeperator 1: timeFormat, 2:seperator, empty string "" to skip parameter
func (v Value) TimeSlice(timeFormatAndSeperator ...string) []time.Time 

// ValueMap ...
func (v Value) ValueMap() map[string]Value 

// ValueSlice ...
func (v Value) ValueSlice(seperator ...string) []Value 
```

