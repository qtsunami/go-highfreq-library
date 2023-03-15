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
		Out:       io.MultiWriter(f, os.Stdout),
		Level:     logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{},
	}

	log.SetReportCaller(true)

	log.Info("Info Level")
	log.Warn("Warn Level")
	log.Error("Error Level")
	log.Fatal("Fatal Level")
}
