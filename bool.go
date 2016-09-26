package conv

// FromBool converts bool to other types
type FromBool bool

// AsInt returns 1 or 0
func (b FromBool) AsInt() int64 {
	if b {
		return 1
	}
	return 0
}

// AsUint returns 1 or 0
func (b FromBool) AsUint() uint64 {
	if b {
		return 1
	}
	return 0
}

// AsFloat returns 1 or 0
func (b FromBool) AsFloat() float64 {
	if b {
		return 1.0
	}
	return 0.0
}

// AsString returns "true" or "false"
func (b FromBool) AsString() string {
	if b {
		return "true"
	}
	return "false"
}

// AsBool returns itself.
func (b FromBool) AsBool() bool {
	return bool(b)
}

func (b FromBool) IsBool() bool   { return true }
func (b FromBool) IsInt() bool    { return false }
func (b FromBool) IsUint() bool   { return false }
func (b FromBool) IsFloat() bool  { return false }
func (b FromBool) IsString() bool { return false }
