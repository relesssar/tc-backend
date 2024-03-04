package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	tc "tc_kaztranscom_backend_go"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user tc.User) (string, error) {
	user_id := uuid.New()
	query := fmt.Sprintf("INSERT INTO %s (id,name,email,phone,password_hash) values ($1,$2,$3,$4,$5)", usersTable)
	r.db.Query(query, user_id.String(), user.Name, user.Email, user.Phone, user.Password)
	//Если есть модули, добавляю
	if user.ModuleName != "" {
		query = fmt.Sprintf("INSERT INTO %s  values ($1,$2,$3,$4)", userModuleTable)
		for _, v := range strings.Split(user.ModuleName, ",") {
			id := uuid.New()
			r.db.Query(query, id.String(), user_id, v, v)
		}
	}
	return user_id.String(), nil
}

func (r *AuthPostgres) GetUser(email, password string) (tc.User, error) {
	logrus.Info("GetUser: start")
	var user tc.User
	query := fmt.Sprintf("SELECT id,name,email,phone FROM %s WHERE email=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)
	logrus.Infof("%s \n -> %v", email, user)
	return user, err
}
