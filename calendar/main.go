package main

import (
	"sync"

	"github.com/yijia-cc/grouplive/calendar/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yijia-cc/grouplive/calendar/db"
	"github.com/yijia-cc/grouplive/calendar/gql"
)

func main() {
	cfg := config.FromEnv()
	database, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	err = db.Migrate(database, cfg)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		gql.StartServer(cfg, database)
	}()
	wg.Wait()
}
