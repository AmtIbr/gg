package main

import (
	"fmt"
	"gg/server/models"
	"strconv"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Cart{}, &models.Forum{})

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

	// r.POST("/test", func(c *gin.Context) {
	// 	login := c.PostForm("login")
	// 	name := c.PostForm("name")
	// 	password := c.PostForm("password")
	// 	check := c.PostForm("check")
	// 	db.Create(&models.User{
	// 		Check:    check,
	// 		Login:    login,
	// 		Name:     name,
	// 		Password: password,
	// 	})
	// 	c.String(200, fmt.Sprintf("Спасибо за регистрацию!"))
	// })

	r.POST("/test", func(c *gin.Context) {
		var (
			login    = c.PostForm("login")
			name     = c.PostForm("name")
			password = c.PostForm("password")
			agreed   = c.PostForm("check") == "on"
		)
		if login == "" || password == "" || !agreed {
			c.String(400, "Заполните все поля и согласитесь с правилами")
			return
		}
		user := models.User{Login: login, Name: name, Password: password, Check: agreed}
		if err := db.Create(&user).Error; err != nil {
			c.String(500, fmt.Sprintf("Не удалось создать пользователя: %v", err))
			return
		}

		c.SetCookie("login", login, 1200, "/", "localhost", false, false)
		c.SetCookie("password", password, 1200, "/", "localhost", false, false)

		c.Redirect(302, "/katalog")
	})

	r.POST("/login", func(c *gin.Context) {
		login := c.PostForm("login")
		password := c.PostForm("password")

		var user models.User
		if err := db.Where("login = ? and password = ?", login, password).First(&user).Error; err != nil {
			c.String(401, "Неверный логин или пароль")
			return
		}

		c.SetCookie("login", login, 1200, "/", "localhost", false, false)
		c.SetCookie("password", password, 1200, "/", "localhost", false, false)

		c.Redirect(302, "/katalog")
	})

	r.POST("/cart/add", func(c *gin.Context) {
		login, Error := c.Cookie("login")
		if Error != nil {
			c.Redirect(302, "/Regist")
		}
		password, Error := c.Cookie("password")
		if Error != nil {
			c.Redirect(302, "/Regist")
		}
		var user models.User
		if err := db.Where("login = ? and password = ?", login, password).First(&user); err != nil {
			c.Redirect(302, "/Regist")
		}

		tovar := c.PostForm("tovar")
		priceS := c.PostForm("price")
		price, err := strconv.ParseFloat(priceS, 64)
		if err != nil {
			c.String(401, "Неверный формат цены")
			return
		}

		var cart models.Cart
		if err := db.Where("UserID = ?", user.ID).First(&cart).Error; err != nil {
			cart := models.Cart{UserID: user.ID}
			if err := db.Create(&cart).Error; err != nil {
				c.String(500, fmt.Sprintf("Не удалось создать корзину: %v", err))
				return
			}
		}
		product := models.Product{CartID: cart.ID, Tovar: tovar, Price: price}
		if err := db.Create(&product).Error; err != nil {
			c.String(501, fmt.Sprintf("Не удалось добавить товар: %v", err))
			return
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
