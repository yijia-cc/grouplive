package main

import (
	"github.com/yijia-cc/grouplive/auth/config"
	"github.com/yijia-cc/grouplive/auth/controller"
	"github.com/yijia-cc/grouplive/auth/model"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	cfg := config.LoadEnv()

	db := model.DBConn(cfg)
	defer db.Close()

	router := controller.StartUp(cfg)

	log.Fatal(http.ListenAndServe(":8080", router))
}