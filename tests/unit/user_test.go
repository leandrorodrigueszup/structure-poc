package unit

import (
	"poc/internal/repository"
	"poc/internal/use_case/user"
	mocks "poc/tests/unit/mocks/repository"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite
	findUserById user.FindUserById
	gormDB       *mocks.Database
}

func (u *UserSuite) SetupSuite() {
	u.gormDB = &mocks.Database{}
	userRep, _ := repository.NewUserRepository(u.gormDB, "../../internal/repository/queries")
	u.findUserById = user.NewFindUserById(userRep)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (u *UserSuite) TestGetByID() {
	mockDB(u)
	a, err := u.findUserById.Execute(uuid.New())

	require.NotNil(u.T(), a)
	require.Nil(u.T(), err)
}

func mockDB(u *UserSuite) {
	u.gormDB.On("Model", mock.Anything).Return(u.gormDB)
	u.gormDB.On("Where", mock.Anything, mock.Anything).Return(u.gormDB)
	u.gormDB.On("First", mock.Anything).Return(u.gormDB)
	u.gormDB.On("GetError").Return(nil)
}

// func (u *UserSuite) TestErrorGetByID() {
// 	u.gormDB.On("First", mock.Anything).Return(models.User{}, logging.NewError("error", nil, nil))
// 	_, err := u.findUserById.Execute(uuid.New())

// 	require.Error(u.T(), err)
// }
