package conv

// FromBool converts bool to other types
type FromBool bool

// AsInt converts bool to int64.
func (b FromBool) AsInt() int64 {
	if b {
		return 1
	}
	return 0
}

// AsUint converts bool to uint64.
func (b FromBool) AsUint() uint64 {
	if b {
		return 1
	}
	return 0
}

// AsFloat converts bool to float64.
func (b FromBool) AsFloat() float64 {
	if b {
		return 1.0
	}
	return 0.0
}

// AsString converts bool to string.
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
