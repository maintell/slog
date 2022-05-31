package slog_test

import (
	"testing"
	"time"

	"github.com/maintell/slog"
	"github.com/maintell/slog/handler"
)

// https://github.com/maintell/slog/issues/27
func TestIssues_27(t *testing.T) {
	defer slog.Reset()

	count := 0
	for {
		if count >= 6 {
			break
		}
		slog.Infof("info log %d", count)
		time.Sleep(time.Second)
		count++
	}
}

// https://github.com/maintell/slog/issues/31
func TestIssues_31(t *testing.T) {
	defer slog.Reset()
	defer slog.MustFlush()

	// slog.DangerLevels equals slog.Levels{slog.PanicLevel, slog.PanicLevel, slog.ErrorLevel, slog.WarnLevel}
	h1 := handler.MustFileHandler("testdata/error_issue31.log", handler.WithLogLevels(slog.DangerLevels))

	infoLevels := slog.Levels{slog.InfoLevel, slog.NoticeLevel, slog.DebugLevel, slog.TraceLevel}
	h2 := handler.MustFileHandler("testdata/info_issue31.log", handler.WithLogLevels(infoLevels))

	slog.PushHandler(h1)
	slog.PushHandler(h2)

	// add logs
	slog.Info("info message text")
	slog.Error("error message text")
}

// https://github.com/maintell/slog/issues/52
func TestIssues_52(t *testing.T) {
	testTemplate := "[{{datetime}}] [{{level}}] {{message}} {{data}} {{extra}}"
	slog.SetLogLevel(slog.ErrorLevel)
	slog.GetFormatter().(*slog.TextFormatter).SetTemplate(testTemplate)

	slog.Error("Error message")
	slog.Reset()

	// dump.P(slog.GetFormatter())
}
