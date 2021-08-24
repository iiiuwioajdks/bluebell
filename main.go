package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/routes"
	"bluebell/settings"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// go web 较通用的脚手架模板

func main() {
	// 加载配置文件
	err := settings.Init()
	if err != nil {
		fmt.Println("init settings failed err:", err)
		return
	}

	// 初始化日志
	err = logger.Init(viper.GetString("app.mode"))
	if err != nil {
		fmt.Println("init logger failed err:", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	// 初始化 mysql
	err = mysql.Init()
	if err != nil {
		fmt.Println("init mysql failed err:", err)
		return
	}
	defer mysql.Close()
	//雪花算法初始化
	if err = snowflake.Init(viper.GetString("app.startTime"), viper.GetInt64("app.machineId")); err != nil {
		fmt.Println("init snowflake failed err:", err)
		return
	}

	// 初始化 redis
	err = redis.Init()
	if err != nil {
		fmt.Println("init redis failed err:", err)
		return
	}
	defer redis.Close()

	// 注册路由
	router := routes.SetUpRouter()

	// 启动服务(优雅关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: router,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
