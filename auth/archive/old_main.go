package archive

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

	//cfg := config.LoadEnv()
	//
	//db := model.DBConn(cfg)
	//defer db.Close()
	//
	//router := controller.StartUp(cfg)
	//
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	log.Fatal(http.ListenAndServe(":8080", router))
	//}()
	//fmt.Println("Service started at :8080")
	//wg.Wait()

	//routingServer := routing.NewServer()
	//lis := net.Li
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		entry.StartServer(cfg, sqlDB)
	}()
	wg.Wait()
}
