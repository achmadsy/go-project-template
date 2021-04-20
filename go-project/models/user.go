package models

type Address struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	AddressLine string `json:"addressLine"`
	UserID      uint   `json:"userId"`
}

type User struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Name       string    `json:"name"`
	Addressses []Address `json:"addresses"`
}
