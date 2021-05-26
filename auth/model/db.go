package model

import (
	"database/sql"
	"fmt"
	"github.com/yijia-cc/grouplive/auth/config"
	"log"
)

// A database connection pool used by the model
var db *sql.DB

// DBConn establishes a database connection and returns a connection pool
func DBConn() *sql.DB {
	// dsn (data source name): username:password@protocol(hostname:port)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.UserDBConfig.Username, config.UserDBConfig.Password, config.UserDBConfig.Hostname, config.UserDBConfig.Port, config.UserDBConfig.DBname)
	var err error
	if db, err = sql.Open(config.UserDBConfig.Driver, dsn); err != nil {
		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
		log.Fatal(fmt.Sprintf("Error %v when connecting to DB", err))
	}
	return db
}


func CreateUserTable() error {
	sql := `CREATE TABLE IF NOT EXISTS users (
				username VARCHAR(30) PRIMARY KEY,
				password VARCHAR(30) NOT NULL,
				first_name VARCHAR(50) NOT NULL,
				last_name VARCHAR(50) NOT NULL,
				email VARCHAR(100) NOT NULL,
				apt VARCHAR(10) NOT NULL, 
				role ENUM('resident', 'admin') NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			)`

	_, err := db.Exec(sql)
	return err
}