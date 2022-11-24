package main

import (
	"github.com/HasegawaShuhei/golang_sample_api/controller"
	"github.com/HasegawaShuhei/golang_sample_api/database"
	"github.com/HasegawaShuhei/golang_sample_api/middleware"
	"github.com/HasegawaShuhei/golang_sample_api/repository"
	"github.com/HasegawaShuhei/golang_sample_api/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = database.OpenDB()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer database.CloseDB(db)

	r := gin.Default()
	authRoutes := r.Group("v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("v1/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}
	r.Run()
}
