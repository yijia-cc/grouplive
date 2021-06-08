package main

import (
	"sync"

	"github.com/yijia-cc/grouplive/calendar/gql/gqlentry"
	"github.com/yijia-cc/grouplive/calendar/obs"

	"github.com/yijia-cc/grouplive/calendar/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yijia-cc/grouplive/calendar/db"
)

func main() {
	logger := obs.NewLogger(obs.Info)
	cfg := config.FromEnv()
	database, err := db.Connect(cfg)
	if err != nil {
		logger.Fatal(nil, err)
		panic(err)
	}
	defer database.Close()

	err = db.Migrate(database, cfg)
	if err != nil {
		logger.Fatal(nil, err)
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		gqlentry.StartServer(cfg, logger.NextLayer(), database)
	}()
	wg.Wait()
}
