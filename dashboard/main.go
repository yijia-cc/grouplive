package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/yijia-cc/grouplive/dashboard/config"
    "github.com/yijia-cc/grouplive/dashboard/db"
    "github.com/yijia-cc/grouplive/dashboard/server"
    "sync"
)

func main() {
    cfg := config.FromEnv()
    err := db.Conn(cfg)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    wg := sync.WaitGroup{}
    wg.Add(1)
    go func() {
        defer wg.Done()
        server.StartUp(cfg)
    }()

    wg.Wait()
}