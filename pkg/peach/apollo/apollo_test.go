package apollo

import (
	"fmt"
	"testing"
)

func TestEnv(t *testing.T) {
	var d appoloDriver = appoloDriver{}
	c, err := d.New()
	if nil != err {
		fmt.Print("vss")
	}

	fmt.Println(c)
	v := c.Get("test")
	if nil == v {
		fmt.Print("v")
	}

}
