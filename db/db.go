package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v9"
)

type User struct {
	Username string `json:"username"`
	Points int `json:"points"`
	Rank int `json:"rank"`

}

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx = context.TODO()
)


func(db *Database) SaveUser(user *User) error {
	member := &redis.Z{
		Score: float64(user.Points),
		Member: user.Username,
	}

	pipe := db.Client.TxPipeline()
	pipe.ZAdd(Ctx,"leaderboard",*member)
	rank := pipe.ZRank(Ctx,"LKEY",user.Username)
	_,err := pipe.Exec(Ctx)
	if err != nil {
		return err
	}

	fmt.Println(rank.Val(),err)
	user.Rank = int(rank.Val())
	return nil



}



func NewDatabase(address string) (*Database,error) {
	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
	})

	if err := client.Ping(Ctx).Err();err != nil {
		return nil,err
	}

	return &Database{
		Client: client,
	},nil

}