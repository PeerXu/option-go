package option_test

import (
	"testing"

	"github.com/PeerXu/option-go"
)

const OPTION_VAL = "val"

var WithVal, GetVal = option.New[int](OPTION_VAL)
var _, MustGetVal = option.NewMust[int](OPTION_VAL)

func f(opts ...option.ApplyOption) (int, error) {
	o := option.Apply(opts...)
	return GetVal(o)
}

func TestOption(t *testing.T) {
	opts := []option.ApplyOption{WithVal(1)}
	ctx := option.Apply(opts...)
	val, err := GetVal(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if val != 1 {
		t.Fatalf("val not equal 1")
	}
}

func TestMust(t *testing.T) {
	opts := []option.ApplyOption{WithVal(1)}
	ctx := option.Apply(opts...)
	val := MustGetVal(ctx)
	if val != 1 {
		t.Fatalf("val not equal 1")
	}
}

func TestMustPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("failed to recover panic")
		}
	}()

	o := option.Apply()
	MustGetVal(o)
}
