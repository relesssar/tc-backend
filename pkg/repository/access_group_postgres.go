package repository

import (
	//"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	"time"

	//"strings"
	tc "tc_kaztranscom_backend_go"
	//"time"
)

type AccessGroupPostgres struct {
	db *sqlx.DB
}

func (r *AccessGroupPostgres) InsertFilterAccessGroup(fro tc.FilterAccessGroup) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	//Если значение можно распарсить через ; то значит нужно вставить множество значений.
	val_m := strings.Split(fro.Filter_val, ";")
	//logrus.Infof("Множественная вставка, %q",val_m)
	id := uuid.New()
	for _, val := range val_m {
		//logrus.Infof("%v", fro)
		filter_val := strings.Trim(val, " ")
		if filter_val == "" {
			continue
		}
		id = uuid.New()
		query := fmt.Sprintf("INSERT INTO %s (id,filter_type,filter_val,filter_desc,user_id,router_name) values ($1,$2,$3,$4,$5,$6)", filterAccessGroup)
		_, err = tx.Exec(query, id.String(), fro.Filter_type, filter_val, fro.Filter_desc, fro.User_id, fro.Router_name)
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}
	return id.String(), tx.Commit()

}
func (r *AccessGroupPostgres) GetFilterAccessGroup(filter tc.FilterAccessGroupSearch) ([]tc.FilterAccessGroupSearch, error) {
	var lists []tc.FilterAccessGroupSearch
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if filter.Id != "" {
		setValues = append(setValues, fmt.Sprintf("f.id=$%d", argId))
		args = append(args, filter.Id)
		argId++
	}

	if filter.Router_name != "" {
		setValues = append(setValues, fmt.Sprintf("f.router_name=$%d", argId))
		args = append(args, filter.Filter_type)
		argId++
	}
	if filter.Filter_type != "" {
		setValues = append(setValues, fmt.Sprintf("f.filter_type=$%d", argId))
		args = append(args, filter.Filter_type)
		argId++
	}
	if filter.Filter_val != "" {
		setValues = append(setValues, fmt.Sprintf("f.filter_val=$%d", argId))
		args = append(args, filter.Filter_val)
		argId++
	}
	if filter.Filter_desc != "" {
		setValues = append(setValues, fmt.Sprintf("f.filter_desc=$%d", argId))
		args = append(args, filter.Filter_desc)
		argId++
	}
	if filter.User_id != "" {
		setValues = append(setValues, fmt.Sprintf("f.user_id=$%d", argId))
		args = append(args, filter.User_id)
		argId++
	}
	if filter.User_name != "" {
		setValues = append(setValues, fmt.Sprintf("u.name=$%d", argId))
		args = append(args, filter.User_name)
		argId++
	}
	if filter.Created_at != "" {
		setValues = append(setValues, fmt.Sprintf("to_char(f.created_at,'YYYY-MM-DD')=$%d", argId))
		args = append(args, filter.Created_at)
		argId++
	}
	if len(setValues) == 0 {
		setValues = append(setValues, fmt.Sprintf("1=$%d", argId))
		args = append(args, "1")
		argId++
	}
	setQuery := strings.Join(setValues, " and ")

	query := fmt.Sprintf("SELECT f.*,u.name FROM %s as f INNER JOIN %s as u ON f.user_id=u.id WHERE %s  order by f.created_at DESC", filterAccessGroup, usersTable, setQuery)
	err := r.db.Select(&lists, query, args...)
	logrus.Infof("%s %v", query, args)

	/*
		setQuery := strings.Join(setValues, ", ")
		query := fmt.Sprintf("UPDATE %s set %s  WHERE  id=$%d and user_id=$%d", categoryTable, setQuery, argId, argId+1)

		_, err := r.db.Exec(query, args...)
	*/
	return lists, err
}
func (r *AccessGroupPostgres) DeleteFilter(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", filterAccessGroup)
	_, err := r.db.Exec(query, id)
	return err
}

func NewAccessGroupPostgres(db *sqlx.DB) *AccessGroupPostgres {
	return &AccessGroupPostgres{db: db}
}
func (r *AccessGroupPostgres) GetAllAccessGroupByQuery(data map[string]string) ([]tc.AccessGroup, error) {
	var lists []tc.AccessGroup
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if data["date"] == "" {
		ttime := time.Now()
		data["date"] = ttime.Format("2006-01-02")
		setValues = append(setValues, fmt.Sprintf("to_char(t.created_at,'YYYY-MM-DD')<=$%d", argId))
		args = append(args, data["date"])
		argId++
	} else {
		setValues = append(setValues, fmt.Sprintf("to_char(t.created_at,'YYYY-MM-DD')=$%d", argId))
		args = append(args, data["date"])
		argId++
	}

	if data["access_status"] != "" {
		setValues = append(setValues, fmt.Sprintf("t.access_status=$%d", argId))
		args = append(args, data["access_status"])
		argId++
	}
	if data["router_name"] != "" {
		val_m := strings.Split(data["router_name"], ",")

		var s_tmp string
		for _, val := range val_m {
			val_tmp := strings.Trim(val, " ")
			if val_tmp == "" {
				continue
			}
			if s_tmp != "" {
				s_tmp += ","
			}
			s_tmp += fmt.Sprintf("$%d", argId)
			args = append(args, val_tmp)
			argId++
		}
		setValues = append(setValues, fmt.Sprintf("t.router_name IN (%s)", s_tmp))

	}

	setQuery := strings.Join(setValues, " and ")
	//Если не учитывать фильтр
	if data["check_filter"] == "false" {
		query := fmt.Sprintf("SELECT * FROM %s as t WHERE %s  order by t.router_name,t.iface_name", accessGroup, setQuery)
		err := r.db.Select(&lists, query, args...)
		logrus.Infof("%s %v", query, args)

		return lists, err
	}
	//Тут учитываю фильтр
	sql := `SELECT t.* FROM %s as t 
		WHERE %s AND t.id NOT IN 
	(SELECT t2.id FROM %s as fro INNER JOIN %s as t2 ON t2.iface_name=fro.filter_val 
		 AND fro.filter_type='interface_name'
	) AND  t.id NOT IN 
	(SELECT t3.id FROM %s as fro2 INNER JOIN %s as t3 ON t3.ip=fro2.filter_val 
		 AND fro2.filter_type='ip'
	) order by t.router_name, t.dognum`

	query := fmt.Sprintf(sql, accessGroup, setQuery, filterAccessGroup, accessGroup, filterAccessGroup, accessGroup)
	err := r.db.Select(&lists, query, args...)
	logrus.Infof("%s %v", query, args)
	return lists, err
}

func (r *AccessGroupPostgres) InsertAccessGroup(ag tc.AccessGroup) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	id := uuid.New()

	query := fmt.Sprintf("INSERT INTO %s (id,router_name,iface_host,ip,iface_name,iface_desc,client_status,in_policy,out_policy,access_group,dognum,clsrv,access_status) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)", accessGroup)
	_, err = tx.Exec(query, id.String(), ag.Router_Name, ag.Iface_host, ag.Ip, ag.Iface_name, ag.Iface_desc, ag.Client_status, ag.In_policy, ag.Out_policy, ag.Access_group, ag.Dognum, ag.Clsrv, ag.Access_status)
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
		return "", err
	}
	return id.String(), tx.Commit()
}

func (r *AccessGroupPostgres) UpdateAccessGroup(id, userId string, input tc.UpdateAccessGroupInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Access_status != "" {
		setValues = append(setValues, fmt.Sprintf("access_status=$%d", argId))
		args = append(args, input.Access_status)
		argId++
	}
	setValues = append(setValues, fmt.Sprintf("updated_at=$%d", argId))
	args = append(args, "now()")
	argId++

	setQuery := strings.Join(setValues, ", ")
	var query string

	args = append(args, id)
	query = fmt.Sprintf("UPDATE %s set %s  WHERE  id=$%d", accessGroup, setQuery, argId)

	logrus.Infof(" %s %v", query, args)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		logrus.Infof("%v %s %v", err, query, args)
		return err
	}
	//Добавляю сразу в историю изменений
	data := tc.AccessGroupHistory{Msg: input.Msg, New_val: input.Access_status, Access_group_id: input.Id, User_id: userId, Old_val: input.Access_status_old}
	_, err = r.InsertAccessGroupHistory(data)
	return err
}
func (r *AccessGroupPostgres) InsertAccessGroupHistory(data tc.AccessGroupHistory) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	//logrus.Infof("%v", proh)
	id := uuid.New()
	query := fmt.Sprintf("INSERT INTO %s values ($1,$2,$3,$4,$5,$6)", accessGroupHistory)
	_, err = tx.Exec(query, id.String(), data.Access_group_id, data.User_id, data.Old_val, data.New_val, data.Msg)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return id.String(), tx.Commit()

}
func (r *AccessGroupPostgres) GetAccessGroupHistory(input tc.GetAccessGroupHistory) ([]tc.AccessGroupHistory, error) {
	var lists []tc.AccessGroupHistory
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Id != "" {
		setValues = append(setValues, fmt.Sprintf("agh.id=$%d", argId))
		args = append(args, input.Id)
		argId++
	}
	if len(input.Ids) != 0 {
		setValues = append(setValues, fmt.Sprintf("agh.id IN($%d)", argId))
		args = append(args, input.Ids)
		argId++
	}
	if len(setValues) == 0 {
		setValues = append(setValues, fmt.Sprintf("1=$%d", argId))
		args = append(args, "1")
		argId++
	}
	setQuery := strings.Join(setValues, " and ")

	query := fmt.Sprintf("SELECT agh.*,u.name"+
		" FROM %s as agh INNER JOIN %s as u ON agh.user_id=u.id WHERE %s  order by agh.access_group_id,agh.created_at,agh.user_id", accessGroupHistory, usersTable, setQuery)
	err := r.db.Select(&lists, query, args...)

	logrus.Infof("%s %v", query, args)

	return lists, err
}
