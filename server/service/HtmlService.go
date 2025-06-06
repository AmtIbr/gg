package service

import (
	"fmt"
	"gg/server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HtmlService struct {
	DB *gorm.DB
}

func (h HtmlService) Tovar1(c *gin.Context) {
	c.HTML(200, "tovar1.html", gin.H{})
}
func (h HtmlService) Tovar2(c *gin.Context) {
	c.HTML(200, "tovar2.html", gin.H{})
}
func (h HtmlService) Tovar3(c *gin.Context) {
	c.HTML(200, "tovar3.html", gin.H{})
}
func (h HtmlService) Tovar4(c *gin.Context) {
	c.HTML(200, "tovar4.html", gin.H{})
}
func (h HtmlService) Tovar5(c *gin.Context) {
	c.HTML(200, "tovar5.html", gin.H{})
}

func (h HtmlService) About(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{})
}

func (h HtmlService) Home(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{})
}

func (h HtmlService) Pereputie(c *gin.Context) {
	c.HTML(200, "pereputie.html", gin.H{})
}

func (h HtmlService) Message(c *gin.Context) {
	c.HTML(200, "message.html", gin.H{})
}

func (h HtmlService) Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func (h HtmlService) Registration(c *gin.Context) {
	c.HTML(200, "registration.html", gin.H{})
}

func (h HtmlService) Agreement(c *gin.Context) {
	login := c.Query("login")
	password := c.Query("password")
	name := c.Query("name")

	if login == "" || password == "" {
		c.String(400, "Заполните все поля")
		return
	}

	c.HTML(200, "agreement.html", gin.H{
		"login":    login,
		"password": password,
		"name":     name,
		// Заход на страницу. Пропихиваем данные на агримент.html
	})
}

func (h HtmlService) Catalog(c *gin.Context) {
	c.HTML(200, "catalog.html", gin.H{})
}

func (s HtmlService) Forum(c *gin.Context) {
	theme := c.Param("theme")
	// Парам - это то, что после форум внутри строки

	allowedThemes := map[string]bool{
		"product-quality": true,
		"offers":          true,
		"complaints":      true,
	}

	if !allowedThemes[theme] {
		c.String(404, "Тема форума не найдена")
		return
	}

	var posts []models.Forum
	if err := s.DB.Order("created_at DESC").Where("theme = ?", theme).Find(&posts).Error; err != nil {
		// Ордер - сортировка, креатед ат по времени создания сообщения
		c.String(500, "Не удалось загрузить посты форума")
		return
	}

	c.HTML(200, "forum.html", gin.H{
		"Posts": posts,
		"Theme": theme,
	})
}

func (h HtmlService) Cart(c *gin.Context) {
	login, err := c.Cookie("login")
	if err != nil {
		c.Redirect(302, "/regist")
		return
	}
	password, err := c.Cookie("password")
	if err != nil {
		c.Redirect(302, "/regist")
		return
	}

	var user models.User
	if err := h.DB.Where("login = ? AND password = ?", login, password).First(&user).Error; err != nil {
		c.String(500, fmt.Sprintf("Ошибка входа: %v", err))
		return
	}

	var cart models.Cart
	if err := h.DB.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
		c.String(500, "Корзина не найдена")
		return
	}

	var products []models.Product
	if err := h.DB.Where("cart_id = ?", cart.ID).Find(&products).Error; err != nil {
		c.String(500, "Не удалось получить товары")
		return
	}

	c.HTML(200, "cart.html", gin.H{
		"Products": products,
		"Total":    cart.Value,
	})
}

func (h HtmlService) SubmitMessage(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	message := c.PostForm("message")

	msg := models.Message{
		Name:    name,
		Email:   email,
		Message: message,
	}

	if err := h.DB.Create(&msg).Error; err != nil {
		c.String(500, "Ошибка при сохранении сообщения")
		return
	}

	c.HTML(200, "message.html", gin.H{
		"Success": true,
	})
}
