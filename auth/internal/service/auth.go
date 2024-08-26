package service

import (
	"context"

	pb "auth/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	return &pb.Response{
		Code: 0,
		Msg:  "",
		Data: &pb.Response_LoginResponse{LoginResponse: &pb.LoginResponse{Token: "123"}},
	}, nil
}
func (s *AuthService) Authentication(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}
func (s *AuthService) Logout(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}
func (s *AuthService) OpenLoginInfo(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}
func (s *AuthService) OpenLogin(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	return &pb.Response{}, nil
}
