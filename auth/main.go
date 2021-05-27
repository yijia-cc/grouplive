package main

import (
	"fmt"
	"github.com/yijia-cc/grouplive/auth/config"
	"github.com/yijia-cc/grouplive/auth/controller"
	"github.com/yijia-cc/grouplive/auth/model"
	"log"
	"net/http"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	cfg := config.LoadEnv()

	db := model.DBConn(cfg)
	defer db.Close()

	router := controller.StartUp(cfg)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(http.ListenAndServe(":8080", router))
	}()
	fmt.Println("Service started at :8080")
	wg.Wait()
}