package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/wangzupeng12061/we-book/internal/domain"
	"time"
)

var ErrKeyNotFound = redis.Nil

type UserCache struct {
	client     redis.Cmdable
	expiration time.Duration
}

func NewUserCache(client redis.Cmdable) *UserCache {
	return &UserCache{
		client:     client,
		expiration: 5 * time.Minute,
	}
}
func (uc *UserCache) Set(ctx context.Context, u domain.User) error {
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	key := uc.key(u.ID)
	return uc.client.Set(ctx, key, val, uc.expiration).Err()
}
func (uc *UserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := uc.key(id)
	val, err := uc.client.Get(ctx, key).Result()
	if err != nil {
		return domain.User{}, err
	}
	var u domain.User
	err = json.Unmarshal([]byte(val), &u)
	return u, err
}

func (uc *UserCache) key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}
