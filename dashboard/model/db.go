package model

import (
	"database/sql"
	"fmt"
	"github.com/yijia-cc/grouplive/dashboard/config"
)

var userDB, dashDB *sql.DB

func DBConn(userCfg *config.Config, dashCfg *config.Config) (*sql.DB, *sql.DB) {
	var err error
	if userDB, err = conn(userCfg); err != nil {
		panic(err)
	}

	if dashDB, err = conn(dashCfg); err != nil {
		panic(err)
	}

	return userDB, dashDB
}

func conn(cfg *config.Config) (*sql.DB, error) {
	// dsn (data source name): username:password@protocol(hostname:port)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DBName)

	var db *sql.DB
	var err error
	if db, err = sql.Open(cfg.DbDriver, dsn); err != nil {
		return db, err
	}

	// Verify the database connection is successful
	err = db.Ping()
	if err != nil {
		return db, fmt.Errorf("database is disconnected due to %v", err)
	}

	return db, err
}