package database

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var rdb *redis.Storage
var SessionStore *session.Store

func InitSession() *session.Store {
	rdb = redis.New(redis.Config{
		Port: 6379,
	})
	log.Println("Redis cache initialized!")

	SessionStore = session.New(session.Config{
		Storage: rdb,
	})
	log.Println("Session started!")
	return SessionStore
}
