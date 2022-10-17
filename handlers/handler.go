package handlers

import (
	"mgmt/api/mgmt"
	"log"

	"github.com/gin-gonic/gin"
)

var err error

func Handler() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	if gin.Mode() == "debug" {
		r.Use(gin.Logger())
	}

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	err = r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

  mgmt.Routes(r) // Add gin Engine to api/mgmt.go

	return r
}
