package api

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/pkg/openapi"
)

type MockApiService struct {
}

func NewMockApiService() openapi.DefaultApiServicer {
	return &MockApiService{}
}

func (s *MockApiService) CreateUser(createUserInput openapi.CreateUserInput) (interface{}, error) {
	return openapi.CreateUserOutput{
		Id:           1,
		EmailAddress: createUserInput.Address,
		LastName:     createUserInput.LastName,
		FirstName:    createUserInput.FirstName,
		Birthday:     createUserInput.Birthday,
		Address:      createUserInput.Address,
	}, nil
}

func (s *MockApiService) GetUser(id int64) (interface{}, error) {
	return openapi.GetUserOutput{
		Id:           id,
		EmailAddress: "example@example.com",
		LastName:     "田中",
		FirstName:    "太郎",
		Birthday:     "2000-01-01",
		Address:      "東京都新宿区西新宿２丁目８−１",
	}, nil
}

func (s *MockApiService) UpdateUser(updateUserInput openapi.UpdateUserInput) (interface{}, error) {
	return openapi.UpdateUserOutput{
		Id:           updateUserInput.Id,
		EmailAddress: updateUserInput.EmailAddress,
		LastName:     updateUserInput.LastName,
		FirstName:    updateUserInput.FirstName,
		Birthday:     updateUserInput.Birthday,
		Address:      updateUserInput.Address,
	}, nil
}
