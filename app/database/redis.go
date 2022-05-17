package database

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var Rdb *redis.Storage
var SessionStore *session.Store

func InitRedis() {
	Rdb = redis.New(redis.Config{
		Port: 6379,
	})

	log.Println("Redis cache connected!")
}

func CreateSession() *session.Store {
	SessionStore = session.New(session.Config{
		Storage: Rdb,
	})
	log.Println("Session started!")
	return SessionStore
}
