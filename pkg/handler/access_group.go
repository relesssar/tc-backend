package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"

	//"os"
	tc "tc_kaztranscom_backend_go"
)

type getAllAccessGroupByQueryResponse struct {
	Data []tc.AccessGroup `json:"data"`
}

type getAllFilterAccessGroupSearchByQueryResponse struct {
	Data []tc.FilterAccessGroupSearch `json:"data"`
}
type getAllAccessGroupHistoryByQueryResponse struct {
	Data []tc.AccessGroupHistory `json:"data"`
}

// @Summary Добавление данных в filter(Исключения)
// @Security ApiKeyAuth
// @Tags Filter Access Group
// @Description Добавление в базу новых данных, история изменений, разные данные
// @ID insert-filter-access-group
// @Accept  json
// @Produce  json
// @Param input body tc.FilterAccessGroup true "Данные для внесения в базу"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/access_group/filter_access_group [post]
func (h *Handler) insertFilterAccessGroup(c *gin.Context) {
	var input tc.FilterAccessGroup
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user_id, err := getUserId(c)
	if err != nil {
		return
	}
	input.User_id = user_id
	id, err := h.services.AccessGroup.InsertFilterAccessGroup(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Получение данных
// @Security ApiKeyAuth
// @Tags Filter Access Group
// @Description Получение  данных, Фильтры, разные данные
// @ID get-filter-access-group
// @Accept  json
// @Produce  json
// @Param input query tc.FilterAccessGroupSearch true "Любые данные для поиск в базе"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/access_group/filter_access_group [get]
func (h *Handler) getFilterAccessGroup(c *gin.Context) {
	var input tc.FilterAccessGroupSearch
	input.Id = c.Query("id")
	input.Filter_type = c.Query("filter_type")
	input.Filter_val = c.Query("filter_val")
	input.Filter_desc = c.Query("filter_desc")
	input.User_id = c.Query("user_id")
	input.Created_at = c.Query("created_at")
	logrus.Infof("%v", input)
	lists, err := h.services.AccessGroup.GetFilterAccessGroup(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllFilterAccessGroupSearchByQueryResponse{
		Data: lists,
	})
}

// @Summary Удаление
// @Security ApiKeyAuth
// @Tags Filter Access Group
// @Description Удаление Фильтра
// @ID delete-filter-access-group
// @Accept  json
// @Produce  json
// @Param id path string true "uuid фильтра"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/access_group/filter_access_group/{id} [delete]
func (h *Handler) deleteFilterAccessGroup(c *gin.Context) {

	id := c.Param("id")
	err := h.services.AccessGroup.DeleteFilter(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Список данных
// @Security ApiKeyAuth
// @Tags Access Group
// @Description Список данных Access Group по переданным параметрам
// @ID get-access-group-by-query
// @Accept  json
// @Produce  json
// @Param date query string false "в формате yyyy-mm-dd"
// @Param check_filter query string false "Учитывать или нет фильтр"
// @Param access_status query string false "Закрыто, В обработке"
// @Param router_name query string false "можно указать несколько, через запятую "
// @Success 200 {object} getAllAccessGroupByQueryResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/access_group [get]
func (h *Handler) getAllAccessGroupByQuery(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	data := make(map[string]string)

	data["date"] = c.Query("date")
	data["check_filter"] = c.Query("check_filter")
	data["access_status"] = c.Query("access_status")
	data["router_name"] = c.Query("router_name")

	//logrus.Infof("getCategoryById(%s)",id)
	lists, err := h.services.AccessGroup.GetAllAccessGroupByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getAllAccessGroupByQueryResponse{
		Data: lists,
	})
}

// @Summary Добавление данных
// @Security ApiKeyAuth
// @Tags Access Group
// @Description Добавление в базу новых данных
// @ID insert-access-group
// @Accept  json
// @Produce  json
// @Param input body tc.AccessGroup true "Данные для внесения в базу"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/access_group [post]
func (h *Handler) insertAccessGroup(c *gin.Context) {
	logrus.Info("\ninsertAccessGroup()")
	var input tc.AccessGroup
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.AccessGroup.InsertAccessGroup(input)
	if err != nil {
		logrus.Errorln(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	logrus.Infof("ok %s\n", id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Добавление данных истории изменений
// @Security ApiKeyAuth
// @Tags Access Group
// @Description Добавление в базу новых данных, история изменений, разные данные
// @ID insert-access-group-history
// @Accept  json
// @Produce  json
// @Param input body tc.ProblemRouterOnymaHistory true "Данные для внесения в базу"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/access_group/access_group_history [post]
func (h *Handler) insertAccessGroupHistory(c *gin.Context) {
	var input tc.AccessGroupHistory
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user_id, err := getUserId(c)
	if err != nil {
		return
	}
	input.User_id = user_id
	id, err := h.services.AccessGroup.InsertAccessGroupHistory(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Получение данных
// @Security ApiKeyAuth
// @Tags Access Group
// @Description Получение  данных, история изменений, разные данные
// @ID get-access-group-history
// @Accept  json
// @Produce  json
// @Param input query tc.GetAccessGroupHistory true "Любые данные для поиск в базе"
// @Success 200 {string} string uuid
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/access_group/access_group_history [post]
func (h *Handler) getAccessGroupHistory(c *gin.Context) {
	var input tc.GetAccessGroupHistory
	input.Ids, _ = c.GetQueryArray("ids")
	input.Id = c.Query("id")
	/*
		input.Access_group_id = c.Query("access_group_id")
		input.Old_val = c.Query("old_val")
		input.New_val = c.Query("new_val")
		input.Msg = c.Query("msg")
		input.User_id = c.Query("user_id")
		input.Created_at = c.Query("created_at")*/
	logrus.Infof("%v", input)
	lists, err := h.services.AccessGroup.GetAccessGroupHistory(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllAccessGroupHistoryByQueryResponse{
		Data: lists,
	})
}

// @Summary Редактирование
// @Security ApiKeyAuth
// @Tags Access Group
// @Description Редактирование информации
// @ID update-access-group
// @Accept  json
// @Produce  json
// @Param id path string true "uuid для редактирования"
// @Param input body tc.UpdateAccessGroupInput true "Данные для редактирования"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/access_group/{id} [put]
func (h *Handler) updateAccessGroup(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id := c.Param("id")

	var input tc.UpdateAccessGroupInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.services.AccessGroup.UpdateAccessGroup(id, userId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Скачивание файла списка Access Group
// @Tags Access Group
// @Description Скачивание файла статистики Форма 1(Excell)
// @ID download-excell-file-access-group
// @Accept  json
// @Produce  json
// @Param date query string false "в формате yyyy-mm-dd"
// @Param check_filter query string false "Учитывать или нет фильтр"
// @Param access_status query string false "Закрыто, В обработке"
// @Param router_name query string false "можно указать несколько, через запятую "
// @Success 200 {string} string status
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/access_group/download [get]
func (h *Handler) getExcellAccessGroup(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	data := make(map[string]string)

	data["date"] = c.Query("date")
	data["check_filter"] = c.Query("check_filter")
	data["access_status"] = c.Query("access_status")
	data["router_name"] = c.Query("router_name")

	//logrus.Infof("getCategoryById(%s)",id)
	lists, err := h.services.AccessGroup.GetAllAccessGroupByQuery(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	//создаю файл ексель и возвращаю имя файла
	filename, err := h.services.AccessGroup.CreateExcellAccessGroup(lists)
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
