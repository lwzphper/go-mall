package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	logger *zap.Logger
)

// 项目启动命令

var serviceCmd = &cobra.Command{
	Use:   "start",
	Short: "go-mail API服务",
	Long:  "go-mail API服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局配置
		if err := loadGlobalConfig(); err != nil {
			return err
		}

		// 初始化全局日志配置
		var err error
		logger, err = loadGlobalLogger()
		if err != nil {
			return err
		}

		// 初始化全部应用
		if err := initAllApp(); err != nil {
			return err
		}

		// 监听中断信号，优雅关闭服务
		ch := make(chan os.Signal, 1)
		defer close(ch)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

		// 初始化服务
		svr, err := newService()
		if err != nil {
			return err
		}

		// 等待信号
		go svr.waitSign(ch)

		// 启动服务
		if err = svr.start(); err != nil {
			if !strings.Contains(err.Error(), "http: Server closed") {
				return err
			}
		}
		return nil
	},
}

type service struct {
	//http *protocol.HTTPService
	//grpc *protocol.GRPCService

	log *zap.Logger
}

// 初始化服务
func newService() (*service, error) {
	return &service{
		log: logger,
	}, nil
}

func (s *service) waitSign(sign chan os.Signal) {
	for sg := range sign {
		switch v := sg.(type) {
		// todo 根据不同的类型，做不同的处理
		default:
			s.log.Info(fmt.Sprintf("receive signal '%v', start graceful shutdown", v.String()))
			/*if err := s.grpc.Stop(); err != nil {
				s.log.Error(fmt.Sprintf("grpc graceful shutdown err: %s, force exit", err))
			} else {
				s.log.Info("grpc service stop complete")
			}

			if err := s.http.Stop(); err != nil {
				s.log.Error(fmt.Sprintf("http graceful shutdown err: %s, force exit", err))
			} else {
				s.log.Info("http service stop complete")
			}*/
			return
		}
	}
}

func (s *service) start() error {
	// 加载各模块的配置
	//s.log.Info("loaded grpc app: %s", app.LoadedGrpcApp())
	//s.log.Info("loaded http app: %s", app.LoadedGinApp())

	// 启动各模块的服务
	/*go s.grpc.Start()
	return s.http.Start()*/

	return nil
}

// 初始化全局日志配置
func loadGlobalLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	return cfg.Build()
}

// 初始化全部应用
func initAllApp() error {
	/*for _, api := range internalApps {
		if err := api.Config(); err != nil {
			return err
		}
	}

	for _, api := range grpcApps {
		if err := api.Config(); err != nil {
			return err
		}
	}

	for _, api := range restfulApps {
		if err := api.Config(); err != nil {
			return err
		}
	}

	for _, api := range httpApps {
		if err := api.Config(); err != nil {
			return err
		}
	}

	for _, api := range ginApps {
		if err := api.Config(); err != nil {
			return err
		}
	}*/
	return nil
}

func init() {
	RootCmd.AddCommand(serviceCmd)
}
