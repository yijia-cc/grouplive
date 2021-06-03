package main

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yijia-cc/grouplive/auth/config"
	"github.com/yijia-cc/grouplive/auth/db"
	"github.com/yijia-cc/grouplive/auth/routing/entry"
)

func main() {
	cfg := config.FromEnv()
	sqlDB, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	err = db.Migrate(sqlDB, cfg)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		entry.StartServer(cfg, sqlDB)
	}()
	wg.Wait()
}
