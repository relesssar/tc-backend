package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	tc "tc_kaztranscom_backend_go"
	"time"
)

type ADLogPostgres struct {
	db         *sqlx.DB
	dbContract *sqlx.DB
}

func NewADLogPostgres(db, dbContract *sqlx.DB) *ADLogPostgres {
	return &ADLogPostgres{db: db, dbContract: dbContract}
}

func (r *ADLogPostgres) UpdateDepartmentInfo() error {

	var lists []string
	//смотрю которые не имеют инфу о департаменте
	query := fmt.Sprintf("SELECT DISTINCT ad_login from %s where department is null", adlogTable)
	err := r.db.Select(&lists, query)
	if err != nil {
		logrus.Errorf("Ошибка %s", err)
		return err
	}
	logrus.Infof("Найдено %d записей без информации о департаменте", len(lists))
	for _, v := range lists {

		var res struct {
			Fullname   string `db:"fullname"`
			Department string `db:"depname"`
		}
		//Узнаю имена департаменов и ФИО
		query = fmt.Sprintf("SELECT u.fullname,d.depname from %s as u INNER JOIN %s as d ON u.depId=d.depid where lower(u.login)=? LIMIT 1", usersConntracts, deplistConntracts)
		err = r.dbContract.QueryRow(query, v).Scan(&res.Fullname, &res.Department)
		if err != nil {
			continue
		}

		logrus.Infof("%s Найдена информация %v", v, res)
		query = fmt.Sprintf("UPDATE %s SET department=$1,full_name=$2  where ad_login=$3", adlogTable)
		_, err = r.db.Exec(query, res.Department, res.Fullname, v)
		if err != nil {
			logrus.Errorf("Ошибка(53) %s", err)
			return err
		}
	}
	return err
}

func (r *ADLogPostgres) InsertADLog(log tc.ADLog) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	id := uuid.New()
	query := fmt.Sprintf("INSERT INTO %s (id,ad_login,date_time,eventid,ip,str) values ($1,$2,$3,$4,$5,$6)", adlogTable)
	_, err = tx.Exec(query, id.String(), strings.ToLower(log.Ad_login), log.DateTime, log.Eventid, log.Ip, log.Str)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return id.String(), tx.Commit()

}

func (r *ADLogPostgres) GetAllByQuery(data map[string]string) ([]tc.ADLog, error) {
	var lists []tc.ADLog
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if len(data["date"]) == 0 && data["from_date"] == "" {
		ttime := time.Now()
		data["date"] = ttime.Format("2006-01")
		setValues = append(setValues, fmt.Sprintf("to_char(date_time,'YYYY-MM')=$%d", argId))
		args = append(args, data["date"])
		argId++
	}
	if len(data["date"]) != 0 && data["from_date"] == "" {
		setValues = append(setValues, fmt.Sprintf("to_char(date_time,'YYYY-MM')=$%d", argId))
		args = append(args, data["date"])
		argId++
	}

	if data["from_date"] != "" {
		setValues = append(setValues, fmt.Sprintf("to_char(date_time,'YYYY-MM-DD')>=$%d", argId))
		args = append(args, data["from_date"])
		argId++
	}
	if data["to_date"] != "" {
		setValues = append(setValues, fmt.Sprintf("to_char(date_time,'YYYY-MM-DD')<=$%d", argId))
		args = append(args, data["to_date"])
		argId++
	}
	if data["ad_login"] != "" {
		setValues = append(setValues, fmt.Sprintf("(lower(ad_login) like($%d) OR lower(full_name) LIKE ($%d) OR lower(department) LIKE ($%d))", argId, argId+1, argId+2))
		args = append(args, "%"+strings.ToLower(data["ad_login"])+"%")
		args = append(args, "%"+strings.ToLower(data["ad_login"])+"%")
		args = append(args, "%"+strings.ToLower(data["ad_login"])+"%")
		argId += 3
	}
	setQuery := strings.Join(setValues, " and ")
	query := fmt.Sprintf("SELECT DISTINCT * FROM %s WHERE %s  order by ad_login,date_time,eventid,ip", adlogTable, setQuery)
	err := r.db.Select(&lists, query, args...)
	//logrus.Infof("%s %v", query, args)
	return lists, err

}
