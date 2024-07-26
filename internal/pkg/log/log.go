// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

package log

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/summingyu/miniblog/internal/pkg/known"
)

// Logger 定义了 miniblog 项目的日志接口. 该接口只包含了支持的日志记录方法.
type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync()
}

// zapLogger 是 Logger 接口的具体实现. 它底层封装了 zap.Logger.
type zapLogger struct {
	z *zap.Logger
}

var _ Logger = &zapLogger{}

var (
	mu sync.Mutex
	// std 定义了默认的全局Logger.
	std = NewLogger(NewOptions())
)

// Init 使用指定的选项初始化 Logger.
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()
	std = NewLogger(opts)
}

// NewLogger 根据传入的opts创建Logger.
func NewLogger(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}
	// 将文本格式的日志级别, 例如 info 转换成为 zap.Level 类型.
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	// 创建一个默认的encoder配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 自定义 MessageKey 为 message, message 语意更明确.
	encoderConfig.MessageKey = "message"
	// 自定义 TimeKey 为 timestamp, timestamp 语意更明确.
	encoderConfig.TimeKey = "timestamp"
	// 指定时间序列化函数，将时间序列化为 `2006-01-02 15:04:05.000` 格式，更易读
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	// 指定 time.Duration 序列化函数，将 time.Duration 序列化为经过的毫秒数的浮点数
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	// 创建构建zap.Logger的配置
	cfg := &zap.Config{
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Encoding:          opts.Format,
		EncoderConfig:     encoderConfig,
		OutputPaths:       opts.OutputPaths,
		ErrorOutputPaths:  []string{"stderr"},
	}
	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger := &zapLogger{z: z}

	// 把标准库的 log.Logger 替换成 zap.Logger
	zap.RedirectStdLog(z)

	return logger
}

// Sync implements Logger.
func Sync() { std.Sync() }

func (l *zapLogger) Sync() {
	_ = l.z.Sync()
}

// Debugw implements Logger.
func Debugw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Debugw(msg, keysAndValues...)
}

// Infow implements Logger.
func Infow(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Infow(msg, keysAndValues...)
}

func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Infow(msg, keysAndValues...)
}

// Warnw implements Logger.
func Warnw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Warnw(msg, keysAndValues...)
}

// Errorw implements Logger.
func Errorw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Errorw(msg, keysAndValues...)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Errorw(msg, keysAndValues...)
}

// Panicw implements Logger.
func Panicw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Panicw(msg, keysAndValues...)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Panicw(msg, keysAndValues...)
}

// Fatalw implements Logger.
func Fatalw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Fatalw(msg, keysAndValues...)
}

// C 解析传入的context, 尝试提取关注的键值, 并添加到zap.Logger的结构化日志中.
func C(ctx context.Context) *zapLogger {
	return std.C(ctx)
}

// C 返回一个带有当前请求ID的zapLogger副本
//
// 如果给定的ctx中包含了已知的请求ID键（known.XRequestIDKey），
// 则会在新的zapLogger中添加该请求ID作为日志字段
//
// 参数：
// ctx：包含请求ID的上下文
//
// 返回值：
// *zapLogger：带有当前请求ID的zapLogger副本
func (l *zapLogger) C(ctx context.Context) *zapLogger {
	// 复制当前 zapLogger 实例
	lc := l.clone()

	// 从 ctx 中获取 requestID
	if requestID := ctx.Value(known.XRequestIDKey); requestID != nil {
		// 如果 requestID 存在，则在日志中添加 request_id 字段
		lc.z = lc.z.With(zap.String("request_id", requestID.(string)))
	}

	if userID := ctx.Value(known.XUsernameKey); userID != nil {
		lc.z = lc.z.With(zap.Any(known.XUsernameKey, userID))
	}
	// 返回克隆后的 zapLogger 实例
	return lc
}

// clone 是zapLogger类型的方法，用于克隆当前的zapLogger实例
// 返回一个指向新克隆的zapLogger实例的指针
func (l *zapLogger) clone() *zapLogger {
	// 创建一个新的zapLogger结构体lc，并将l的内容复制到lc中
	lc := *l
	// 返回lc的地址
	return &lc
}
