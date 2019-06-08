package logging

import (
	"github.com/evalphobia/go-log-wrapper/log"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
)


var once sync.Once
var logger = new(log.Logger)
func GetInstance() *log.Logger {
	//
	once.Do(func() {
		logger = log.NewLogger()
		//logger.SetLevel(logrus.DebugLevel)
		// log.Formatter.(*logrus.TextFormatter).DisableColors = true // remove colors
		logger.Level = logrus.WarnLevel
		logger.Out = os.Stdout

		logger.SetOutput(&lumberjack.Logger{
			Filename:   "audiomemory.log",
			MaxSize:    50, // megabytes
			MaxBackups: 30,
			MaxAge:     60, // days
			// Compress:   true, // disabled by default
		})
	})
	return logger
}
