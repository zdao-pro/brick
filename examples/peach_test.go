package peach

import (
	"brick/pkg/peach"
	"brick/pkg/peach/apollo"
	"fmt"
	"testing"
)

func TestApollo(t *testing.T) {
	peach.Init("/", apollo.PeachDriverApollo)
	a, _ := peach.Get("test").String()
	fmt.Println(a)
}
