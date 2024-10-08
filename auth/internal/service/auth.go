package service

import (
	"auth/internal/biz"
	"auth/internal/conf"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
	"strings"
	"time"
	v1 "user/api/helloworld/v1"
	"util/jwt"
	"util/pwd"

	pb "auth/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	user *biz.AuthUsecase
	log  *log.Helper

	greeterClient v1.GreeterClient

	auth          *conf.Auth
	openLoginInfo *conf.OpenLoginList
}

func NewAuthService(stu *biz.AuthUsecase, logger log.Logger, auth *conf.Auth, openLoginInfo *conf.OpenLoginList, etcdClient registry.Discovery) *AuthService {
	// 服务发现
	connGRPC, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user"),
		grpc.WithDiscovery(etcdClient),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &AuthService{
		user:          stu,
		log:           log.NewHelper(logger),
		auth:          auth,
		openLoginInfo: openLoginInfo,
		greeterClient: v1.NewGreeterClient(connGRPC),
	}
}

// Login 登录
func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {

	s.log.Info("[HTTP] --> login")
	response := &pb.Response{}
	// 检查用户
	user, err := s.user.GetUser(ctx, req.UserName)
	if err != nil {
		//return nil, errors.New(http.StatusUnauthorized, "UserNotFound", "用户名或密码错误")
		return nil, pb.ErrorUserNotFound("用户名或密码错误")
	}
	// 检查密码
	if !pwd.CheckPwd(user.Pwd, req.PassWord) {
		return nil, pb.ErrorPasswordError("用户名或密码错误")
	}
	log.Debug("Expire:", s.auth.AccessExpire)
	// 生成 jwt
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		UserID:   user.ID,
		Nickname: user.Nickname,
		Role:     user.Role,
	}, s.auth.AccessSecret, int(s.auth.AccessExpire))
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "TokenError", "生成 token 失败")
	}
	response.Msg = "登录成功"
	response.LoginResponse = &pb.Response_LoginResponse{
		Token: token,
	}
	return response, nil
}
func (s *AuthService) Authentication(ctx context.Context, empty *emptypb.Empty) (*pb.Response, error) {
	s.log.Infof("[HTTP] --> authentication")
	// 从请求头中拿到token
	if tr, ok := transport.FromServerContext(ctx); ok {
		token := tr.RequestHeader().Get("Authorization")
		if token == "" {
			return nil, pb.ErrorTokenEmpty("认证失败")
		}
		payload, err := jwt.ParseToken(token, s.auth.AccessSecret)
		if err != nil {
			return nil, pb.ErrorParseTokenError("认证失败")
		}
		var build strings.Builder
		build.WriteString("auth:logout:")
		build.WriteString(strconv.Itoa(int(payload.UserID)))
		getToken, err := s.user.GetToken(ctx, build.String())
		// 说明Redis中有Token, 说明用户已退出,所以认证失败
		if err == nil {
			return nil, pb.ErrorAuthenticationFailed("认证失败")
		}
		fmt.Print(getToken)
	}
	return &pb.Response{
		Code:              200,
		LoginResponse:     nil,
		OpenLoginResponse: nil,
		Msg:               "认证成功",
	}, nil
}
func (s *AuthService) Logout(ctx context.Context, empty *emptypb.Empty) (*pb.Response, error) {
	s.log.Info("[HTTP] --> logout")
	// 从请求头中拿到token
	if tr, ok := transport.FromServerContext(ctx); ok {
		token := tr.RequestHeader().Get("Authorization")
		if token == "" {
			return nil, pb.ErrorTokenEmpty("token 为空")
		}
		payload, err := jwt.ParseToken(token, s.auth.AccessSecret)
		if err != nil {
			return nil, pb.ErrorParseTokenError("token 解析失败")
		}
		// 过期时间
		expiration := payload.ExpiresAt.Time.Sub(time.Now())
		var key strings.Builder
		key.WriteString("auth:logout:")
		key.WriteString(strconv.Itoa(int(payload.UserID)))

		if err = s.user.SetToken(ctx, key.String(), "", expiration); err != nil {
			return nil, pb.ErrorSetTokenError("设置 token 失败")
		}
	}
	return &pb.Response{
		Code:              200,
		LoginResponse:     nil,
		OpenLoginResponse: nil,
		Msg:               "注销成功",
	}, nil
}
func (s *AuthService) OpenLoginInfo(ctx context.Context, empty *emptypb.Empty) (*pb.Response, error) {
	//hello, _ := s.greeterClient.SayHello(ctx, &v1.HelloRequest{
	//	Name: "openLogin",
	//})
	//fmt.Println(hello)
	s.log.Info("[HTTP] --> openLoginInfo")
	response := &pb.Response{
		OpenLoginResponse: &pb.Response_OpenLoginResponses{
			OpenLoginInfo: make([]*pb.OpenLoginResponse, 0, 2),
		},
	}
	for _, item := range s.openLoginInfo.Items {
		response.OpenLoginResponse.OpenLoginInfo = append(response.OpenLoginResponse.OpenLoginInfo, &pb.OpenLoginResponse{
			Name: item.Name,
			Icon: item.Icon,
			Href: item.Href,
		})
	}
	response.Code = 200
	response.Msg = "获取第三方登录信息成功"
	return response, nil
}
func (s *AuthService) OpenLogin(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	fmt.Print("openLogin")
	return &pb.Response{}, nil
}
