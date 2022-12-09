package option_test

import (
	"testing"

	"github.com/PeerXu/option-go"
)

const OPTION_VAL = "val"

var WithVal, GetVal = option.New[int](OPTION_VAL)

func f(opts ...option.ApplyOption) (int, error) {
	o := option.Apply(opts...)
	return GetVal(o)
}

func TestOption(t *testing.T) {
	val, err := f(WithVal(1))
	if err != nil {
		t.Fatal(err)
	}
	if val != 1 {
		t.Fatalf("val not equal 1")
	}
}
