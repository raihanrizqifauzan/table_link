package redis

import (
	"encoding/json"
	"table_link/internal/domain/entity"
	"time"

	"gopkg.in/redis.v5"
)

type AuthCache interface {
	StoreUserSession(token string, user *entity.User, expiration time.Duration) error
	GetUserSession(token string) (*entity.User, error)
	DeleteUserSession(token string) error
}

type authCache struct {
	redisClient *redis.Client
}

func NewAuthCache(redisClient *redis.Client) AuthCache {
	return &authCache{}
}

func (c *authCache) StoreUserSession(token string, user *entity.User, expiration time.Duration) error {
	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return c.redisClient.Set(token, userData, expiration).Err()
}

func (c *authCache) GetUserSession(token string) (*entity.User, error) {
	vall, err := c.redisClient.Get(token).Result()
	if err != nil {
		return nil, err
	}

	var user entity.User
	if err := json.Unmarshal([]byte(vall), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *authCache) DeleteUserSession(token string) error {
	return c.redisClient.Del(token).Err()
}
