package option_test

import (
	"testing"

	"github.com/PeerXu/option-go/v2"
)

var OPTION_VAL1, WithVal1, GetVal1 = option.New[int]("val1")
var OPTION_VAL2, WithVal2, MustGetVal2 = option.NewMust[int]("val2")

func TestOption(t *testing.T) {
	opts := []option.ApplyOption{WithVal1(1)}
	ctx := option.Apply(opts...)
	val, err := GetVal1(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if val != 1 {
		t.Fatalf("val not equal 1")
	}
}

func TestMust(t *testing.T) {
	opts := []option.ApplyOption{WithVal2(1)}
	ctx := option.Apply(opts...)
	val := MustGetVal2(ctx)
	if val != 1 {
		t.Fatalf("val not equal 1")
	}
}
