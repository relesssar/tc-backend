package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "tc_kaztranscom_backend_go/docs"
	"tc_kaztranscom_backend_go/pkg/service"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	gin.ForceConsoleColor()
	router := gin.New()

	/*
		Разрешает делать запросы с других хостов
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			//return origin == "http://note_frontend_vue3.loc"
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
	/*  END Разрешает делать запросы с других хостов */

	/*docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}*/
	/* Для свагера запустить в корне swag init -g cmd/main.go */
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	/* Для Запуска проверки строк router_onyma на проблемные */
	router.GET("api/router_onyma/problem/check_router_onyma_speed", h.checkRouterOnymaSpeed)

	/* Для Запуска проверки и грузки лог файла Актив Директори, в базу данных*/
	router.GET("api/ad_log/check_log_file_csv", h.checkLogFileCSV)

	/* Только для пользователе прошедших проверку токена h.userIdentity
	TODO Чтобы создать первого суперадмина, нужно убрать проверку
	auth := router.Group("/auth")
	*/
	router.POST("/auth/sign-in", h.signIn)
	auth := router.Group("/auth", h.userIdentity)
	{
		auth.POST("/sign-up", h.signUp)
	}

	/* Только для пользователе прошедших проверку токена h.userIdentity */
	api := router.Group("/api", h.userIdentity)
	{
		access_group := api.Group("/access_group")
		{
			access_group.POST("/", h.insertAccessGroup)
			access_group.POST("/access_group_history", h.insertAccessGroupHistory)
			access_group.POST("/get_access_group_history", h.getAccessGroupHistory)
			access_group.GET("/", h.getAllAccessGroupByQuery)
			access_group.GET("/download", h.getExcellAccessGroup)

			access_group.POST("/filter_access_group", h.insertFilterAccessGroup)
			access_group.GET("/filter_access_group", h.getFilterAccessGroup)
			access_group.DELETE("/filter_access_group/:id", h.deleteFilterAccessGroup)
			access_group.PUT("/:id", h.updateAccessGroup)
		}
		ad_log := api.Group("/ad_log")
		{
			ad_log.GET("/download", h.getExcellADLog)
			ad_log.GET("/", h.getAllByQuery)
		}

		user := api.Group("/user")
		{
			user.GET("/list", h.getUsersList)
			user.GET("/:id", h.getUserById)
			user.GET("/module/:id", h.getUserModuleByUserId)
			user.PUT("/:id", h.updateUser)
			user.DELETE("/:id", h.deleteUser)
			user.GET("/module/", h.getUserModule)

		}
		router_onyma := api.Group("/router_onyma")
		{
			{
				router_onyma.POST("/speeds", h.insertRouterOnymaSpeeds)
				router_onyma.POST("/speeds_all", h.insertRouterOnymaSpeedsAll)
				router_onyma.GET("/speeds/:date", h.getAllRouterOnymaSpeedsByDate)
				router_onyma.GET("/speeds", h.getAllRouterOnymaSpeedsByQuery)
				router_onyma.GET("/statistic/group_by_router", h.getStatisticRouterOnymaSpeedsGroupByRouterByQuery)

				router_onyma.POST("/router_onyma_history", h.insertRouterOnymaHistory)
				router_onyma.GET("/router_onyma_history", h.getRouterOnymaHistory)

				router_onyma.POST("/filter_router_onyma", h.insertFilterRouterOnyma)
				router_onyma.GET("/filter_router_onyma", h.getFilterRouterOnyma)
				router_onyma.DELETE("/filter_router_onyma/:id", h.deleteFilter)

				router_onyma.GET("/problem/date_list", h.getDateListProblemRouterOnymaSpeed)
				//router_onyma.GET("/problem/check_router_onyma_speed", h.checkRouterOnymaSpeed)
				router_onyma.POST("/problem", h.insertProblemRouterOnymaSpeeds)
				router_onyma.GET("/problem", h.getAllProblemRouterOnymaSpeedsByQuery)
				router_onyma.GET("/problem/count", h.getCountAllProblemRouterOnymaSpeedsByQuery)
				router_onyma.PUT("/problem/:id", h.updateProblemRouterOnymaSpeeds)
				router_onyma.PUT("/problem/edit/:id", h.updateProblemStatusByInterface)
				router_onyma.GET("/problem/download", h.getExcellProblemRouterOnymaSpeeds)

				router_onyma.GET("/control_time_pause/download", h.getExcellControlTimePause)
				router_onyma.GET("/control_time_pause", h.getControlTimePauseByQuery)
				//Statistic
				router_onyma.GET("/problem/statistic/status_all", h.getStatisticStatusByQuery)
				router_onyma.GET("/problem/statistic/form1", h.getExcellProblemRouterOnymaForm1)

				/* Для Запуска проверки проблемных строк problem_router_onyma_speed на не проблемные
				Тоесть исправленым строкам ставлю статус ОК, от пользователя Bot */
				router_onyma.GET("/problem/check_close_problem_router_onyma_speed", h.checkCloseProblemRouterOnymaSpeed)

				/* Для Запуск проверки Контроль временного отключения , от пользователя Bot */
				router_onyma.GET("/problem/check_control_time_pause", h.checkControlTimePause)

				//router_onyma.GET("/start_search_problem", h.startSearchProblem)
			}

		}

		return router
	}
}
