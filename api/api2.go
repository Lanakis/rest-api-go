package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(middlewareCORS)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	// Example routes
	router.GET("/app/streams", handleAppStreams)
	router.GET("/streams", handleStreams)

	// Start server
	go func() {
		if err := router.Run(":2325"); err != nil {
			panic(err)
		}
	}()
	return router
}

func handleAppStreams(c *gin.Context) {
	// Your logic for "/app/streams" route
	c.JSON(http.StatusOK, gin.H{"message": "Hello from app/streams"})
}

func handleStreams(c *gin.Context) {
	// Your logic for "/streams" route
	c.JSON(http.StatusOK, gin.H{"message": "Hello from /streams"})
}

func middlewareCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
	c.Next()
}
