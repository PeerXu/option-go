package option_test

import (
	"errors"
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

func TestWithDefaultValue(t *testing.T) {
	KEY_CTX, WithVal, GetVal := option.New[string]("test-key")
	defaultValue := func() option.Option {
		return option.NewOption(map[option.OptionKey]any{
			KEY_CTX: "default",
		})
	}

	result := option.ApplyWithDefault(defaultValue())
	if result == nil {
		t.Fatal("ApplyWithDefault returned nil")
	}

	val, err := GetVal(result)
	if err != nil {
		t.Fatal(err)
	}
	if val != "default" {
		t.Fatalf("val not equal 'default', got %s", val)
	}

	val2, err := GetVal(option.ApplyWithDefault(defaultValue(), WithVal("provided")))
	if err != nil {
		t.Fatal(err)
	}
	if val2 != "provided" {
		t.Fatalf("val not equal 'provided', got %s", val2)
	}
}

func TestGetMissingOption(t *testing.T) {
	_, _, GetVal := option.New[int]("missing")
	opts := option.Apply()
	_, err := GetVal(opts)
	if err == nil {
		t.Fatal("expected error for missing option, got nil")
	}
	if !errors.Is(err, option.ErrOptionRequired) {
		t.Fatalf("expected ErrOptionRequired, got %v", err)
	}
}

func TestGetUnexpectedType(t *testing.T) {
	_, WithInt, GetInt := option.New[int]("type-mismatch")
	_, _, GetString := option.New[string]("type-mismatch")
	opts := option.Apply(WithInt(42))
	_, err := GetString(opts)
	if err == nil {
		t.Fatal("expected error for type mismatch, got nil")
	}
	if !errors.Is(err, option.ErrUnexpectedType) {
		t.Fatalf("expected ErrUnexpectedType, got %v", err)
	}
	_, err = GetInt(opts)
	if err != nil {
		t.Fatalf("expected no error for correct type, got %v", err)
	}
}

func TestMultipleOptionsSameType(t *testing.T) {
	WithA, GetA := option.LightNew[int]()
	WithB, GetB := option.LightNew[int]()

	opts := option.Apply(WithA(10), WithB(20))
	a, err := GetA(opts)
	if err != nil {
		t.Fatal(err)
	}
	b, err := GetB(opts)
	if err != nil {
		t.Fatal(err)
	}
	if a != 10 {
		t.Fatalf("expected 10, got %d", a)
	}
	if b != 20 {
		t.Fatalf("expected 20, got %d", b)
	}
}

func TestMustPanicsOnMissingOption(t *testing.T) {
	_, MustGetVal := option.LightMustNew[int]()
	opts := option.Apply()

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic for missing option, got none")
		}
		err, ok := r.(error)
		if !ok {
			t.Fatalf("expected error from panic, got %T", r)
		}
		if !errors.Is(err, option.ErrOptionRequired) {
			t.Fatalf("expected ErrOptionRequired, got %v", err)
		}
	}()

	MustGetVal(opts)
}

func TestNewOptionNoAliasing(t *testing.T) {
	orig := map[option.OptionKey]any{}
	key, WithVal, _ := option.New[string]("alias-test")

	orig[key] = "original"
	copied := option.NewOption(orig)
	WithVal("modified")(copied)

	val := orig[key].(string)
	if val != "original" {
		t.Fatalf("NewOption aliased the input map: expected 'original', got %q", val)
	}
}
