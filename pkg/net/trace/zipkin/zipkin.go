package zipkin

import (
	"github.com/opentracing/opentracing-go"
)

// 第一步: 开一个全局变量
var zkTracer opentracing.Tracer
