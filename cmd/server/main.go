package main

import (
	"net/http"
	"os"

	"github.com/HasegawaShuhei/golang_sample_api/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.OpenDB(os.Getenv("DRIVER"), os.Getenv("DSN"))
	defer database.CloseDB(db)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.Run()
}
