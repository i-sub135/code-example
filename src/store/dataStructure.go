package store

import "time"

type (
	Tables interface {
		TableName() string
	}

	TProduk struct {
		ID          uint      `json:"id" gorm:"primaryKey"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Rating      int       `json:"rating"`
		Image       string    `json:"image"`
		CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
		UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	}
)

func (TProduk) TableName() string {
	return "t_produk"
}
