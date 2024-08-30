package service

import (
	"auth/internal/biz"
	"auth/internal/conf"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
	"util/jwt"
	"util/pwd"

	pb "auth/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	user *biz.UserUsecase
	log  *log.Helper
	auth *conf.Auth
}

func NewAuthService(stu *biz.UserUsecase, logger log.Logger, auth *conf.Auth) *AuthService {
	return &AuthService{
		user: stu,
		log:  log.NewHelper(logger),
		auth: auth,
	}
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	response := &pb.Response{}
	// 检查用户
	user, err := s.user.GetUser(ctx, req.UserName)
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "UserNotFound", "用户名或密码错误")
	}
	// 检查密码
	if !pwd.CheckPwd(user.Pwd, req.PassWord) {
		return nil, errors.New(http.StatusUnauthorized, "PasswordError", "用户名或密码错误")
	}
	// 生成 jwt
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		UserID:   user.ID,
		Nickname: user.Nickname,
		Role:     user.Role,
	}, s.auth.AccessSecret, int(s.auth.AccessExpire))
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "TokenError", "生成 token 失败")
	}
	response.Data = &pb.Response_LoginResponse{
		LoginResponse: &pb.LoginResponse{
			Token: token,
		},
	}
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
