package option_test

import (
	"fmt"
	"testing"

	"github.com/PeerXu/option-go/v2"
)

var WithName, GetName = option.LightNew[string]()

func exampleGreeting(aos ...option.ApplyOption) (string, error) {
	opts := option.Apply(aos...)
	name, err := GetName(opts)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Hello, %s", name), nil
}

func TestGreeting(t *testing.T) {
	message, err := exampleGreeting(WithName("World"))
	if err != nil {
		t.Fatal(err)
	}
	if message != "Hello, World" {
		t.Fatalf("message not equal Hello, World")
	}
}

func Example() {
	var WithName, GetName = option.LightNew[string]()

	greeting := func(aos ...option.ApplyOption) (string, error) {
		opts := option.Apply(aos...)
		name, err := GetName(opts)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Hello, %s", name), nil
	}

	msg, err := greeting(WithName("World"))
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
	// Output: Hello, World
}

func ExampleNew() {
	key, WithTimeout, GetTimeout := option.New[int]("timeout")

	opts := option.Apply(WithTimeout(30))
	val, err := GetTimeout(opts)
	if err != nil {
		panic(err)
	}
	fmt.Println(key, val)
	// Output: timeout 30
}

func ExampleApplyWithDefault() {
	KEY_CTX, WithVal, GetVal := option.New[string]("ctx")
	defaults := option.NewOption(map[option.OptionKey]any{KEY_CTX: "default"})

	// With defaults only
	opts := option.ApplyWithDefault(defaults)
	val, _ := GetVal(opts)
	fmt.Println(val)

	// With override
	opts = option.ApplyWithDefault(defaults, WithVal("override"))
	val, _ = GetVal(opts)
	fmt.Println(val)
	// Output:
	// default
	// override
}
