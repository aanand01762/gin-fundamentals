package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFile("/", "./index.html")

	//sending files
	//c.File("./a_tale_of_two_cities.txt")

	// fs := gin.Dir("/path", true)
	// c.FileFromDS("filename", fs)

	//c.FileAttachment("/path", "name after download")

	router.GET("/tale_of_two_cities", func(c *gin.Context) {
		//c.File("./a_tale_of_two_cities.txt")
		f, err := os.Open("./a_tale_of_two_cities.txt")
		defer f.Close()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		data, err := io.ReadAll(f)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.Data(http.StatusOK, "text/plain", data)

	})

	router.GET("/great_expectations", func(c *gin.Context) {
		f, err := os.Open("./great_expectations.txt")
		defer f.Close()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		fi, err := f.Stat()

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.DataFromReader(
			http.StatusOK,
			fi.Size(),
			"text/plain",
			f,
			map[string]string{
				"Content-Deposition": "attachment;filename=/great_expectations.txt",
			})
	})
	log.Fatal(router.Run(":3000"))
}
