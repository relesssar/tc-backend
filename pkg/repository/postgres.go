package repository

import (
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//Имена таблиц базы данных
const (
	usersTable                = "users"
	userModuleTable           = "user_module"
	routerOnymaTable          = "router_onyma_speeds"
	problemRouterOnymaTable   = "problem_router_onyma_speeds"
	problemRouterOnymaHistory = "problem_router_onyma_history"
	filterRouterOnyma         = "filter_router_onyma"
	filterAccessGroup         = "filter_access_groups"
	adlogTable                = "ad_log"
	accessGroup               = "access_groups"
	accessGroupHistory        = "access_group_history"
	controlTimePause          = "control_time_pause"
	controlTimePauseHistory   = "control_time_pause_history"
	dpppAdressBaseTable       = "dppp_adress_base"
)

//Настройки базы данных
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	/*db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))*/
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
