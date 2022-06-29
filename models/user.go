package models

type User struct {
	Id       uint   `json:"id"`
	Username string `gorm:"unique"`
	Password []byte `json:"-"`
	Balance  int    `json:"balance"`
	Address  []byte `json:"-"`
}
