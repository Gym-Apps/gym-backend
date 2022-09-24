package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

type ICache interface {
	Get(ctx context.Context, key string, list interface{}) bool
	GetInt(ctx context.Context, key string) (int, bool)
	GetHash(ctx context.Context, key string) (map[string]string, bool)
	SetHash(ctx context.Context, key string, hash map[string]string)
	GetString(ctx context.Context, key string) (string, bool)
	Set(ctx context.Context, key string, list interface{}, ex int64)
	SetNoEx(ctx context.Context, key string, list interface{})
	DeleteFromPattern(ctx context.Context, pattern string)
	Delete(ctx context.Context, key string)
	Exists(ctx context.Context, key string) bool
}

type Cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) ICache {
	return &Cache{client: client}
}

func (c *Cache) Get(ctx context.Context, key string, list interface{}) bool {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	err = json.Unmarshal([]byte(value), &list)
	if err != nil {
		fmt.Printf("unmarshal : %s", value)
		return false
	}

	return true
}

func (c *Cache) GetInt(ctx context.Context, key string) (int, bool) {
	value, err := c.client.Get(ctx, key).Int()
	if err != nil {
		return 0, false
	}

	return value, true
}

func (c *Cache) GetHash(ctx context.Context, key string) (map[string]string, bool) {
	value, err := c.client.Do(ctx, "get", key).Result()
	if err != nil {
		return make(map[string]string, 0), false
	}

	result := value.(map[string]string)

	return result, true
}

func (c *Cache) SetHash(ctx context.Context, key string, hash map[string]string) {
	for k, v := range hash {
		c.client.Do(ctx, "HSET", key, k, v)
	}
}

func (c *Cache) GetString(ctx context.Context, key string) (string, bool) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", false
	}
	return value, true
}

func (c *Cache) Set(ctx context.Context, key string, list interface{}, ex int64) {
	jsondata, err := json.Marshal(list)
	if err != nil {
		fmt.Println("json converter")
	}
	err = c.client.Set(ctx, key, jsondata, time.Duration(ex*int64(time.Second))).Err()
	if err != nil {
		fmt.Println(err)
	}

}

func (c *Cache) SetNoEx(ctx context.Context, key string, list interface{}) {
	jsondata, err := json.Marshal(list)
	if err != nil {
		fmt.Println("json converter")
	}
	err = c.client.Do(ctx, "set", key, jsondata).Err()
	if err == nil {
		fmt.Println(err)
	}
}

func (c *Cache) DeleteFromPattern(ctx context.Context, pattern string) {
	val, err := c.client.Do(ctx, "keys", pattern).StringSlice()
	if err == nil {
		for _, item := range val {
			c.client.Do(ctx, "del", item)
		}
	}
}

func (c *Cache) Delete(ctx context.Context, key string) {
	err := c.client.Do(ctx, "del", key).Err()
	if err != nil {
		fmt.Println(err)
	}
}

// eğer rediste verilen key değerine göre bir değer varsa true döner, yoksa false.
func (c *Cache) Exists(ctx context.Context, key string) bool {
	val, err := c.client.Do(ctx, "exists", key).Bool()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return val
}
