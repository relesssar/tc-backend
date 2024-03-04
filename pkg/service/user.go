package service

import (
	"crypto/sha1"
	"fmt"
	tc "tc_kaztranscom_backend_go"
	"tc_kaztranscom_backend_go/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsersList() ([]tc.UsersList, error) {
	return s.repo.GetUsersList()
}

func (s *UserService) GetModuleByUserId(userId string) ([]tc.UserModule, error) {
	return s.repo.GetModuleByUserId(userId)
}

func (s *UserService) Update(data tc.UpdateUserInput) error {
	if data.Password != "" {
		data.Password = s.GeneratePasswordHash(data.Password)
	}
	return s.repo.Update(data)
}
func (s *UserService) Delete(id string) error {

	return s.repo.Delete(id)
}

/*
func (s *UserService) Create(userId string, category note.InsertCategoryInput) (string, error) {

	return s.repo.Create(userId, category)
}
func (s *UserService) GetAll(userId string) ([]note.Category, error) {

	return s.repo.GetAll(userId)
}*/
func (s *UserService) GetById(userId string) (tc.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
