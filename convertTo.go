// TO-DO case of time.Time for all convertions

// Copyright © 2018 Ernestas Leliuga.
//
// Code borrowed some parts from https://github.com/spf13/cast/blob/master/caste.go by Steve Francia <spf@spf13.com>

package value

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var errNegativeNotAllowed = errors.New("unable to cast negative value")

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirectToStringerOrError returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error,
func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// ToStringE casts an interface to a string type.
func ToStringE(i interface{}, defaultValue ...string) (value string, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirectToStringerOrError(i)
	value = ""
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case string:
		value = s
	case bool:
		value = strconv.FormatBool(s)
	case float64:
		value = strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		value = strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		value = strconv.Itoa(s)
	case int64:
		value = strconv.FormatInt(s, 10)
	case int32:
		value = strconv.Itoa(int(s))
	case int16:
		value = strconv.FormatInt(int64(s), 10)
	case int8:
		value = strconv.FormatInt(int64(s), 10)
	case uint:
		value = strconv.FormatInt(int64(s), 10)
	case uint64:
		value = strconv.FormatInt(int64(s), 10)
	case uint32:
		value = strconv.FormatInt(int64(s), 10)
	case uint16:
		value = strconv.FormatInt(int64(s), 10)
	case uint8:
		value = strconv.FormatInt(int64(s), 10)
	case []byte:
		value = string(s)
	case template.HTML:
		value = string(s)
	case template.URL:
		value = string(s)
	case template.JS:
		value = string(s)
	case template.CSS:
		value = string(s)
	case template.HTMLAttr:
		value = string(s)
	case nil:
	case fmt.Stringer:
		value = s.String()
	case error:
		value = s.Error()
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
	return
}
func ToString(i interface{}, defaultValue ...string) string {
	v, _ := ToStringE(i, defaultValue...)
	return v
}

// ToBoolE casts an interface to a bool type.
func ToBoolE(i interface{}, boolTrue ...string) (value bool, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	value = false
	err = nil
	i = indirect(i)

	switch b := i.(type) {
	case bool:
		value = b
	case nil:
	case float32, float64, uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64:
		if len(boolTrue) > 0 {
			if ToString(b) == boolTrue[0] {
				value = true
				return
			}
			value = false
			return
		}

		if ToString(b) != "0" {
			value = true
		}
	case string:
		if len(boolTrue) > 0 {
			if b == boolTrue[0] {
				value = true
				return
			}
			value = false
			return
		}
		value, err = strconv.ParseBool(b)
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to bool", i, i)
	}
	return
}
func ToBool(i interface{}, boolTrue ...string) bool {
	v, _ := ToBoolE(i, boolTrue...)
	return v
}

// ToFloat64E casts an interface to a float64 type.
func ToFloat64E(i interface{}, defaultValue ...float64) (value float64, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = float64(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case float64:
		value = s
	case float32:
		value = float64(s)
	case int:
		value = float64(s)
	case int64:
		value = float64(s)
	case int32:
		value = float64(s)
	case int16:
		value = float64(s)
	case int8:
		value = float64(s)
	case uint:
		value = float64(s)
	case uint64:
		value = float64(s)
	case uint32:
		value = float64(s)
	case uint16:
		value = float64(s)
	case uint8:
		value = float64(s)
	case string:
		v, e := strconv.ParseFloat(s, 64)
		if e == nil {
			value = v
		} else {
			err = fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
		}
	case bool:
		if s {
			value = 1
		}
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
	}
	return
}
func ToFloat64(i interface{}, defaultValue ...float64) float64 {
	v, _ := ToFloat64E(i, defaultValue...)
	return v
}

// ToFloat32E casts an interface to a float32 type.
func ToFloat32E(i interface{}, defaultValue ...float32) (value float32, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = float32(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case float64:
		value = float32(s)
	case float32:
		value = s
	case int:
		value = float32(s)
	case int64:
		value = float32(s)
	case int32:
		value = float32(s)
	case int16:
		value = float32(s)
	case int8:
		value = float32(s)
	case uint:
		value = float32(s)
	case uint64:
		value = float32(s)
	case uint32:
		value = float32(s)
	case uint16:
		value = float32(s)
	case uint8:
		value = float32(s)
	case string:
		v, e := strconv.ParseFloat(s, 32)
		if e == nil {
			value = float32(v)
		} else {
			err = fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
		}
	case bool:
		if s {
			value = 1
		}
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
	}
	return
}
func ToFloat32(i interface{}, defaultValue ...float32) float32 {
	v, _ := ToFloat32E(i, defaultValue...)
	return v
}

// ToInt64E casts an interface to an int64 type.
func ToInt64E(i interface{}, defaultValue ...int64) (value int64, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = int64(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case int:
		value = int64(s)
	case int64:
		value = s
	case int32:
		value = int64(s)
	case int16:
		value = int64(s)
	case int8:
		value = int64(s)
	case uint:
		value = int64(s)
	case uint64:
		value = int64(s)
	case uint32:
		value = int64(s)
	case uint16:
		value = int64(s)
	case uint8:
		value = int64(s)
	case float64:
		value = int64(s)
	case float32:
		value = int64(s)
	case string:
		v, e := strconv.ParseInt(s, 0, 0)
		if e == nil {
			value = v
		} else {
			err = fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case bool:
		if s {
			value = 1
		}
	case nil:
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	}
	return
}
func ToInt64(i interface{}, defaultValue ...int64) int64 {
	v, _ := ToInt64E(i, defaultValue...)
	return v
}

// ToInt32E casts an interface to an int32 type.
func ToInt32E(i interface{}, defaultValue ...int32) (value int32, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = int32(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case int:
		value = int32(s)
	case int64:
		value = int32(s)
	case int32:
		value = s
	case int16:
		value = int32(s)
	case int8:
		value = int32(s)
	case uint:
		value = int32(s)
	case uint64:
		value = int32(s)
	case uint32:
		value = int32(s)
	case uint16:
		value = int32(s)
	case uint8:
		value = int32(s)
	case float64:
		value = int32(s)
	case float32:
		value = int32(s)
	case string:
		v, e := strconv.ParseInt(s, 0, 0)
		if e == nil {
			value = int32(v)
		} else {
			err = fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
		}
	case bool:
		if s {
			value = 1
		}
	case nil:
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
	}
	return
}
func ToInt32(i interface{}, defaultValue ...int32) int32 {
	v, _ := ToInt32E(i, defaultValue...)
	return v
}

// ToInt16E casts an interface to an int16 type.
func ToInt16E(i interface{}, defaultValue ...int16) (value int16, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = int16(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case int:
		value = int16(s)
	case int64:
		value = int16(s)
	case int32:
		value = int16(s)
	case int16:
		value = s
	case int8:
		value = int16(s)
	case uint:
		value = int16(s)
	case uint64:
		value = int16(s)
	case uint32:
		value = int16(s)
	case uint16:
		value = int16(s)
	case uint8:
		value = int16(s)
	case float64:
		value = int16(s)
	case float32:
		value = int16(s)
	case string:
		v, e := strconv.ParseInt(s, 0, 0)
		if e == nil {
			value = int16(v)
		} else {
			err = fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
		}
	case bool:
		if s {
			value = 1
		}
	case nil:
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	}
	return
}
func ToInt16(i interface{}, defaultValue ...int16) int16 {
	v, _ := ToInt16E(i, defaultValue...)
	return v
}

// ToInt8E casts an interface to an int8 type.
func ToInt8E(i interface{}, defaultValue ...int8) (value int8, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = int8(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case int:
		value = int8(s)
	case int64:
		value = int8(s)
	case int32:
		value = int8(s)
	case int16:
		value = int8(s)
	case int8:
		value = s
	case uint:
		value = int8(s)
	case uint64:
		value = int8(s)
	case uint32:
		value = int8(s)
	case uint16:
		value = int8(s)
	case uint8:
		value = int8(s)
	case float64:
		value = int8(s)
	case float32:
		value = int8(s)
	case string:
		v, e := strconv.ParseInt(s, 0, 0)
		if e == nil {
			value = int8(v)
		} else {
			err = fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case bool:
		if s {
			value = 1
		}
	case nil:
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	}
	return
}
func ToInt8(i interface{}, defaultValue ...int8) int8 {
	v, _ := ToInt8E(i, defaultValue...)
	return v
}

// ToIntE casts an interface to an int type.
func ToIntE(i interface{}, defaultValue ...int) (value int, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = int(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case int:
		value = s
	case int64:
		value = int(s)
	case int32:
		value = int(s)
	case int16:
		value = int(s)
	case int8:
		value = int(s)
	case uint:
		value = int(s)
	case uint64:
		value = int(s)
	case uint32:
		value = int(s)
	case uint16:
		value = int(s)
	case uint8:
		value = int(s)
	case float64:
		value = int(s)
	case float32:
		value = int(s)
	case string:
		v, e := strconv.ParseInt(s, 0, 0)
		if e == nil {
			value = int(v)
		} else {
			err = fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case bool:
		if s {
			value = 1
		}
	case nil:
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
	return
}
func ToInt(i interface{}, defaultValue ...int) int {
	v, _ := ToIntE(i, defaultValue...)
	return v
}

// ToUint64E casts an interface to a uint64 type.
func ToUint64E(i interface{}, defaultValue ...uint64) (value uint64, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = uint64(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case string:
		v, e := strconv.ParseUint(s, 0, 64)
		if e == nil {
			value = v
		} else {
			err = fmt.Errorf("unable to cast %#v to uint64: %s", i, err)
		}
	case int:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint64(s)
		}
	case int64:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint64(s)
		}
	case int32:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint64(s)
		}
	case int16:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint64(s)
		}
	case int8:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint64(s)
		}
	case uint:
		value = uint64(s)
	case uint64:
		value = s
	case uint32:
		value = uint64(s)
	case uint16:
		value = uint64(s)
	case uint8:
		value = uint64(s)
	case float32:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint64(s)
		}
	case float64:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint64(s)
		}
	case bool:
		if s {
			value = 1
		}
	case nil:
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
	}
	return
}
func ToUint64(i interface{}, defaultValue ...uint64) uint64 {
	v, _ := ToUint64E(i, defaultValue...)
	return v
}

// ToUint32E casts an interface to a uint32 type.
func ToUint32E(i interface{}, defaultValue ...uint32) (value uint32, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = uint32(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case string:
		v, e := strconv.ParseUint(s, 0, 32)
		if e == nil {
			value = uint32(v)
		} else {
			err = fmt.Errorf("unable to cast %#v to uint32: %s", i, err)
		}
	case int:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint32(s)
		}
	case int64:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint32(s)
		}
	case int32:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint32(s)
		}
	case int16:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint32(s)
		}
	case int8:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint32(s)
		}
	case uint:
		value = uint32(s)

	case uint64:
		value = uint32(s)
	case uint32:
		value = s
	case uint16:
		value = uint32(s)
	case uint8:
		value = uint32(s)
	case float64:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint32(s)
		}
	case float32:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint32(s)
		}
	case bool:
		if s {
			value = 1
		}
	case nil:
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
	}
	return
}
func ToUint32(i interface{}, defaultValue ...uint32) uint32 {
	v, _ := ToUint32E(i, defaultValue...)
	return v
}

// ToUint16E casts an interface to a uint16 type.
func ToUint16E(i interface{}, defaultValue ...uint16) (value uint16, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = uint16(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case string:
		v, e := strconv.ParseUint(s, 0, 16)
		if e == nil {
			value = uint16(v)
		} else {
			err = fmt.Errorf("unable to cast %#v to uint16: %s", i, err)
		}
	case int:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint16(s)
		}
	case int64:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint16(s)
		}
	case int32:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint16(s)
		}
	case int16:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint16(s)
		}
	case int8:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint16(s)
		}
	case uint:
		value = uint16(s)
	case uint64:
		value = uint16(s)
	case uint32:
		value = uint16(s)
	case uint16:
		value = s
	case uint8:
		value = uint16(s)
	case float64:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint16(s)
		}
	case float32:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint16(s)
		}
	case bool:
		if s {
			value = 1
		}
	case nil:
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to uint16", i, i)
	}
	return
}
func ToUint16(i interface{}, defaultValue ...uint16) uint16 {
	v, _ := ToUint16E(i, defaultValue...)
	return v
}

// ToUint8E casts an interface to a uint type.
func ToUint8E(i interface{}, defaultValue ...uint8) (value uint8, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = uint8(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case string:
		v, e := strconv.ParseUint(s, 0, 8)
		if e == nil {
			value = uint8(v)
		} else {
			err = fmt.Errorf("unable to cast %#v to uint8: %s", i, err)
		}
	case int:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint8(s)
		}
	case int64:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint8(s)
		}
	case int32:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint8(s)
		}
	case int16:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint8(s)
		}
	case int8:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint8(s)
		}
	case uint:
		value = uint8(s)
	case uint64:
		value = uint8(s)
	case uint32:
		value = uint8(s)
	case uint16:
		value = uint8(s)
	case uint8:
		value = s
	case float64:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint8(s)
		}
	case float32:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint8(s)
		}
	case bool:
		if s {
			value = 1
		}
	case nil:
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
	}
	return
}
func ToUint8(i interface{}, defaultValue ...uint8) uint8 {
	v, _ := ToUint8E(i, defaultValue...)
	return v
}

// ToUintE casts an interface to a uint type.
func ToUintE(i interface{}, defaultValue ...uint) (value uint, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = uint(0)
	err = nil
	if len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	switch s := i.(type) {
	case string:
		v, e := strconv.ParseUint(s, 0, 0)
		if e == nil {
			value = uint(v)
		} else {
			err = fmt.Errorf("unable to cast %#v to uint: %s", i, err)
		}
	case int:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint(s)
		}
	case int64:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint(s)
		}
	case int32:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint(s)
		}
	case int16:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint(s)
		}
	case int8:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint(s)
		}
	case uint:
		value = s
	case uint64:
		value = uint(s)
	case uint32:
		value = uint(s)
	case uint16:
		value = uint(s)
	case uint8:
		value = uint(s)
	case float64:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint(s)
		}
	case float32:
		if s < 0 {
			err = errNegativeNotAllowed
		} else {
			value = uint(s)
		}
	case bool:
		if s {
			value = 1
		}
	case nil:
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
	}
	return
}
func ToUint(i interface{}, defaultValue ...uint) uint {
	v, _ := ToUintE(i, defaultValue...)
	return v
}

// StringToDate attempts to parse a string into a time.Time type using a
// predefined list of formats.  If no suitable format is found, an error is
// returned.
func StringToDate(s string, timeFormat ...string) (time.Time, error) {
	if len(timeFormat) > 0 && timeFormat[0] != "" {
		return time.Parse(timeFormat[0], s)
	}

	return parseDateWith(s, []string{
		time.RFC3339,
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02",                              // MySQL Date
		"15:04:05",                                // MySQL time
		"02 Jan 2006",
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00", // RFC3339 without T
		"2006-01-02 15:04:05",       // MySQL timestamp
		"2006-01-02 15:04",
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	})
}

func parseDateWith(s string, dates []string) (d time.Time, e error) {
	for _, dateType := range dates {
		if d, e = time.Parse(dateType, s); e == nil {
			return
		}
	}
	return d, fmt.Errorf("unable to parse date: %s", s)
}

// ToTimeE casts an interface to a time.Time type.
func ToTimeE(i interface{}, timeFormat ...string) (value time.Time, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = time.Time{}
	err = nil

	switch v := i.(type) {
	case time.Time:
		value = v
	case string:
		value, err = StringToDate(v, timeFormat...)
	case int:
		value = time.Unix(int64(v), 0)
	case int64:
		value = time.Unix(v, 0)
	case int32:
		value = time.Unix(int64(v), 0)
	case uint:
		value = time.Unix(int64(v), 0)
	case uint64:
		value = time.Unix(int64(v), 0)
	case uint32:
		value = time.Unix(int64(v), 0)
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to Time", i, i)
	}
	return
}
func ToTime(i interface{}, timeFormat ...string) time.Time {
	v, _ := ToTimeE(i, timeFormat...)
	return v
}

// ToTimeStringE casts an interface to a time string. timeFormat[0] : format of time string, timeFormat[1] : if interface string optionly provide specific format
func ToTimeStringE(i interface{}, timeFormat ...string) (value string, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	i = indirect(i)
	value = ""
	err = nil

	switch v := i.(type) {
	case time.Time:
		if len(timeFormat) > 0 && timeFormat[0] != "" {
			value = v.Format(timeFormat[0])
		} else {
			value = v.String()
		}
	case string:
		var d time.Time
		var e error
		if len(timeFormat) > 1 && timeFormat[1] != "" {
			d, e = StringToDate(v, timeFormat[1])
			if e != nil {
				err = e
				return
			}
		} else {
			d, e = StringToDate(v)
			if e != nil {
				err = e
				return
			}
		}
		if len(timeFormat) > 0 && timeFormat[0] != "" {
			value = d.Format(timeFormat[0])
		} else {
			value = d.String()
		}

	case int:
		if len(timeFormat) > 0 && timeFormat[0] != "" {
			value = time.Unix(int64(v), 0).Format(timeFormat[0])
		} else {
			value = time.Unix(int64(v), 0).String()
		}
	case int64:
		if len(timeFormat) > 0 && timeFormat[0] != "" {
			value = time.Unix(v, 0).Format(timeFormat[0])
		} else {
			value = time.Unix(v, 0).String()
		}
	case int32:
		if len(timeFormat) > 0 && timeFormat[0] != "" {
			value = time.Unix(int64(v), 0).Format(timeFormat[0])
		} else {
			value = time.Unix(int64(v), 0).String()
		}
	case uint:
		if len(timeFormat) > 0 && timeFormat[0] != "" {
			value = time.Unix(int64(v), 0).Format(timeFormat[0])
		} else {
			value = time.Unix(int64(v), 0).String()
		}
	case uint64:
		if len(timeFormat) > 0 && timeFormat[0] != "" {
			value = time.Unix(int64(v), 0).Format(timeFormat[0])
		} else {
			value = time.Unix(int64(v), 0).String()
		}
	case uint32:
		if len(timeFormat) > 0 && timeFormat[0] != "" {
			value = time.Unix(int64(v), 0).Format(timeFormat[0])
		} else {
			value = time.Unix(int64(v), 0).String()
		}
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to Time string", i, i)
	}
	return
}

// ToTimeString. timeFormat[0] : format of time string, timeFormat[1] : if interface string optionly provide specific format
func ToTimeString(i interface{}, timeFormat ...string) string {
	v, _ := ToTimeStringE(i, timeFormat...)
	return v
}

// ToStringSliceE casts an interface to a []string type.
func ToStringSliceE(i interface{}, seperator ...string) (value []string, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	value = []string{}
	err = nil

	switch v := i.(type) {
	case []interface{}:
		for _, u := range v {
			value = append(value, ToString(u))
		}
	case []Value:
		for _, u := range v {
			value = append(value, u.String())
		}
	case []string:
		value = v
	case []int:
		for _, n := range v {
			value = append(value, ToString(n))
		}
	case []int32:
		for _, n := range v {
			value = append(value, ToString(n))
		}
	case []int64:
		for _, n := range v {
			value = append(value, ToString(n))
		}
	case []float32:
		for _, n := range v {
			value = append(value, ToString(n))
		}
	case []float64:
		for _, n := range v {
			value = append(value, ToString(n))
		}
	case []uint:
		for _, n := range v {
			value = append(value, ToString(n))
		}
	case []uint32:
		for _, n := range v {
			value = append(value, ToString(n))
		}
	case []uint64:
		for _, n := range v {
			value = append(value, ToString(n))
		}
	case string:
		if len(seperator) > 0 {
			value = strings.Split(v, seperator[0])
		} else {
			f := func(c rune) bool {
				return !unicode.IsLetter(c) && !unicode.IsNumber(c)
			}
			value = strings.FieldsFunc(v, f)
		}

	case interface{}:
		str, e := ToStringE(v)
		if e != nil {
			err = fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
			return
		}
		value, err = ToStringSliceE(str)
	case float64, float32:
		value, err = ToStringSliceE(ToString(v))
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
	}
	return
}
func ToStringSlice(i interface{}, seperator ...string) []string {
	v, _ := ToStringSliceE(i, seperator...)
	return v
}

// ToMapE casts an interface to a map[string]interface{} type.
func ToMapE(i interface{}) (value map[string]interface{}, err error) {
	if v, ok := i.(Value); ok {
		i = v.value
	}
	value = map[string]interface{}{}
	err = nil

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			value[ToString(k)] = val
		}
	case map[string]interface{}:
		value = v
	case map[string]Value:
		for k, val := range v {
			value[k] = val.Interface()
		}
	case map[Value]Value:
		for k, val := range v {
			value[k.String()] = val.Interface()
		}
	case map[Value]interface{}:
		for k, val := range v {
			value[k.String()] = val
		}
	case string:
		err = json.Unmarshal([]uint8(v), &value)
	case []uint8:
		err = json.Unmarshal(v, &value)
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to map[string]interface{}", i, i)
	}
	return
}
func ToMap(i interface{}) map[string]interface{} {
	v, _ := ToMapE(i)
	return v
}

func ToMapSliceE(i interface{}) (value []map[string]interface{}, err error) {
	if v, ok := i.(Value); ok {
		i = v.value
	}
	value = []map[string]interface{}{}
	err = nil

	switch v := i.(type) {
	case []map[string]interface{}:
		value = v
	case []interface{}:
		for _, val := range v {
			value = append(value, ToMap(val))
		}
	case string:
		err = json.Unmarshal([]uint8(v), &value)
	case []uint8:
		err = json.Unmarshal(v, &value)
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to map[string]interface{}", i, i)
	}
	return
}
func ToMapSlice(i interface{}) []map[string]interface{} {
	v, _ := ToMapSliceE(i)
	return v
}

func ToMapGetE(i interface{}, path string) (Value, error) {
	keys := strings.Split(path, ".")
	parent, err := ToMapE(i)
	if err != nil {
		return New(nil), err
	}
	var v interface{}
	for i, key := range keys {
		var ok bool
		v, ok = parent[key]
		if !ok {
			return New(nil), errors.New("path '" + path + "' part '" + key + "' not exist in map")
		}

		if i+1 < len(keys) {
			parent, ok = v.(map[string]interface{})
			if !ok {
				return New(nil), errors.New("Part '" + key + "' in path '" + path + "' is not 'map[string]interface{}' type")
			}
		}
	}

	return New(v), nil
}

func ToMapGet(i interface{}, path string) Value {
	v, _ := ToMapGetE(i, path)
	return v
}

// ToSliceE casts an interface to a []interface{} type.
func ToSliceE(i interface{}, seperator ...string) (value []interface{}, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	value = []interface{}{}
	err = nil

	switch v := i.(type) {
	case []interface{}:
		value = v
	case map[string]interface{}:
		for k, val := range v {
			value = append(value, k)
			value = append(value, val)
		}
	case string, float64, float32, int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
		strArr, e := ToStringSliceE(v, seperator...)
		if e != nil {
			err = fmt.Errorf("unable to cast %#v of type %T to []interface{}", i, i)
			return
		}
		value = make([]interface{}, len(strArr))
		for i, inter := range strArr {
			value[i] = interface{}(inter)
		}
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to []interface{}", i, i)
	}
	return
}
func ToSlice(i interface{}, seperator ...string) []interface{} {
	v, _ := ToSliceE(i, seperator...)
	return v
}

// ToIntSliceE casts an interface to a []int type.
func ToIntSliceE(i interface{}, seperator ...string) (value []int, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	value = []int{}
	err = nil
	if i == nil {
		err = fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
		return
	}

	switch v := i.(type) {
	case []int:
		value = v
		return
	case string, float64, float32, int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
		strArr, e := ToStringSliceE(v, seperator...)
		if e != nil {
			err = fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
			return
		}
		a := make([]int, len(strArr))
		for i, inter := range strArr {
			val, e := ToIntE(inter)
			if e != nil {
				err = fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
				return
			}
			a[i] = val
		}
		value = a
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, e := ToIntE(s.Index(j).Interface())
			if e != nil {
				err = fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
				return
			}
			a[j] = val
		}
		value = a
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	}
	return
}
func ToIntSlice(i interface{}, seperator ...string) []int {
	v, _ := ToIntSliceE(i, seperator...)
	return v
}

// ToBoolSliceE casts an interface to a []bool type.
// boolTrueAndSeperator 1: boolTrue, 2:seperator, empty string "" to skip parameter
func ToBoolSliceE(i interface{}, boolTrueAndSeperator ...string) (value []bool, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	boolTrue := []string{}
	seperator := []string{}
	if len(boolTrueAndSeperator) > 0 && boolTrueAndSeperator[0] != "" {
		boolTrue = append(boolTrue, boolTrueAndSeperator[0])
	}
	if len(boolTrueAndSeperator) > 1 && boolTrueAndSeperator[1] != "" {
		seperator = append(seperator, boolTrueAndSeperator[1])
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	value = []bool{}
	err = nil
	if i == nil {
		err = fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}

	switch v := i.(type) {
	case []bool:
		value = v
		return
	case string, float64, float32, int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
		strArr, e := ToStringSliceE(v, seperator...)
		if e != nil {
			err = fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
			return
		}
		a := make([]bool, len(strArr))
		for i, inter := range strArr {
			val, e := ToBoolE(inter, boolTrue...)
			if e != nil {
				err = fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
				return
			}
			a[i] = val
		}
		value = a
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]bool, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, e := ToBoolE(s.Index(j).Interface(), boolTrue...)
			if e != nil {
				err = fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
				return
			}
			a[j] = val
		}
		value = a
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}
	return
}
func ToBoolSlice(i interface{}, boolTrueAndSeperator ...string) []bool {
	v, _ := ToBoolSliceE(i, boolTrueAndSeperator...)
	return v
}

// ToTimeSliceE casts an interface to a []time.Time type.
// timeFormatAndSeperator 1: timeFormat, 2:seperator, empty string "" to skip parameter
func ToTimeSliceE(i interface{}, timeFormatAndSeperator ...string) (value []time.Time, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	timeFormat := []string{}
	seperator := []string{}
	if len(timeFormatAndSeperator) > 0 && timeFormatAndSeperator[0] != "" {
		timeFormat = append(timeFormat, timeFormatAndSeperator[0])
	}
	if len(timeFormatAndSeperator) > 1 && timeFormatAndSeperator[1] != "" {
		seperator = append(seperator, timeFormatAndSeperator[1])
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	value = []time.Time{}
	err = nil
	if i == nil {
		err = fmt.Errorf("unable to cast %#v of type %T to []time.Time", i, i)
	}

	switch v := i.(type) {
	case []time.Time:
		value = v
		return
	case string, float64, float32, int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
		strArr, e := ToStringSliceE(v, seperator...)
		if e != nil {
			err = fmt.Errorf("unable to cast %#v of type %T to []time.Time", i, i)
			return
		}
		a := make([]time.Time, len(strArr))
		for i, inter := range strArr {
			val, e := ToTimeE(inter, timeFormat...)
			if e != nil {
				err = fmt.Errorf("unable to cast %#v of type %T to []time.Time", i, i)
				return
			}
			a[i] = val
		}
		value = a
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]time.Time, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToTimeE(s.Index(j).Interface(), timeFormat...)
			if err != nil {
				return []time.Time{}, fmt.Errorf("unable to cast %#v of type %T to []time.Time", i, i)
			}
			a[j] = val
		}
		value = a
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to []time.Time", i, i)
	}
	return
}
func ToTimeSlice(i interface{}, timeFormatAndSeperator ...string) []time.Time {
	v, _ := ToTimeSliceE(i, timeFormatAndSeperator...)
	return v
}

// ToValueMapE casts an interface to a map[string]interface{} type.
func ToValueMapE(i interface{}) (value map[string]Value, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	value = map[string]Value{}
	err = nil

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			value[ToString(k)] = New(val)
		}
	case map[string]interface{}:
		for k, val := range v {
			value[k] = New(val)
		}
	case map[string]Value:
		value = v
	case map[Value]Value:
		for k, val := range v {
			value[k.String()] = val
		}
	case map[Value]interface{}:
		for k, val := range v {
			value[k.String()] = New(val)
		}
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to map[string]Value", i, i)
	}
	return
}
func ToValueMap(i interface{}) map[string]Value {
	v, _ := ToValueMapE(i)
	return v
}

// ToSliceE casts an interface to a []interface{} type.
func ToValueSliceE(i interface{}, seperator ...string) (value []Value, err error) {
	if v, ok := i.([]uint8); ok {
		i = string(v)
	}
	if v, ok := i.(Value); ok {
		i = v.value
	}
	value = []Value{}
	err = nil

	switch v := i.(type) {
	case []Value:
		value = v
	case []interface{}:
		for _, val := range v {
			value = append(value, New(val))
		}
	case map[string]interface{}:
		for k, val := range v {
			value = append(value, New(k))
			value = append(value, New(val))
		}
	case string, float64, float32, int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
		strArr, e := ToStringSliceE(v, seperator...)
		if e != nil {
			err = fmt.Errorf("unable to cast %#v of type %T to []Value", i, i)
			return
		}
		value = make([]Value, len(strArr))
		for i, inter := range strArr {
			value[i] = New(inter)
		}
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to []Value", i, i)
	}
	return
}
func ToValueSlice(i interface{}, seperator ...string) []Value {
	v, _ := ToValueSliceE(i, seperator...)
	return v
}
