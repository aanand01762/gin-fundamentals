package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/params/*rest", func(c *gin.Context) {
		url := c.Request.URL.String()
		headers := c.Request.Header
		cookies := c.Request.Cookies()
		c.IndentedJSON(http.StatusOK, gin.H{
			"url":     url,
			"headers": headers,
			"cookies": cookies,
		})
	})

	router.GET("/query/*rest", func(ctx *gin.Context) {
		username := ctx.Query("username")
		year := ctx.DefaultQuery("year", "2023")
		months := ctx.QueryArray("month")

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"username": username,
			"year":     year,
			"months":   months,
		})
	})

	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	type TimeoffRequest struct {
		Date   time.Time `json:"date" form:"date" binding:"-" time_format:"2006-01-02"`
		Amount float64   `json:"amount" form:"amount" binding:"-"`
	}

	router.POST("employee", func(c *gin.Context) {
		var timeoff TimeoffRequest
		if err := c.ShouldBind(&timeoff); err == nil {
			c.IndentedJSON(http.StatusOK, timeoff)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
		/*
			date := c.PostForm("date")
			amount := c.PostForm("amount")
			username := c.DefaultPostForm("user", "john")
			c.IndentedJSON(http.StatusOK, gin.H{
				"username": username,
				"amount":   amount,
				"date":     date,
			})*/

	})

	apiGroup := router.Group("api")

	apiGroup.POST("timeoff", func(c *gin.Context) {
		var timeoff TimeoffRequest
		if err := c.ShouldBind(&timeoff); err == nil {
			c.IndentedJSON(http.StatusOK, timeoff)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}

	})
	log.Fatal(router.Run(":3000"))
}
