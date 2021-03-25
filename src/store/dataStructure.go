package store

import "time"

type (
	Tables interface {
		TableName() string
	}

	TProduk struct {
		ID          uint      `groups:"id" json:"id" gorm:"primaryKey"`
		Title       string    `groups:"title" json:"title"`
		Description string    `groups:"description" json:"description"`
		Rating      int       `groups:"rating" json:"rating"`
		Image       string    `groups:"image" json:"image"`
		CreatedAt   time.Time `groups:"create" json:"created_at" gorm:"autoCreateTime"`
		UpdatedAt   time.Time `groups:"update" json:"updated_at" gorm:"autoUpdateTime"`
	}

	AuthParams struct {
		Phone string `json:"phone"`
		Otp   string `json:"otp"`
	}
)

func (TProduk) TableName() string {
	return "t_produk"
}
