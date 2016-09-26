package conv

import "strconv"

// FromFloat converts float64 to other types
type FromFloat float64

// AsInt converts float64 to int64 by type-casting.
func (f FromFloat) AsInt() int64 {
	return int64(f)
}

// AsUint converts float64 to uint64 by type-casting.
func (f FromFloat) AsUint() uint64 {
	return uint64(f)
}

// AsFloat returns itself
func (f FromFloat) AsFloat() float64 {
	return float64(f)
}

// AsString converts float64 to string by calling strconv.FormatFloat. Returns 0 on error.
func (f FromFloat) AsString() string {
	return strconv.FormatFloat(float64(f), 'G', -1, 64)
}

// AsBool returns false when 0.
func (f FromFloat) AsBool() bool {
	return f != 0
}

func (b FromFloat) IsBool() bool   { return false }
func (b FromFloat) IsInt() bool    { return false }
func (b FromFloat) IsUint() bool   { return false }
func (b FromFloat) IsFloat() bool  { return true }
func (b FromFloat) IsString() bool { return false }
