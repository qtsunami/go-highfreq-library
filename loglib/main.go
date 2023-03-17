package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	logFile := "log.txt"
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + logFile)
		panic(err)
	}
	defer f.Close()

	log := &logrus.Logger{
		Out:   io.MultiWriter(f, os.Stdout),
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05",
			ForceColors:            true,
			DisableLevelTruncation: true,
			ForceQuote:             true,
		},
	}

	log.SetReportCaller(true)

	log.WithFields(logrus.Fields{
		"Lang":    "Zh-cn",
		"Country": "Chinese",
	}).Info("Info Level")
	log.Warn("Warn Level")
	log.Error("Error Level")
	log.Fatal("Fatal Level")
}
