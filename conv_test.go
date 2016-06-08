package conv

import (
	"database/sql"
	"testing"
)

func TestConv(t *testing.T) {
	a := int64(-1)
	u1 := uint64(a)
	tbl := []struct {
		src interface{}
		i   int64
		u   uint64
		f   float64
		s   string
		b   bool
	}{
		// bool
		{true, 1, 1, 1, "true", true},
		{false, 0, 0, 0, "false", false},

		// int
		{int(1), 1, 1, 1, "1", true},
		{int(0), 0, 0, 0, "0", false},
		{int(-1), -1, u1, -1, "-1", true},

		// int64
		{int64(1), 1, 1, 1, "1", true},
		{int64(0), 0, 0, 0, "0", false},
		{int64(-1), -1, u1, -1, "-1", true},

		// uint
		{uint(1), 1, 1, 1, "1", true},
		{uint(0), 0, 0, 0, "0", false},

		// uint64
		{uint64(1), 1, 1, 1, "1", true},
		{uint64(0), 0, 0, 0, "0", false},

		// float
		{float64(1.5), 1, 1, 1.5, "1.5", true},
		{float64(0), 0, 0, 0, "0", false},
		{float64(-1.5), -1, u1, -1.5, "-1.5", true},

		// string
		{"1", 1, 1, 1, "1", true},
		{"0", 0, 0, 0, "0", false},
		{"true", 0, 0, 0, "true", true},
		{"false", 0, 0, 0, "false", false},
		{"t", 0, 0, 0, "t", true},
		{"f", 0, 0, 0, "f", false},
		{"1.5", 0, 0, 1.5, "1.5", true},
		{"-1", -1, 0, -1, "-1", true},
		{"Learn Go, instead of just stringing random characters together until it compiles",
			0, 0, 0,
			"Learn Go, instead of just stringing random characters together until it compiles",
			true},
		{"0xfeed", 0, 0, 0, "0xfeed", true},
		{"017", 17, 17, 17, "017", true},
		{"", 0, 0, 0, "", false},

		// sql nullable values
		{sql.NullBool{true, true}, 1, 1, 1, "true", true},
		{sql.NullBool{false, true}, 0, 0, 0, "false", false},
		{sql.NullBool{true, false}, 0, 0, 0, "false", false},
		{sql.NullBool{false, false}, 0, 0, 0, "false", false},
		{sql.NullInt64{1, true}, 1, 1, 1, "1", true},
		{sql.NullInt64{1, false}, 0, 0, 0, "0", false},
		{sql.NullFloat64{1, true}, 1, 1, 1, "1", true},
		{sql.NullFloat64{1, false}, 0, 0, 0, "0", false},
		{sql.NullString{"1", true}, 1, 1, 1, "1", true},
		{sql.NullString{"1", false}, 0, 0, 0, "", false},
	}

	for _, data := range tbl {
		a := FromAny(data.src)

		if actual := a.AsInt(); actual != data.i {
			t.Errorf("%T[%v].AsInt() = %d", data.src, data.src, actual)
		}

		if actual := a.AsUint(); actual != data.u {
			t.Errorf("%T[%v].AsUint() = %d", data.src, data.src, actual)
		}

		if actual := a.AsFloat(); actual != data.f {
			t.Errorf("%T[%v].AsFloat() = %G", data.src, data.src, actual)
		}

		if actual := a.AsString(); actual != data.s {
			t.Errorf("%T[%v].AsString() = %s", data.src, data.src, actual)
		}

		if actual := a.AsBool(); actual != data.b {
			t.Errorf("%T[%v].AsBool() = %t", data.src, data.src, actual)
		}
	}
}
