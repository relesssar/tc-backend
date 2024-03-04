package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Имена таблиц базы данных
const (
	usersConntracts   = "users"
	deplistConntracts = "deplist"
)

//Настройки базы данных
type ConfigContracts struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewMysqlDBContracts(cfg ConfigContracts) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
