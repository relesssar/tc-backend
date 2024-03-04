package service

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"

	//"time"
	"os"
	//"strconv"
	"fmt"
	"strings"
	tc "tc_kaztranscom_backend_go"
	"tc_kaztranscom_backend_go/pkg/repository"
)

type AccessGroupService struct {
	repo repository.AccessGroup
}

func NewAccessGroupService(repo repository.AccessGroup) *AccessGroupService {
	return &AccessGroupService{repo: repo}
}

func (s *AccessGroupService) InsertAccessGroupHistory(fro tc.AccessGroupHistory) (string, error) {
	return s.repo.InsertAccessGroupHistory(fro)
}
func (s *AccessGroupService) GetAccessGroupHistory(input tc.GetAccessGroupHistory) ([]tc.AccessGroupHistory, error) {
	return s.repo.GetAccessGroupHistory(input)
}
func (s *AccessGroupService) UpdateAccessGroup(id, userId string, data tc.UpdateAccessGroupInput) error {

	return s.repo.UpdateAccessGroup(id, userId, data)
}

func (s *AccessGroupService) InsertFilterAccessGroup(fro tc.FilterAccessGroup) (string, error) {
	return s.repo.InsertFilterAccessGroup(fro)
}
func (s *AccessGroupService) GetFilterAccessGroup(filter tc.FilterAccessGroupSearch) ([]tc.FilterAccessGroupSearch, error) {

	return s.repo.GetFilterAccessGroup(filter)
}
func (s *AccessGroupService) DeleteFilter(id string) error {

	return s.repo.DeleteFilter(id)
}

func (s *AccessGroupService) InsertAccessGroup(ag tc.AccessGroup) (string, error) {
	return s.repo.InsertAccessGroup(ag)
}
func (s *AccessGroupService) GetAllAccessGroupByQuery(data map[string]string) ([]tc.AccessGroup, error) {
	return s.repo.GetAllAccessGroupByQuery(data)
}

/*создание ексель файла для выгрузки с интерфейса Форма 1*/
func (s *AccessGroupService) CreateExcellAccessGroup(data []tc.AccessGroup) (string, error) {

	f := excelize.NewFile()
	sh := "access group"
	index := f.NewSheet(sh)
	// Set value of a cell.
	rows := 1
	line := "Дата проверки;Статус;Router;Сеть;Айпи;Интерфейс;Описание;Состояние;In Policy;Out Policy;Access Group;Клиент в Ониме Dognum;Клиент в Ониме Clsrv"
	line_m := strings.Split(line, ";")
	char_m := strings.Split("A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X", ", ")
	for i, v := range line_m {
		f.SetCellValue(sh, fmt.Sprintf("%s%d", char_m[i], rows), v)
		//rows+=1
	}
	rows += 1

	for _, v := range data {
		t, _ := time.Parse("2006-01-02T15:04:05Z", v.Created_at)

		f.SetCellValue(sh, fmt.Sprintf("A%d", rows), t.Format("02.01.2006"))
		f.SetCellValue(sh, fmt.Sprintf("B%d", rows), v.Access_status)
		f.SetCellValue(sh, fmt.Sprintf("C%d", rows), v.Access_group)
		f.SetCellValue(sh, fmt.Sprintf("D%d", rows), v.Router_Name)
		f.SetCellValue(sh, fmt.Sprintf("E%d", rows), v.Iface_host)
		f.SetCellValue(sh, fmt.Sprintf("F%d", rows), v.Ip)
		f.SetCellValue(sh, fmt.Sprintf("G%d", rows), v.Iface_name)
		f.SetCellValue(sh, fmt.Sprintf("H%d", rows), v.Iface_desc)
		f.SetCellValue(sh, fmt.Sprintf("I%d", rows), v.Client_status)
		f.SetCellValue(sh, fmt.Sprintf("J%d", rows), v.In_policy)
		f.SetCellValue(sh, fmt.Sprintf("K%d", rows), v.Out_policy)
		f.SetCellValue(sh, fmt.Sprintf("L%d", rows), v.Access_group)
		f.SetCellValue(sh, fmt.Sprintf("M%d", rows), v.Dognum)
		f.SetCellValue(sh, fmt.Sprintf("N%d", rows), v.Clsrv)
		//f.SetCellValue(sh, fmt.Sprintf("K%d", rows), data_new[i]["test"])
		//f.SetCellValue(sh, fmt.Sprintf("L%d", rows), data_new[i]["close"])
		rows += 1
	}
	f.SetActiveSheet(index)

	// Set active sheet of the workbook.
	//f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	id := uuid.New()
	filename := fmt.Sprintf("%s.xlsx", id.String())
	path_file := fmt.Sprintf("/%s/%s", os.Getenv("PATH_DOWNLOAD"), filename)
	logrus.Infof("CreateExcellAccessGroup:сохраняю файл %s", path_file)

	if err := f.SaveAs(path_file); err != nil {
		logrus.Errorf("CreateExcellAccessGroup: ошибка записи на диск %s", path_file)
		logrus.Errorf("CreateExcellAccessGroup:  %s", err)
		return "", err
	}

	logrus.Infof("CreateExcellAccessGroup: файл сохранён на диск %s", path_file)
	return filename, error(nil)
}
