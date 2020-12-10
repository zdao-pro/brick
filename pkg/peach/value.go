package peach

import (
	"errors"
)

// ErrNotExist value key not exist.
var (
	ErrNotExist      = errors.New("peach: value key not exit")
	ErrDifferentType = errors.New("peach: value type different")
	ErrTypeAssertion = errors.New("paladin: value type assertion no match")
)

//Value is config value
type Value struct {
	val interface{}
	raw string
}

//NewValue return new value
func NewValue(val interface{}, raw string) *Value {
	return &Value{
		val: val,
		raw: raw,
	}
}

//Bool return bool value
func (v *Value) Bool() (bool, error) {
	if nil == v.val {
		return false, ErrNotExist
	}

	a, ok := v.val.(bool)
	if !ok {
		return false, ErrTypeAssertion
	}

	return a, nil

}

//Int return Int value
func (v *Value) Int() (int, error) {
	i, err := v.Int64()
	return int(i), err

}

// Int64 return int64 value.
func (v *Value) Int64() (int64, error) {
	if v.val == nil {
		return 0, ErrNotExist
	}
	i, ok := v.val.(int64)
	if !ok {
		return 0, ErrTypeAssertion
	}
	return i, nil
}

// Int32 return int32 value.
func (v *Value) Int32() (int32, error) {
	i, err := v.Int64()
	return int32(i), err
}

// Float32 return float32 value.
func (v *Value) Float32() (float32, error) {
	f, err := v.Float64()
	if err != nil {
		return 0.0, err
	}
	return float32(f), nil
}

// Float64 return float64 value.
func (v *Value) Float64() (float64, error) {
	if v.val == nil {
		return 0.0, ErrNotExist
	}
	f, ok := v.val.(float64)
	if !ok {
		return 0.0, ErrTypeAssertion
	}
	return f, nil
}

// String return string value.
func (v *Value) String() (string, error) {
	if v.val == nil {
		return "", ErrNotExist
	}
	s, ok := v.val.(string)
	if !ok {
		return "", ErrTypeAssertion
	}
	return s, nil
}

// Raw return raw value.
func (v *Value) Raw() (string, error) {
	if v.val == nil {
		return "", ErrNotExist
	}
	return v.raw, nil
}
