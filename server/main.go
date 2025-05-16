package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.StaticFile("/static/css/style.css", "server/css/style.css")

	r.LoadHTMLGlob("html/*.html")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "Login.html", gin.H{
			"message": "pong",
		})
	})

	r.GET("/regist", func(c *gin.Context) {
		c.HTML(200, "Regist.html", gin.H{
			"message": "pong",
		})
	})

	r.GET("/katalog", func(c *gin.Context) {
		c.HTML(200, "verstkaPOKUPKI.html", gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
