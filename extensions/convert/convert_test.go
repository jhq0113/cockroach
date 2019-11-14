package convert

import (
	"testing"
)

func TestString2Int(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{
			"123",
			123,
		},
		{
			"12Hello",
			0,
		},
		{
			"Hello",
			0,
		},
	}

	for _, ca := range cases {
		got := String2Int(ca.in)
		if got != ca.want {
			t.Errorf("String2Int(%q) == %d, want %d", ca.in, got, ca.want)
		}
	}
}

func TestString2Int64(t *testing.T) {
	cases := []struct {
		in   string
		want int64
	}{
		{
			"123",
			123,
		},
		{
			"12Hello",
			0,
		},
		{
			"Hello",
			0,
		},
	}

	for _, ca := range cases {
		got := String2Int64(ca.in)
		if got != ca.want {
			t.Errorf("String2Int64(%q) == %d, want %d", ca.in, got, ca.want)
		}
	}
}

func TestInt2String(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{
			123,
			"123",
		},
		{
			0,
			"0",
		},
		{
			-1,
			"-1",
		},
	}

	for _, ca := range cases {
		got := Int2String(ca.in)
		if got != ca.want {
			t.Errorf("Int2String(%d) == %q, want %q", ca.in, got, ca.want)
		}
	}
}

func TestInt64ToString(t *testing.T) {
	cases := []struct {
		in   int64
		want string
	}{
		{
			123,
			"123",
		},
		{
			0,
			"0",
		},
		{
			-1,
			"-1",
		},
	}

	for _, ca := range cases {
		got := Int64ToString(ca.in)
		if got != ca.want {
			t.Errorf("Int2String(%d) == %q, want %q", ca.in, got, ca.want)
		}
	}
}
