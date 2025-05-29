package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login    string
	Name     string
	Password string
	Check    bool
}

type Cart struct {
	gorm.Model
	UserID   uint
	Value    float64   // общая стоимость (можно пересчитывать в хуке AfterSave)
	Products []Product `gorm:"foreignKey:CartID"` // Cart.ID → Product.CartID
}

type Product struct {
	gorm.Model
	CartID uint // внешний ключ
	Cart   Cart `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // optional: поведение FK
	Tovar  string
	Price  float64
}

type Forum struct {
	gorm.Model
	Login   string
	Title   string
	Content string
}
