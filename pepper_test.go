package pepper

import (
	"testing"
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

	}

	p := New(config)
	if p == nil{
		t.Error("Pepper is nil")
	}
	p.Info("Just is an info message")
	p.Error("This is an error message")
	p.Debug("This is a debug message")

}