package arquitecture

import (
	"runtime"
	"testing"
)

func TestDepends(t *testing.T){
	t.Parallel() //Can be executed in a parallel form.
	if runtime.GOARCH == "amd64"{ //Can't be executed in amd64
		t.Skip("Not working with amd64")
	}
	t.Fail()
}
