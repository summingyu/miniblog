package miniblog

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/summingyu/miniblog/internal/pkg/log"
	mw "github.com/summingyu/miniblog/internal/pkg/middleware"
	"github.com/summingyu/miniblog/pkg/version/verflag"
)

var cfgFile string

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "miniblog",
		// 命令的简短描述
		Short: "A good Go practical project",
		// 命令的详细描述
		Long: `A good Go practical project, used to create user with basic information.

Find more miniblog information at:
https://github.com/marmotedu/miniblog#readme`,

		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			// 初始化日志
			log.Init(logOptions())
			// Sync 将缓冲区中的日志数据写入到文件中
			defer log.Sync()
			// 如果 `--version=true`，则打印版本并退出
			verflag.PrintAndExitIfRequested()
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q dose not take any arguments, go %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)
	// 在这里您将定义标志和配置设置。

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.miniblog.yaml)")
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// 添加 --version 标志
	verflag.AddFlags(cmd.PersistentFlags())
	return cmd
}

// run 函数是实际的业务代码入口函数.
func run() error {
	// 初始化 store 层
	if err := initStore(); err != nil {
		return err
	}
	// 设置 Gin 运行模式
	gin.SetMode(viper.GetString("runmode"))

	// 创建 Gin 引擎
	g := gin.New()

	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.RequestID()}

	g.Use(mws...)
	if err := installRouters(g); err != nil {
		return err
	}
	// 创建HTTP Server 实例
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}
	// 启动HTTP Server
	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw("Failed to start http server", "error", err.Error())
		}
	}()

	// 等待信号退出
	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infow("Shutting down server...")

	// 创建 ctx 用于通知服务器goroutine, 它有10s 完成当前处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 10s 内关闭服务器
	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Fatalw("Server forced to shutdown", "error", err)
		return err
	}

	log.Infow("Server exiting")
	return nil
}
