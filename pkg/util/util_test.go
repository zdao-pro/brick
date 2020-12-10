package util

import (
	"fmt"
	"testing"
)

func TestFuncName(t *testing.T) {
	fmt.Println(GetFuncName())
}

func TestLocalAddress(t *testing.T) {
	fmt.Println(GetLocalAddress())
}
