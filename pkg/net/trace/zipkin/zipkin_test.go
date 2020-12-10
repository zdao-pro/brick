package zipkin

import (
	"log"
	"testing"
	"time"

	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func TestZipkin(t *testing.T) {
	{
		reporter := zkHttp.NewReporter("http://localhost:9411/api/v2/spans")
		defer reporter.Close()
		endpoint, err := zipkin.NewEndpoint("main3", "localhost")
		if err != nil {
			log.Fatalf("unable to create local endpoint: %+v\n", err)
		}
		nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		if err != nil {
			log.Fatalf("unable to create tracer: %+v\n", err)
		}
		zkTracer = zkOt.Wrap(nativeTracer)
		opentracing.SetGlobalTracer(zkTracer)

		span := zkTracer.StartSpan("zipkin_test")
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
}
