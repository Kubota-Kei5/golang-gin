package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func testResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

func slowHandler(c *gin.Context) {
	time.Sleep(800 * time.Millisecond)
	c.Status(http.StatusOK)
}

func main() {
	r := gin.New()

	r.GET("/slow", timeout.New(
		timeout.WithTimeout(500*time.Millisecond),
		timeout.WithResponse(testResponse),
	), slowHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}