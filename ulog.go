package ulog

import (
	"io"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config 可用在配置文件中
type Config struct {
	Filename   string            // 日志文件
	MaxSize    int               // megabytes
	MaxBackups int               // MaxBackups
	MaxAge     int               // days
	Fields     map[string]string // slog的初始化字段(session)
	IsJSON     bool              // 默认是非json格式
}

// NewLog 日常使用的log, 建议使用json效率更高
func NewLog(c Config) *zerolog.Logger {
	// init zerolog format
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000"
	zerolog.DurationFieldInteger = true
	zerolog.TimestampFieldName = "timestamp"
	zerolog.DurationFieldUnit = time.Millisecond

	out := &lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    c.MaxSize,    // megabytes
		MaxBackups: c.MaxBackups, // MaxBackups
		MaxAge:     c.MaxAge,     // days
		LocalTime:  true,         // 这个需要设置, 不然日志文件的名字就是UTC时间
	}

	var zout io.Writer
	if c.IsJSON {
		zout = out
	} else {
		w := zerolog.NewConsoleWriter()
		w.TimeFormat = "2006-01-02T15:04:05.000"
		w.NoColor = true
		w.FormatLevel = formatLevel
		w.Out = out
		zout = w
	}

	zc := zerolog.New(zout).With().Timestamp()
	for k, v := range c.Fields {
		zc = zc.Str(k, v)
	}
	slog := zc.Logger()
	return &slog
}

// NewLogConsole 创建终端日志格式
func NewLogConsole() *zerolog.Logger {
	w := zerolog.NewConsoleWriter()
	w.TimeFormat = "01-02T15:04:05"
	w.NoColor = true
	w.FormatLevel = formatLevel
	log := zerolog.New(w).With().Timestamp().Logger()
	return &log
}

// formatLevel 不用默认的, 用比较直观的
// 参考: https://github.com/rs/zerolog/blob/master/console.go#L315
func formatLevel(i interface{}) string {
	if ll, ok := i.(string); ok {
		return ll
	}
	return "?????" // 异常时候用的
}
