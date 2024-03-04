package service

import (
	tc "tc_kaztranscom_backend_go"
	"tc_kaztranscom_backend_go/pkg/repository"
)

type Authorization interface {
	CreateUser(user tc.User) (string, error)
	GenerateToken(email, password string) (string, error)
	GeneratePasswordHash(password string) string
	ParseToken(token string) (string, error)
}
type AccessGroup interface {
	InsertAccessGroup(ac tc.AccessGroup) (string, error)
	GetAllAccessGroupByQuery(data map[string]string) ([]tc.AccessGroup, error)
	InsertFilterAccessGroup(accessGroup tc.FilterAccessGroup) (string, error)
	GetFilterAccessGroup(data tc.FilterAccessGroupSearch) ([]tc.FilterAccessGroupSearch, error)
	DeleteFilter(id string) error
	CreateExcellAccessGroup(data []tc.AccessGroup) (string, error)

	GetAccessGroupHistory(input tc.GetAccessGroupHistory) ([]tc.AccessGroupHistory, error)
	InsertAccessGroupHistory(agh tc.AccessGroupHistory) (string, error)
	UpdateAccessGroup(id, userId string, data tc.UpdateAccessGroupInput) error
}
type User interface {
	GetUsersList() ([]tc.UsersList, error)
	GetById(userId string) (tc.User, error)
	GetModuleByUserId(userId string) ([]tc.UserModule, error)
	Update(data tc.UpdateUserInput) error
	Delete(id string) error
}
type ADLog interface {
	CheckLogFileCSV() error
	GetAllByQuery(data map[string]string) ([]tc.ADLog, error)
	CreateExcell(data []tc.ADLog) (string, error)
}

type RouterOnyma interface {
	CheckRouterOnymaSpeed() (string, error)
	CheckCloseProblemRouterOnymaSpeed(userId string) (string, error)

	CheckControlTimePause(userId string) (string, error)
	GetControlTimePauseByQuery(data map[string]string) ([]tc.ControlTimePause, error)

	CreateExcellProblemRouterOnymaSpeed(data []tc.ProblemRouterOnymaQuery) (string, error)
	CreateExcellProblemRouterOnymaForm1(data []tc.ProblemRouterOnymaQuery) (string, error)
	CreateExcellControlTimePause(data []tc.ControlTimePause) (string, error)

	InsertFilterRouterOnyma(routerOnyma tc.FilterRouterOnyma) (string, error)
	GetFilterRouterOnyma(data tc.FilterRouterOnymaSearch) ([]tc.FilterRouterOnymaSearch, error)
	DeleteFilter(id string) error

	InsertRouterOnymaHistory(routerOnyma tc.ProblemRouterOnymaHistory) (string, error)
	InsertRouterOnymaSpeeds(routerOnyma tc.RouterOnyma) (string, error)
	InsertRouterOnymaSpeedsAll(routerOnyma []tc.RouterOnyma) (string, error)
	InsertProblemRouterOnymaSpeeds(problemRouterOnyma tc.ProblemRouterOnyma) (string, error)

	GetAllRouterOnymaSpeedsByDate(date string) ([]tc.RouterOnyma, error)
	GetAllRouterOnymaSpeedsByQuery(data map[string]string) ([]tc.RouterOnyma, error)

	GetAllProblemRouterOnymaSpeedsByQuery(data map[string]string) ([]tc.ProblemRouterOnymaQuery, error)
	GetCountAllProblemRouterOnymaSpeedsByQuery(data map[string]string) (int, error)
	GetDateListProblemRouterOnymaSpeed() ([]string, error)
	GetRouterOnymaHistory(data map[string]string) ([]tc.ProblemRouterOnymaHistorySearch, error)
	UpdateProblemRouterOnymaSpeeds(id, userId string, data tc.UpdateProblemRouterOnymaSpeedInput) error
	UpdateProblemStatusByInterface(id, userId string, data tc.UpdateProblemRouterOnymaSpeedInput) error

	GetStatisticStatatusAll(data []tc.ProblemRouterOnymaQuery) (map[string]map[string]int, error)
	GetStatisticRouterOnymaSpeedsGroupByRouterByQuery(data map[string]string) ([]tc.RouterOnymaGroupByRouter, error)

	GetStatisticRouterOnymaSpeedsGroupByInsertDateByQuery(data map[string]string) ([]tc.RouterOnymaGroupByInsertDate, error)
}

type Service struct {
	Authorization
	AccessGroup
	ADLog
	User
	RouterOnyma

	//Category
	//Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		AccessGroup:   NewAccessGroupService(repos.AccessGroup),
		ADLog:         NewADLogService(repos.ADLog),
		User:          NewUserService(repos.User),
		RouterOnyma:   NewRouterOnymaService(repos.RouterOnyma),
	}
}
