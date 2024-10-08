package service

import (
	"context"
	"user/internal/biz"

	pb "user/api/user/v1"
)

type UsersService struct {
	pb.UnimplementedUsersServer

	uc *biz.UserUsecase
}

func NewUsersService(uc *biz.UserUsecase) *UsersService {
	return &UsersService{
		uc: uc,
	}
}

func (s *UsersService) CreateUser(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {

	return &pb.UserInfoResponse{}, nil
}
