package conv

import "strconv"

// FromString converts string to other types
type FromString string

// AsInt converts string to int64 by calling strconv.ParseInt.
// It supports decimal only and will return 0 on format error.
func (s FromString) AsInt() int64 {
	if ret, err := strconv.ParseInt(string(s), 10, 64); err == nil {
		return ret
	}
	return 0
}

// AsUint converts string to uint64 by calling strconv.ParseUint.
// It supports decimal only and will return 0 on format error.
func (s FromString) AsUint() uint64 {
	if ret, err := strconv.ParseUint(string(s), 10, 64); err == nil {
		return ret
	}
	return 0
}

// AsFloat converts string to float64 by calling strconv.ParseFloat.
// It will return 0.0 on format error.
func (s FromString) AsFloat() float64 {
	if ret, err := strconv.ParseFloat(string(s), 64); err == nil {
		return ret
	}
	return 0.0
}

// AsString returns the string itself.
func (s FromString) AsString() string {
	return string(s)
}

// AsBool converts string to bool by calling strconv.ParseBool.
// It will return false on format error.
func (s FromString) AsBool() bool {
	if s.AsInt() != 0 {
		return true
	}
	if s.AsFloat() != 0 {
		return true
	}
	if ret, err := strconv.ParseBool(string(s)); err == nil {
		return ret
	}
	return string(s) != ""
}
