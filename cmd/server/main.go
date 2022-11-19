package main

import (
	"net/http"
	"os"

	"github.com/HasegawaShuhei/golang_sample_api/config"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.OpenDB(os.Getenv("DRIVER"), os.Getenv("DSN"))
	defer config.CloseDB(db)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.Run()
}
