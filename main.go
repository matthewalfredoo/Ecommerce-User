package main

import (
	"Ecommerce-User/conn"
	"Ecommerce-User/controller"
	"Ecommerce-User/repository"
	"Ecommerce-User/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = conn.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer conn.CloseDatabaseConnection(db)
	router := gin.Default()

	routes := router.Group("/api/user")
	{
		routes.POST("/register", userController.Register)
		routes.POST("/login", userController.Login)
		routes.PUT("/profile", userController.Update)
		routes.GET("/profile", userController.Profile)
	}

	err := router.Run("192.168.100.8:8080")
	if err != nil {
		return
	}
}
