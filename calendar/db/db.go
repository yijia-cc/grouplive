package db

import (
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/yijia-cc/grouplive/calendar/config"
)

const dbType = "mysql"

func Connect(cfg config.Config) (*sql.DB, error){
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DBName)
	return sql.Open(dbType, dbSource)
}

func Migrate(db *sql.DB, cfg config.Config) error{
	migrations := migrate.FileMigrationSource{Dir: cfg.DbMigrationDir}
	_, err := migrate.Exec(db, dbType, migrations, migrate.Up)
	return err
}