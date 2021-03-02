package repository

import (
	"fmt"
	"poc/internal/domain"
	"poc/internal/logging"
	"poc/internal/repository/models"

	"github.com/google/uuid"
	"github.com/nleof/goyesql"
)

type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindAllCustom() ([]domain.User, error)
	Create(user domain.User) (domain.User, error)
	GetByID(id uuid.UUID) (domain.User, error)
	Update(id uuid.UUID, user domain.User) (domain.User, error)
	Delete(id uuid.UUID) error
}

type userRepository struct {
	db      Database
	queries goyesql.Queries
}

func NewUserRepository(db Database, queriesPath string) (UserRepository, error) {
	queries, err := goyesql.ParseFile(fmt.Sprintf("%s/%s", queriesPath, "user_queries.sql"))
	if err != nil {
		return userRepository{}, err
	}

	return userRepository{db: db, queries: queries}, nil
}

func (r userRepository) FindAll() ([]domain.User, error) {
	var users []models.User

	if res := r.db.Find(&users); res.GetError() != nil {
		return nil, logging.NewError("Find all users failed", res.GetError(), nil, "repository.FindAll.Find")
	}

	usersFound := make([]domain.User, 0)
	for _, u := range users {
		usersFound = append(usersFound, domain.User(u))
	}

	return usersFound, nil
}

func (r userRepository) FindAllCustom() ([]domain.User, error) {
	var users []models.User

	if res := r.db.Raw(r.queries["find-all-custom"]).Scan(&users); res.GetError() != nil {
		return nil, logging.NewError("Find all users failed", res.GetError(), nil, "repository.FindAllCustom.Find")
	}

	usersFound := make([]domain.User, 0)
	for _, u := range users {
		usersFound = append(usersFound, domain.User(u))
	}

	return usersFound, nil
}

func (r userRepository) Create(user domain.User) (domain.User, error) {
	user.ID = uuid.New()
	userToSave := models.User(user)
	if res := r.db.Save(&userToSave); res.GetError() != nil {
		return domain.User{}, logging.NewError("Save user failed", res.GetError(), nil, "repository.Create.Save")
	}

	return user, nil
}

func (r userRepository) GetByID(id uuid.UUID) (domain.User, error) {
	var user models.User

	if res := r.db.Model(models.User{}).Where("id = ?", id).First(&user); res.GetError() != nil {
		return domain.User{}, logging.NewError("Find user failed", res.GetError(), nil, "repository.GetById.First")
	}

	return domain.User(user), nil
}

func (r userRepository) Update(id uuid.UUID, user domain.User) (domain.User, error) {
	userToUpdate := models.User(user)
	if res := r.db.Model(models.User{}).Where("id = ?", id).Updates(&userToUpdate); res.GetError() != nil {
		return domain.User{}, logging.NewError("Update user failed", res.GetError(), nil, "repository.Update.Updates")
	}

	return user, nil
}

func (r userRepository) Delete(id uuid.UUID) error {
	if res := r.db.Delete(models.User{}, id); res.GetError() != nil {
		return logging.NewError("Delete user failed", res.GetError(), nil, "repository.Delete.Delete")
	}

	return nil
}
