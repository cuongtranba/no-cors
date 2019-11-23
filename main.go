package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	jsoniter "github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "no-cors")
	})
	request := gorequest.New().Timeout(time.Second * 4)

	r.Any("/do", func(c *gin.Context) {
		url := c.Query("url")
		switch c.Request.Method {
		case "POST":
			return
		case "GET":
			_, body, err := request.Get(url).End()
			if err != nil {
				c.JSON(http.StatusBadRequest, body)
				return
			}
			c.JSON(http.StatusOK, body)
			return
		case "PUT":
			return
		case "DELETE":
			return
		default:
			c.JSON(http.StatusBadRequest, nil)
		}
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(os.Getenv("PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
