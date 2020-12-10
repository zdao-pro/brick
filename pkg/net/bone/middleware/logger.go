package middleware

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"brick/pkg/log"

	"github.com/gin-gonic/gin"
)

//LogConfig log configure
type LogConfig struct {
	OutPut io.Writer
}

//ParamData ...
type ParamData struct {
	Request  *http.Request
	TimeNow  time.Time
	Latency  time.Duration
	HTTPCode int
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string
}

func render(param ParamData) string {
	return fmt.Sprintf("%v %v %d %s %s %s %s", param.TimeNow, param.Latency.Milliseconds(), param.HTTPCode, param.ClientIP, param.Method, param.Path, param.ErrorMessage)
}

//GetLogger retunrn access logger
func GetLogger(conf LogConfig) gin.HandlerFunc {
	out := conf.OutPut
	if nil != out {
		gin.DefaultWriter = out
	}

	return func(c *gin.Context) {
		param := ParamData{
			Request: c.Request,
		}

		start := time.Now()

		c.Next()

		param.TimeNow = time.Now()
		param.Latency = param.TimeNow.Sub(start)
		param.Path = c.Request.URL.Path
		param.Method = c.Request.Method
		param.HTTPCode = c.Writer.Status()
		param.ClientIP = c.ClientIP()
		log.Access(render(param))
	}
}
