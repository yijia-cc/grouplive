package main

import (
	"fmt"
	"github.com/yijia-cc/grouplive/auth/controller"
	"github.com/yijia-cc/grouplive/auth/model"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Print(123)
	db := model.DBConn()
	defer db.Close()
	jwtMiddleware := controller.StartupJWT()
	router := controller.StartupHttpRouter(jwtMiddleware)
	log.Fatal(http.ListenAndServe(":8080", router))
}
