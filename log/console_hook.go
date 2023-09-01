package log

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// AddConsoleOut ...
func AddConsoleOut(level int) {
	DisableDefaultConsole()
	logrus.AddHook(newConsoleHook(level))
}

type consoleHook struct {
	formatter logrus.Formatter
	levels    []logrus.Level
}

// Fire event
func (c *consoleHook) Fire(entry *logrus.Entry) error {
	for key, value := range defaultFieldMap {
		if _, ok := entry.Data[key]; !ok {
			entry.Data[key] = value
		}
	}
	formatBytes, err := c.formatter.Format(entry)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to fortmat the log line on consoleHook %s", err)
		return err
	}
	fmt.Print(string(formatBytes))
	return nil
}

func (c *consoleHook) Levels() []logrus.Level {
	return c.levels
}

func newConsoleHook(level int) *consoleHook {
	plainFormatter := &jsonFormatter{
		TimestampFormat:  time.RFC3339Nano,
		DisableTimestamp: false,
		CallerPrettyfier: callerPrettyfier,
		FieldMap: fieldMap{
			logrus.FieldKeyMsg: "message",
		},
	}
	return &consoleHook{plainFormatter, getHookLevel(level)}
}
