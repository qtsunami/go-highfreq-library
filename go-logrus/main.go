package main

import (
	"github.com/go-highfreq-library/go-logrus/util"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logger := util.NewLogger()

	logger.SetOutput(os.Stdout)
	logger.SetLevel(util.DebugLevel)

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//logger.SetFormatter(&util.CustomFormatter{})

	logger.Warnln("Hello")
	//logger.WithFields(map[string]interface{}{
	//	"Name": "rr",
	//}).Info("HHHH")
	//
	//httpLog := logger.LogHandler(map[string]interface{}{
	//	"name": "Lee",
	//	"age":  29,
	//})
	//
	//httpLog.Warn("request failed!")

}

func logtofile() {
	//logFile := "log.txt"
	//f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	//if err != nil {
	//	fmt.Println("Failed to create logfile" + logFile)
	//	panic(err)
	//}
	//defer f.Close()
	//
	//log := &logrus.Logger{
	//	Out:   io.MultiWriter(f, os.Stdout),
	//	Level: logrus.DebugLevel,
	//	Formatter: &logrus.TextFormatter{
	//		FullTimestamp:          true,
	//		TimestampFormat:        "2006-01-02 15:04:05",
	//		ForceColors:            true,
	//		DisableLevelTruncation: true,
	//		ForceQuote:             true,
	//	},
	//}
	//
	//log.SetReportCaller(true)
	//
	//log.WithFields(logrus.Fields{
	//	"request_id": "887B4F43-D330-5213-94DF-1B14FA3388C",
	//	"ip":         "101.202.112.12",
	//	"user_id":    10022,
	//}).Info("Info Level")
	//log.Warn("Warn Level")
	//log.Error("Error Level")
	//log.Fatal("Fatal Level")
}
