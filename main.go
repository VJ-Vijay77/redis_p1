package main

import (
	"log"

	"github.com/VJ-Vijay77/redis_cache_p1/db"
	"github.com/VJ-Vijay77/redis_cache_p1/router"
	"github.com/VJ-Vijay77/redis_cache_p1/api"
)

var (
	ListenAddr = "localhost:8080"
	RedisAddr  = "localhost:6379"
)

func main() {
	database, err := db.NewDatabase(RedisAddr)
	if err != nil {
		log.Fatalln("Failed to connect to redis")
	}

	e := router.InitRouter(database)
	api.API(e)
	e.Logger.Fatal(e.Start(ListenAddr))
}
