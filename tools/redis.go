package tools

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	*redis.Client
}

type AppChatInfo struct {
	LastMsgNumber uint `json:"last_msg_number"`
	ChatId        uint `json:"chat_id"`
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
	lastChatNum, _ := rc.Client.Get(appToken).Result()
	if lastChatNum == "" {
		log.Println("No prevoius chats")
		return 0
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

// get application chat info (chat id, last message number in this chat) using chat key (compination of app token and chat number)
func (rc *RedisClient) GetAppChatInfo(chatKey string) (lastMsgNum uint, chatId uint) {
	appChat, _ := rc.Client.Get(chatKey).Result()

	if appChat == "" {
		log.Println("No prevoius messages")
		return 0, 0
	} else {
		var appChatInfo AppChatInfo
		err := json.Unmarshal([]byte(appChat), &appChatInfo)
		if err != nil {
			log.Println(err)
		}
		return appChatInfo.LastMsgNumber, appChatInfo.ChatId
	}
}

// update application chat info
func (rc *RedisClient) SetAppChatInfo(chatKey string, msgNum, chatId uint) error {
	appChatInfo, err := json.Marshal(AppChatInfo{LastMsgNumber: msgNum, ChatId: chatId})
	if err != nil {
		log.Println(err)
	}

	err = rc.Client.Set(chatKey, appChatInfo, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
