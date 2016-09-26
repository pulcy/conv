package conv

import "strconv"

// FromUint converts uint to other types
type FromUint uint

// AsInt converts uint to int by type-casting.
func (u FromUint) AsInt() int64 {
	return int64(u)
}

// AsUint returns itself
func (u FromUint) AsUint() uint64 {
	return uint64(u)
}

// AsFloat converts uint to float64 by type-casting.
func (u FromUint) AsFloat() float64 {
	return float64(u)
}

// AsString converts uint to string by calling strconv.FormatUint.
func (u FromUint) AsString() string {
	return strconv.FormatUint(uint64(u), 10)
}

// AsBool returns false when 0.
func (u FromUint) AsBool() bool {
	return u != 0
}

// FromUint64 converts uint64 to other types
type FromUint64 uint64

// AsInt converts uint64 to int64 by type-casting.
func (u FromUint64) AsInt() int64 {
	return int64(u)
}

// AsUint returns itself
func (u FromUint64) AsUint() uint64 {
	return uint64(u)
}

// AsFloat converts uint64 to float64 by type-casting.
func (u FromUint64) AsFloat() float64 {
	return float64(u)
}

// AsString converts uint64 to string by calling strconv.FormatUint.
func (u FromUint64) AsString() string {
	return strconv.FormatUint(uint64(u), 10)
}

// AsBool returns false when 0.
func (u FromUint64) AsBool() bool {
	return u != 0
}

func (b FromUint64) IsBool() bool   { return false }
func (b FromUint64) IsInt() bool    { return false }
func (b FromUint64) IsUint() bool   { return true }
func (b FromUint64) IsFloat() bool  { return false }
func (b FromUint64) IsString() bool { return false }

func (b FromUint) IsBool() bool   { return false }
func (b FromUint) IsInt() bool    { return false }
func (b FromUint) IsUint() bool   { return true }
func (b FromUint) IsFloat() bool  { return false }
func (b FromUint) IsString() bool { return false }
