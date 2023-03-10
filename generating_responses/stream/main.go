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

	router.GET("/tale_of_two_cities", func(c *gin.Context) {
		f, err := os.Open("./a_tale_of_two_cities.txt")
		defer f.Close()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.Stream(streamer(f))
	})

	log.Fatal(router.Run(":3000"))
}

func streamer(r io.Reader) func(io.Writer) bool {
	return func(step io.Writer) bool {
		for {
			buf := make([]byte, 4*2^10)
			if _, err := r.Read(buf); err == nil {
				_, err := step.Write(buf)
				return err == nil
			} else {
				return false
			}
		}
	}
}
