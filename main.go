package main

import (
	"net/http"

	"golang-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// HelloWorldExample godoc
// @Summary hello World
// @Schemes
// @Description Hello
// @Tags Hello World
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /hello [get]
func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func main() {
	route := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := route.Group("/api/v1")
	{
		eg := v1.Group("/")
		{
			eg.GET("/hello", HelloWorld)
		}
	}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	route.Run(":8080")
}