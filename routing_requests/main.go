package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed public/*
var f embed.FS

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})

	router.StaticFile("/index", "./public/index.html")
	router.Static("/public", "./public")
	router.StaticFS("/fs", http.FileSystem(http.FS(f)))

	router.GET("/employee", func(ctx *gin.Context) {
		ctx.File("./public/employee.html")
	})

	router.POST("/employee", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "New request posted suceessfully!")
	})

	router.GET("/employees/:username/*rest", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"usename":  ctx.Param("username"),
			"rest":     ctx.Param("rest"),
			"fullpath": ctx.FullPath(),
		})
	})
	adminGroup := router.Group("/admin")
	adminGroup.GET("/users", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "admin page to users")
	})
	adminGroup.GET("/policies", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "admin page to policies")
	})
	adminGroup.GET("/roles", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "admin page to roles")
	})

	log.Fatal(router.Run(":3000"))
}
