package db

import (
	"database/sql"
	"fmt"
	"github.com/yijia-cc/grouplive/dashboard/config"
)

var (
	UserDB *sql.DB
	DashDB *sql.DB
)

func Conn(cfg *config.Config) error {
	var err error
	if UserDB, err = connect(cfg.Db["user"]); err != nil {
		return err
	}

	if DashDB, err = connect(cfg.Db["dash"]); err != nil {
		return err
	}

	return nil
}

func Close() {
	UserDB.Close()
	DashDB.Close()
}

func connect(dbConfig *config.DbConfig) (*sql.DB, error) {
	// dsn (data source name): username:password@protocol(hostname:port)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbConfig.DbUser, dbConfig.DbPassword, dbConfig.DbHost, dbConfig.DbPort, dbConfig.DBName)

	var db *sql.DB
	var err error
	if db, err = sql.Open(dbConfig.DbDriver, dsn); err != nil {
		return db, fmt.Errorf("invalid database configuration: %v", err)
	}

	// Verify the database connection is successful
	err = db.Ping()
	if err != nil {
		return db, fmt.Errorf("database connection failed due to %v", err)
	}

	return db, err
}

