package main

import (
	"github.com/HasegawaShuhei/golang_sample_api/controller"
	"github.com/HasegawaShuhei/golang_sample_api/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = database.OpenDB()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer database.CloseDB(db)

	r := gin.Default()
	authRoutes := r.Group("v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run()
}
