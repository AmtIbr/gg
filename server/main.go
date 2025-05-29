package main

import (
	"gg/server/models"
	"gg/server/service"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Cart{}, &models.Forum{}, &models.Product{})

	r := gin.Default()

	htmlService := service.HtmlService{DB: db}
	userService := service.UserService{DB: db}
	cartService := service.CartService{DB: db}
	forumService := service.ForumService{DB: db}

	r.StaticFile("/static/css/style.css", "server/css/style.css")

	r.LoadHTMLGlob("html/*.html")

	r.GET("/login", htmlService.Login)
	r.GET("/regist", htmlService.Registration)
	r.GET("/catalog", htmlService.Catalog)
	r.GET("/cart", htmlService.Cart)
	r.GET("/forum/:theme", htmlService.Forum)

	r.POST("/test", userService.Registration)
	r.POST("/login", userService.Login)
	r.POST("/cart/add", cartService.Add)
	r.POST("/forum/:theme", forumService.AddPost)

	r.Run()
}
