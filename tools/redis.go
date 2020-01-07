package tools

import (
	"log"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	*redis.Client
}

var redisClient *RedisClient

var once sync.Once

func NewRedisClient() *RedisClient {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		redisClient = &RedisClient{Client: client}
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	log.Println("Redis connection ping response : " + pong)
	return redisClient
}

// get last chat number for given app
func (rc *RedisClient) GetAppChatNumber(appToken string) uint {
	lastChatNum, err := rc.Client.Get(appToken).Result()
	if err != redis.Nil {
		log.Println("No prevoius chats")
		return 0
	} else if err != nil {
		panic(err)
	} else {
		lastNum, _ := strconv.Atoi(lastChatNum)
		return uint(lastNum)
	}
}

// update last chat number for given app
func (rc *RedisClient) SetAppChatNumber(appToken string, num uint) error {
	err := rc.Client.Set(appToken, num, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
