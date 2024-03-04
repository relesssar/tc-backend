package repository

import (
	"github.com/jmoiron/sqlx"
	tc "tc_kaztranscom_backend_go"
)

type Authorization interface {
	CreateUser(user tc.User) (string, error)
	GetUser(email, password string) (tc.User, error)
}
type AccessGroup interface {
	InsertAccessGroup(ag tc.AccessGroup) (string, error)
	GetAllAccessGroupByQuery(data map[string]string) ([]tc.AccessGroup, error)

	GetFilterAccessGroup(onyma tc.FilterAccessGroupSearch) ([]tc.FilterAccessGroupSearch, error)
	InsertFilterAccessGroup(proh tc.FilterAccessGroup) (string, error)
	DeleteFilter(id string) error

	InsertAccessGroupHistory(proh tc.AccessGroupHistory) (string, error)
	GetAccessGroupHistory(input tc.GetAccessGroupHistory) ([]tc.AccessGroupHistory, error)

	UpdateAccessGroup(id, userId string, data tc.UpdateAccessGroupInput) error
}
type ADLog interface {
	InsertADLog(log tc.ADLog) (string, error)
	GetAllByQuery(data map[string]string) ([]tc.ADLog, error)
	UpdateDepartmentInfo() error
}
type Contracts interface {
	//error
}
type User interface {
	GetById(userId string) (tc.User, error)
	GetModuleByUserId(userId string) ([]tc.UserModule, error)
	GetUsersList() ([]tc.UsersList, error)
	Update(data tc.UpdateUserInput) error
	Delete(id string) error
}
type RouterOnyma interface {
	GetProblemRouterOnymaSpeedsInfoByObjectNoClose(tc.ProblemRouterOnyma) ([]tc.ProblemRouterOnyma, error)
	CheckRouterOnymaSpeed() ([]tc.RouterOnyma, error)
	GetFilterRouterOnyma(onyma tc.FilterRouterOnymaSearch) ([]tc.FilterRouterOnymaSearch, error)
	InsertFilterRouterOnyma(proh tc.FilterRouterOnyma) (string, error)
	DeleteFilter(id string) error

	InsertRouterOnymaHistory(proh tc.ProblemRouterOnymaHistory) (string, error)
	InsertControlTimePauseHistory(ctph []tc.ControlTimePauseHistory) (string, error)
	InsertControlTimePause(ctph []tc.ControlTimePauseInsert) (string, error)

	InsertRouterOnymaSpeeds(routerOnyma tc.RouterOnyma) (string, error)
	InsertRouterOnymaSpeedsAll(routerOnyma []tc.RouterOnyma) (string, error)
	InsertProblemRouterOnymaSpeeds(pRo tc.ProblemRouterOnyma) (string, error)

	GetAllRouterOnymaSpeedsByDate(date string) ([]tc.RouterOnyma, error)
	GetAllRouterOnymaSpeedsByQuery(data map[string]string) ([]tc.RouterOnyma, error)
	GetAllProblemRouterOnymaSpeedsByQuery(data map[string]string) ([]tc.ProblemRouterOnymaQuery, error)
	GetCountAllProblemRouterOnymaSpeedsByQuery(data map[string]string) (int, error)
	GetDateListProblemRouterOnymaSpeed() ([]string, error)

	GetControlTimePauseByQuery(data map[string]string) ([]tc.ControlTimePause, error)
	GetRouterOnymaHistory(data map[string]string) ([]tc.ProblemRouterOnymaHistorySearch, error)
	UpdateProblemRouterOnymaSpeeds(id, userId string, data tc.UpdateProblemRouterOnymaSpeedInput) error
	UpdateProblemStatusByInterface(id, userId string, data tc.UpdateProblemRouterOnymaSpeedInput) error
	GetStatisticRouterOnymaSpeedsGroupByRouterByQuery(data map[string]string) ([]tc.RouterOnymaGroupByRouter, error)
	GetStatisticRouterOnymaSpeedsGroupByInsertDateByQuery(data map[string]string) ([]tc.RouterOnymaGroupByInsertDate, error)
}

type Repository struct {
	Authorization
	AccessGroup
	User
	RouterOnyma
	ADLog
}

func NewRepository(db, dbContract *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		AccessGroup:   NewAccessGroupPostgres(db),
		//ADLog:  	 	NewADLogPostgres(db),
		User:        NewUserPostgres(db),
		RouterOnyma: NewRouterOnymaPostgres(db),
		ADLog:       NewADLogPostgres(db, dbContract),
	}
}
