package main

import (
	"auth/internal/conf"
	"flag"
	z "github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampMilli}
	// 日志输出时间格式
	zerolog.TimeFieldFormat = time.StampMilli
	zerologger := zerolog.New(output).With().Timestamp().CallerWithSkipFrameCount(4).
		Logger()

	// 日志输出文件行号
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		var builder strings.Builder
		builder.WriteString(filepath.Base(file))
		builder.WriteString(":")
		builder.WriteString(strconv.Itoa(line))
		return builder.String()
		//return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	logger := z.NewLogger(&zerologger)

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	// 注入 auth 配置
	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Auth, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
