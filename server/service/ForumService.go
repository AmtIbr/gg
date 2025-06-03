package service

import (
	"gg/server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ForumService struct {
	DB *gorm.DB
}

func (f ForumService) AddPost(c *gin.Context) {
	login := c.PostForm("username")
	content := c.PostForm("message")
	theme := c.PostForm("theme")

	forum := models.Forum{
		Login:   login,
		Theme:   theme,
		Content: content,
	}

	if err := f.DB.Create(&forum).Error; err != nil {
		c.String(500, "Не удалось сохранить сообщение")
		return
	}

	c.Redirect(302, "/forum/"+theme)
}
