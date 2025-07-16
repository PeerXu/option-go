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
