package store

import (
	"fmt"
	"time"
)

type (
	JSONTime time.Time

	Tables interface {
		TableName() string
	}

	TProduk struct {
		ID          uint      `json:"id,omitempty" gorm:"primaryKey"`
		Title       string    `json:"title,omitempty"`
		Description string    `json:"description,omitempty"`
		Rating      int       `json:"rating,omitempty"`
		Image       string    `json:"image,omitempty"`
		CreatedAt   time.Time `json:"created_at,omitempty"  gorm:"autoCreateTime"`
		UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	}
	OutputProduk struct {
		ID          uint     `groups:"id" json:"id,omitempty"`
		Title       string   `groups:"title" json:"title,omitempty"`
		Description string   `groups:"description" json:"description,omitempty"`
		Rating      int      `groups:"rating" json:"rating,omitempty"`
		Image       string   `groups:"image" json:"image,omitempty"`
		CreatedAt   JSONTime `groups:"create" json:"created_at,omitempty"`
		UpdatedAt   JSONTime `groups:"update" json:"updated_at,omitempty"`
	}

	AuthParams struct {
		Phone string `json:"phone"`
		Otp   string `json:"otp"`
	}
)

func (TProduk) TableName() string {
	return "t_produk"
}
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
