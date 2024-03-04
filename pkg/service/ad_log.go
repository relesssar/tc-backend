package service

import (
	"bufio"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	tc "tc_kaztranscom_backend_go"
	"tc_kaztranscom_backend_go/pkg/repository"
	"time"
)

type ADLogService struct {
	repo repository.ADLog
}

func NewADLogService(repo repository.ADLog) *ADLogService {
	return &ADLogService{repo: repo}
}

//В Windows 2008 событие успешного входа имеет идентификатор Event ID 4624, а logoff — Event ID 4634
// LogParser.exe "SELECT TimeGenerated,EventID,Strings INTO C:\log_ad.csv FROM C:\log_ad.evtx WHERE EventID=4634 OR EventID=4624" -i:EVT

func (s *ADLogService) CheckLogFileCSV() error {
	//Проверяю если есть файл, то гружу в базу и удаляю после загрузки
	filename := os.Getenv("AD_LOG_FILE")
	file, err := os.Open(filename)
	if err != nil {
		logrus.Infof("%s %s", filename, err)
		return err
	}
	logrus.Infof("Найден файл лога AD, загружаю данные в базу. %s", filename)

	fileScanner := bufio.NewScanner(file)

	lineCount := 0
	// read line by line
	for fileScanner.Scan() {
		lineCount++
		line := fileScanner.Text()
		str := strings.Split(line, ",")
		var login, date_time_str, ip string
		var eventid int
		for i, v := range str {
			if i == 0 {
				t, _ := time.Parse("2006-01-02 15:04:05", v)
				date_time_str = t.Format("2006-01-02 15:04:05")
				continue
			}
			if i == 1 {
				eventid, _ = strconv.Atoi(v)
				continue
			}
			if i == 2 {
				if eventid == 4624 {
					login_a := strings.Split(v, "|")
					login = login_a[5]
					ip = login_a[18]
				}
				if eventid == 4634 {
					login_a := strings.Split(v, "|")
					login = login_a[1]
				}
				continue
			}
		}
		if len(login) == 0 {
			continue
		}
		if login == "ANONYMOUS LOGON" || login == "FILE_SERVER$" || login[:3] == "IUS" || login[:3] == "ASS" || login[:3] == "ALK" || login[:3] == "ASN" || login[:3] == "UFS" || login[:3] == "VPR" || login[:3] == "pro" || login[:3] == "PRO" || login[:3] == "KMM" || login[:3] == "ZFS" || login[:3] == "ALM" || login[:3] == "LOC" || login[:3] == "KMN" || login[:3] == "ALS" || login[:3] == "UFP" || login[:3] == "AKF" || login[:3] == "AGN" || login[:3] == "AGP" || login[:3] == "ALP" || login[:3] == "AKB" || login[:3] == "VFP" || login[:3] == "ASP" || login[:3] == "MFP" || login[:3] == "ZFP" || login[:3] == "ALN" || login[:3] == "YFP" || login[:3] == "ccr" {
			continue
		}
		//Сохроняю в базу
		id, err := s.repo.InsertADLog(tc.ADLog{DateTime: date_time_str, Ad_login: login, Eventid: eventid, Ip: ip, Str: line})
		fmt.Printf("ok-%s %s\n", id, err)
	}
	//fmt.Printf("\n[%q]",line_parse)
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		fmt.Printf("Error while reading file: %s", err)
	}
	file.Close()

	os.Remove(filename)
	fmt.Printf("прочитано %d строк", lineCount)
	logrus.Info("Удалил файл лога AD, данные загружены в базу.")
	logrus.Info("Заношу информацию о департаменте и ФИО, для новых записей.")
	err = s.repo.UpdateDepartmentInfo()
	return err
}

func (s *ADLogService) GetAllByQuery(data map[string]string) ([]tc.ADLog, error) {
	return s.repo.GetAllByQuery(data)
}

/*создание ексель файла для выгрузки с интерфейса*/
func (s *ADLogService) CreateExcell(data []tc.ADLog) (string, error) {
	//fmt.Printf("%v",data)
	f := excelize.NewFile()

	// Create a new sheet.
	sh := "tc.ktc.kz_ad_log"

	index := f.NewSheet(sh)
	f.SetColWidth(sh, "A", "Z", 20)
	// Set value of a cell.
	//line := "ФИО;Департамент;Тип подключения \n(удаленно/офис);дата;Время входа;Время выхода"
	char_m := strings.Split("A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Z", ", ")

	cols := make(map[int]string)
	i := 0
	for _, v := range char_m {
		i++
		cols[i] = v
	}
	for _, v := range char_m {
		for _, v2 := range char_m {
			i++
			cols[i] = v + v2
		}
	}
	var date_str string
	data_m := make(map[string]map[string]map[string]string)
	data_user_info_m := make(map[string]map[string]string)
	i = 4
	//Формирование данных в мапу
	for _, v := range data {
		ad_login := strings.ToLower(v.Ad_login)
		_, ok := data_m[ad_login]
		if !ok {
			m := make(map[string]map[string]string)
			data_m[ad_login] = m
			i = 4

			date_parse, _ := time.Parse("2006-01-02T15:04:05Z", v.DateTime)
			t, _ := time.Parse("2006-01-02", fmt.Sprintf("%s-01", date_parse.Format("2006-01")))
			//Формирование матрицы по дням
			for d := 0; d < 31; d++ {
				date_tmp_str := t.AddDate(0, 0, d)
				date_str = date_tmp_str.Format("2006.01.02")
				_, ok = data_m[ad_login][date_str]
				if !ok {
					mm := make(map[string]string)
					data_m[ad_login][date_str] = mm
				}
				data_m[ad_login][date_str]["cols"] = strconv.Itoa(i)
				i = i + 4
			}
		}
		_, ok = data_user_info_m[ad_login]
		if !ok {
			mm := make(map[string]string)
			data_user_info_m[ad_login] = mm
		}
		//logrus.Infof("%s %s %s", ad_login, v.Fullname.String,v.Department.String)
		if v.Fullname.Valid {
			data_user_info_m[ad_login]["fullname"] = v.Fullname.String
		}
		if v.Department.Valid {
			data_user_info_m[ad_login]["department"] = v.Department.String
		}
		//logrus.Infof("%v", data_user_info_m[ad_login])
		date_parse, _ := time.Parse("2006-01-02T15:04:05Z", v.DateTime)
		date_str = date_parse.Format("2006.01.02")

		time_tmp := date_parse.Format("15:04:05")
		//LogIn
		if v.Eventid == 4624 {
			//Если уже есть такое время , то не ставлю , чтобы не было одиноковых
			if strings.Contains(data_m[ad_login][date_str]["in"], time_tmp) == false {
				if len(data_m[ad_login][date_str]["in"]) != 0 {
					data_m[ad_login][date_str]["in"] += ", "
				}
				data_m[ad_login][date_str]["in"] = data_m[ad_login][date_str]["in"] + time_tmp
			}

		}
		//LogOut
		if v.Eventid == 4634 {
			//Если уже есть такое время , то не ставлю , чтобы не было одиноковых
			if strings.Contains(data_m[ad_login][date_str]["out"], time_tmp) == false {
				if len(data_m[ad_login][date_str]["out"]) != 0 {
					data_m[ad_login][date_str]["out"] += ", "
				}
				data_m[ad_login][date_str]["out"] = data_m[ad_login][date_str]["out"] + time_tmp
			}
		}

		type_connect := ""
		if v.Ip != "" && len(v.Ip) >= 3 {
			type_connect = "Удалённо " + v.Ip
			if v.Ip[:3] == "192" || v.Ip[:3] == "172" || v.Ip[:3] == "10." {
				type_connect = "Офис или VPN" + v.Ip
			}
		}
		//Если уже есть такой IP , то не ставлю , чтобы не было одиноковых
		if strings.Contains(data_m[v.Ad_login][date_str]["type_connect"], type_connect) == false {
			if len(data_m[ad_login][date_str]["type_connect"]) != 0 {
				data_m[ad_login][date_str]["type_connect"] += ", "
			}
			data_m[ad_login][date_str]["type_connect"] = type_connect
		}

		//fmt.Printf("\n[%s] %d \n [in]=%s\n [out]=%s\n [cols]=%s",date_str,v.Eventid,data_m[v.Ad_login][date_str]["in"],data_m[v.Ad_login][date_str]["out"],data_m[v.Ad_login][date_str]["cols"])

		//fmt.Printf("\n****\n%v",data_m)

	}
	//fmt.Printf("%v",data_m)
	c := 3
	r := 2
	f.SetCellValue(sh, fmt.Sprintf("A%d", 1), "ФИО")
	f.SetCellValue(sh, fmt.Sprintf("B%d", 1), "Департамент")

	for i, v := range data_m {
		if data_user_info_m[i]["fullname"] != "" {
			f.SetCellValue(sh, fmt.Sprintf("A%d", r), data_user_info_m[i]["fullname"])
		} else {
			f.SetCellValue(sh, fmt.Sprintf("A%d", r), i)
		}
		f.SetCellValue(sh, fmt.Sprintf("B%d", r), data_user_info_m[i]["department"])
		for date_str, v2 := range v {

			c, _ = strconv.Atoi(v2["cols"])

			f.SetCellValue(sh, fmt.Sprintf("%s%d", cols[c], 1), "Тип подключения \\n(удаленно/офис)")
			f.SetCellValue(sh, fmt.Sprintf("%s%d", cols[c+1], 1), "Дата")
			f.SetCellValue(sh, fmt.Sprintf("%s%d", cols[c+2], 1), "Время входа")
			f.SetCellValue(sh, fmt.Sprintf("%s%d", cols[c+3], 1), "Время выхода")

			f.SetCellValue(sh, fmt.Sprintf("%s%d", cols[c], r), v2["type_connect"])
			f.SetCellValue(sh, fmt.Sprintf("%s%d", cols[c+1], r), date_str)
			f.SetCellValue(sh, fmt.Sprintf("%s%d", cols[c+2], r), v2["in"])
			f.SetCellValue(sh, fmt.Sprintf("%s%d", cols[c+3], r), v2["out"])

		}
		r++
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	id := uuid.New()
	filename := fmt.Sprintf("%s.xlsx", id.String())
	path_file := fmt.Sprintf("/%s/%s", os.Getenv("PATH_DOWNLOAD"), filename)
	logrus.Infof("CreateExceAdlog:сохраняю файл %s", path_file)

	if err := f.SaveAs(path_file); err != nil {
		logrus.Errorf("CreateExceAdlog: ошибка запси на диск %s", path_file)
		logrus.Errorf("CreateExceAdlog:  %s", err)
		return "", err
	}

	logrus.Infof("CreateExceAdlog: файл сохранён на диск %s", path_file)
	return filename, error(nil)
}
