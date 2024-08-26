package server

import (
	v1 "auth/api/auth/v1"
	"auth/internal/conf"
	"auth/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
//func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
//	var opts = []http.ServerOption{
//		http.Middleware(
//			recovery.Recovery(),
//		),
//	}
//	if c.Http.Network != "" {
//		opts = append(opts, http.Network(c.Http.Network))
//	}
//	if c.Http.Addr != "" {
//		opts = append(opts, http.Address(c.Http.Addr))
//	}
//	if c.Http.Timeout != nil {
//		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
//	}
//	srv := http.NewServer(opts...)
//	v1.RegisterGreeterHTTPServer(srv, greeter)
//	return srv
//}

func NewHTTPServer(c *conf.Server, auth *service.AuthService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterAuthHTTPServer(srv, auth)
	return srv
}

//func customResponseEncoder(w http.ResponseWriter, r *http.Request, data interface{}) error {
//	// 创建一个 map 来存储响应数据
//	response := make(map[string]interface{})
//
//	// 设置默认值
//	response["code"] = 0    // 默认状态码
//	response["msg"] = "OK"  // 默认消息
//	response["data"] = data // 默认数据
//
//	// 如果需要，你可以根据 data 的类型或者具体的业务逻辑来更改 code 和 msg 的值
//	// 例如:
//	// if data == nil {
//	//     response["code"] = 1
//	//     response["msg"] = "No data found"
//	// }
//
//	// 将 response 序列化为 JSON
//	jsonData, err := json.Marshal(response)
//	if err != nil {
//		return err
//	}
//
//	// 设置 Content-Type 标头
//	w.Header().Set("Content-Type", "application/json")
//
//	// 写入 HTTP 响应
//	w.Write(jsonData)
//
//	return nil
//}
