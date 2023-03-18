package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(os.Stderr)
	logrus.Trace("Trace Level")
	logrus.Debug("Debug Level")
	logrus.Info("Info Level")
	logrus.Warn("Warn Level")
	logrus.Error("Error Level")
	logrus.Fatal("Fatal Level")
	logrus.Panic("Panic Level")
}
