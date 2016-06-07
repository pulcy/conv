package conv

import "strconv"

// FromInt converts int to other types
type FromInt int

// AsInt return int itself
func (i FromInt) AsInt() int64 {
	return int64(i)
}

// AsUint converts int to uint64.
func (i FromInt) AsUint() uint64 {
	return uint64(i)
}

// AsFloat converts int to float64.
func (i FromInt) AsFloat() float64 {
	return float64(i)
}

// AsString converts int to string by calling strconv.Itoa.
func (i FromInt) AsString() string {
	return strconv.Itoa(int(i))
}

// AsBool returns false when 0.
func (i FromInt) AsBool() bool {
	return i != 0
}

// FromInt64 converts int64 to other types
type FromInt64 int64

// AsInt return int64 itself
func (i FromInt64) AsInt() int64 {
	return int64(i)
}

// AsUint converts int64 to uint64.
func (i FromInt64) AsUint() uint64 {
	return uint64(i)
}

// AsFloat converts int64 to float64.
func (i FromInt64) AsFloat() float64 {
	return float64(i)
}

// AsString converts int64 to string by calling strconv.FormatInt.
func (i FromInt64) AsString() string {
	return strconv.FormatInt(int64(i), 10)
}

// AsBool returns false when 0.
func (i FromInt64) AsBool() bool {
	return i != 0
}
