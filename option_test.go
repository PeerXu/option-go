package option_test

import (
	"testing"

	"github.com/PeerXu/option-go/v2"
)

func TestLightNew(t *testing.T) {
	var WithVal1, GetVal1 = option.LightNew[int]()

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

func TestLightMustNew(t *testing.T) {
	var WithVal2, MustGetVal2 = option.LightMustNew[string]("val2")

	opts := []option.ApplyOption{WithVal2("hello")}
	ctx := option.Apply(opts...)
	val := MustGetVal2(ctx)
	if val != "hello" {
		t.Fatalf("val not equal 1")
	}
}

func TestNew(t *testing.T) {
	// Test with no key provided
	_, WithVal3, GetVal3 := option.New[float64]()
	
	opts := []option.ApplyOption{WithVal3(3.14)}
	ctx := option.Apply(opts...)
	
	// Test getter functionality
	val, err := GetVal3(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if val != 3.14 {
		t.Fatalf("val not equal 3.14, got %v", val)
	}
	
	// Test with explicit key
	key2, WithVal4, GetVal4 := option.New[[]string]("string-slice")
	testSlice := []string{"a", "b", "c"}
	
	opts2 := []option.ApplyOption{WithVal4(testSlice)}
	ctx2 := option.Apply(opts2...)
	
	// Verify key is set correctly
	if key2.String() != "string-slice" {
		t.Fatalf("key not set correctly, expected 'string-slice', got %v", key2.String())
	}
	
	// Test getter functionality with explicit key
	val2, err := GetVal4(ctx2)
	if err != nil {
		t.Fatal(err)
	}
	if len(val2) != len(testSlice) {
		t.Fatalf("slice length mismatch, expected %d, got %d", len(testSlice), len(val2))
	}
}

func TestNewMust(t *testing.T) {
	// Test with no key provided
	_, WithVal5, MustGetVal5 := option.NewMust[bool]()
	
	opts := []option.ApplyOption{WithVal5(true)}
	ctx := option.Apply(opts...)
	
	// Test must getter functionality
	val := MustGetVal5(ctx)
	if !val {
		t.Fatalf("val not equal true")
	}
	
	// Test with explicit key
	key2, WithVal6, MustGetVal6 := option.NewMust[map[string]int]("map-key")
	testMap := map[string]int{"one": 1, "two": 2}
	
	opts2 := []option.ApplyOption{WithVal6(testMap)}
	ctx2 := option.Apply(opts2...)
	
	// Verify key is set correctly
	if key2.String() != "map-key" {
		t.Fatalf("key not set correctly, expected 'map-key', got %v", key2.String())
	}
	
	// Test must getter functionality with explicit key
	val2 := MustGetVal6(ctx2)
	if len(val2) != len(testMap) {
		t.Fatalf("map size mismatch, expected %d, got %d", len(testMap), len(val2))
	}
	if val2["one"] != 1 {
		t.Fatalf("map value mismatch, expected 1 for key 'one', got %d", val2["one"])
	}
}
