package initialize

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lwzphper/go-mall/bff/global"
	"github.com/lwzphper/go-mall/pkg/common/config/app"
	"github.com/lwzphper/go-mall/pkg/server"
	"log"
	"net/http"
	"os"
	"time"
)

func InitGin() {
	gin.SetMode(getGinModeByEnv())
	router := Routers()
	srv := &http.Server{
		Addr:    global.C.App.Addr,
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.L.Fatalf("Server listen error: %s\n", err)
		}
	}()

	hook := server.NewHook()
	hook.Close(func(sg os.Signal) {
		log.Println("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			global.L.Fatalf("Server Shutdown error:%v", err)
		}
		log.Println("Server exiting")
	})
}

func getGinModeByEnv() string {
	var ginMode string
	switch global.C.App.Env {
	case app.EnvDevelopment:
		ginMode = gin.DebugMode
	case app.EnvTest:
		ginMode = gin.TestMode
	default:
		ginMode = gin.ReleaseMode
	}
	return ginMode
}
