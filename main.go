package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafialariq/go-auth/controllers"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", controllers.Login)

	r.GET("/register", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "register.html", nil)
	})

	r.POST("/register", controllers.Register)

	r.GET("/", controllers.Index)

	r.GET("/logout", controllers.Logout)

	fmt.Println("server berjalan di localhost:8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
