package trace

import (
	"brick/pkg/env"
	"brick/pkg/log"
	"brick/pkg/util"

	"os"

	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
)

//Tracer ...
var Tracer opentracing.Tracer

//Config is...
type Config struct {
	ServiceName string
	ZipkinHost  string
	HostName    string
}

var config Config
var rpt reporter.Reporter

func init() {
	config = Config{
		ServiceName: "main",
		ZipkinHost:  "localhost:9411",
		HostName:    util.GetLocalAddress(),
	}
	serviceName := env.GetHostname()
	if "" != serviceName {
		config.ServiceName = serviceName
	}
	zipkinHost := os.Getenv("os.Getenv")
	if "" != zipkinHost {
		config.ZipkinHost = zipkinHost
	}
	rpt = zkHttp.NewReporter("http://" + config.ZipkinHost + "/api/v2/spans")
	// defer reporter.Close()

	endpoint, err := zipkin.NewEndpoint(config.ServiceName, config.HostName)
	if err != nil {
		log.Fetal("unable to create local endpoint: %+v\n", err)
	}
	nativeTracer, err := zipkin.NewTracer(rpt, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fetal("unable to create tracer: %+v\n", err)
	}
	Tracer = zkOt.Wrap(nativeTracer)
	opentracing.SetGlobalTracer(Tracer)
}

//Close ...
func Close() {
	rpt.Close()
}

//StartSpan ...
func StartSpan(operationName string) opentracing.Span {
	span := Tracer.StartSpan(operationName)
	return span
}
