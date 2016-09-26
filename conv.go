// Package conv provides auto converter for primitive types
package conv

// Conv defines converting interface
type Conv interface {
	IsBool() bool
	AsBool() bool
	IsInt() bool
	AsInt() int64
	IsUint() bool
	AsUint() uint64
	IsFloat() bool
	AsFloat() float64
	IsString() bool
	AsString() string
}
