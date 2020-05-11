package service

import (
	"context"

	"github.com/d0riven/HelloApiSchema/server/golang/proto/pkg/pb"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Create(ctx context.Context, input *pb.CreateUserInput) (*pb.CreateUserOutput, error) {
	return &pb.CreateUserOutput{
		Id:           1,
		EmailAddress: input.GetEmailAddress(),
		LastName:     input.GetLastName(),
		FirstName:    input.GetFirstName(),
		Birthday:     input.GetBirthday(),
		Address:      input.GetAddress(),
	}, nil
}

func (u *UserService) Update(ctx context.Context, input *pb.UpdateUserInput) (*pb.UpdateUserOutput, error) {
	return &pb.UpdateUserOutput{
		Id:           input.GetId(),
		EmailAddress: input.GetEmailAddress(),
		LastName:     input.GetLastName(),
		FirstName:    input.GetFirstName(),
		Birthday:     input.GetBirthday(),
		Address:      input.GetAddress(),
	}, nil
}

func (u *UserService) Get(ctx context.Context, input *pb.GetUserInput) (*pb.GetUserOutput, error) {
	return &pb.GetUserOutput{
		Id:           input.GetId(),
		EmailAddress: "example@example.com",
		LastName:     "山田",
		FirstName:    "太郎",
		Birthday:     "2000-01-01",
		Address:      "東京都新宿区西新宿２丁目８−１",
	}, nil
}
