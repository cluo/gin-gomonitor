package gomonitor

import (
	"fmt"

	"github.com/gin-gonic/gin"
	mon "gopkg.in/mcuadros/go-monitor.v1"
	"gopkg.in/mcuadros/go-monitor.v1/aspects"
)

// Metrics exposes metrics of the famous
// https://github.com/gopkg.in/mcuadros/go-monitor.v1 package to your
// [gin](https://github.com/gin-gonic/gin) based webapp.  It supports
// custom metrics which are not triggered using the middleware
// context. If you want to add a page counter please see the example.
// Metrics() get a port to expose monitoring data to and a slice of
// aspects.Aspect defined by the user.
//
// Example:
//    package main
//
//    import (
//    	"net/http"
//
//    	"github.com/gin-gonic/gin"
//    	"github.com/zalando-techmonkeys/gin-gomonitor"
//	"github.com/zalando-techmonkeys/gin-gomonitor/aspects"
//    	"gopkg.in/mcuadros/go-monitor.v1/aspects"
//    )
//
//    func main() {
//    	router := gin.New()
//    	counterAspect := &ginmon.CounterAspect{0}
//    	asps := []aspects.Aspect{counterAspect}
//    	// curl http://localhost:9000/Counter
//    	router.Use(ginmon.CounterHandler(counterAspect))
//    	// curl http://localhost:9000/
//    	router.Use(gomonitor.Metrics(9000, asps))
//    	// last middleware
//    	router.Use(gin.Recovery())
//
//    	// each request to all handlers like below will increment the Counter
//    	router.GET("/", func(ctx *gin.Context) {
//    		ctx.JSON(http.StatusOK, gin.H{"title": "Counter - Hello World"})
//    	})
//
//    	//..
//    	router.Run(":8080")
//    }
func Metrics(port int, asps []aspects.Aspect) gin.HandlerFunc {
	var monitor *mon.Monitor = mon.NewMonitor(fmt.Sprintf(":%d", port))
	for _, aspect := range asps {
		monitor.AddAspect(aspect)
	}

	go monitor.Start()
	return func(c *gin.Context) {
	}
}
