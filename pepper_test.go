package pepper

import (
	"testing"
	"os"
)

func TestNew(t *testing.T) {
	prefix := &Prefix{
		FileName: true,
		PackageName: true,
		FunctionName: true,
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

	p.Debug(config)
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


func TestNewDefaultLevel(t *testing.T) {
	spec := NewDefaultLevel(7)

	type obj struct{
		Message string
		Data string
	}
	spec.System("This is a system message")
	spec.Alert("This is an alert message")
	spec.Critical("This is an critical message")
	spec.Error("This is an error message")
	spec.Warning("This is a warning message")
	spec.Notice("This is a notice message")
	spec.Info("Just is an info message")
	spec.Debug("This is a debug message")

	spec.Info(obj{"asdf", "wert"}, "another message")
}