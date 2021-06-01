package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/yijia-cc/grouplive/dashboard/config"
    "github.com/yijia-cc/grouplive/dashboard/controller"
    "github.com/yijia-cc/grouplive/dashboard/model"
    "log"
    "net/http"
    "sync"
)

func main() {
    userDBConfig, dashDBConfig := config.LoadEnv()
    userDB, dashDB := model.DBConn(userDBConfig, dashDBConfig)
    defer userDB.Close()
    defer dashDB.Close()

    router := controller.StartUp(dashDBConfig)

    wg := sync.WaitGroup{}
    wg.Add(1)
    go func() {
        defer wg.Done()
        log.Fatal(http.ListenAndServe(":9090", router))
    }()
    fmt.Println("Dashboard Service started at :9090")
    wg.Wait()
}