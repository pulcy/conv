package conv

import (
	"database/sql"
	"fmt"
	"reflect"
)

// FromAny returns a Conv to converts anything to primitive types
//
// For numeric types, bool, string, see FromXXX types. []rune is treated as string
//
// For sql nullable types, converts to zero value if not valid. If valid, same as above.
//
// For other types, works as FromString when excepts AsBool. Detect zero value for AsBool.
func FromAny(i interface{}) Conv {
	switch v := i.(type) {
	case int:
		return FromInt(v)
	case *int:
		return FromInt(*v)

	case int8:
		return FromInt64(int64(v))
	case int16:
		return FromInt64(int64(v))
	case int32:
		return FromInt64(int64(v))
	case int64:
		return FromInt64(int64(v))

	case *int8:
		return FromInt64(int64(*v))
	case *int16:
		return FromInt64(int64(*v))
	case *int32:
		return FromInt64(int64(*v))
	case *int64:
		return FromInt64(int64(*v))

	case uint:
		return FromUint(v)
	case *uint:
		return FromUint(*v)

	case uint8:
		return FromUint64(uint64(v))
	case uint16:
		return FromUint64(uint64(v))
	case uint32:
		return FromUint64(uint64(v))
	case uint64:
		return FromUint64(uint64(v))

	case *uint8:
		return FromUint64(uint64(*v))
	case *uint16:
		return FromUint64(uint64(*v))
	case *uint32:
		return FromUint64(uint64(*v))
	case *uint64:
		return FromUint64(uint64(*v))

	case float32:
		return FromFloat(float64(v))
	case float64:
		return FromFloat(float64(v))

	case *float32:
		return FromFloat(float64(*v))
	case *float64:
		return FromFloat(float64(*v))

	case bool:
		return FromBool(v)
	case *bool:
		return FromBool(*v)

	case []rune:
		return FromString(string(v))
	case string:
		return FromString(v)
	case *string:
		return FromString(*v)

	case sql.NullInt64:
		if v.Valid {
			return FromInt(v.Int64)
		}
		return FromInt(0)
	case *sql.NullInt64:
		if v.Valid {
			return FromInt(v.Int64)
		}
		return FromInt(0)
	case sql.NullFloat64:
		if v.Valid {
			return FromFloat(v.Float64)
		}
		return FromFloat(0)
	case *sql.NullFloat64:
		if v.Valid {
			return FromFloat(v.Float64)
		}
		return FromFloat(0)
	case sql.NullString:
		if v.Valid {
			return FromString(v.String)
		}
		return FromString("")
	case *sql.NullString:
		if v.Valid {
			return FromString(v.String)
		}
		return FromString("")
	case sql.NullBool:
		return FromBool(v.Bool && v.Valid)
	case *sql.NullBool:
		return FromBool(v.Bool && v.Valid)

	case fmt.Stringer:
		return FromString(v.String())
	case fmt.GoStringer:
		return FromString(v.GoString())
	case error:
		return FromString(v.Error())
	}
	return perper{i}
}

type perper struct {
	i interface{}
}

// AsInt converts anything to int64.
// It will convert to string first.
func (p perper) AsInt() int64 {
	return FromString(fmt.Sprint(p.i)).AsInt()
}

// AsUint converts anything to uint64.
// It will convert to string first.
func (p perper) AsUint() uint64 {
	return FromString(fmt.Sprint(p.i)).AsUint()
}

// AsFloat converts anything to float64.
// It will convert to string first.
func (p perper) AsFloat() float64 {
	return FromString(fmt.Sprint(p.i)).AsFloat()
}

// AsString converts anything to string.
func (p perper) AsString() string {
	return fmt.Sprint(p.i)
}

// AsBool converts anything to bool.
// For primitive types, just same as other converter.
// For others, compares to nil/zero value.
func (p perper) AsBool() bool {
	if p.i == nil || p.i == reflect.Zero(reflect.TypeOf(p.i)).Interface() {
		return false
	}

	return true
}

func (b perper) IsBool() bool   { return false }
func (b perper) IsInt() bool    { return false }
func (b perper) IsUint() bool   { return false }
func (b perper) IsFloat() bool  { return false }
func (b perper) IsString() bool { return false }
