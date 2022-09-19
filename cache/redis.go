package cache

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type ICache interface {
	Get(key string, list interface{}) bool
	GetInt(key string) (int, bool)
	GetHash(key string) (map[string]string, bool)
	SetHash(key string, hash map[string]string)
	GetString(key string) (string, bool)
	Set(key string, list interface{}, ex uint64)
	SetNoEx(key string, list interface{})
	SetString(key, data string, ex uint64)
	DeleteFromPattern(pattern string)
	Delete(key string)
	Exists(key string) bool
	SetSession(key string, val interface{}, ex uint64) error
}

type Cache struct {
	Pool *redis.Pool
}

func (c *Cache) Get(key string, list interface{}) bool {
	var client = c.Pool.Get()
	defer client.Close()
	value, err := redis.String(client.Do("get", key))
	if err != nil {
		return false
	}
	err = json.Unmarshal([]byte(value), &list)
	if err != nil {
		return false
	}

	return true
}

func (c *Cache) GetInt(key string) (int, bool) {
	var client = c.Pool.Get()
	defer client.Close()
	value, err := redis.Int(client.Do("get", key))
	if err != nil {
		return 0, false
	}

	return value, true
}

func (c *Cache) GetHash(key string) (map[string]string, bool) {
	var client = c.Pool.Get()
	defer client.Close()
	value, err := redis.StringMap(client.Do("HGETALL", key))

	if err != nil {
		return make(map[string]string, 0), false
	}
	return value, true
}

func (c *Cache) SetHash(key string, hash map[string]string) {
	var client = c.Pool.Get()
	defer client.Close()

	for k, v := range hash {
		client.Do("HSET", key, k, v)
	}
}

func (c *Cache) GetString(key string) (string, bool) {
	var client = c.Pool.Get()
	defer client.Close()
	value, err := redis.String(client.Do("get", key))
	if err != nil {
		return "", false
	}

	return string(value), true
}

func (c *Cache) Set(key string, list interface{}, ex uint64) {
	var client = c.Pool.Get()
	defer client.Close()
	jsondata, err := json.Marshal(list)
	if err != nil {
		fmt.Println("json converter")
	}
	redis.Int64(client.Do("set", key, string(jsondata), "ex", ex))

}

func (c *Cache) SetNoEx(key string, list interface{}) {
	var client = c.Pool.Get()
	defer client.Close()
	jsondata, err := json.Marshal(list)
	if err != nil {
		fmt.Println("json converter")
	}
	redis.Int64(client.Do("set", key, string(jsondata)))

}

func (c *Cache) SetString(key, data string, ex uint64) {
	var client = c.Pool.Get()
	defer client.Close()

	redis.Int64(client.Do("set", key, data, "ex", ex))
}

func (c *Cache) DeleteFromPattern(pattern string) {
	var client = c.Pool.Get()
	defer client.Close()

	val, err := redis.Strings(client.Do("keys", pattern))
	if err == nil {
		for _, item := range val {
			client.Do("del", item)
		}
	}
}

func (c *Cache) Delete(key string) {
	var client = c.Pool.Get()
	client.Do("del", key)
	defer client.Close()
}

// eğer rediste verilen key değerine göre bir değer varsa true döner, yoksa false.
func (c *Cache) Exists(key string) bool {
	var client = c.Pool.Get()
	defer client.Close()
	val, err := redis.Int64(client.Do("exists", key))
	if err != nil {
		return false
	}

	if val <= 0 {
		return false
	}

	return true
}

func (c *Cache) SetSession(key string, val interface{}, ex uint64) error {
	var client = c.Pool.Get()
	defer client.Close()
	jsondata, err := json.Marshal(val)
	if err != nil {
		fmt.Println("json converter")
	}
	_, err = redis.Int64(client.Do("set", key, string(jsondata), "ex", ex))
	return err
}

func (c *Cache) GetSession(key string) (interface{}, error) {
	var client = c.Pool.Get()
	defer client.Close()
	value, err := client.Do("get", key)
	return value, err
}


