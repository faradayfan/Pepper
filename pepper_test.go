package pepper

import (
	"testing"
)

func TestNew(t *testing.T) {
	p := New()
	if p == nil{
		t.Error("Pepper is nil")
	}
	//p.Info("Just is an info message")
	//p.Error("This is an error message")
	//p.Debug("This is a debug message")

}