// Package conv provides auto converter for primitive types
package conv

// Conv defines converting interface
type Conv interface {
	AsBool() bool
	AsInt() int64
	AsUint() uint64
	AsFloat() float64
	AsString() string
}
