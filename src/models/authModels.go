package models

import (
	"context"
	"time"
)

var ctx = context.Background()

// SetRedis -- set OTP
func (r *RestModels) SetRedis(keys, data string) error {
	err := redisClien.Set(ctx, keys, data, 300*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetOtp -- get Options
func (r *RestModels) GetOtp(keys string) (string, error) {
	val, err := redisClien.Get(ctx, keys).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
