package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	tc "tc_kaztranscom_backend_go"
	//tc "tc_kaztranscom_backend_go"
)

/*
// @Summary Создание
// @Security ApiKeyAuth
// @Tags Категории Заметок
// @Description Создание новой Категории
// @ID create-category
// @Accept  json
// @Produce  json
// @Param input body note.InsertCategoryInput true "Информация для создания"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/category [post]
func (h *Handler) createCategory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	//Данные категории
	var input note.InsertCategoryInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Category.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllCategoriesResponse struct {
	Data []note.Category `json:"data"`
}
*/
/*
// @Summary Список
// @Security ApiKeyAuth
// @Tags Категории Заметок
// @Description Список всех Категорий
// @ID get-all-category
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllCategoriesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/category [get]
func (h *Handler) getAllCategories(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.Category.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	/*c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	logrus.Info("getAllCategories() ")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	} else {
		c.Next()
	}* /
	c.JSON(http.StatusOK, getAllCategoriesResponse{
		Data: lists,
	})

}
*/

type getAllUserModuleResponse struct {
	Data []tc.UserModule `json:"data"`
}
type getAllUsersResponse struct {
	Data []tc.UsersList `json:"data"`
}

// @Summary Информация
// @Security ApiKeyAuth
// @Tags Пользователь
// @Description Информация по пользователю по uuid пользователя
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "uuid пользователя"
// @Success 200 {object} tc.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/{id} [get]
func (h *Handler) getUserById(c *gin.Context) {

	id := c.Param("id")
	//logrus.Infof("getCategoryById(%s)",id)
	list, err := h.services.User.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, list)
}

// @Summary Список всех пользователей
// @Security ApiKeyAuth
// @Tags Пользователь
// @Description Информация по всем пользователям
// @ID get-users-list
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllUsersResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/list [get]
func (h *Handler) getUsersList(c *gin.Context) {

	lists, err := h.services.User.GetUsersList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: lists,
	})
}

// @Summary Информация по модулям
// @Security ApiKeyAuth
// @Tags Пользователь
// @Description Информация по модулям доступа пользователя, по uuid пользователя
// @ID get-user-module-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "uuid пользователя"
// @Success 200 {object} getAllUserModuleResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/module/{id} [get]
func (h *Handler) getUserModuleByUserId(c *gin.Context) {

	id := c.Param("id")
	//logrus.Infof("getCategoryById(%s)",id)
	lists, err := h.services.User.GetModuleByUserId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getAllUserModuleResponse{
		Data: lists,
	})
}

// @Summary Информация по модулям
// @Security ApiKeyAuth
// @Tags Пользователь
// @Description Информация по модулям доступа авторизованного пользователя
// @ID get-user-module
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllUserModuleResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/module [get]
func (h *Handler) getUserModule(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	//logrus.Infof("getCategoryById(%s)",id)
	lists, err := h.services.User.GetModuleByUserId(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, getAllUserModuleResponse{
		Data: lists,
	})
}

// @Summary Редактирование
// @Security ApiKeyAuth
// @Tags Пользователь
// @Description Редактирование информации о пользователе
// @ID update-user
// @Accept  json
// @Produce  json
// @Param id path string true "uuid пользователя"
// @Param input body tc.UpdateUserInput true "Поля для изменений"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {

	_, err := getUserId(c)
	if err != nil {
		return
	}
	//id := c.Param("id")

	var input tc.UpdateUserInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if input.Password != "" {
		input.Password = h.services.Authorization.GeneratePasswordHash(input.Password)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		}
	}
	if err := h.services.User.Update(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

// @Summary Удаление
// @Security ApiKeyAuth
// @Tags Пользователь
// @Description Удаление Пользователя
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param id path string true "uuid пользователя"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {

	id := c.Param("id")
	err := h.services.User.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
