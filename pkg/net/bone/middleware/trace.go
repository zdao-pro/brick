package middleware

import (
	"brick/pkg/net/trace"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

//GetTracer ...
func GetTracer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var s opentracing.Span
		wireContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header))
		if nil != err {
			s = trace.StartSpan(c.Request.URL.Path)
		} else {
			s = opentracing.StartSpan("zip", opentracing.ChildOf(wireContext))
		}

		c.Next()
		s.Finish()
	}
}
