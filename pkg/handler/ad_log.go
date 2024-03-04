package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	tc "tc_kaztranscom_backend_go"
)

type getAllADLogByQueryResponse struct {
	Data []tc.ADLog `json:"data"`
}

// @Summary Список данных Ad log
// @Security ApiKeyAuth
// @Tags Active Directory
// @Description Список данных AD log по переданным параметрам
// @ID get-ad-log-by-query
// @Accept  json
// @Produce  json
// @Param date query string false "в формате yyyy-mm"
// @Param from_date query string false "начало поиска в формате yyyy-mm-dd,(включительно)"
// @Param to_date query string false "конец поиска в формате yyyy-mm-dd,(включительно)"
// @Param ad_login query string false "не строгий поиск, логин или фамилия или департамент"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/ad_log [get]
func (h *Handler) getAllByQuery(c *gin.Context) {
	/*_, err := getUserId(c)
	if err != nil {
		return
	}*/
	data := make(map[string]string)

	data["date"] = c.Query("date")
	data["from_date"] = c.Query("from_date")
	data["to_date"] = c.Query("to_date")
	data["ad_login"] = c.Query("ad_login")
	lists, err := h.services.ADLog.GetAllByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getAllADLogByQueryResponse{
		Data: lists,
	})
}

// @Summary Проверка и загрузка данных
// @Tags Active Directory
// @Description Проверка на существование лог файла в папке, для загрузки в базу данных,
// @Description Путь к папке в файле настроек проекта
// @ID check-log-file-csv
// @Accept  json
// @Produce  json
// @Success 200 {string} string status
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/ad_log/check_log_file_csv [get]
func (h *Handler) checkLogFileCSV(c *gin.Context) {

	err := h.services.ADLog.CheckLogFileCSV()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// @Summary Скачивание файла
// @Tags Active Directory
// @Description Скачивание файла Excell с данными лога
// @ID download-excell-file-ad-log
// @Accept  json
// @Produce  json
// @Success 200 {string} string status
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/ad_log/download [get]
func (h *Handler) getExcellADLog(c *gin.Context) {

	//получаю данные с учётом фильтра
	data := make(map[string]string)

	data["ad_login"] = c.Query("ad_login")
	data["date"] = c.Query("date")
	data["from_date"] = c.Query("from_date")
	data["to_date"] = c.Query("to_date")

	//logrus.Infof("getCategoryById(%s)",id)
	lists, err := h.services.ADLog.GetAllByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	//logrus.Info(lists)
	//создаю файл ексель и возвращаю имя файла
	filename, err := h.services.ADLog.CreateExcell(lists)
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
