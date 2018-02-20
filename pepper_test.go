package pepper

import (
	"testing"
	"os"
)

func TestNew(t *testing.T) {
	prefix := &Prefix{
		FileName: true,
		PackageName: false,
		FunctionName: false,
		LineNumber: true,
	}

	config := &Config{
		Prefix: prefix,
		Output: os.Stdout,
	}

	type obj struct{
		Message string
		Data string
	}
	var asdf = obj{
		"this is just a test obj message.",
		"here is some data",
	}

	p := New(config)
	if p == nil{
		t.Error("Pepper is nil")
	}
	p.Info(asdf)
	p.Info("Just is an info message")
	p.Error("This is an error message")
	p.Debug("This is a debug message")

}

func TestNewDefault(t *testing.T) {
	spec := NewDefault()

	type obj struct{
		Message string
		Data string
	}

	spec.Info("Just is an info message")
	spec.Error("This is an error message")
	spec.Debug("This is a debug message")
	spec.Info(obj{"asdf", "wert"}, "another message")
}