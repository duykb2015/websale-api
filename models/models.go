package models

import (
	"login-api/database"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = database.ConnectDB("selling_phone")

type Admin struct {
	ID       int    `form:"id" db:"id"`
	Name     string `form:"name" db:"name"`
	UserName string `form:"username" db:"username"`
	PassWord string `form:"password" db:"password"`
	Created  string `form:"created" db:"created"`
	Updated  string `form:"updated" db:"updated"`
}

type Category struct {
	ID          int    `form:"id" db:"id"`
	Name        string `form:"name" db:"name"`
	ParentID    int    `form:"parent_id" db:"parent_id"`
	SubCategory []SubCategory
}

type SubCategory struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	ParentID int    `db:"parent_id"`
}

type Product struct {
	ID             int    `form:"id" db:"id"`
	CategoryID     int    `form:"category_id" db:"category_id"`
	Name           string `form:"name" db:"name"`
	Slug           string `form:"slug" db:"slug"`
	Price          int    `form:"price" db:"price"`
	Discount       int    `form:"discount" db:"discount"`
	Specifications string `form:"specifications" db:"specifications"`
	Promotion      string `form:"promotion" db:"promotion"`
	View           int    `form:"view" db:"view"`
	Guarantee      string `form:"guarantee" db:"guarantee"`
	Total          int    `form:"total" db:"total"`
	Buyed          int    `form:"buyed" db:"buyed"`
	ProductImage   string `form:"product_image" db:"product_image"`
	ImageList      string `form:"image_list" db:"image_list"`
	IpAddress      string `form:"ip_address" db:"ip_address"`
	RateCount      int    `form:"rate_count" db:"rate_count"`
	RateTotal      int    `form:"rate_total" db:"rate_total"`
	Status         int    `form:"status" db:"status"`
	Created        string `form:"created" db:"created"`
	Updated        string `form:"updated" db:"updated"`
}

type OrderStorage struct {
}
type Order struct {
	ID            int    `form:"id" db:"name"`
	Status        int    `form:"status"`
	TransactionID int    `db:"transaction_id"`
	ProductID     int    `db:"product_id"`
	Quantity      int    `db:"qty"`
	Amount        int    `db:"amount"`
	Data          string `db:"data"`
	Created       string `db:"created"`
}

type Transaction struct {
	ID                int     `form:"id" db:"id"`
	TransactionStatus string  `form:"transaction_status" db:"transaction_status"`
	UserID            int     `db:"user_id"`
	Amount            float64 `db:"amount"`
	Payment           string  `db:"payment"`
	Comment           string  `db:"comment"`
	Created           string  `db:"created"`
}

//Chứa thông tin của user lấy được từ database

type User struct {
	ID              int    `db:"id"`
	Email           string `json:"email" db:"email"`
	Password        string `json:"password" db:"password"`
	PasswordConfirm string `json:"password_confirm"`
	Name            string `db:"name"`
	Address         string `db:"address"`
	Phone           string `db:"phone"`
	Created         string `db:"created"`
}

//store token
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type ImageSize struct {
	Width  int64 `form:"width"`
	Height int64 `form:"height"`
}
