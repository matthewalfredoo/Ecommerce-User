package routes

import "github.com/gin-gonic/gin"

//StartGin function
func StartGin() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/users")
	}
}
