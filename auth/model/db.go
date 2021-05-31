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
func DBConn(cfg *config.Config) *sql.DB {
	// dsn (data source name): username:password@protocol(hostname:port)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DBName)
	var err error
	if db, err = sql.Open(cfg.DbDriver, dsn); err != nil {
		panic(err)
	}

	// Verify the database connection is successful
	err = db.Ping()
	if err != nil {
		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
		log.Fatal(fmt.Sprintf("Database is disconnected due to %v", err))
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