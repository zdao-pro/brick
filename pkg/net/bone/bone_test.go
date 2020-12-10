package bone

import (
	"brick/pkg/log"
	"testing"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func cli(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func TestApollo(t *testing.T) {
	log.Init(nil)
	e := Init()
	e.GET("/ping", ping)
	e.GET("/cli", cli)
	e.Run()
}
