package service

import (
	"fmt"
	"gg/server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (u UserService) Registration(c *gin.Context) {
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
	if err := u.DB.Create(&user).Error; err != nil {
		c.String(500, fmt.Sprintf("Не удалось создать пользователя: %v", err))
		return
	}

	c.SetCookie("login", login, 1200, "/", "localhost", false, false)
	c.SetCookie("password", password, 1200, "/", "localhost", false, false)
	// Куки - данные, которые отсылаем сервером на браузер. Путь, домен на котором работает печеньки, секьюр и httponly для https 

	c.Redirect(302, "/catalog")
}

func (u UserService) Login(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	var user models.User
	if err := u.DB.Where("login = ? and password = ?", login, password).First(&user).Error; err != nil {
		c.String(401, "Неверный логин или пароль")
		return
	}

	c.SetCookie("login", login, 1200, "/", "localhost", false, false)
	c.SetCookie("password", password, 1200, "/", "localhost", false, false)

	c.Redirect(302, "/catalog")
}

func (u UserService) Agreement(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")
	name := c.PostForm("name")

	redirectURL := fmt.Sprintf("/agreement?login=%s&password=%s&name=%s", login, password, name)
	// ? для параметров. На ссылке страницы агримент видно
	

	c.Redirect(302, redirectURL)
}
