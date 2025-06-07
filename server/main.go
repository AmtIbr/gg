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
	// Подключение к базе данных

	if err != nil {
		panic("failed to connect database")
	}
	// Проверяем подключение к БД

	db.AutoMigrate(&models.User{}, &models.Cart{}, &models.Forum{}, &models.Product{}, &models.Message{})
	// Создание в БД таблички

	r := gin.Default()
	// Инициализируем веб сервер

	htmlService := service.HtmlService{DB: db}
	// Создаем htmlservice с подключением к бд
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
	r.GET("/agreement", htmlService.Agreement)
	r.GET("/about", htmlService.About)
	r.GET("/home", htmlService.Home)
	r.GET("/message", htmlService.Message)
	r.GET("/pereputie", htmlService.Pereputie)
	r.GET("/tovar1", htmlService.Tovar1)
	r.GET("/tovar2", htmlService.Tovar2)
	r.GET("/tovar3", htmlService.Tovar3)
	r.GET("/tovar4", htmlService.Tovar4)
	r.GET("/tovar5", htmlService.Tovar5)

	r.POST("/test", userService.Registration)
	r.POST("/login", userService.Login)
	r.POST("/cart/add", cartService.Add)
	r.POST("/forum", forumService.AddPost)
	r.POST("/agreement", userService.Agreement)
	r.POST("/message", htmlService.SubmitMessage)

	r.Run()
}
