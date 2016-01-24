package main

import (
	"github.com/sergystepanov/bee1/Godeps/_workspace/src/github.com/gin-gonic/contrib/gzip"
	"github.com/sergystepanov/bee1/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	gin.SetMode(gin.ReleaseMode)
	
	app := gin.Default()
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Go away.")
	})

	app.POST("/url", func(c *gin.Context) {
		url := c.PostForm("_url")

		log.Print(url)
		resp, err := http.Get(url)
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		c.Data(http.StatusOK, "text/html; charset=windows-1251", body)
	})

	app.Run(":" + port)
}
