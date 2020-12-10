package trace

import (
	"testing"
	"time"

	"github.com/opentracing/opentracing-go"
)

func TestZipkin(t *testing.T) {
	defer Close()
	span := StartSpan("zipkin_test")
	time.Sleep(time.Duration(1) * time.Second)
	spa2 := opentracing.StartSpan("zip", opentracing.ChildOf(span.Context()))
	spa3 := opentracing.StartSpan("zip2", opentracing.FollowsFrom(spa2.Context()))
	time.Sleep(time.Duration(1) * time.Second)
	spa3.Finish()
	time.Sleep(time.Duration(2) * time.Second)
	spa2.Finish()
	time.Sleep(time.Duration(1) * time.Second)
	span.Finish()
}
