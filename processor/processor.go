package processor

import (
	"os"
	"runtime"

	"github.com/gookit/slog"
)

// AddHostname to record
func AddHostname() slog.Processor {
	hostname,_ := os.Hostname()

	return slog.ProcessorFunc(func(record *slog.Record) {
		record.AddField("hostname", hostname)
	})
}

// MemoryUsage Get memory usage.
var MemoryUsage slog.ProcessorFunc = func(record *slog.Record) {
	stat := new(runtime.MemStats)
	runtime.ReadMemStats(stat)
	record.Extra["memoryUsage"] = stat.Alloc
}
