package service

import (
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

type RouterOnymaService struct {
	repo repository.RouterOnyma
}

func NewRouterOnymaService(repo repository.RouterOnyma) *RouterOnymaService {
	return &RouterOnymaService{repo: repo}
}

func (s *RouterOnymaService) InsertFilterRouterOnyma(fro tc.FilterRouterOnyma) (string, error) {
	return s.repo.InsertFilterRouterOnyma(fro)
}
func (s *RouterOnymaService) GetFilterRouterOnyma(filter tc.FilterRouterOnymaSearch) ([]tc.FilterRouterOnymaSearch, error) {

	return s.repo.GetFilterRouterOnyma(filter)
}
func (s *RouterOnymaService) DeleteFilter(id string) error {

	return s.repo.DeleteFilter(id)
}

func (s *RouterOnymaService) InsertRouterOnymaHistory(proh tc.ProblemRouterOnymaHistory) (string, error) {

	return s.repo.InsertRouterOnymaHistory(proh)
}
func (s *RouterOnymaService) InsertRouterOnymaSpeeds(routerOnyma tc.RouterOnyma) (string, error) {

	return s.repo.InsertRouterOnymaSpeeds(routerOnyma)
}
func (s *RouterOnymaService) InsertRouterOnymaSpeedsAll(routerOnyma []tc.RouterOnyma) (string, error) {

	return s.repo.InsertRouterOnymaSpeedsAll(routerOnyma)
}
func (s *RouterOnymaService) InsertProblemRouterOnymaSpeeds(pRo tc.ProblemRouterOnyma) (string, error) {

	return s.repo.InsertProblemRouterOnymaSpeeds(pRo)
}
func (s *RouterOnymaService) GetAllRouterOnymaSpeedsByDate(date string) ([]tc.RouterOnyma, error) {

	return s.repo.GetAllRouterOnymaSpeedsByDate(date)
}
func (s *RouterOnymaService) GetAllRouterOnymaSpeedsByQuery(data map[string]string) ([]tc.RouterOnyma, error) {

	return s.repo.GetAllRouterOnymaSpeedsByQuery(data)
}
func (s *RouterOnymaService) GetCountAllProblemRouterOnymaSpeedsByQuery(data map[string]string) (int, error) {

	return s.repo.GetCountAllProblemRouterOnymaSpeedsByQuery(data)
}
func (s *RouterOnymaService) GetAllProblemRouterOnymaSpeedsByQuery(data map[string]string) ([]tc.ProblemRouterOnymaQuery, error) {

	return s.repo.GetAllProblemRouterOnymaSpeedsByQuery(data)
}
func (s *RouterOnymaService) GetControlTimePauseByQuery(data map[string]string) ([]tc.ControlTimePause, error) {
	return s.repo.GetControlTimePauseByQuery(data)
}
func (s *RouterOnymaService) GetDateListProblemRouterOnymaSpeed() ([]string, error) {

	return s.repo.GetDateListProblemRouterOnymaSpeed()
}
func (s *RouterOnymaService) GetRouterOnymaHistory(data map[string]string) ([]tc.ProblemRouterOnymaHistorySearch, error) {

	return s.repo.GetRouterOnymaHistory(data)
}
func (s *RouterOnymaService) GetStatisticRouterOnymaSpeedsGroupByRouterByQuery(data map[string]string) ([]tc.RouterOnymaGroupByRouter, error) {

	return s.repo.GetStatisticRouterOnymaSpeedsGroupByRouterByQuery(data)
}
func (s *RouterOnymaService) GetStatisticRouterOnymaSpeedsGroupByInsertDateByQuery(data map[string]string) ([]tc.RouterOnymaGroupByInsertDate, error) {

	return s.repo.GetStatisticRouterOnymaSpeedsGroupByInsertDateByQuery(data)
}
func (s *RouterOnymaService) UpdateProblemRouterOnymaSpeeds(id, userId string, data tc.UpdateProblemRouterOnymaSpeedInput) error {

	return s.repo.UpdateProblemRouterOnymaSpeeds(id, userId, data)
}
func (s *RouterOnymaService) UpdateProblemStatusByInterface(id, userId string, data tc.UpdateProblemRouterOnymaSpeedInput) error {

	return s.repo.UpdateProblemStatusByInterface(id, userId, data)
}

/*создание ексель файла для выгрузки с интерфейса:> Контроль временного отключения*/
func (s *RouterOnymaService) CreateExcellControlTimePause(data []tc.ControlTimePause) (string, error) {
	f := excelize.NewFile()
	// Create a new sheet.
	sh := "Контроль временного отключения"
	index := f.NewSheet(sh)
	// Set value of a cell.
	rows := 1
	line := "№;Branch\nService;Router Name;Date and Time;Shutdown\n1-выкл,0-вкл;Interface Name;Interface Description;Interface\nIP Address;Gap;Branch\nContract;Номер ЛС;ID\nподключения;Статус клиента;Наименование\nкомпании;Onyma\nIP Address"
	line_m := strings.Split(line, ";")
	char_m := strings.Split("A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X", ", ")
	for i, v := range line_m {
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
		//rows+=1
	}
	rows += 1
	for i, v := range data {
		f.SetCellValue(sh, fmt.Sprintf("A%d", rows), fmt.Sprintf("%d", i))
		f.SetCellValue(sh, fmt.Sprintf("B%d", rows), v.Branch_service)
		f.SetCellValue(sh, fmt.Sprintf("C%d", rows), v.Router_name)
		f.SetCellValue(sh, fmt.Sprintf("D%d", rows), v.Insert_datetime)
		f.SetCellValue(sh, fmt.Sprintf("E%d", rows), v.Iface_shutdown_router)
		f.SetCellValue(sh, fmt.Sprintf("F%d", rows), v.Interface_name)
		f.SetCellValue(sh, fmt.Sprintf("G%d", rows), v.Interface_description)
		f.SetCellValue(sh, fmt.Sprintf("H%d", rows), v.Ip_interface)

		f.SetCellValue(sh, fmt.Sprintf("I%d", rows), "")

		f.SetCellValue(sh, fmt.Sprintf("J%d", rows), v.Branch_contract)
		f.SetCellValue(sh, fmt.Sprintf("K%d", rows), v.Dognum)
		f.SetCellValue(sh, fmt.Sprintf("L%d", rows), v.Clsrv)
		f.SetCellValue(sh, fmt.Sprintf("M%d", rows), v.Client_status_onyma)
		f.SetCellValue(sh, fmt.Sprintf("N%d", rows), v.Company_name)
		f.SetCellValue(sh, fmt.Sprintf("O%d", rows), v.Ip_interface)
		/*
			f.SetCellValue(sh, fmt.Sprintf("O%d", rows), v.)
			f.SetCellValue(sh, fmt.Sprintf("P%d", rows), v.)
			f.SetCellValue(sh, fmt.Sprintf("Q%d", rows),
			f.SetCellValue(sh, fmt.Sprintf("R%d", rows), v.In_speed_onyma)
			f.SetCellValue(sh, fmt.Sprintf("S%d", rows), v.Out_speed_onyma)
			f.SetCellValue(sh, fmt.Sprintf("T%d", rows), */
		rows += 1
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	id := uuid.New()
	filename := fmt.Sprintf("%s.xlsx", id.String())
	path_file := fmt.Sprintf("/%s/%s", os.Getenv("PATH_DOWNLOAD"), filename)
	logrus.Infof("CreateExcellControlClientPause:сохраняю файл %s", path_file)

	if err := f.SaveAs(path_file); err != nil {
		logrus.Errorf("CreateExcellControlClientPause: ошибка запси на диск %s", path_file)
		logrus.Errorf("CreateExcellControlClientPause:  %s", err)
		return "", err
	}

	logrus.Infof("CreateExcellControlClientPause: файл сохранён на диск %s", path_file)
	return filename, error(nil)
}

/*создание ексель файла для выгрузки с интерфейса*/
func (s *RouterOnymaService) CreateExcellProblemRouterOnymaSpeed(data []tc.ProblemRouterOnymaQuery) (string, error) {
	f := excelize.NewFile()
	// Create a new sheet.
	sh := "tc.ktc.kz_check_router"
	index := f.NewSheet(sh)
	// Set value of a cell.
	rows := 1
	line := "№;Статус;Branch\nService;Router Name;Date and Time;Interface Name;Interface Description;Input Policy\n(Router);Input Speed;Output Policy\n(Router);Output Speed;Interface\nIP Address;Gap;Branch\nContract;Номер ЛС;ID\nподключения;Наименование\nкомпании;Input Speed\n(Onyma);Output Speed\n(Onyma);Onyma\nIP Address"
	line_m := strings.Split(line, ";")
	char_m := strings.Split("A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X", ", ")
	for i, v := range line_m {
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
		//rows+=1
	}
	rows += 1
	for i, v := range data {
		f.SetCellValue(sh, fmt.Sprintf("A%d", rows), fmt.Sprintf("%d", i))
		f.SetCellValue(sh, fmt.Sprintf("B%d", rows), v.Problem_status)
		f.SetCellValue(sh, fmt.Sprintf("C%d", rows), v.Branch_service)
		f.SetCellValue(sh, fmt.Sprintf("D%d", rows), v.Router_name)
		f.SetCellValue(sh, fmt.Sprintf("E%d", rows), v.Insert_datetime)
		f.SetCellValue(sh, fmt.Sprintf("F%d", rows), v.Interface_name)
		f.SetCellValue(sh, fmt.Sprintf("G%d", rows), v.Interface_description)
		f.SetCellValue(sh, fmt.Sprintf("H%d", rows), v.In_policy_router)
		f.SetCellValue(sh, fmt.Sprintf("I%d", rows), v.In_speed_router)
		f.SetCellValue(sh, fmt.Sprintf("J%d", rows), v.Out_policy_router)
		f.SetCellValue(sh, fmt.Sprintf("K%d", rows), v.Out_speed_router)
		f.SetCellValue(sh, fmt.Sprintf("L%d", rows), v.Ip_interface)
		f.SetCellValue(sh, fmt.Sprintf("M%d", rows), "")
		f.SetCellValue(sh, fmt.Sprintf("N%d", rows), v.Branch_contract)
		f.SetCellValue(sh, fmt.Sprintf("O%d", rows), v.Dognum)
		f.SetCellValue(sh, fmt.Sprintf("P%d", rows), v.Clsrv)
		f.SetCellValue(sh, fmt.Sprintf("Q%d", rows), v.Company_name)
		f.SetCellValue(sh, fmt.Sprintf("R%d", rows), v.In_speed_onyma)
		f.SetCellValue(sh, fmt.Sprintf("S%d", rows), v.Out_speed_onyma)
		f.SetCellValue(sh, fmt.Sprintf("T%d", rows), v.Ip_interface)
		rows += 1
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	id := uuid.New()
	filename := fmt.Sprintf("%s.xlsx", id.String())
	path_file := fmt.Sprintf("/%s/%s", os.Getenv("PATH_DOWNLOAD"), filename)
	logrus.Infof("CreateExcellProblemRouterOnymaSpeed:сохраняю файл %s", path_file)

	if err := f.SaveAs(path_file); err != nil {
		logrus.Errorf("CreateExcellProblemRouterOnymaSpeed: ошибка запси на диск %s", path_file)
		logrus.Errorf("CreateExcellProblemRouterOnymaSpeed:  %s", err)
		return "", err
	}

	logrus.Infof("CreateExcellProblemRouterOnymaSpeed: файл сохранён на диск %s", path_file)
	return filename, error(nil)
}

/*создание ексель файла для выгрузки с интерфейса Форма 1*/
func (s *RouterOnymaService) CreateExcellProblemRouterOnymaForm1(data []tc.ProblemRouterOnymaQuery) (string, error) {

	//Формирую массив для файла
	data_new := make(map[string]map[string]int)
	interface_new := make(map[string]map[int]string)
	parse_interface := make(map[string]string)
	branch_a :=map[int]string {
		0:"МФ",
		1:"ЮФ",
		2:"ЗФ",
		3:"АФ",
		4:"НФ",
		5:"ШФ",
		6:"УФ",
		7:"ВФ",
	}
	for _, vv := range data {
		_, ok := parse_interface[vv.Interface_name]
		if ok { //Уже считали этот интерфейс, переходим к следующей итерации
			continue
		}
		parse_interface[vv.Interface_name] = vv.Interface_name

		_, ok = data_new[vv.Branch_service]
		if !ok {
			data_new[vv.Branch_service] = make(map[string]int)
			data_new[vv.Branch_service]["work"] = 0          //В обработке
			data_new[vv.Branch_service]["check"] = 0         //На проверку
			data_new[vv.Branch_service]["re_check"] = 0      //Дообследование в ДЭ
			data_new[vv.Branch_service]["test"] = 0          //Test
			data_new[vv.Branch_service]["close"] = 0         //Закрыто
			data_new[vv.Branch_service]["no_router"] = 0         //Нет на Роутере
			data_new[vv.Branch_service]["error_speed"] = 0   //* Не соответствие скоростей;
			data_new[vv.Branch_service]["filed_empty"] = 0   //* Пустые значения
			data_new[vv.Branch_service]["no_onyma"] = 0      //* Нет в ONYMA
			data_new[vv.Branch_service]["ip_dublicate"] = 0  //* Пересечение IP адресов
			data_new[vv.Branch_service]["not_allocated"] = 0 //* Не распределённые
		}
		_, ok = interface_new[vv.Branch_service+"work"][len(interface_new[vv.Branch_service+"work"])]
		if !ok {
			interface_new[vv.Branch_service+"work"] = make(map[int]string)
		}
		_, ok = interface_new[vv.Branch_service+"check"][len(interface_new[vv.Branch_service+"check"])]
		if !ok {
			interface_new[vv.Branch_service+"check"] = make(map[int]string)
		}
		_, ok = interface_new[vv.Branch_service+"re_check"][len(interface_new[vv.Branch_service+"re_check"])]
		if !ok {
			interface_new[vv.Branch_service+"re_check"] = make(map[int]string)
		}
		_, ok = interface_new[vv.Branch_service+"test"][len(interface_new[vv.Branch_service+"test"])]
		if !ok {
			interface_new[vv.Branch_service+"test"] = make(map[int]string)
		}
		_, ok = interface_new[vv.Branch_service+"close"][len(interface_new[vv.Branch_service+"close"])]
		if !ok {
			interface_new[vv.Branch_service+"close"] = make(map[int]string)
		}
		_, ok = interface_new[vv.Branch_service+"no_router"][len(interface_new[vv.Branch_service+"no_router"])]
		if !ok {
			interface_new[vv.Branch_service+"no_router"] = make(map[int]string)
		}
		_, ok = interface_new[vv.Branch_service+"error_speed"][len(interface_new[vv.Branch_service+"error_speed"])]
		if !ok {
			interface_new[vv.Branch_service+"error_speed"] = make(map[int]string)
		}
		_, ok = interface_new[vv.Branch_service+"filed_empty"][len(interface_new[vv.Branch_service+"filed_empty"])]
		if !ok {
			interface_new[vv.Branch_service+"filed_empty"] = make(map[int]string)
		}
		_, ok = interface_new[vv.Branch_service+"no_onyma"][len(interface_new[vv.Branch_service+"no_onyma"])]
		if !ok {
			interface_new[vv.Branch_service+"no_onyma"] = make(map[int]string)
		}
		_, ok = interface_new[vv.Branch_service+"ip_dublicate"][len(interface_new[vv.Branch_service+"ip_dublicate"])]
		if !ok {
			interface_new[vv.Branch_service+"ip_dublicate"] = make(map[int]string)
		}
		_, ok = interface_new[vv.Branch_service+"not_allocated"][len(interface_new[vv.Branch_service+"not_allocated"])]
		if !ok {
			interface_new[vv.Branch_service+"not_allocated"] = make(map[int]string)
		}




		switch vv.Problem_status {
			case "В обработке":
				data_new[vv.Branch_service]["work"] += 1
				data_new[vv.Branch_service]["not_allocated"] += 1
				interface_new[vv.Branch_service+"work"][len(interface_new[vv.Branch_service+"work"])+1] = vv.Interface_name
				interface_new[vv.Branch_service+"not_allocated"][len(interface_new[vv.Branch_service+"not_allocated"])+1] = vv.Interface_name
			case "На проверку":
				data_new[vv.Branch_service]["check"] += 1
				interface_new[vv.Branch_service+"check"][len(interface_new[vv.Branch_service+"check"])+1] = vv.Interface_name
			case "Дообследование в ДЭ":
				data_new[vv.Branch_service]["re_check"] += 1
				interface_new[vv.Branch_service+"re_check"][len(interface_new[vv.Branch_service+"re_check"])+1] = vv.Interface_name
			case "Test":
				data_new[vv.Branch_service]["test"] += 1
				interface_new[vv.Branch_service+"test"][len(interface_new[vv.Branch_service+"test"])+1] = vv.Interface_name
			case "Закрыто":
				data_new[vv.Branch_service]["close"] += 1
				interface_new[vv.Branch_service+"close"][len(interface_new[vv.Branch_service+"close"])+1] = vv.Interface_name
			case "Дублирование IP":
				data_new[vv.Branch_service]["ip_dublicate"] += 1
				interface_new[vv.Branch_service+"ip_dublicate"][len(interface_new[vv.Branch_service+"ip_dublicate"])+1] = vv.Interface_name
			case "Нет на Роутере":
				data_new[vv.Branch_service]["no_router"] += 1
				interface_new[vv.Branch_service+"no_router"][len(interface_new[vv.Branch_service+"no_router"])+1] = vv.Interface_name

		}
		//Не соответствие скоростей; error_speed

		if vv.In_speed_onyma != vv.In_speed_router || vv.Out_speed_onyma != vv.Out_speed_router {

			var razi, razo, io, oo float64 //Разница скоростей
			if vv.In_speed_router != "" && vv.In_speed_onyma != "" {
				ir, err := strconv.ParseFloat(vv.In_speed_router, 64)
				if err != nil {
					logrus.Infof("CreateExcellProblemRouterOnymaForm1 %s", err)
					return "error", err
				}
				io, err := strconv.ParseFloat(vv.In_speed_onyma, 64)
				if err != nil {
					logrus.Infof("CreateExcellProblemRouterOnymaForm1 %s", err)
					return "error", err
				}
				if io > 0 {
					razi = ir / io
				}
			}

			if vv.Out_speed_router != "" && vv.Out_speed_onyma != "" {
				or, err := strconv.ParseFloat(vv.Out_speed_router, 64)
				if err != nil {
					logrus.Infof("CreateExcellProblemRouterOnymaForm1 %s", err)
					return "error", err
				}
				oo, err = strconv.ParseFloat(vv.Out_speed_onyma, 64)
				if err != nil {
					logrus.Infof("CreateExcellProblemRouterOnymaForm1 %s", err)
					return "error", err
				}
				if oo > 0 {
					razo = or / oo
				}
			}
			//Скоростя не совпадают
			if razo > 1.1 || razi > 1.1 || io == 0 || oo == 0 {
				data_new[vv.Branch_service]["error_speed"] += 1
				interface_new[vv.Branch_service+"error_speed"][len(interface_new[vv.Branch_service+"error_speed"])+1] = vv.Interface_name

			}
		}
		//Пустые значения
		if vv.In_speed_router == "" || vv.In_speed_onyma == "" || vv.Out_speed_router == "" || vv.Out_speed_onyma == "" || vv.Company_name == "" {
			data_new[vv.Branch_service]["filed_empty"] += 1
			interface_new[vv.Branch_service+"filed_empty"][len(interface_new[vv.Branch_service+"filed_empty"])+1] = vv.Interface_name
		}
		//Нет в ONYMA
		if vv.Dognum == "" || vv.Clsrv == "" {
			data_new[vv.Branch_service]["no_onyma"] += 1
			interface_new[vv.Branch_service+"no_onyma"][len(interface_new[vv.Branch_service+"no_onyma"])+1] = vv.Interface_name
		}
		//Не распределённые
		if false {
			data_new[vv.Branch_service]["not_allocated"] += 1
			interface_new[vv.Branch_service+"not_allocated"][len(interface_new[vv.Branch_service+"not_allocated"])+1] = vv.Interface_name
		}
	}

	//logrus.Infof("CreateExcellProblemRouterOnymaForm1: data_new %v", data_new)

	f := excelize.NewFile()

	sh := "ТС"
	index := f.NewSheet(sh)
	// Set value of a cell.
	rows := 1
	line := "Филиал;Не соответствие скоростей;Пустые значения;Нет в ONYMA;Пересечение IP адресов;Не распределённые;;В обработке;На проверку;Нет на Роутере;Дообследование в ДЭ" //;Test;Закрыто"
	line_m := strings.Split(line, ";")
	char_m := strings.Split("A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X", ", ")
	for i, v := range line_m {
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
		//rows+=1
	}
	rows += 1

	for i, _ := range data_new {
		f.SetCellValue(sh, fmt.Sprintf("A%d", rows), i)
		f.SetCellValue(sh, fmt.Sprintf("B%d", rows), data_new[i]["error_speed"])
		f.SetCellValue(sh, fmt.Sprintf("C%d", rows), data_new[i]["filed_empty"])
		f.SetCellValue(sh, fmt.Sprintf("D%d", rows), data_new[i]["no_onyma"])
		f.SetCellValue(sh, fmt.Sprintf("E%d", rows), data_new[i]["ip_dublicate"])
		f.SetCellValue(sh, fmt.Sprintf("F%d", rows), data_new[i]["not_allocated"])
		f.SetCellValue(sh, fmt.Sprintf("G%d", rows), "")
		f.SetCellValue(sh, fmt.Sprintf("H%d", rows), data_new[i]["work"])
		f.SetCellValue(sh, fmt.Sprintf("I%d", rows), data_new[i]["check"])
		f.SetCellValue(sh, fmt.Sprintf("J%d", rows), data_new[i]["no_router"])
		f.SetCellValue(sh, fmt.Sprintf("K%d", rows), data_new[i]["re_check"])
		//f.SetCellValue(sh, fmt.Sprintf("K%d", rows), data_new[i]["test"])
		//f.SetCellValue(sh, fmt.Sprintf("L%d", rows), data_new[i]["close"])
		rows += 1
	}

	//вкладка В обработке
	for i,b :=range(branch_a){

		sh = "В обработке"
		index = f.NewSheet(sh)
		f.SetActiveSheet(index)
		rows = 1
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), b)
		rows += 1
		for _,v :=range(interface_new[b+"work"]){
			f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
			rows += 1
		}
	}
	//На проверку
	for i,b :=range(branch_a){
		sh = "На проверку"
		index = f.NewSheet(sh)
		f.SetActiveSheet(index)
		rows = 1
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), b)
		rows += 1
		for _,v :=range(interface_new[b+"check"]){
			f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
			rows += 1
		}
	}
	//Дообследование в ДЭ
	for i,b :=range(branch_a){
		sh = "Дообследование в ДЭ"
		index = f.NewSheet(sh)
		f.SetActiveSheet(index)
		rows = 1
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), b)
		rows += 1
		for _,v :=range(interface_new[b+"re_check"]){
			f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
			rows += 1
		}
	}
	//Test
	for i,b :=range(branch_a){
		sh = "Test"
		index = f.NewSheet(sh)
		f.SetActiveSheet(index)
		rows = 1
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), b)
		rows += 1
		for _,v :=range(interface_new[b+"test"]){
			f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
			rows += 1
		}
	}
	//Нет на Роутере
	for i,b :=range(branch_a){
		sh = "Нет на Роутере"
		index = f.NewSheet(sh)
		f.SetActiveSheet(index)
		rows = 1
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), b)
		rows += 1
		for _,v :=range(interface_new[b+"no_router"]){
			f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
			rows += 1
		}
	}
	//Не соответствие скоростей
	for i,b :=range(branch_a){
		sh = "Не соответствие скоростей"
		index = f.NewSheet(sh)
		f.SetActiveSheet(index)
		rows = 1
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), b)
		rows += 1
		for _,v :=range(interface_new[b+"error_speed"]){
			f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
			rows += 1
		}
	}
	//Пустые значения
	for i,b :=range(branch_a){
		sh = "Пустые значения"
		index = f.NewSheet(sh)
		f.SetActiveSheet(index)
		rows = 1
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), b)
		rows += 1
		for _,v :=range(interface_new[b+"filed_empty"]){
			f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
			rows += 1
		}
	}
	//Нет в ONYMA
	for i,b :=range(branch_a){
		sh = "Нет в ONYMA"
		index = f.NewSheet(sh)
		f.SetActiveSheet(index)
		rows = 1
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), b)
		rows += 1
		for _,v :=range(interface_new[b+"no_onyma"]){
			f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
			rows += 1
		}
	}
	//Пересечение IP адресов
	for i,b :=range(branch_a){
		sh = "Пересечение IP адресов"
		index = f.NewSheet(sh)
		f.SetActiveSheet(index)
		rows = 1
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), b)
		rows += 1
		for _,v :=range(interface_new[b+"ip_dublicate"]){
			f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
			rows += 1
		}
	}
	//Не распределённые
	for i,b :=range(branch_a){
		sh = "Не распределённые"
		index = f.NewSheet(sh)
		f.SetActiveSheet(index)
		rows = 1
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), b)
		rows += 1
		for _,v :=range(interface_new[b+"not_allocated"]){
			f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
			rows += 1
		}
	}

	f.SetActiveSheet(index)

	id := uuid.New()
	filename := fmt.Sprintf("%s.xlsx", id.String())
	path_file := fmt.Sprintf("/%s/%s", os.Getenv("PATH_DOWNLOAD"), filename)
	logrus.Infof("CreateExcellProblemRouterOnymaForm1:сохраняю файл %s", path_file)

	if err := f.SaveAs(path_file); err != nil {
		logrus.Errorf("CreateExcellProblemRouterOnymaForm1: ошибка запси на диск %s", path_file)
		logrus.Errorf("CreateExcellProblemRouterOnymaForm1:  %s", err)
		return "", err
	}

	logrus.Infof("CreateExcellProblemRouterOnymaForm1: файл сохранён на диск %s", path_file)
	return filename, error(nil)
}

/*Формирование ответа*/
func (s *RouterOnymaService) GetStatisticStatatusAll(data []tc.ProblemRouterOnymaQuery) (map[string]map[string]int, error) {

	//Формирую массив для файла
	data_new := make(map[string]map[string]int)
	parse_interface := make(map[string]string)
	for _, vv := range data {
		_, ok := parse_interface[vv.Interface_name]
		if ok { //Уже считали этот интерфейс, переходим к следующей итерации
			continue
		}
		//Список для проверки уже пропарсеных
		parse_interface[vv.Interface_name] = vv.Interface_name
		t, _ := time.Parse("2006-01-02T15:04:05Z", vv.Insert_datetime)
		date := t.Format("2006.01.02")
		_, ok = data_new[date]
		if !ok {
			data_new[date] = make(map[string]int)
			data_new[date]["work"] = 0          //В обработке
			data_new[date]["check"] = 0         //На проверку
			data_new[date]["re_check"] = 0      //Дообследование в ДЭ
			data_new[date]["test"] = 0          //Test
			data_new[date]["close"] = 0         //Закрыто
			data_new[date]["no_router"] = 0     //Нет на Роутере
			data_new[date]["error_speed"] = 0   //* Не соответствие скоростей;
			data_new[date]["filed_empty"] = 0   //* Пустые значения
			data_new[date]["no_onyma"] = 0      //* Нет в ONYMA
			data_new[date]["ip_dublicate"] = 0  //* Пересечение IP адресов
			data_new[date]["not_allocated"] = 0 //* Не распределённые
		}
		switch vv.Problem_status {
		case "В обработке":
			data_new[date]["work"] += 1
			data_new[date]["not_allocated"] += 1
		case "На проверку":
			data_new[date]["check"] += 1
		case "Дообследование в ДЭ":
			data_new[date]["re_check"] += 1
		case "Test":
			data_new[date]["test"] += 1
		case "Закрыто":
			data_new[date]["close"] += 1
		case "Дублирование IP":
			data_new[date]["ip_dublicate"] += 1
		case "Нет на Роутере":
			data_new[date]["no_router"] += 1

		}
		//Не соответствие скоростей; error_speed

		if vv.In_speed_onyma != vv.In_speed_router || vv.Out_speed_onyma != vv.Out_speed_router {
			var razi, razo, io, oo float64 //Разница скоростей
			if vv.In_speed_router != "" && vv.In_speed_onyma != "" {
				ir, err := strconv.ParseFloat(vv.In_speed_router, 64)
				if err != nil {
					logrus.Infof("getStatisticStatatusAll %s", err)
					return data_new, err
				}
				io, err := strconv.ParseFloat(vv.In_speed_onyma, 64)
				if err != nil {
					logrus.Infof("getStatisticStatatusAll %s", err)
					return data_new, err
				}
				if io > 0 {
					razi = ir / io
				}
			}

			if vv.Out_speed_router != "" && vv.Out_speed_onyma != "" {
				or, err := strconv.ParseFloat(vv.Out_speed_router, 64)
				if err != nil {
					logrus.Infof("getStatisticStatatusAll %s", err)
					return data_new, err
				}
				oo, err = strconv.ParseFloat(vv.Out_speed_onyma, 64)
				if err != nil {
					logrus.Infof("getStatisticStatatusAll %s", err)
					return data_new, err
				}
				if oo > 0 {
					razo = or / oo
				}
			}
			//Скоростя не совпадают
			if razo > 1.1 || razi > 1.1 || io == 0 || oo == 0 {
				data_new[date]["error_speed"] += 1
			}
		}
		//Пустые значения
		if vv.In_speed_router == "" || vv.In_speed_onyma == "" || vv.Out_speed_router == "" || vv.Out_speed_onyma == "" || vv.Company_name == "" {
			data_new[date]["filed_empty"] += 1
		}
		//Нет в ONYMA
		if vv.Dognum == "" || vv.Clsrv == "" {
			data_new[date]["no_onyma"] += 1
		}
		//Не распределённые
		if false {
			data_new[date]["not_allocated"] += 1
		}
	}

	//Сортировка по ключу

	/*keys := make([]string, 0, len(data_new))
	for k := range data_new {
		keys = append(keys, k)
	}
	sort.Strings(keys)*/
	/*for _, k := range keys {
		//logrus.Infoln(k, data_new[k])
	}*/

	return data_new, error(nil)
}

/* Проверка на соответствие записей полученых с Роутера и Онимы*/
//@Description Записываю в таблицу проблемных записей
func (s *RouterOnymaService) CheckRouterOnymaSpeed() (string, error) {
	ttime := time.Now()
	//Мапа для проверки дубликатов айпи адреса на разных Лицевых счетах
	//Отдельно ставлю статус "Дублирование IP"
	double_gognum_for_ip := make(map[string]map[string]string)

	//Все строки с router_onyma_speeds
	data, err := s.repo.GetAllRouterOnymaSpeedsByDate(ttime.Format("2006-01-02"))
	if err != nil {
		return "error", err
	}
	for _, v := range data {


		_, ok := double_gognum_for_ip[v.Ip_interface]
		if !ok {
			m := make(map[string]string)
			double_gognum_for_ip[v.Ip_interface] = m
		}

		if v.Dognum != "" {
			double_gognum_for_ip[v.Ip_interface][v.Dognum] = v.Dognum
		}
		//if v.Dognum == "15972" || v.Dognum == "29126" {
			//logrus.Infof("CheckRouterOnymaSpeed %v", v)
			//logrus.Infof("CheckRouterOnymaSpeed %v", double_gognum_for_ip)
		//}

		insert_b := false
		//if k>1000{continue}
		var razi, razo, ir, io, or, oo float64 //Разница скоростей
		if v.In_speed_router != "" && v.In_speed_onyma != "" {
			if ir, err = strconv.ParseFloat(v.In_speed_router, 64); err != nil {
				logrus.Infof("CheckRouterOnymaSpeed %s", err)
				return "error", err
			}
			if io, err = strconv.ParseFloat(v.In_speed_onyma, 64); err != nil {
				logrus.Infof("CheckRouterOnymaSpeed %s", err)
				return "error", err
			}
		}

		if v.Out_speed_router != "" && v.Out_speed_onyma != "" {
			if or, err = strconv.ParseFloat(v.Out_speed_router, 64); err != nil {
				logrus.Infof("CheckRouterOnymaSpeed %s", err)
				return "error", err
			}
			if oo, err = strconv.ParseFloat(v.Out_speed_onyma, 64); err != nil {
				logrus.Infof("CheckRouterOnymaSpeed %s", err)
				return "error", err
			}
		}

		if io > 0 {
			razi = ir / io
		}
		if oo > 0 {
			razo = or / oo
		}

		//Скоростя не совпадают
		if razo > 1.1 || razi > 1.1 || io == 0 || oo == 0 {
			insert_b = true
		}
		if v.In_speed_router == "" || v.In_speed_onyma == "" || v.Out_speed_router == "" || v.Out_speed_onyma == "" || v.Company_name == "" {
			insert_b = true
			/* Если совпадения новых записей с предыдущими записями в сводной таблице нет, они заносятся в эту же сводную таблицу
			и информация о новых записях появляется в WEB интерфейсе программы у ДЭ и ДКБ.
			При этом каждой записи автоматически присваивается статус ”В обработке”.
			Если в тексте описания интерфейса есть ### INET или ### ID и  запись проблемная, то сразу "На проверку", это клиенты
			*/
		}
		//Если нет проблем топследующая итерация
		if !insert_b {
			continue
		}
		//Проверяю существование в проблемной таблице (чтоб не было дубликатов)
		//(статус "Закрыто" не учитываются)
		checkTable, err := s.repo.GetProblemRouterOnymaSpeedsInfoByObjectNoClose(tc.ProblemRouterOnyma{Router_name: v.Router_name, Interface_description: v.Interface_description, Interface_name: v.Interface_name, Ip_interface: v.Ip_interface})

		if err != nil {
			return "error", err
		}
		//Есть такая запись, не добавляю.
		if len(checkTable) != 0 {
			//logrus.Infof("continue - уже вносилась запись ранее %v",check_table)
			continue
		}
		//Новая проблемная запись
		problem := tc.ProblemRouterOnyma{
			Router_onyma_speed_id: v.Id,
			Branch_service:        v.Branch_service,
			Router_name:           v.Router_name,
			Interface_name:        v.Interface_name,
			Interface_description: v.Interface_description,

			In_policy_router:    v.In_policy_router,
			In_speed_router:     v.In_speed_router,
			Out_policy_router:   v.Out_policy_router,
			Out_speed_router:    v.Out_speed_router,
			Ip_interface:        v.Ip_interface,
			Branch_contract:     v.Branch_contract,
			Dognum:              v.Dognum,
			Clsrv:               v.Clsrv,
			Company_name:        v.Company_name,
			In_speed_onyma:      v.In_speed_onyma,
			Out_speed_onyma:     v.Out_speed_onyma,
			Problem_status:      "В обработке",
			Client_status_onyma: v.Client_status_onyma,
		}
		if strings.HasPrefix(problem.Interface_description, "### INET") || strings.HasPrefix(problem.Interface_description, "###INET") {
			problem.Problem_status = "На проверку"
		}
		if strings.HasPrefix(problem.Interface_description, "### ID") || strings.HasPrefix(problem.Interface_description, "###ID") {
			problem.Problem_status = "На проверку"
		}
		if strings.HasPrefix(problem.Interface_description, "### TEST") || strings.HasPrefix(problem.Interface_description, "###TEST") {
			problem.Problem_status = "Test"
		}
		pro_id, err := s.repo.InsertProblemRouterOnymaSpeeds(problem)

		var msg string
		if razo > 1.1 {
			msg += "Out скорости не совпадают"
		}
		if razi > 1.1 {
			if len(msg) != 0 {
				msg += "\n"
			}
			msg += "In скорости не совпадают"
		}
		if io == 0 || oo == 0 || v.In_speed_router == "" || v.In_speed_onyma == "" || v.Out_speed_router == "" || v.Out_speed_onyma == "" {
			if len(msg) != 0 {
				msg += "\n"
			}
			msg += "Скорость не указана"
		}
		if v.Company_name == "" {
			if len(msg) != 0 {
				msg += "\n"
			}
			msg += "Имя компании не указано."
		}

		id := uuid.New()
		s.repo.InsertRouterOnymaHistory(tc.ProblemRouterOnymaHistory{Id: id.String(),
			Problem_router_onyma_speed_id: pro_id,
			User_id:                       "cb91c368-02d9-4cbd-961b-9bd6cb8c6af2", //TODO Хард код Юзер CRON
			Old_val:                       "",
			New_val:                       "",
			Msg:                           msg,
		})
		logrus.Printf("Скорости не совпадают In_speed_router != In_speed_onyma : %s != %s", v.In_speed_router, v.In_speed_onyma)

	}
	//logrus.Infof("CheckRouterOnymaSpeed %q", double_gognum_for_ip)
	//Мапа для проверки дубликатов айпи адреса на разных Лицевых счетах
	//Отдельно ставлю статус "Дублирование IP"
	for _, v := range data {
		//Если на одном айпи адресе один ЛС, то нормально всё.Пропускаем итерацию
		if len(double_gognum_for_ip[v.Ip_interface]) < 2 {
			continue
		}

		//Проверяю существование в проблемной таблице (чтоб не было дубликатов)
		//(статус "Закрыто" не учитываются)
		checkTable, err := s.repo.GetProblemRouterOnymaSpeedsInfoByObjectNoClose(tc.ProblemRouterOnyma{Problem_status: "Дублирование IP", Router_name: v.Router_name, Interface_description: v.Interface_description, Interface_name: v.Interface_name, Ip_interface: v.Ip_interface})

		if err != nil {
			return "error", err
		}
		//Есть такая запись, не добавляю.
		if len(checkTable) != 0 {
			logrus.Infof("continue - уже вносилась запись ранее %v", checkTable)
			continue
		}

		logrus.Infof("double_gognum_for_ip %s %q", v.Ip_interface, double_gognum_for_ip[v.Ip_interface])
		//Новая проблемная запись
		problem := tc.ProblemRouterOnyma{
			Router_onyma_speed_id: v.Id,
			Branch_service:        v.Branch_service,
			Router_name:           v.Router_name,
			Interface_name:        v.Interface_name,
			Interface_description: v.Interface_description,

			In_policy_router:    v.In_policy_router,
			In_speed_router:     v.In_speed_router,
			Out_policy_router:   v.Out_policy_router,
			Out_speed_router:    v.Out_speed_router,
			Ip_interface:        v.Ip_interface,
			Branch_contract:     v.Branch_contract,
			Dognum:              v.Dognum,
			Clsrv:               v.Clsrv,
			Company_name:        v.Company_name,
			In_speed_onyma:      v.In_speed_onyma,
			Out_speed_onyma:     v.Out_speed_onyma,
			Problem_status:      "Дублирование IP",
			Client_status_onyma: v.Client_status_onyma,
		}

		pro_id, err := s.repo.InsertProblemRouterOnymaSpeeds(problem)
		if err != nil {
			return "err", err
		}
		var msg string
		for _, m := range double_gognum_for_ip[v.Ip_interface] {
			if msg != "" {
				msg += ","
			}
			msg += "ЛС " + m
		}
		id := uuid.New()
		s.repo.InsertRouterOnymaHistory(tc.ProblemRouterOnymaHistory{Id: id.String(),
			Problem_router_onyma_speed_id: pro_id,
			User_id:                       "cb91c368-02d9-4cbd-961b-9bd6cb8c6af2", //TODO Хард код Юзер CRON
			Old_val:                       "",
			New_val:                       "",
			Msg:                           "Дублирование IP: " + msg,
		})

	}



	return "ok", err
}

// Проверка исправленных записей problem router на сегоднящний день, ставлю статус "Закрыто"
// Проверка, если уже нет записи на роутере, а в проблемных висит, ставлю статус "Нет на Роутере"

// /problem/check_close_problem_router_onyma_speed
func (s *RouterOnymaService) CheckCloseProblemRouterOnymaSpeed(userId string) (string, error) {
	ttime := time.Now()
	//Все строки с router_onyma_speeds на сегодня
	all_ros, err := s.repo.GetAllRouterOnymaSpeedsByDate(ttime.Format("2006-01-02"))
	if err != nil {
		return "error", err
	}

	//Айпишники ip_interface
	var ip_interface_not_in string

	for _, v := range all_ros {
		//Список проверенных айдишников с Router Onyma Speed
		if ip_interface_not_in!=""{
			ip_interface_not_in += ","
		}
		ip_interface_not_in += v.Ip_interface


		//флаг проблемная строка или нет
		problem_data := false
		/*if v.Ip_interface !="188.127.33.3" {
			continue
		}*/
		//logrus.Infof("%s;  ", v.Ip_interface)
		//logrus.Infof("data= %v", v)
		var razi, razo, ir, io, or, oo float64 //Разница скоростей
		if v.In_speed_router != "" && v.In_speed_onyma != "" {
			if ir, err = strconv.ParseFloat(v.In_speed_router, 64); err != nil {
				logrus.Infof("CheckCloseProblemRouterOnymaSpeed %s", err)
				return "error", err
			}
			if io, err = strconv.ParseFloat(v.In_speed_onyma, 64); err != nil {
				logrus.Infof("CheckCloseProblemRouterOnymaSpeed %s", err)
				return "error", err
			}
		}
		if v.Out_speed_router != "" && v.Out_speed_onyma != "" {
			if or, err = strconv.ParseFloat(v.Out_speed_router, 64); err != nil {
				logrus.Infof("CheckCloseProblemRouterOnymaSpeed %s", err)
				return "error", err
			}
			if oo, err = strconv.ParseFloat(v.Out_speed_onyma, 64); err != nil {
				logrus.Infof("CheckCloseProblemRouterOnymaSpeed %s", err)
				return "error", err
			}
		}

		if io > 0 {
			razi = ir / io
		}
		if oo > 0 {
			razo = or / oo
		}

		//проблемная запись, Скоростя не совпадают или нет Имени
		if razo > 1.1 || razi > 1.1 || io == 0 || oo == 0 || v.In_speed_router == "" || v.In_speed_onyma == "" || v.Out_speed_router == "" || v.Out_speed_onyma == "" || v.Company_name == "" {
			//logrus.Infof("Проблемная %s", v.Ip_interface)
			problem_data = true
		}

		//Ищю записи в проблемной таблице что бы поставть комментарий
		query := map[string]string{
			"router_name": v.Router_name,
			//"dognum":                v.Dognum, закоментил чтобы закрыть исправленные с пустыми dognum
			//"clsrv":                 v.Clsrv,  закоментил чтобы закрыть исправленные с пустыми clsrv
			"ip_interface":          v.Ip_interface,
			"problem_status_not_in": "Закрыто,Test,Дублирование IP,Нет на Роутере",
			"check_filter":          "true",
			"like":                  "false",
		}
		//logrus.Infof("Ищю записи в проблемной таблице что бы поставть статус закрыто или сделайть апдейт комментария, если данные изминились %v", query)
		//logrus.Infof("CheckCloseProblemRouterOnymaSpeed делаю запрос: GetAllProblemRouterOnymaSpeedsByQuery %v",query)
		problem, err := s.repo.GetAllProblemRouterOnymaSpeedsByQuery(query)
		if err != nil {
			logrus.Infof("CheckCloseProblemRouterOnymaSpeed err=%v", err)
			return "error", err
		}

		//logrus.Infof("CheckCloseProblemRouterOnymaSpeed router_onyma_speed_id=%s problem=%v", v.Id, problem)
		for _, p_v := range problem {
			//Проверяю что изменилось и заношу комментарий,
			msg := ""
			if v.In_speed_router != p_v.In_speed_router {
				msg += fmt.Sprintf("Изменена Входящая скорость на роутере, \nбыло(%s) \nстало: %s", p_v.In_speed_router, v.In_speed_router)
			}
			if v.Out_speed_router != p_v.Out_speed_router {
				if len(msg) != 0 {
					msg += "\n"
				}
				msg += fmt.Sprintf("Изменена Исходящая скорость на роутере, \nбыло(%s) \nстало: %s", p_v.Out_speed_router, v.Out_speed_router)
			}
			if v.In_speed_onyma != p_v.In_speed_onyma {
				if len(msg) != 0 {
					msg += "\n"
				}
				msg += fmt.Sprintf("Изменена Входящая скорость Онимы, \nбыло(%s) \nстало: %s", p_v.In_speed_onyma, v.In_speed_onyma)
			}
			if v.Out_speed_onyma != p_v.Out_speed_onyma {
				if len(msg) != 0 {
					msg += "\n"
				}
				msg += fmt.Sprintf("Изменена Исходящая скорость Онимы, \nбыло(%s) \nстало: %s", p_v.Out_speed_onyma, v.Out_speed_onyma)
			}
			if v.Interface_name != p_v.Interface_name {
				if len(msg) != 0 {
					msg += "\n"
				}
				msg += fmt.Sprintf("Изменено имя интерфейса, \nбыло(%s) \nстало: %s", p_v.Interface_name, v.Interface_name)
			}
			if v.Interface_description != p_v.Interface_description {
				if len(msg) != 0 {
					msg += "\n"
				}
				msg += fmt.Sprintf("Изменено имя описание интерфейса, \nбыло(%s) \nстало: %s", p_v.Interface_description, v.Interface_description)
			}
			if v.Ip_interface != p_v.Ip_interface {
				if len(msg) != 0 {
					msg += "\n"
				}
				msg += fmt.Sprintf("Изменен IP, \nбыло(%s) \nстало: %s", p_v.Ip_interface, v.Ip_interface)
			}
			if v.Branch_contract != p_v.Branch_contract {
				if len(msg) != 0 {
					msg += "\n"
				}
				msg += fmt.Sprintf("Изменен Филиал общества, \nбыло(%s) \nстало: %s", p_v.Branch_contract, v.Branch_contract)
			}
			if v.Company_name != p_v.Company_name {
				if len(msg) != 0 {
					msg += "\n"
				}
				msg += fmt.Sprintf("Изменен Филиал общества, \nбыло(%s) \nстало: %s", p_v.Company_name, v.Company_name)
			}
			if v.Dognum != p_v.Dognum {
				if len(msg) != 0 {
					msg += "\n"
				}
				msg += fmt.Sprintf("Изменен Лицевой счёт, \nбыло(%s) \nстало: %s", p_v.Dognum, v.Dognum)
			}
			if v.Clsrv != p_v.Clsrv {
				if len(msg) != 0 {
					msg += "\n"
				}
				msg += fmt.Sprintf("Изменен ID подключения, \nбыло(%s) \nстало: %s", p_v.Clsrv, v.Clsrv)
			}

			//Если так же осталась проблемной и есть изменения, заношу изменения, заношу комментарий, но статус не меняю
			if problem_data && msg != "" {
				up_data := tc.UpdateProblemRouterOnymaSpeedInput{
					Id:               p_v.Id,
					Dognum:           v.Dognum,
					Clsrv:            v.Clsrv,
					Company_name:     v.Company_name,
					In_speed_onyma:   v.In_speed_onyma,
					Out_speed_onyma:  v.Out_speed_onyma,
					In_speed_router:  v.In_speed_router,
					Out_speed_router: v.Out_speed_router,
					Interface_name:   v.Interface_name,
					Branch_contract:  v.Branch_contract,

					Problem_status_old: p_v.Problem_status,
					Problem_status:     p_v.Problem_status,
					Msg:                msg,
				}
				err = s.repo.UpdateProblemRouterOnymaSpeeds(p_v.Id, userId, up_data)
				if err != nil {
					logrus.Infof("(549)UpdateProblemRouterOnymaSpeeds err=%v, msg ", err, msg)
					return "error", err
				}

			}
			if !problem_data && msg != "" {
				//Значит не проблемная закрываю
				up_data := tc.UpdateProblemRouterOnymaSpeedInput{
					Id:                 p_v.Id,
					Problem_status_old: p_v.Problem_status,
					Problem_status:     "Закрыто",
					Msg:                msg,
				}
				err = s.repo.UpdateProblemRouterOnymaSpeeds(p_v.Id, userId, up_data)
				if err != nil {
					logrus.Infof("(549)UpdateProblemRouterOnymaSpeeds err=%v, msg ", err, msg)
					return "error", err
				}
			}

		}

	}

	/* Отдельно ставлю статус "Нет на Роутере"
	Проверяю данные с роутера за несколько дней, беру среднее значение для проверки
	, чтобы проставить Нет на роутере
	---

	 */
	//10 дней это 240 часов
	date_from:=ttime.Add(-240*time.Hour).Format("2006-01-02")
	date_to:=ttime.Format("2006-01-02")
	logrus.Infof("date_from %s date_to %s",date_from,date_to)
	data_group, err:=s.repo.GetStatisticRouterOnymaSpeedsGroupByInsertDateByQuery(map[string]string{"date_from":date_from,"date_to":date_to})

	if len(data_group)!=0{
		//Среднее, сумму делим на колво дней
		sr:=getSumCount(data_group)/10 // 10-колво дней.
		logrus.Infof("CheckCloseProblemRouterOnymaSpeed Среднее кол-во строк за 10 дней = %d", sr)
		//если среднее значение за 10 дней минус сегоднешнее колво данных не больше чем 1.1 то проверяю данные на "Нет на Роутере"
		//all_ros - Все строки с router_onyma_speeds на сегодня
		if float64(sr/len(all_ros)) <1.1{
			logrus.Infof("CheckCloseProblemRouterOnymaSpeed Среднее %d/%d кол-во строк за сегодня = %f это меньше 1.1", sr,len(all_ros), float64(sr/len(all_ros)))
			logrus.Info("CheckCloseProblemRouterOnymaSpeed Делаю проверку на статус 'Нет на Роутере'")

			// Вытаскиваю все записи проблемных, которых нет на роутере 188.0.146.223 Bundle-Ether30.21010465
			//		Ищю записи в проблемной таблице что бы поставть комментарий
			query := map[string]string{
				"ip_interface_not_in": ip_interface_not_in,//кроме сегодняшних данных
				"check_filter":          "true",
				"like":                  "false",
				"problem_status_not_in": "Закрыто,Test,Дублирование IP,Нет на Роутере",
			}
			problem, err := s.repo.GetAllProblemRouterOnymaSpeedsByQuery(query)
			if err != nil {
				logrus.Infof("GetAllProblemRouterOnymaSpeedsByQuery err=%v", err)
				return "error", err
			}
			/*for _,v := range(problem){
				//Отдельно ставлю статус "Нет на Роутере"
				up_data := tc.UpdateProblemRouterOnymaSpeedInput{
					Id:                 v.Id,
					Problem_status_old: v.Problem_status,
					Problem_status:     "Нет на Роутере",
					Msg:                "Нет на роутере",
				}
				err = s.repo.UpdateProblemRouterOnymaSpeeds(v.Id, userId, up_data)
				if err != nil {
					logrus.Infof("(1130)UpdateProblemRouterOnymaSpeedInput err=%v, msg ", err, "Нет на Роутере")
					return "error", err
				}
			}*/
			logrus.Infof("\nВсе данные с роутера прочитаны, в мапе проблемных остались: %d",len(problem))
			//logrus.Infof("%v",problem)
			for _,v:=range problem{
				logrus.Infof(",%v",v.Ip_interface)
			}
		}
	}

	logrus.Info("CheckCloseProblemRouterOnymaSpeed - ok")
	return "ok", err
}
func getSumCount(data []tc.RouterOnymaGroupByInsertDate) int{
	var summ int
	for _,v:= range data {
		summ+=v.Count
	}
	return summ
}
// Поиск проблемных записей, в ежедневном срезе данных в таблице router_onyma_speeds, с разными статусами на Роутере и в Ониме,
// создаю запись в control_time_pause и ставлю статус "1"
func (s *RouterOnymaService) CheckControlTimePause(userId string) (string, error) {
	ttime := time.Now()
	insert_ctp_m := []tc.ControlTimePauseInsert{}
	insert_ctph_m := []tc.ControlTimePauseHistory{}
	//Все строки с router_onyma_speeds на текущую дату
	data, err := s.repo.GetAllRouterOnymaSpeedsByDate(ttime.Format("2006-01-02"))
	if err != nil {
		return "error", err
	}
	for _, v := range data {
		//Если статусы выкл совпадают, то пропускаю итерацию
		if v.Iface_shutdown_router == 1 && v.Client_status_onyma == 3 {
			continue
		}
		//Если статусы вкл совпадают, то пропускаю итерацию
		if v.Iface_shutdown_router == 0 && v.Client_status_onyma == 0 {
			continue
		}

		//Все остальные данные записываю в таблицу и ставлю комментарий о новой записи.

		//Чтобы не было дублей проверяю на существование такой строки с открытым для проверки статуса.
		query := map[string]string{
			"router_name":    v.Router_name,
			"dognum":         v.Dognum,
			"clsrv":          v.Clsrv,
			"ip_interface":   v.Ip_interface,
			"control_status": "1", //1-проблема. 0-нет проблем
		}
		//logrus.Infof("CheckCloseProblemRouterOnymaSpeed делаю запрос: GetAllProblemRouterOnymaSpeedsByQuery %v",query)

		res, err := s.repo.GetControlTimePauseByQuery(query)
		if err != nil {
			logrus.Infof("GetControlTimePauseByQuery err=%v", err)
			return "error", err
		}
		if len(res) != 0 { //Есть такия строка и с проблемным статусом. пропускаю итерацию
			logrus.Infof("GetControlTimePauseByQuery, пропуск, Запись уже есть в проблемных router_onyma_speed_id=%s", res[0].Router_onyma_speed_id)
			continue
		}
		id := uuid.New()
		insert_ctp_m = append(insert_ctp_m, tc.ControlTimePauseInsert{Id: id.String(), Router_onyma_speed_id: v.Id, Control_status: 1, Created_at: "now()"})
		id_h := uuid.New()
		insert_ctph_m = append(insert_ctph_m, tc.ControlTimePauseHistory{Id: id_h.String(), ControlTimePauseId: id.String(), UserId: userId, Msg: "Статусы не совпадают"})
	}

	logrus.Infof("%q", insert_ctp_m)
	//logrus.Infof("GetControlTimePauseByQuery Новых строк %d",len(insert_ctp_m))
	res_ctp, err := s.repo.InsertControlTimePause(insert_ctp_m)
	if err != nil {
		logrus.Infof("InsertControlTimePause err=%v", err)
		return "error", err
	}
	logrus.Infof("InsertControlTimePause %s", res_ctp)
	res_h, err := s.repo.InsertControlTimePauseHistory(insert_ctph_m)
	if err != nil {
		logrus.Infof("InsertControlTimePauseHistory err=%v", err)
		return "error", err
	}
	logrus.Infof("InsertControlTimePauseHistory %s", res_h)

	logrus.Info("GetControlTimePauseByQuery - ok")
	return "ok", err
}
