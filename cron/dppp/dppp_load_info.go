/*
Скрипт для запуска по крону, проверка на испавленные данные,
подробнее https://tc.kaztranscom.kz:9000/swagger/index.html#/Problem_Router_Onyma/check_close_problem_router_onyma_speed
*/

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
)

type RealIp struct {
	Real_ip string
	Name    string
}

var schema = `
CREATE TABLE dppp_adress_base (
    id text,
    reap_ip text,
    name text
);`

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//TODO тут Хард кор пути файла настроек на сервере
	env_path := "/home/ubuntu/go/src/tc_kaztranscom_backend_go/.env"

	if err := godotenv.Load(env_path); err != nil {
		logrus.Fatalf("Ошибка чтения конфиг файла: %s %v", env_path, err)
	}

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSLMODE")))

	if err != nil {
		logrus.Fatalf("Ошибка подключения: %s", err.Error())
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatalf("Ошибка подключения: %s", err.Error())
		os.Exit(1)
	}

	dbDppp, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", os.Getenv("DB_USER_DPPP"), os.Getenv("DB_PASS_DPPP"), os.Getenv("DB_HOST_DPPP"), os.Getenv("DB_PORT_DPPP"), os.Getenv("DB_NAME_DPPP")))

	if err != nil {
		logrus.Fatalf("Ошибка подключения DPPP: %v", err)
		os.Exit(1)
	}

	err = dbDppp.Ping()
	if err != nil {
		logrus.Fatalf("Ошибка подключения DPPP: %s", err.Error())
		os.Exit(1)
	}

	real_ip := []RealIp{}
	dbDppp.Select(&real_ip, "select distinct real_ip, name from (select real_ip, name from adress_base.tableses union all select real_ip, name from adress_base.sub_tableses ) a order by real_ip")

	if len(real_ip) == 0 {
		logrus.Info("Нет данных, выход")
		os.Exit(2)
	}
	//db.MustExec(schema)
	tx := db.MustBegin()
	for _, v := range real_ip {
		if v.Name == "" {
			continue
		}
		tx.MustExec("INSERT INTO dppp_adress_base (real_ip, dppp_name) VALUES ($1, $2)", v.Real_ip, v.Name)
	}
	tx.Commit()
	logrus.Infof("%v", real_ip)
	fmt.Println("Скрипт успешно отработал")

}
