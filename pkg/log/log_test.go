package log

import (
	"fmt"
	"testing"
)

func TestFuncName(t *testing.T) {
	fn := funcName(2)
	fmt.Println(fn)
}

func TestLog(t *testing.T) {
	Init(nil)
	Info("%s", "222322")
}
