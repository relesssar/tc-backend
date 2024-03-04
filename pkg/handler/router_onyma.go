package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	tc "tc_kaztranscom_backend_go"
)

type getAllRouterOnymaSpeedsByQueryResponse struct {
	Data []tc.RouterOnyma `json:"data"`
}
type getAllRouterOnymaSpeedsByDateResponse struct {
	Data []tc.RouterOnyma `json:"data"`
}
type getAllProblemRouterOnymaSpeedsByQueryResponse struct {
	Data []tc.ProblemRouterOnymaQuery `json:"data"`
}
type getControlTimePauseByQueryResponse struct {
	Data []tc.ControlTimePause `json:"data"`
}
type getDateListProblemRouterOnymaSpeedsByResponse struct {
	Data []string `json:"data"`
}
type getAllRouterOnymaHistorySearchByQueryResponse struct {
	Data []tc.ProblemRouterOnymaHistorySearch `json:"data"`
}
type getAllFilterRouterOnymaSearchByQueryResponse struct {
	Data []tc.FilterRouterOnymaSearch `json:"data"`
}

// @Summary Скачивание файла
// @Tags Problem Router Onyma
// @Description Скачивание файла Excell с проблемными строками
// @ID download-excell-file-problev-router-onyma
// @Accept  json
// @Produce  json
// @Success 200 {string} string status
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem/download [get]
func (h *Handler) getExcellProblemRouterOnymaSpeeds(c *gin.Context) {

	//получаю данные с учётом фильтра
	data := make(map[string]string)
	//d_map := []string{"date","updated_at", "branch_service", "problem_status", "router_onyma_speed_id", "interface_name", "ip_interface", "check_filter"}
	d_map := []string{"check_filter", "clsrv", "router_name", "problem_status",
		"dognum", "company_name", "ip_interface", "interface_name", "router_onyma_speed_id",
		"branch_service", "date", "updated_at", "client_status_onyma","onyma_speed_null",
		"onyma_speed_error","onyma_dognum_null",}
	for _, v := range d_map {
		data[v] = c.Query(v)
	}

	//logrus.Infof("getCategoryById(%s)",id)
	lists, err := h.services.RouterOnyma.GetAllProblemRouterOnymaSpeedsByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	//создаю файл ексель и возвращаю имя файла
	filename, err := h.services.RouterOnyma.CreateExcellProblemRouterOnymaSpeed(lists)
	path_file := fmt.Sprintf("/%s/%s", os.Getenv("PATH_DOWNLOAD"), filename)
	fi, err := os.Stat(path_file)
	if err != nil {
		logrus.Infof("%s %s", path_file, err)
	}
	// get the size
	size := fi.Size()
	logrus.Infof("размер счачеваемого файла состовляет %d", size)

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)) //fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Length", fmt.Sprintf("%d", size))
	//c.Writer.Header().Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNTI3OGE5MWEtMDlkMS00ZGJlLWJmN2QtMzRmZjFkYzQzNzBmIiwiZXhwIjoxNjIzMzkwOTc4LCJpYXQiOjE2MjMyMTgxNzgsImlzcyI6IlRvdGFsIENvbnRyb2wgKzc3NzczNzg1NjMxLDc3NzM3ODU2MzFAbWFpbC5ydSJ9.LW8lTNMObc4p-pYRUtfG6pSbdsyJOVJG3jEVyX439Es")
	c.File(path_file)

	/*status, err := h.services.RouterOnyma.GetExcellProblemRouterOnymaSpeeds()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})*/

}

// @Summary Скачивание файла, Форма 1
// @Tags Problem Router Onyma
// @Description Скачивание файла статистики Форма 1(Excell)
// @ID download-excell-file-problev-router-onyma-form1
// @Accept  json
// @Produce  json
// @Param date_end query string false "все записи до текущей даты, дата в формате yyyy-mm-dd"
// @Success 200 {string} string status
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem/statistic/form1 [get]
func (h *Handler) getExcellProblemRouterOnymaForm1(c *gin.Context) {

	//получаю данные с учётом фильтра
	data := make(map[string]string)
	data["date_end"] = c.Query("date_end")
	data["check_filter"] = "true"
	data["problem_status_not_in"] = "Закрыто,Test"

	//logrus.Infof("getCategoryById(%s)",id)
	lists, err := h.services.RouterOnyma.GetAllProblemRouterOnymaSpeedsByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	//создаю файл ексель и возвращаю имя файла
	filename, err := h.services.RouterOnyma.CreateExcellProblemRouterOnymaForm1(lists)
	path_file := fmt.Sprintf("/%s/%s", os.Getenv("PATH_DOWNLOAD"), filename)
	fi, err := os.Stat(path_file)
	if err != nil {
		logrus.Infof("%s %s", path_file, err)
	}
	// get the size
	size := fi.Size()
	logrus.Infof("размер счачеваемого файла состовляет %d", size)

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)) //fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Length", fmt.Sprintf("%d", size))
	//c.Writer.Header().Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNTI3OGE5MWEtMDlkMS00ZGJlLWJmN2QtMzRmZjFkYzQzNzBmIiwiZXhwIjoxNjIzMzkwOTc4LCJpYXQiOjE2MjMyMTgxNzgsImlzcyI6IlRvdGFsIENvbnRyb2wgKzc3NzczNzg1NjMxLDc3NzM3ODU2MzFAbWFpbC5ydSJ9.LW8lTNMObc4p-pYRUtfG6pSbdsyJOVJG3jEVyX439Es")
	c.File(path_file)

	/*status, err := h.services.RouterOnyma.GetExcellProblemRouterOnymaSpeeds()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})*/

}

// @Summary Статистика
// @Tags Problem Router Onyma
// @Description Данные по Статусам
// @ID statistic-status-all
// @Accept  json
// @Produce  json
// @Param date_from query string false "c текущей даты, дата в формате yyyy-mm-dd"
// @Param date_to query string false "по текущую дату, дата в формате yyyy-mm-dd"
// @Success 200 {string} string status
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem/statistic/status_all [get]
func (h *Handler) getStatisticStatusByQuery(c *gin.Context) {

	//получаю данные с учётом фильтра
	data := make(map[string]string)
	data["date_from"] = c.Query("date_from")
	data["date_to"] = c.Query("date_to")
	data["check_filter"] = "true"

	//logrus.Infof("getCategoryById(%s)",id)
	lists, err := h.services.RouterOnyma.GetAllProblemRouterOnymaSpeedsByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	//Формирование данных для вывода
	res, err := h.services.RouterOnyma.GetStatisticStatatusAll(lists)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)

}// @Summary Статистика

// @Tags Router Onyma
// @Security ApiKeyAuth
// @Description Данные сгрупперованны по роутерам
// @ID statistic-router-onyma-group-by-router
// @Accept  json
// @Produce  json
// @Param date_from query string false "c текущей даты, дата в формате yyyy-mm-dd"
// @Param date_to query string false "по текущую дату, дата в формате yyyy-mm-dd"
// @Success 200 {string} string status
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/statistic/group_by_router [get]
func (h *Handler) getStatisticRouterOnymaSpeedsGroupByRouterByQuery(c *gin.Context) {

	//получаю данные с учётом фильтра
	data := make(map[string]string)
	data["date_from"] = c.Query("date_from")
	data["date_to"] = c.Query("date_to")


	//logrus.Infof("getCategoryById(%s)",id)
	lists, err := h.services.RouterOnyma.GetStatisticRouterOnymaSpeedsGroupByRouterByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, lists)

}

// @Summary Скачивание файла
// @Tags Router Onyma
// @Description Контроль временного отключения
// @Description Скачивание файла Excell с проблемными статусами Роутера и Онимы
// @ID download-excell-file-ccp-router-onyma
// @Accept  json
// @Produce  json
// @Success 200 {string} string status
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/control_time_pause/download [get]
func (h *Handler) getExcellControlTimePause(c *gin.Context) {

	//получаю данные с учётом фильтра
	data := make(map[string]string)
	d_map := []string{"control_status", "iface_shutdown_router", "client_status_onyma"}
	for _, v := range d_map {
		data[v] = c.Query(v)
	}

	//logrus.Infof("getCategoryById(%s)",id)
	lists, err := h.services.RouterOnyma.GetControlTimePauseByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	//создаю файл ексель и возвращаю имя файла
	filename, err := h.services.RouterOnyma.CreateExcellControlTimePause(lists)
	path_file := fmt.Sprintf("/%s/%s", os.Getenv("PATH_DOWNLOAD"), filename)
	fi, err := os.Stat(path_file)
	if err != nil {
		logrus.Infof("%s %s", path_file, err)
	}
	// get the size
	size := fi.Size()
	logrus.Infof("размер счачеваемого файла состовляет %d", size)

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)) //fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Length", fmt.Sprintf("%d", size))
	//c.Writer.Header().Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNTI3OGE5MWEtMDlkMS00ZGJlLWJmN2QtMzRmZjFkYzQzNzBmIiwiZXhwIjoxNjIzMzkwOTc4LCJpYXQiOjE2MjMyMTgxNzgsImlzcyI6IlRvdGFsIENvbnRyb2wgKzc3NzczNzg1NjMxLDc3NzM3ODU2MzFAbWFpbC5ydSJ9.LW8lTNMObc4p-pYRUtfG6pSbdsyJOVJG3jEVyX439Es")
	c.File(path_file)

	/*status, err := h.services.RouterOnyma.GetExcellProblemRouterOnymaSpeeds()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})*/

}

// @Summary Поиск проблемных строк
// @Tags Problem Router Onyma
// @Description Проверка на соответствие данных полученых с Роутеров и Онимы
// @Description по умолчанию статус "В обработке".
// @Description Если в тексте описания интерфейса есть ### INET или ### ID и  запись проблемная, то сразу "На проверку", это клиенты
// @ID check-router-onyma-speed
// @Accept  json
// @Produce  json
// @Success 200 {string} string status
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem/check_router_onyma_speed [get]
func (h *Handler) checkRouterOnymaSpeed(c *gin.Context) {

	status, err := h.services.RouterOnyma.CheckRouterOnymaSpeed()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}

// @Summary Проверка на соответствие
// @Security ApiKeyAuth
// @Tags Problem Router Onyma
// @Description Проверка на соответствие данных полученых с Роутеров и Онимы
// @Description Если проблемная строка перестала быть проблемной, то ставлю "Закрыто", это исправленная строка
// @ID check-control-time-pause
// @Accept  json
// @Produce  json
// @Success 200 {string} string status
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem/check_control_time_pause [get]
func (h *Handler) checkControlTimePause(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.Infof("(177) user id не найден, %s", err)
		return
	}
	status, err := h.services.RouterOnyma.CheckControlTimePause(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}

// @Summary Проверка на соответствие
// @Security ApiKeyAuth
// @Tags Problem Router Onyma
// @Description Проверка на соответствие данных полученых с Роутеров и Онимы
// @Description Если проблемная строка перестала быть проблемной, то ставлю "Закрыто", это исправленная строка
// @ID check-close-problem-router-onyma-speed
// @Accept  json
// @Produce  json
// @Success 200 {string} string status
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem/check_close_problem_router_onyma_speed [get]
func (h *Handler) checkCloseProblemRouterOnymaSpeed(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.Infof("(204) user id не найден, %s", err)
		return
	}
	status, err := h.services.RouterOnyma.CheckCloseProblemRouterOnymaSpeed(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}

// @Summary Список дат Problem Router Onyma Speeds
// @Security ApiKeyAuth
// @Tags Problem Router Onyma
// @Description Список дат Problem Router Onyma Speeds, даты на которые есть проблемные записи
// @ID get-date-list-problem-router-onyma-speeds
// @Accept  json
// @Produce  json
// @Success 200 {object} getDateListProblemRouterOnymaSpeedsByResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem/date_list [get]
func (h *Handler) getDateListProblemRouterOnymaSpeed(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.RouterOnyma.GetDateListProblemRouterOnymaSpeed()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getDateListProblemRouterOnymaSpeedsByResponse{
		Data: lists,
	})
}

// @Summary Список данных Control Time Pause
// @Security ApiKeyAuth
// @Tags Problem Router Onyma
// @Description Контроль временных отключений
// @Description Список данных Control Time Pause по переданным параметрам
// @ID get-ctpq
// @Accept  json
// @Produce  json
// @Param date query string false "в формате yyyy-mm-dd"
// @Param control_status query string false "1 или 0, 1-проблема,0-закрыто"
// @Param router_onyma_speed_id query string false "Идентификатор связи с router onyma speed"
// @Param interface_name query string false "Имя интерфейса"
// @Param ip_interface query string false "Айпи интерфейса"
// @Param dognum query string false "Номер Лицевого счёта Онимы"
// @Param router_name query string false "Имя роутера"
// @Param clsrv query string false "Айди подключения в Ониме"
// @Success 200 {object} getControlTimePauseByQueryResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/control_time_pause [get]
func (h *Handler) getControlTimePauseByQuery(c *gin.Context) {
	/*_, err := getUserId(c)
	if err != nil {
		return
	}*/
	data := make(map[string]string)
	d_map := []string{"control_status", "clsrv", "router_name", "dognum", "ip_interface", "interface_name", "router_onyma_speed_id", "branch_service", "date", "iface_shutdown_router", "client_status_onyma"}
	for _, v := range d_map {
		data[v] = c.Query(v)
	}

	lists, err := h.services.RouterOnyma.GetControlTimePauseByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getControlTimePauseByQueryResponse{
		Data: lists,
	})
}

// @Summary Список данных Problem Router Onyma Speeds
// @Security ApiKeyAuth
// @Tags Problem Router Onyma
// @Description Список данных Problem Router Onyma Speeds по переданным параметрам, в формате yyyy-mm-dd
// @ID get-problem-router-onyma-speeds-by-query
// @Accept  json
// @Produce  json
// @Param date query string false "в формате yyyy-mm-dd"
// @Param problem_status query string false "Статус записи"
// @Param router_onyma_speed_id query string false "Идентификатор связи с router onyma speed"
// @Param interface_name query string false "Имя интерфейса"
// @Param ip_interface query string false "Айпи интерфейса"
// @Param dognum query string false "Номер Лицевого счёта Онимы"
// @Param router_name query string false "Имя роутера"
// @Param clsrv query string false "Айди подключения в Ониме"
// @Success 200 {object} getAllProblemRouterOnymaSpeedsByQueryResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem [get]
func (h *Handler) getAllProblemRouterOnymaSpeedsByQuery(c *gin.Context) {
	/*_, err := getUserId(c)
	if err != nil {
		return
	}*/
	data := make(map[string]string)
	d_map := []string{"check_filter", "clsrv", "router_name", "problem_status",
		"dognum", "company_name", "ip_interface", "interface_name", "router_onyma_speed_id",
		"branch_service", "date", "updated_at", "client_status_onyma","not_in_ros_ids","onyma_speed_null",
		"onyma_speed_error","onyma_dognum_null",
	}
	for _, v := range d_map {
		data[v] = c.Query(v)
	}

	lists, err := h.services.RouterOnyma.GetAllProblemRouterOnymaSpeedsByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getAllProblemRouterOnymaSpeedsByQueryResponse{
		Data: lists,
	})
}

// @Summary Количество строк данных Problem Router Onyma Speeds
// @Security ApiKeyAuth
// @Tags Count Problem Router Onyma
// @Description Количество строк данных Problem Router Onyma Speeds по переданным параметрам, в формате yyyy-mm-dd
// @ID get-count-problem-router-onyma-speeds-by-query
// @Accept  json
// @Produce  json
// @Param date query string false "в формате yyyy-mm-dd"
// @Param problem_status query string false "Статус записи"
// @Param router_onyma_speed_id query string false "Идентификатор связи с router onyma speed"
// @Param interface_name query string false "Имя интерфейса"
// @Param ip_interface query string false "Айпи интерфейса"
// @Param dognum query string false "Номер Лицевого счёта Онимы"
// @Param router_name query string false "Имя роутера"
// @Param clsrv query string false "Айди подключения в Ониме"
// @Success 200 {string} int count
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem/count [get]
func (h *Handler) getCountAllProblemRouterOnymaSpeedsByQuery(c *gin.Context) {
	/*_, err := getUserId(c)
	if err != nil {
		return
	}*/
	data := make(map[string]string)
	d_map := []string{"check_filter", "clsrv", "router_name", "problem_status",
		"dognum", "company_name", "ip_interface", "interface_name", "router_onyma_speed_id",
		"branch_service", "date", "client_status_onyma", "problem_status_not_in"}
	for _, v := range d_map {
		data[v] = c.Query(v)
	}

	count, err := h.services.RouterOnyma.GetCountAllProblemRouterOnymaSpeedsByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"count": count,
	})
}

// @Summary Список данных Router Onyma Speeds by Date
// @Security ApiKeyAuth
// @Tags Router Onyma
// @Description Список данных Router Onyma Speeds на определённую дату, в формате yyyy-mm-dd
// @ID get-router-onyma-speeds-by-date
// @Accept  json
// @Produce  json
// @Param date path string true "на какую дату ищем?"
// @Success 200 {object} getAllRouterOnymaSpeedsByDateResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/speeds/{date} [get]
func (h *Handler) getAllRouterOnymaSpeedsByDate(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	date := c.Param("date")
	//logrus.Infof("getCategoryById(%s)",id)
	lists, err := h.services.RouterOnyma.GetAllRouterOnymaSpeedsByDate(date)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getAllRouterOnymaSpeedsByDateResponse{
		Data: lists,
	})
}

// @Summary Список данных Router Onyma Speeds
// @Security ApiKeyAuth
// @Tags Router Onyma
// @Description Список данных  Router Onyma Speeds по переданным параметрам
// @ID get-router-onyma-speeds-by-query
// @Accept  json
// @Produce  json
// @Param id query string false "Идентификатор uuid router onyma speed"
// @Param date query string false "в формате yyyy-mm-dd"
// @Param branch_service query string false "Филиал общества в которо предоставляется услуга"
// @Param interface_name query string false "Имя интерфейса"
// @Param ip_interface query string false "Айпи интерфейса"
// @Param router_name query string false "Имя роутера"
// @Param dognum query string false "Номер Лицевого счёта Онимы"
// @Param clsrv query string false "Айди подключения в Ониме"
// @Param iface_shutdown_router query int false "на роутере 1-shutdown, 0-включено"
// @Param client_status_onyma query int false "Статус клиента в ониме"
// @Success 200 {object} getAllRouterOnymaSpeedsByQueryResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/speeds [get]
func (h *Handler) getAllRouterOnymaSpeedsByQuery(c *gin.Context) {
	/*_, err := getUserId(c)
	if err != nil {
		return
	}*/
	data := make(map[string]string)

	data["id"] = c.Query("id")
	data["date"] = c.Query("date")
	data["branch_service"] = c.Query("branch_service")
	data["interface_name"] = c.Query("interface_name")
	data["ip_interface"] = c.Query("ip_interface")
	data["check_filter"] = c.Query("check_filter")
	data["dognum"] = c.Query("dognum")
	data["router_name"] = c.Query("router_name")
	data["clsrv"] = c.Query("clsrv")
	data["iface_shutdown_router"] = c.Query("iface_shutdown_router")
	data["client_status_onyma"] = c.Query("client_status_onyma")
	lists, err := h.services.RouterOnyma.GetAllRouterOnymaSpeedsByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getAllRouterOnymaSpeedsByQueryResponse{
		Data: lists,
	})
}

// @Summary Добавление данных в filter(Исключения)
// @Security ApiKeyAuth
// @Tags Filter Router Onyma
// @Description Добавление в базу новых данных, история изменений, разные данные
// @ID insert-filter-router-onyma
// @Accept  json
// @Produce  json
// @Param input body tc.FilterRouterOnyma true "Данные для внесения в базу"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/filter_router_onyma [post]
func (h *Handler) insertFilterRouterOnyma(c *gin.Context) {
	var input tc.FilterRouterOnyma
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user_id, err := getUserId(c)
	if err != nil {
		return
	}
	input.User_id = user_id
	id, err := h.services.RouterOnyma.InsertFilterRouterOnyma(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Получение данных
// @Security ApiKeyAuth
// @Tags Filter Router Onyma
// @Description Получение  данных, Фильтры, разные данные
// @ID get-filter-router-onyma
// @Accept  json
// @Produce  json
// @Param input query tc.FilterRouterOnymaSearch true "Любые данные для поиск в базе"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/filter_router_onyma [get]
func (h *Handler) getFilterRouterOnyma(c *gin.Context) {
	var input tc.FilterRouterOnymaSearch
	input.Id = c.Query("id")
	input.Filter_type = c.Query("filter_type")
	input.Filter_val = c.Query("filter_val")
	input.Filter_desc = c.Query("filter_desc")
	input.User_id = c.Query("user_id")
	input.Created_at = c.Query("created_at")
	logrus.Infof("%v", input)
	lists, err := h.services.RouterOnyma.GetFilterRouterOnyma(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllFilterRouterOnymaSearchByQueryResponse{
		Data: lists,
	})
}

// @Summary Удаление
// @Security ApiKeyAuth
// @Tags Filter Router Onyma
// @Description Удаление Фильтра
// @ID delete-filter
// @Accept  json
// @Produce  json
// @Param id path string true "uuid фильтра"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/filter_router_onyma/{id} [delete]
func (h *Handler) deleteFilter(c *gin.Context) {

	id := c.Param("id")
	err := h.services.RouterOnyma.DeleteFilter(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Получение данных
// @Security ApiKeyAuth
// @Tags History Router Onyma
// @Description Получение  данных, история изменений, разные данные
// @ID get-router-onyma-history
// @Accept  json
// @Produce  json
// @Param id query string false "id"
// @Param problem_router_onyma_speed_id query string false "problem_router_onyma_speed_id"
// @Param old_val query string false "старое значение"
// @Param new_val query string false "новое значение"
// @Param msg query string false "сообщение"
// @Param user_id query string false "Айди юзера"
// @Param branch_service query string false "Филиал общества"
// @Param problem_status query string false "Статус проблемной записи"
// @Param client_status_onyma query string false "Статус Клиента в Ониме"
// @Param company_name query string false "Название компании в Ониме"
// @Param ip_interface query string false "IP интерфейса"
// @Param interface_name query string false "Название интерфейса"
// @Param created_at query string false "дата создания сообщения"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/router_onyma_history [get]
func (h *Handler) getRouterOnymaHistory(c *gin.Context) {
	/*var input tc.ProblemRouterOnymaHistorySearch
	input.Id = c.Query("id")
	input.Problem_router_onyma_speed_id = c.Query("problem_router_onyma_speed_id")
	input.Old_val = c.Query("old_val")
	input.New_val = c.Query("new_val")
	input.Msg = c.Query("msg")
	input.User_id = c.Query("user_id")
	input.Created_at = c.Query("created_at")
	logrus.Infof("%v", input)*/
	data := make(map[string]string)
	d_map := []string{"id", "problem_router_onyma_speed_id", "old_val", "new_val", "interface_name",
		"msg", "user_id", "created_at", "branch_service", "problem_status", "client_status_onyma", "company_name", "ip_interface"}
	for _, v := range d_map {
		data[v] = c.Query(v)
	}

	lists, err := h.services.RouterOnyma.GetRouterOnymaHistory(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllRouterOnymaHistorySearchByQueryResponse{
		Data: lists,
	})
}

// @Summary Добавление данных истории изменений
// @Security ApiKeyAuth
// @Tags History Router Onyma
// @Description Добавление в базу новых данных, история изменений, разные данные
// @ID insert-router-onyma-history
// @Accept  json
// @Produce  json
// @Param input body tc.ProblemRouterOnymaHistory true "Данные для внесения в базу"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/router_onyma_history [post]
func (h *Handler) insertRouterOnymaHistory(c *gin.Context) {
	var input tc.ProblemRouterOnymaHistory
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user_id, err := getUserId(c)
	if err != nil {
		return
	}
	input.User_id = user_id
	id, err := h.services.RouterOnyma.InsertRouterOnymaHistory(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Добавление данных
// @Security ApiKeyAuth
// @Tags Router Onyma
// @Description Добавление в базу новых данных
// @ID insert-router-onyma-speeds
// @Accept  json
// @Produce  json
// @Param input body tc.RouterOnyma true "Данные для внесения в базу"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/speeds [post]
func (h *Handler) insertRouterOnymaSpeeds(c *gin.Context) {
	var input tc.RouterOnyma
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.RouterOnyma.InsertRouterOnymaSpeeds(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Добавление Массива данных
// @Security ApiKeyAuth
// @Tags Router Onyma
// @Description Добавление в базу новых данных, сразу несколько строк
// @ID insert-router-onyma-speeds-all
// @Accept  json
// @Produce  json
// @Param input body []tc.RouterOnyma true "Данные для внесения в базу"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/speeds_all [post]
func (h *Handler) insertRouterOnymaSpeedsAll(c *gin.Context) {
	var input []tc.RouterOnyma
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.RouterOnyma.InsertRouterOnymaSpeedsAll(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Добавление данных
// @Security ApiKeyAuth
// @Tags Problem Router Onyma
// @Description Добавление проблемных данных в таблицу с проблемными данными
// @ID insert-problem-router-onyma-speeds
// @Accept  json
// @Produce  json
// @Param input body tc.ProblemRouterOnyma true "Проблемные Данные для внесения в базу"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem [post]
func (h *Handler) insertProblemRouterOnymaSpeeds(c *gin.Context) {
	var input tc.ProblemRouterOnyma
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.RouterOnyma.InsertProblemRouterOnymaSpeeds(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	/*c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
	*/
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Редактирование по id
// @Security ApiKeyAuth
// @Tags Problem Router Onyma
// @Description Редактирование информации проблемной строки.
// @ID update-problem-router-onyma-speed
// @Accept  json
// @Produce  json
// @Param id path string true "uuid для редактирования"
// @Param input body tc.UpdateProblemRouterOnymaSpeedInput true "Данные для редактирования"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem/{id} [put]
func (h *Handler) updateProblemRouterOnymaSpeeds(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id := c.Param("id")

	var input tc.UpdateProblemRouterOnymaSpeedInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.services.RouterOnyma.UpdateProblemRouterOnymaSpeeds(id, userId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Редактирование по interface_name
// @Security ApiKeyAuth
// @Tags Problem Router Onyma
// @Description Редактирование информации проблемной строки.
// @ID update-problem-router-onyma-speed-by-intrface-name
// @Accept  json
// @Produce  json
// @Param id path string true "uuid для редактирования"
// @Param input body tc.UpdateProblemRouterOnymaSpeedInput true "Данные для редактирования"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/problem/edit/{id} [put]
func (h *Handler) updateProblemStatusByInterface(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id := c.Param("id")

	var input tc.UpdateProblemRouterOnymaSpeedInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.services.RouterOnyma.UpdateProblemStatusByInterface(id, userId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

/*
// @Summary Старт сверки данных
// @Security ApiKeyAuth
// @Tags Router Onyma
// @Description Поиск проблем по скорости и не заполненых данных
// @ID start-search-problem
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/router_onyma/start_search_problem [get]
func (h *Handler) startSearchProblem(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	list, err := h.services.RouterOnyma.StartSearchProblem()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, list)
}
*/
