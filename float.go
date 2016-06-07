package conv

import "strconv"

// FromFloat converts float64 to other types
type FromFloat float64

// AsInt converts float64 to int64
func (f FromFloat) AsInt() int64 {
	return int64(f)
}

// AsUint converts float64 to uint64.
func (f FromFloat) AsUint() uint64 {
	return uint64(f)
}

// AsFloat returns itself
func (f FromFloat) AsFloat() float64 {
	return float64(f)
}

// AsString converts float64 to string by calling strconv.FormatFloat.
func (f FromFloat) AsString() string {
	return strconv.FormatFloat(float64(f), 'G', -1, 64)
}

// AsBool returns false when 0.
func (f FromFloat) AsBool() bool {
	return f != 0
}
