package model

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	ID          int          `gorm:"primaryKey"`
	Fullname    string       `json:"fullname"`
	Birthday    string       `json:"birthday"`
	Address     string       `json:"address"`
	Province    string       `json:"province"`
	City        string       `json:"city"`
	Zipcode     string       `json:"zipcode"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
	CreditCards []CreditCard `gorm:"foreignKey:MemberID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Orders      []Order      `gorm:"foreignKey:MemberID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Members struct {
	Members []Member `json:"members"`
}

type CreditCard struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	CardName string `json:"cardname"`
	CardNo   string `json:"cardno"`
	Expire   string `json:"expire"`
	Cvv      string `json:"cvv"`
	MemberID int    `json:"memberID"`
}

type CreditCards struct {
	CreditCards []CreditCard `json:"creditcards"`
}

type Order struct {
	gorm.Model
	ID          int     `gorm:"primaryKey"`
	Totalamount float64 `json:"totalamount"`
	Orderdate   string  `json:"orderdate"`
	MemberID    int     `json:"memberID"`
	Books       []Book  `gorm:"foreignKey:OrderID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Orders struct {
	Orders []Order `json:"orders"`
}

type Book struct {
	gorm.Model
	ID          int     `gorm:"primaryKey"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	OrderID     int     `json:"orderID"`
	AuthorID    int     `json:"authorID"`
	PublisherID int     `json:"publisherID"`
}

type Books struct {
	Books []Book `json:"books"`
}

type Author struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Fullname  string `json:"fullname"`
	Biography string `json:"biography"`
	Books     []Book `gorm:"foreignKey:AuthorID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Authors struct {
	Authors []Author `json:"authors"`
}

type Publisher struct {
	gorm.Model
	ID      int    `gorm:"primaryKey"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Books   []Book `gorm:"foreignKey:PublisherID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Publishers struct {
	Publishers []Publisher `json:"publishers"`
}
