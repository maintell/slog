package main

import (
	"fmt"
	"github.com/maintell/slog"
	"os"
	"path/filepath"
)

// profile run:
//
// go build -gcflags '-m -l' simple.go
func main() {
	// stackIt()
	// _ = stackIt2()
	slogTest()
}

//go:noinline
func stackIt() int {
	y := 2
	return y * 2
}

//go:noinline
func stackIt2() *int {
	y := 2
	res := y * 2
	return &res
}

func slogTest() {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	LogPath := filepath.Join(exPath, "log")
	os.MkdirAll(LogPath, os.ModePerm)
	//h2 := handler.MustRotateFile("log/test.log", handler.EveryDay, handler.WithBuffSize(1024))
	//slog.PushHandler(h2)

	var msg = "The quick brown fox jumps over the lazy dog"

	slog.Info("rate", "15", "low", 16, "high", 123.2, msg)
	slog.Info("rate", "15", "low", 16, "high", 123.2, msg)
	slog.Info("rate", "15", "low", 16, "high", 123.2, msg)
	slog.Info("rate", "15", "low", 16, "high", 123.2, msg)
	slog.Info("rate", "15", "low", 16, "high", 123.2, msg)
	// slog.WithFields(slog.M{
	// 	"omg":    true,
	// 	"number": 122,
	// }).Infof("slog %s", "message message")
}
