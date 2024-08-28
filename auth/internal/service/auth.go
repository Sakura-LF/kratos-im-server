package service

import (
	"auth/internal/biz"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"

	pb "auth/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	user *biz.UserUsecase
	log  *log.Helper
}

func NewAuthService(stu *biz.UserUsecase, logger log.Logger) *AuthService {
	return &AuthService{
		user: stu,
		log:  log.NewHelper(logger),
	}
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	response := &pb.Response{}
	fmt.Println("Test")
	return response, nil
}
func (s *AuthService) Authentication(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	fmt.Println("Test")
	return &pb.Response{}, nil
}
func (s *AuthService) Logout(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	fmt.Println("logout")
	return &pb.Response{}, nil
}
func (s *AuthService) OpenLoginInfo(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	fmt.Println("openLoginInfo")
	return &pb.Response{}, nil
}
func (s *AuthService) OpenLogin(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	fmt.Print("openLogin")
	return &pb.Response{}, nil
}
