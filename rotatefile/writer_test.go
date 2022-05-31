package rotatefile_test

import (
	"github.com/maintell/slog"
	"github.com/maintell/slog/handler"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/fsutil"
	"github.com/maintell/slog/rotatefile"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cfg := rotatefile.NewConfig("testdata/test.log")

	assert.Equal(t, rotatefile.DefaultBackNum, cfg.BackupNum)
	assert.Equal(t, rotatefile.DefaultBackTime, cfg.BackupTime)
	assert.Equal(t, rotatefile.EveryHour, cfg.RotateTime)
	assert.Equal(t, rotatefile.DefaultMaxSize, cfg.MaxSize)

	dump.P(cfg)
}

func TestNewWriter(t *testing.T) {
	testFile := "testdata/test.log"
	assert.NoError(t, fsutil.DeleteIfExist(testFile))

	h2 := handler.MustRotateFile("testdata/test.log", handler.EveryDay, handler.WithBuffSize(1024))
	slog.PushHandler(h2)

	slog.Print("fsfasdfsdafsafsdafasdf")
	slog.Print("fsfasdfsdafsafsdafasdf")
	slog.Print("fsfasdfsdafsafsdafasdf")
	slog.Print("fsfasdfsdafsafsdafasdf")


}
