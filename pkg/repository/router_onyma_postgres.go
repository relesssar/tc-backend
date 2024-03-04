package repository

import (
	//"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	tc "tc_kaztranscom_backend_go"
	"time"
)

type RouterOnymaPostgres struct {
	db *sqlx.DB
}

func NewRouterOnymaPostgres(db *sqlx.DB) *RouterOnymaPostgres {
	return &RouterOnymaPostgres{db: db}
}

func (r *RouterOnymaPostgres) InsertRouterOnymaHistory(proh tc.ProblemRouterOnymaHistory) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	//Проверяю если такой записи нет делаю инсерт.
	var lists []tc.ProblemRouterOnymaHistory
	query := fmt.Sprintf(`SELECT * from %s WHERE user_id=$1 and msg=$2`, problemRouterOnymaHistory)
	err = r.db.Select(&lists, query, proh.User_id,proh.Msg)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	//Если есть такая запись то выход
	if len(lists)!=0{
		tx.Rollback()
		return "", error(nil)
	}

	//logrus.Infof("%v", proh)
	id := uuid.New()
	query = fmt.Sprintf("INSERT INTO %s values ($1,$2,$3,$4,$5,$6)", problemRouterOnymaHistory)
	_, err = tx.Exec(query, id.String(), proh.Problem_router_onyma_speed_id, proh.User_id, proh.Old_val, proh.New_val, proh.Msg)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return id.String(), tx.Commit()

}
func (r *RouterOnymaPostgres) InsertControlTimePause(ctp []tc.ControlTimePauseInsert) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}
	//TODO Тут переделать на множественный инсерт
	for _, v := range ctp {
		query := fmt.Sprintf("INSERT INTO %s (id,router_onyma_speed_id,control_status,created_at) values ($1,$2,$3,$4)", controlTimePause)
		_, err = tx.Exec(query, v.Id, v.Router_onyma_speed_id, v.Control_status, v.Created_at)
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}
	return "ok", tx.Commit()
}
func (r *RouterOnymaPostgres) InsertControlTimePauseHistory(ctph []tc.ControlTimePauseHistory) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}
	//TODO Тут переделать на множественный инсерт
	for _, v := range ctph {
		query := fmt.Sprintf("INSERT INTO %s (id,control_time_pause_id,user_id,msg) values ($1,$2,$3,$4)", controlTimePauseHistory)
		_, err = tx.Exec(query, v.Id, v.ControlTimePauseId, v.UserId, v.Msg)
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}
	return "ok", tx.Commit()
}
func (r *RouterOnymaPostgres) GetRouterOnymaHistory(data map[string]string) ([]tc.ProblemRouterOnymaHistorySearch, error) {
	var lists []tc.ProblemRouterOnymaHistorySearch

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if data["id"] != "" {
		setValues = append(setValues, fmt.Sprintf("id=$%d", argId))
		args = append(args, data["id"])
		argId++
	}
	if data["problem_router_onyma_speed_id"] != "" {
		setValues = append(setValues, fmt.Sprintf("proh.problem_router_onyma_speed_id=$%d", argId))
		args = append(args, data["problem_router_onyma_speed_id"])
		argId++
	}
	if data["branch_service"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.branch_service=$%d", argId))
		args = append(args, data["branch_service"])
		argId++
	}
	if data["problem_status"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.problem_status=$%d", argId))
		args = append(args, data["problem_status"])
		argId++
	}
	if data["client_status_onyma"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.client_status_onyma=$%d", argId))
		args = append(args, data["client_status_onyma"])
		argId++
	}
	if data["company_name"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.company_name LIKE ($%d)", argId))
		args = append(args, "%"+data["company_name"]+"%")
		argId++
	}
	if data["ip_interface"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.ip_interface LIKE ($%d)", argId))
		args = append(args, data["ip_interface"]+"%")
		argId++
	}
	if data["interface_name"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.interface_name LIKE ($%d)", argId))
		args = append(args, "%"+data["interface_name"]+"%")
		argId++
	}
	if data["old_val"] != "" {
		setValues = append(setValues, fmt.Sprintf("proh.old_val=$%d", argId))
		args = append(args, data["old_val"])
		argId++
	}
	if data["new_val"] != "" {
		setValues = append(setValues, fmt.Sprintf("proh.new_val=$%d", argId))
		args = append(args, data["new_val"])
		argId++
	}
	if data["msg"] != "" {
		setValues = append(setValues, fmt.Sprintf("proh.msg=$%d", argId))
		args = append(args, data["msg"])
		argId++
	}
	if data["user_id"] != "" {
		setValues = append(setValues, fmt.Sprintf("u.user_id=$%d", argId))
		args = append(args, data["user_id"])
		argId++
	}

	//если нет ни одного критерия поиска, то указываю дату на сегодня, чтобы не повесить запрос.
	var qwery_exists bool
	for _, v := range data {
		if v != "" {
			qwery_exists = true
		}
	}
	if !qwery_exists {
		ttime := time.Now()
		data["created_at"] = ttime.Format("2006-01-02")
	}
	if data["created_at"] != "" {
		setValues = append(setValues, fmt.Sprintf("proh.created_at>=$%d and proh.created_at<=$%d", argId, argId+1))
		args = append(args, data["created_at"]+" 00:00:00")
		args = append(args, data["created_at"]+" 23:59:59")
		argId++
		argId++
	}
	if len(setValues) == 0 {
		setValues = append(setValues, fmt.Sprintf("1=$%d", argId))
		args = append(args, "1")
		argId++
	}
	setQuery := strings.Join(setValues, " and ")

	query := fmt.Sprintf("SELECT proh.*,u.name,pro.id as problem_router_onyma_id,pro.router_onyma_speed_id,pro.branch_service"+
		",pro.router_name,pro.interface_name,pro.interface_description,pro.in_policy_router,pro.in_speed_router,pro.out_policy_router"+
		",pro.out_speed_router,pro.ip_interface,pro.branch_contract,pro.dognum,pro.clsrv,pro.company_name,pro.in_speed_onyma,pro.out_speed_onyma"+
		",pro.problem_status,pro.insert_datetime,pro.updated_at"+
		" FROM %s as proh"+
		" INNER JOIN %s as u ON proh.user_id=u.id"+
		" INNER JOIN %s as pro ON proh.problem_router_onyma_speed_id=pro.id"+
		" WHERE %s  order by proh.created_at,proh.user_id", problemRouterOnymaHistory, usersTable, problemRouterOnymaTable, setQuery)
	err := r.db.Select(&lists, query, args...)

	logrus.Infof("%s %v", query, args)

	/*
		setQuery := strings.Join(setValues, ", ")
		query := fmt.Sprintf("UPDATE %s set %s  WHERE  id=$%d and user_id=$%d", categoryTable, setQuery, argId, argId+1)

		_, err := r.db.Exec(query, args...)
	*/
	return lists, err
}

func (r *RouterOnymaPostgres) InsertRouterOnymaSpeeds(router_onyma tc.RouterOnyma) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	//logrus.Infof("%v", router_onyma)
	id := uuid.New()
	query := fmt.Sprintf("INSERT INTO %s values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)", routerOnymaTable)
	_, err = tx.Exec(query, id.String(), router_onyma.Branch_service, router_onyma.Router_name, router_onyma.Interface_name, router_onyma.Interface_description, router_onyma.In_policy_router, router_onyma.In_speed_router, router_onyma.Out_policy_router, router_onyma.Out_speed_router, router_onyma.Ip_interface, router_onyma.Branch_contract, router_onyma.Dognum, router_onyma.Clsrv, router_onyma.Company_name, router_onyma.In_speed_onyma, router_onyma.Out_speed_onyma, "now()", router_onyma.Iface_shutdown_router, router_onyma.Client_status_onyma)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return id.String(), tx.Commit()

}
func (r *RouterOnymaPostgres) InsertRouterOnymaSpeedsAll(router_onyma []tc.RouterOnyma) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}
	//Множественный insert | insert multiple
	/*arg :=[] map[string]interface{}{
		{
		"Name": "Basa",
		"Email":"mail@mail",
		"Phone":"7777777777",
		"Password":"Password",
	},
	{	"Name": "Basa2",
		"Email":"mail@mail2",
		"Phone":"7222",
		"Password":"Pass2222word",}}

	query, args, err := sqlx.Named("INSERT  INTO users (name,email,phone,password_hash) values (:Name,:Email,:Phone,:Password)", arg)
	logrus.Info(query)

	query, args, err = sqlx.In(query, args...)*/

	query, args, err := sqlx.Named(fmt.Sprintf("INSERT INTO %s (branch_service,router_name,interface_name,interface_description"+
		",in_policy_router,in_speed_router,out_policy_router,out_speed_router,ip_interface,branch_contract,dognum,clsrv"+
		",company_name,in_speed_onyma,out_speed_onyma,iface_shutdown_router,client_status_onyma)"+
		" values (:branch_service,:router_name,:interface_name,"+
		" :interface_description,:in_policy_router,:in_speed_router,:out_policy_router,:out_speed_router,:ip_interface"+
		",:branch_contract,:dognum,:clsrv,:company_name,:in_speed_onyma,:out_speed_onyma,:iface_shutdown_router"+
		",:client_status_onyma)", routerOnymaTable), router_onyma)
	//logrus.Info(query)
	if err != nil {
		return "", err
	}
	query, args, err = sqlx.In(query, args...)

	if err != nil {
		return "", err
	}
	query = r.db.Rebind(query)
	//r.db.Query(query, args...)
	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	//logrus.Info(query)
	/*
		//logrus.Infof("%v", router_onyma)
		id := uuid.New()
		query := fmt.Sprintf("INSERT INTO %s values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)", routerOnymaTable)
		_, err = tx.Exec(query, id.String(), router_onyma.Branch_service, router_onyma.Router_name, router_onyma.Interface_name, router_onyma.Interface_description, router_onyma.In_policy_router, router_onyma.In_speed_router, router_onyma.Out_policy_router, router_onyma.Out_speed_router, router_onyma.Ip_interface, router_onyma.Branch_contract, router_onyma.Dognum, router_onyma.Clsrv, router_onyma.Company_name, router_onyma.In_speed_onyma, router_onyma.Out_speed_onyma, "now()", router_onyma.Iface_shutdown_router, router_onyma.Client_status_onyma)*/
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return "ok", tx.Commit()

}
func (r *RouterOnymaPostgres) InsertProblemRouterOnymaSpeeds(pRo tc.ProblemRouterOnyma) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	//logrus.Infof("%v", pRo)
	id := uuid.New()
	query := fmt.Sprintf("INSERT INTO %s values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21)", problemRouterOnymaTable)
	_, err = tx.Exec(query, id.String(), pRo.Router_onyma_speed_id, pRo.Branch_service, pRo.Router_name, pRo.Interface_name, pRo.Interface_description, pRo.In_policy_router, pRo.In_speed_router, pRo.Out_policy_router, pRo.Out_speed_router, pRo.Ip_interface, pRo.Branch_contract, pRo.Dognum, pRo.Clsrv, pRo.Company_name, pRo.In_speed_onyma, pRo.Out_speed_onyma, pRo.Problem_status, "now()", "now()", pRo.Client_status_onyma)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return id.String(), tx.Commit()
}

//problem_status<>Закрыто problem_status<>Дублирование IP
func (r *RouterOnymaPostgres) GetProblemRouterOnymaSpeedsInfoByObjectNoClose(pRo tc.ProblemRouterOnyma) ([]tc.ProblemRouterOnyma, error) {

	var lists []tc.ProblemRouterOnyma
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	setValues = append(setValues, fmt.Sprintf("problem_status<>$%d", argId))
	args = append(args, "Закрыто")
	argId++

	if pRo.Router_onyma_speed_id != "" {
		setValues = append(setValues, fmt.Sprintf("router_onyma_speed_id=$%d", argId))
		args = append(args, pRo.Router_onyma_speed_id)
		argId++
	}
	if pRo.Dognum != "" {
		setValues = append(setValues, fmt.Sprintf("dognum=$%d", argId))
		args = append(args, pRo.Dognum)
		argId++
	}
	if pRo.Clsrv != "" {
		setValues = append(setValues, fmt.Sprintf("clsrv=$%d", argId))
		args = append(args, pRo.Clsrv)
		argId++
	}
	if pRo.Router_name != "" {
		setValues = append(setValues, fmt.Sprintf("router_name=$%d", argId))
		args = append(args, pRo.Router_name)
		argId++
	}
	if pRo.Interface_name != "" {
		setValues = append(setValues, fmt.Sprintf("interface_name=$%d", argId))
		args = append(args, pRo.Interface_name)
		argId++
	}
	if pRo.Ip_interface != "" {
		setValues = append(setValues, fmt.Sprintf("ip_interface=$%d", argId))
		args = append(args, pRo.Ip_interface)
		argId++
	}
	if pRo.Insert_datetime != "" {
		setValues = append(setValues, fmt.Sprintf("to_char(insert_datetime,'YYYY-MM-DD')=to_char($%d,'YYYY-MM-DD')", argId))
		args = append(args, pRo.Insert_datetime)
		argId++
	}
	if pRo.In_speed_onyma != "" {
		setValues = append(setValues, fmt.Sprintf("in_speed_onyma=$%d", argId))
		args = append(args, pRo.In_speed_onyma)
		argId++
	}
	if pRo.Out_speed_onyma != "" {
		setValues = append(setValues, fmt.Sprintf("out_speed_onyma=$%d", argId))
		args = append(args, pRo.Out_speed_onyma)
		argId++
	}
	if pRo.Problem_status != "" {
		setValues = append(setValues, fmt.Sprintf("problem_status=$%d", argId))
		args = append(args, pRo.Problem_status)
		argId++
	}
	setQuery := strings.Join(setValues, " and ")

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s  order by router_name,company_name", problemRouterOnymaTable, setQuery)
	err := r.db.Select(&lists, query, args...)
	return lists, err
}

/**/
func (r *RouterOnymaPostgres) CheckRouterOnymaSpeed() ([]tc.RouterOnyma, error) {
	var lists []tc.RouterOnyma
	query := fmt.Sprintf("SELECT * FROM %s WHERE to_char(insert_datetime,'YYYY-MM-DD')=to_char(now(),'YYYY-MM-DD')  order by router_name, interface_name,company_name", routerOnymaTable)
	err := r.db.Select(&lists, query)

	return lists, err
}
/// statistic/group_by_router
func (r *RouterOnymaPostgres) GetStatisticRouterOnymaSpeedsGroupByRouterByQuery(data map[string]string) ([]tc.RouterOnymaGroupByRouter, error) {
	var lists []tc.RouterOnymaGroupByRouter
	query := fmt.Sprintf(`SELECT router_name,to_char(insert_datetime,'YYYY-MM-DD') as insert_date,count(*) from %s
    where to_char(insert_datetime,'YYYY-MM-DD')>=$1 and to_char(insert_datetime,'YYYY-MM-DD')<=$2
    group by router_name,to_char(insert_datetime,'YYYY-MM-DD');`, routerOnymaTable)
	err := r.db.Select(&lists, query, data["date_from"],data["date_to"])

	return lists, err
}
///
func (r *RouterOnymaPostgres) GetStatisticRouterOnymaSpeedsGroupByInsertDateByQuery(data map[string]string) ([]tc.RouterOnymaGroupByInsertDate, error) {
	var lists []tc.RouterOnymaGroupByInsertDate
	query := fmt.Sprintf(`SELECT to_char(insert_datetime,'YYYY-MM-DD') as insert_date,count(*) from %s
    where to_char(insert_datetime,'YYYY-MM-DD')>=$1 and to_char(insert_datetime,'YYYY-MM-DD')<=$2
    group by to_char(insert_datetime,'YYYY-MM-DD')`, routerOnymaTable)
	err := r.db.Select(&lists, query, data["date_from"],data["date_to"])

	return lists, err
}
func (r *RouterOnymaPostgres) GetAllRouterOnymaSpeedsByDate(date string) ([]tc.RouterOnyma, error) {
	var lists []tc.RouterOnyma
	query := fmt.Sprintf(`SELECT ros.*,dppp.dppp_name FROM %s as ros 
	LEFT OUTER JOIN %s as dppp ON ros.ip_interface = dppp.real_ip and to_char(ros.insert_datetime,'YYYY-MM-DD')=to_char(dppp.created_at,'YYYY-MM-DD')
	WHERE ros.insert_datetime>=$1 and ros.insert_datetime<=$2 order by ros.router_name, ros.interface_name,ros.company_name`, routerOnymaTable,dpppAdressBaseTable)
	err := r.db.Select(&lists, query, date+" 00:00:00", date+" 23:59:59")
	logrus.Printf("%s %s",query, date)
	return lists, err
}

/*func (r *RouterOnymaPostgres) GetDoubleDogNumForIp() ([]tc.RouterOnyma, error) {
	var lists []tc.RouterOnyma
	query := fmt.Sprintf("SELECT DISTINCT ros.dognum,ros.ip_interface,ros.company_name from %s as ros" +
		" INNER join %s AS ros2 ON ros2.dognum<>ros.dognum and ros2.ip_interface=ros.ip_interface  and ros2.dognum<>''" +
		" WHERE ros.dognum<>'' and to_char(ros.insert_datetime,'YYYY-MM-DD') = to_char(now(),'YYYY-MM-DD') order by ros.ip_interface;", routerOnymaTable, routerOnymaTable)
	err := r.db.Select(&lists, query)
	//logrus.Printf("%s %s",query, date)
	return lists, err
}*/
func (r *RouterOnymaPostgres) GetAllProblemRouterOnymaSpeedsByQuery(data map[string]string) ([]tc.ProblemRouterOnymaQuery, error) {
	//logrus.Infof("GetAllProblemRouterOnymaSpeedsByQuery data=%v",data)
	like, err := data["like"] //Если нет переменной то значит лайк по умолчанию
	if !err {
		like = "true"
	}

	var lists []tc.ProblemRouterOnymaQuery
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if data["id"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.id=$%d", argId))
		args = append(args, data["id"])
		argId++
	}

	if len(data["date"]) == 0 {
		//ttime := time.Now()
		//data["date"] = ttime.Format("2006-01-02")
		//setValues = append(setValues, fmt.Sprintf("pro.insert_datetime<=$%d", argId))
		//args = append(args, data["date"]+" 23:59:59")
		//argId++
		setValues = append(setValues, fmt.Sprintf("0=$%d", argId))
		args = append(args, "0")
		argId++
	} else {
		setValues = append(setValues, fmt.Sprintf("pro.insert_datetime>=$%d and pro.insert_datetime<=$%d", argId, argId+1))
		args = append(args, data["date"]+" 00:00:00")
		args = append(args, data["date"]+" 23:59:59")
		argId++
		argId++
	}

	if data["updated_at"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.updated_at>=$%d and pro.updated_at<=$%d", argId, argId+1))
		args = append(args, data["updated_at"]+" 00:00:00")
		args = append(args, data["updated_at"]+" 23:59:59")
		argId++
		argId++
	}
	if data["date_from"] != "" && data["date_to"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.insert_datetime>=$%d and pro.insert_datetime<=$%d", argId, argId+1))
		args = append(args, data["date_from"]+" 00:00:00")
		args = append(args, data["date_to"]+" 23:59:59")
		argId++
		argId++
	}
	if data["date_end"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.insert_datetime<=$%d ", argId))
		args = append(args, data["date_end"]+" 23:59:59")
		argId++
	}

	if data["onyma_dognum_null"] == "true" {
		setValues = append(setValues, "pro.dognum=''")
		//Если нет данных онима, нет смысла проверять скорости
		data["onyma_speed_null"] = "false"
		data["onyma_speed_error"] = "false"
	}

	if data["onyma_speed_null"] == "true" {
		setValues = append(setValues, "pro.dognum<>'' and (pro.in_speed_onyma='' OR pro.out_speed_onyma='' OR pro.in_speed_onyma='0' OR pro.out_speed_onyma='0')")

	}
	if data["onyma_speed_error"] == "true" {
		setValues = append(setValues, "pro.dognum<>'' and (pro.in_speed_onyma<>pro.in_speed_router OR pro.out_speed_onyma<>pro.out_speed_router) and pro.in_speed_onyma<>'' and pro.out_speed_onyma<>'' and pro.in_speed_onyma<>'0' and pro.out_speed_onyma<>'0'")
	}

	if data["client_status_onyma"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.client_status_onyma=$%d", argId))
		args = append(args, data["client_status_onyma"])
		argId++
	}
	if data["router_onyma_speed_id"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.router_onyma_speed_id=$%d", argId))
		args = append(args, data["router_onyma_speed_id"])
		argId++
	}
	if data["company_name"] != "" {
		if like == "true" {
			setValues = append(setValues, fmt.Sprintf("pro.company_name LIKE ($%d)", argId))
			args = append(args, "%"+data["company_name"]+"%")
		} else {
			setValues = append(setValues, fmt.Sprintf("pro.company_name = $%d", argId))
			args = append(args, data["company_name"])
		}

		argId++
	}
	if data["interface_name"] != "" {
		if like == "true" {
			setValues = append(setValues, fmt.Sprintf("pro.interface_name LIKE ($%d)", argId))
			args = append(args, "%"+data["interface_name"]+"%")
		} else {
			setValues = append(setValues, fmt.Sprintf("pro.interface_name=$%d", argId))
			args = append(args, data["interface_name"])
		}
		argId++
	}
	if data["ip_interface"] != "" {
		if like == "true" {
			setValues = append(setValues, fmt.Sprintf("pro.ip_interface LIKE ($%d)", argId))
			args = append(args, "%"+data["ip_interface"]+"%")
		} else {
			setValues = append(setValues, fmt.Sprintf("pro.ip_interface=$%d", argId))
			args = append(args, data["ip_interface"])
		}
		argId++
	}
	if data["problem_status"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.problem_status=$%d", argId))
		args = append(args, data["problem_status"])
		argId++
	}
	//Через запятую, без кавычек и без пробелов
	//Например: Закрыто,Test,Дублирование IP
	if data["ip_interface_not_in"] != "" {
		val_m := strings.Split(data["ip_interface_not_in"], ",")
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
		setValues = append(setValues, fmt.Sprintf("pro.ip_interface NOT IN (%s)", s_tmp))
	}


	//Через запятую, без кавычек и без пробелов
	//Например: Закрыто,Test,Дублирование IP
	if data["problem_status_in"] != "" {
		val_m := strings.Split(data["problem_status_in"], ",")
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
		setValues = append(setValues, fmt.Sprintf("pro.problem_status IN (%s)", s_tmp))
	}

	//Через запятую, без кавычек и без пробелов
	//Например: 2942d369-209a-4fa6-aca2-832ca4e51465,ad30ab32-981a-4c80-9d51-fa1f8629b389
	if len(data["not_in_ros_ids"]) != 0 {
		val_m := strings.Split(data["not_in_ros_ids"], ",")
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
		setValues = append(setValues, fmt.Sprintf("pro.router_onyma_speed_id NOT IN (%s)", s_tmp))
	}
	//Через запятую, без кавычек и без пробелов
	//Например: Закрыто,Test,Дублирование IP
	if data["problem_status_not_in"] != "" {
		val_m := strings.Split(data["problem_status_not_in"], ",")
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
		setValues = append(setValues, fmt.Sprintf("pro.problem_status NOT IN (%s)", s_tmp))
	}

	if data["branch_service"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.branch_service=$%d", argId))
		args = append(args, data["branch_service"])
		argId++
	}
	if data["dognum"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.dognum=$%d", argId))
		args = append(args, data["dognum"])
		argId++
	}
	if data["router_name"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.router_name=$%d", argId))
		args = append(args, data["router_name"])
		argId++
	}
	if data["clsrv"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.clsrv=$%d", argId))
		args = append(args, data["clsrv"])
		argId++
	}
	setQuery := strings.Join(setValues, " and ")
	//Если не учитывать фильтр
	if data["check_filter"] == "false" || data["check_filter"] == "" {
		sql := `SELECT DISTINCT pro.*, proh.problem_router_onyma_speed_id,proh.user_id,proh.old_val,proh.new_val,proh.msg,proh.created_at,
		u.name,dppp.dppp_name FROM %s as pro 
		LEFT OUTER JOIN %s as proh ON proh.problem_router_onyma_speed_id = pro.id
		LEFT OUTER JOIN %s as u ON proh.user_id = u.id
		LEFT OUTER JOIN %s as dppp ON pro.ip_interface = dppp.real_ip and to_char(pro.insert_datetime,'YYYY-MM-DD')=to_char(dppp.created_at,'YYYY-MM-DD')
			WHERE %s  order by pro.insert_datetime`
		query := fmt.Sprintf(sql, problemRouterOnymaTable, problemRouterOnymaHistory, usersTable, dpppAdressBaseTable, setQuery)
		err := r.db.Select(&lists, query, args...)
		//logrus.Infof("%s %v", query, args)
		if err != nil {
			return lists, err
		}
	}

	//Тут учитываю фильтр
	if data["check_filter"] == "true" {
		sql := `SELECT DISTINCT pro.*, proh.problem_router_onyma_speed_id,proh.user_id,proh.old_val,proh.new_val,proh.msg,proh.created_at,
		u.name,dppp.dppp_name FROM %s as pro 
		LEFT OUTER JOIN %s as proh ON proh.problem_router_onyma_speed_id = pro.id
		LEFT OUTER JOIN %s as u ON proh.user_id = u.id
		LEFT OUTER JOIN %s as dppp ON pro.ip_interface = dppp.real_ip  and to_char(pro.insert_datetime,'YYYY-MM-DD')=to_char(dppp.created_at,'YYYY-MM-DD')
			WHERE %s AND pro.id NOT IN 
		(SELECT DISTINCT pro2.id FROM %s as fro INNER JOIN %s as pro2 ON pro2.interface_name=fro.filter_val 
			 AND fro.filter_type='interface_name')
 	AND pro.id NOT IN 
		(SELECT DISTINCT pro3.id FROM %s as fro2 INNER JOIN %s as pro3 ON pro3.ip_interface=fro2.filter_val 
			 AND fro2.filter_type ='ip') order by pro.insert_datetime`

		query := fmt.Sprintf(sql, problemRouterOnymaTable, problemRouterOnymaHistory, usersTable, dpppAdressBaseTable, setQuery, filterRouterOnyma, problemRouterOnymaTable, filterRouterOnyma, problemRouterOnymaTable)
		err := r.db.Select(&lists, query, args...)
		logrus.Infof("%s %v", query, args)
		if err != nil {
			return lists, err
		}

	}

	return lists, error(nil)

}
func (r *RouterOnymaPostgres) GetCountAllProblemRouterOnymaSpeedsByQuery(data map[string]string) (int, error) {
	lists, err := r.GetAllProblemRouterOnymaSpeedsByQuery(data)
	if err != nil {
		return 0, err
	}
	return len(lists), error(nil)
}

func (r *RouterOnymaPostgres) GetControlTimePauseByQuery(data map[string]string) ([]tc.ControlTimePause, error) {
	//logrus.Infof("GetAllProblemRouterOnymaSpeedsByQuery data=%v",data)
	var lists []tc.ControlTimePause
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if data["date"] != "" {
		setValues = append(setValues, fmt.Sprintf("ctp.created_at>=$%d and ctp.created_at<=$%d", argId, argId+1))
		args = append(args, data["date"]+" 00:00:00")
		args = append(args, data["date"]+" 23:59:59")
		argId++
		argId++
	}

	if data["control_status"] != "" {
		setValues = append(setValues, fmt.Sprintf("ctp.control_status=$%d", argId))
		cs, err := strconv.Atoi(data["control_status"])
		if err != nil {
			return lists, err
		}
		args = append(args, cs)
		argId++
	}
	if data["router_onyma_speed_id"] != "" {
		setValues = append(setValues, fmt.Sprintf("ctp.router_onyma_speed_id=$%d", argId))
		args = append(args, data["router_onyma_speed_id"])
		argId++
	}

	if data["interface_name"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.interface_name=$%d", argId))
		args = append(args, data["interface_name"])
		argId++
	}
	if data["ip_interface"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.ip_interface=$%d", argId))
		args = append(args, data["ip_interface"])
		argId++
	}

	if data["branch_service"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.branch_service=$%d", argId))
		args = append(args, data["branch_service"])
		argId++
	}
	if data["dognum"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.dognum=$%d", argId))
		args = append(args, data["dognum"])
		argId++
	}
	if data["router_name"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.router_name=$%d", argId))
		args = append(args, data["router_name"])
		argId++
	}
	if data["clsrv"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.clsrv=$%d", argId))
		args = append(args, data["clsrv"])
		argId++
	}
	if data["iface_shutdown_router"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.iface_shutdown_router=$%d", argId))
		args = append(args, data["iface_shutdown_router"])
		argId++
	}
	if data["client_status_onyma"] != "" {
		setValues = append(setValues, fmt.Sprintf("pro.client_status_onyma=$%d", argId))
		args = append(args, data["client_status_onyma"])
		argId++
	}
	setQuery := strings.Join(setValues, " and ")
	sql := `SELECT ctp.*, pro.branch_service,pro.router_name,pro.interface_name,pro.interface_description
	,pro.in_policy_router,pro.in_speed_router,pro.out_policy_router,pro.out_speed_router,pro.ip_interface,pro.branch_contract
	,pro.dognum,pro.clsrv,pro.company_name,pro.in_speed_onyma,pro.out_speed_onyma,pro.insert_datetime,pro.iface_shutdown_router
	,pro.client_status_onyma FROM %s as ctp 
		INNER JOIN %s as pro ON ctp.router_onyma_speed_id = pro.id
			WHERE %s  order by  pro.router_name, pro.company_name,pro.insert_datetime`
	query := fmt.Sprintf(sql, controlTimePause, routerOnymaTable, setQuery)
	err := r.db.Select(&lists, query, args...)
	//logrus.Infof("%s %v", query, args)
	if err != nil {
		return lists, err
	}
	return lists, error(nil)

}

func (r *RouterOnymaPostgres) GetAllRouterOnymaSpeedsByQuery(data map[string]string) ([]tc.RouterOnyma, error) {
	var lists []tc.RouterOnyma
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if len(data["date"]) == 0 {
		ttime := time.Now()
		data["date"] = ttime.Format("2006-01-02")
	}
	setValues = append(setValues, fmt.Sprintf("ros.insert_datetime>=$%d and ros.insert_datetime<=$%d", argId, argId+1))
	args = append(args, data["date"]+" 00:00:00")
	args = append(args, data["date"]+" 23:59:59")
	argId++
	argId++

	if data["id"] != "" {
		setValues = append(setValues, fmt.Sprintf("ros.id=$%d", argId))
		args = append(args, data["id"])
		argId++
	}
	if data["interface_name"] != "" {
		setValues = append(setValues, fmt.Sprintf("ros.interface_name=$%d", argId))
		args = append(args, data["interface_name"])
		argId++
	}
	if data["ip_interface"] != "" {
		setValues = append(setValues, fmt.Sprintf("ros.ip_interface=$%d", argId))
		args = append(args, data["ip_interface"])
		argId++
	}

	if data["branch_service"] != "" {
		setValues = append(setValues, fmt.Sprintf("ros.branch_service=$%d", argId))
		args = append(args, data["branch_service"])
		argId++
	}
	if data["dognum"] != "" {
		setValues = append(setValues, fmt.Sprintf("ros.dognum=$%d", argId))
		args = append(args, data["dognum"])
		argId++
	}
	if data["router_name"] != "" {
		setValues = append(setValues, fmt.Sprintf("ros.router_name=$%d", argId))
		args = append(args, data["router_name"])
		argId++
	}
	if data["clsrv"] != "" {
		setValues = append(setValues, fmt.Sprintf("ros.clsrv=$%d", argId))
		args = append(args, data["clsrv"])
		argId++
	}
	if data["iface_shutdown_router"] != "" {
		setValues = append(setValues, fmt.Sprintf("ros.iface_shutdown_router=$%d", argId))
		isr, err := strconv.Atoi(data["iface_shutdown_router"])
		if err != nil {
			return lists, err
		}
		args = append(args, isr)
		argId++
	}
	if data["client_status_onyma"] != "" {
		setValues = append(setValues, fmt.Sprintf("ros.client_status_onyma=$%d", argId))
		cso, err := strconv.Atoi(data["client_status_onyma"])
		if err != nil {
			return lists, err
		}
		args = append(args, cso)
		argId++
	}
	setQuery := strings.Join(setValues, " and ")
	sql := `SELECT ros.*,dppp.dppp_name  FROM %s as ros 
		LEFT OUTER JOIN %s as dppp ON ros.ip_interface = dppp.real_ip and to_char(ros.insert_datetime,'YYYY-MM-DD')=to_char(dppp.created_at,'YYYY-MM-DD')
		WHERE %s  order by ros.router_name, ros.company_name,ros.insert_datetime`
	query := fmt.Sprintf(sql, routerOnymaTable, dpppAdressBaseTable, setQuery)
	err := r.db.Select(&lists, query, args...)
	logrus.Infof("%s %v", query, args)
	if err != nil {
		return lists, err
	}
	return lists, error(nil)

}
func (r *RouterOnymaPostgres) GetDateListProblemRouterOnymaSpeed() ([]string, error) {
	var lists []string
	query := fmt.Sprintf("SELECT DISTINCT to_char(insert_datetime,'YYYY-MM-DD') FROM %s order by to_char(insert_datetime,'YYYY-MM-DD')", problemRouterOnymaTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *RouterOnymaPostgres) UpdateProblemRouterOnymaSpeeds(id, userId string, input tc.UpdateProblemRouterOnymaSpeedInput) error {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Dognum != "" {
		setValues = append(setValues, fmt.Sprintf("dognum=$%d", argId))
		args = append(args, input.Dognum)
		argId++
	}
	if input.Clsrv != "" {
		setValues = append(setValues, fmt.Sprintf("clsrv=$%d", argId))
		args = append(args, input.Clsrv)
		argId++
	}
	if input.Company_name != "" {
		setValues = append(setValues, fmt.Sprintf("company_name=$%d", argId))
		args = append(args, input.Company_name)
		argId++
	}

	if input.In_speed_onyma != "" {
		setValues = append(setValues, fmt.Sprintf("in_speed_onyma=$%d", argId))
		args = append(args, input.In_speed_onyma)
		argId++
	}

	if input.Out_speed_onyma != "" {
		setValues = append(setValues, fmt.Sprintf("out_speed_onyma=$%d", argId))
		args = append(args, input.Out_speed_onyma)
		argId++
	}

	if input.In_speed_router != "" {
		setValues = append(setValues, fmt.Sprintf("in_speed_router=$%d", argId))
		args = append(args, input.In_speed_router)
		argId++
	}

	if input.Out_speed_router != "" {
		setValues = append(setValues, fmt.Sprintf("out_speed_router=$%d", argId))
		args = append(args, input.Out_speed_router)
		argId++
	}

	if input.Interface_name != "" {
		setValues = append(setValues, fmt.Sprintf("interface_name=$%d", argId))
		args = append(args, input.Interface_name)
		argId++
	}
	if input.Branch_contract != "" {
		setValues = append(setValues, fmt.Sprintf("branch_contract=$%d", argId))
		args = append(args, input.Branch_contract)
		argId++
	}
	if input.Problem_status != "" {
		setValues = append(setValues, fmt.Sprintf("problem_status=$%d", argId))
		args = append(args, input.Problem_status)
		argId++
		if input.Problem_status_old != "" {
			setValues = append(setValues, fmt.Sprintf("problem_status_old=$%d", argId))
			args = append(args, input.Problem_status_old)
			argId++
		}

	}
	setValues = append(setValues, fmt.Sprintf("updated_at=$%d", argId))
	args = append(args, "now()")
	argId++

	args = append(args, id)
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s set %s  WHERE  id=$%d", problemRouterOnymaTable, setQuery, argId)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	//Добавляю сразу в историю изменений
	data := tc.ProblemRouterOnymaHistory{Msg: input.Msg, New_val: input.Problem_status, Problem_router_onyma_speed_id: input.Id, User_id: userId, Old_val: input.Problem_status_old}
	_, err = r.InsertRouterOnymaHistory(data)
	return err
}
func (r *RouterOnymaPostgres) UpdateProblemStatusByInterface(id, userId string, input tc.UpdateProblemRouterOnymaSpeedInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	//Информация по проблемной записи
	info, err := r.GetAllProblemRouterOnymaSpeedsByQuery(map[string]string{"id": id})
	if err != nil {
		return err
	}
	//Сохроняю старый статус
	setValues = append(setValues, fmt.Sprintf("problem_status_old=$%d", argId))
	args = append(args, info[0].Problem_status)
	argId++

	if input.Problem_status != "" {
		setValues = append(setValues, fmt.Sprintf("problem_status=$%d", argId))
		args = append(args, input.Problem_status)
		argId++
	}
	setValues = append(setValues, fmt.Sprintf("updated_at=$%d", argId))
	args = append(args, "now()")
	argId++

	args = append(args, info[0].Interface_name)
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s set %s  WHERE interface_name=$%d", problemRouterOnymaTable, setQuery, argId)

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	//Добавляю сразу в историю изменений
	data := tc.ProblemRouterOnymaHistory{Msg: input.Msg, New_val: input.Problem_status, Problem_router_onyma_speed_id: input.Id, User_id: userId, Old_val: info[0].Problem_status}
	_, err = r.InsertRouterOnymaHistory(data)
	return err
}

func (r *RouterOnymaPostgres) InsertFilterRouterOnyma(fro tc.FilterRouterOnyma) (string, error) {
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
		query := fmt.Sprintf("INSERT INTO %s (id,filter_type,filter_val,filter_desc,user_id,router_name) values ($1,$2,$3,$4,$5,$6)", filterRouterOnyma)
		_, err = tx.Exec(query, id.String(), fro.Filter_type, filter_val, fro.Filter_desc, fro.User_id, fro.Router_name)
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}
	return id.String(), tx.Commit()

}
func (r *RouterOnymaPostgres) GetFilterRouterOnyma(filter tc.FilterRouterOnymaSearch) ([]tc.FilterRouterOnymaSearch, error) {
	var lists []tc.FilterRouterOnymaSearch
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

	query := fmt.Sprintf("SELECT f.*,u.name FROM %s as f INNER JOIN %s as u ON f.user_id=u.id WHERE %s  order by f.created_at DESC", filterRouterOnyma, usersTable, setQuery)
	err := r.db.Select(&lists, query, args...)
	logrus.Infof("%s %v", query, args)

	/*
		setQuery := strings.Join(setValues, ", ")
		query := fmt.Sprintf("UPDATE %s set %s  WHERE  id=$%d and user_id=$%d", categoryTable, setQuery, argId, argId+1)

		_, err := r.db.Exec(query, args...)
	*/
	return lists, err
}
func (r *RouterOnymaPostgres) DeleteFilter(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", filterRouterOnyma)
	_, err := r.db.Exec(query, id)
	return err
}
