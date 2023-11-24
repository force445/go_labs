package model

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	ID          int          `gorm:"primaryKey"`
	Fullname    string       `json:"username"`
	Birthday    string       `json:"birthday"`
	Address     string       `json:"address"`
	Province    string       `json:"province"`
	City        string       `json:"city"`
	Zipcode     string       `json:"zipcode"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
	CreatedAt   string       `json:"createdat"`
	UpdatedAt   string       `json:"updatedat"`
	CreditCards []CreditCard `gorm:"foreignKey:MemberID"`
}

type Members struct {
	Members []Member `json:"members"`
}

type CreditCard struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	CardName  string `json:"cardname"`
	CardNo    string `json:"cardno"`
	Expire    string `json:"expire"`
	Cvv       string `json:"cvv"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
	MemberID  int    `json:"memberid"`
}

type CreditCards struct {
	CreditCards []CreditCard `json:"creditcards"`
}

type Order struct {
	gorm.Model
	ID          int     `gorm:"primaryKey"`
	Totalamount float64 `json:"totalamount"`
	Orderdate   string  `json:"orderdate"`
	MemberID    int     `json:"memberid"`
	BookID      int     `json:"bookid"`
}

type Orders struct {
	Orders []Order `json:"orders"`
}

type Book struct {
	gorm.Model
	ID          int         `gorm:"primaryKey"`
	Title       string      `json:"title"`
	Price       float64     `json:"price"`
	Stock       int         `json:"stock"`
	AuthorID    []Author    `gorm:"foreignKey:BookID"`
	PublisherID []Publisher `gorm:"foreignKey:BookID"`
}

type Books struct {
	Books []Book `json:"books"`
}

type Author struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Fullname  string `json:"fullname"`
	Biography string `json:"biography"`
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
}

type Publishers struct {
	Publishers []Publisher `json:"publishers"`
}
