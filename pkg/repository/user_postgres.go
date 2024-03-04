package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	tc "tc_kaztranscom_backend_go"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

/*
func (r *UserPostgres) Create(userId string, category tc.InsertCategoryInput) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}
	if len(category.ParentId) == 0 {
		category.ParentId = "00000000-0000-0000-0000-000000000000"
	}
	logrus.Infof("%v", category)
	id := uuid.New()
	query := fmt.Sprintf("INSERT INTO %s (id,parent_id,user_id,name) values ($1,$2,$3,$4)", categoryTable)
	_, err = tx.Exec(query, id.String(), category.ParentId, userId, category.Name)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return id.String(), tx.Commit()

}

func (r *UserPostgres) GetAll(userId string) ([]note.Category, error) {

	var lists []note.Category
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1  order by name", categoryTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}
*/
func (r *UserPostgres) GetModuleByUserId(userId string) ([]tc.UserModule, error) {

	var lists []tc.UserModule
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1  order by module_name", userModuleTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *UserPostgres) GetUsersList() ([]tc.UsersList, error) {

	var lists []tc.UsersList
	query := fmt.Sprintf("SELECT u.id,u.name,u.email,u.phone, string_agg(um.module_name,', ') as module_name FROM %s as u LEFT OUTER JOIN %s as um ON u.id=um.user_id  GROUP BY u.id,u.name,u.email,u.phone", usersTable, userModuleTable)
	err := r.db.Select(&lists, query)
	return lists, err
}

func (r *UserPostgres) GetById(userId string) (tc.User, error) {

	var user tc.User
	query := fmt.Sprintf("SELECT id,name,email,phone FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, userId)
	return user, err
}
func (r *UserPostgres) Update(input tc.UpdateUserInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	logrus.Infof("User.Update: id=%s", input.Id)
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Name != "" {
		logrus.Infof("User.Update: Name=%s", input.Name)
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}
	if input.Email != "" {
		logrus.Infof("User.Update: Email=%s", input.Email)
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, input.Email)
		argId++
	}
	if input.Phone != "" {
		logrus.Infof("User.Update: Phone=%s", input.Phone)
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argId))
		args = append(args, input.Phone)
		argId++
	}
	if input.Password != "" {

		logrus.Infof("User.Update: Password=%s", input.Password)
		setValues = append(setValues, fmt.Sprintf("password_hash=$%d", argId))
		args = append(args, input.Password)
		argId++
	}
	if input.ModuleName != "" {
		logrus.Infof("User.Update: DELETE ModuleName where user_id=%s", input.Id)
		r.db.Query(fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", userModuleTable), input.Id)
		logrus.Infof("User.Update: INSERT ModuleName=%s ", input.ModuleName)
		query_module := fmt.Sprintf("INSERT INTO %s  values ($1,$2,$3,$4)", userModuleTable)
		for _, v := range strings.Split(input.ModuleName, ",") {
			id := uuid.New()

			r.db.Query(query_module, id.String(), input.Id, v, v)
		}
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s set %s  WHERE id=$%d", usersTable, setQuery, argId)
	args = append(args, input.Id)
	argId++
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *UserPostgres) Delete(userId string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	//Проверяю есть ещё юзеры или нет
	var id []string
	query := fmt.Sprintf("SELECT u.id  FROM %s as u INNER JOIN %s as um ON um.user_id=u.id and um.module_name='users' WHERE u.id<>$1 ", usersTable, userModuleTable)
	err = r.db.Select(&id, query, userId)
	logrus.Printf("Проверяю есть ещё юзеры или нет? id=%s", id)
	if len(id) == 0 {
		logrus.Errorf("Нельзя удалить Последнего юзера! id=", id)
		return err
	}
	//END Проверяю
	query = fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", userModuleTable)
	_, err = r.db.Exec(query, userId)
	query = fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	_, err = r.db.Exec(query, userId)

	return tx.Commit()
}
