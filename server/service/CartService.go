package service

import (
	"errors"
	"fmt"
	"gg/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartService struct {
	DB *gorm.DB
}


func (crt CartService) Add(c *gin.Context) {
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
	if err := crt.DB.Where("login = ? and password = ?", login, password).First(&user).Error; err != nil {
		// Создаем запрос, First - первая запись
		c.String(500, fmt.Sprintf("Не удалось создать корзину: %v", err))
		return
	}

	tovar := c.PostForm("tovar")
	priceS := c.PostForm("price")
	price, err := strconv.ParseFloat(priceS, 64)
	// strconv - парсим строку, из строки в цифры
	if err != nil {
		c.String(401, "Неверный формат цены")
		return
	}

	var cart models.Cart
	if err := crt.DB.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		// Проверяем на наличие корзины
			cart = models.Cart{UserID: user.ID}
			if err := crt.DB.Create(&cart).Error; err != nil {
				c.String(500, fmt.Sprintf("Не удалось создать корзину: %v", err))
				return
			// Создаем корзину
			}
		} else {
			c.String(500, fmt.Sprintf("Ошибка базы данных: %v", err))
			return
		}
	}
	product := models.Product{CartID: cart.ID, Tovar: tovar, Price: price}
	if err := crt.DB.Create(&product).Error; err != nil {
		c.String(501, fmt.Sprintf("Не удалось добавить товар: %v", err))
		return
	}

	var totalPrice float64
	if err := crt.DB.Model(&models.Product{}).Where("cart_id = ?", cart.ID).Select("SUM(price)").Scan(&totalPrice).Error; err != nil {
		c.String(500, fmt.Sprintf("Не удалось пересчитать стоимость корзины: %v", err))
		return
	}
	cart.Value = totalPrice
	if err := crt.DB.Save(&cart).Error; err != nil {
		c.String(500, fmt.Sprintf("Не удалось обновить корзину: %v", err))
		return
	}
	c.String(200, "Товар успешно добавлен в корзину")
}
