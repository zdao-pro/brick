package env

import (
	"fmt"
	"testing"
)

func TestEnv(t *testing.T) {
	a := GetAppID()
	fmt.Println(a)

	e := GetEnv()
	fmt.Println(e)

	h := GetHostname()
	fmt.Println(h)
}
