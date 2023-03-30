package util

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	gray = 32
)

type CustomFormatter struct {
}

func (c *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	// entry.Data 获取 withFields的参数
	data := make(logrus.Fields)

	for k, v := range entry.Data {
		data[k] = v
	}

	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	//fmt.Println(entry.Data)
	//fmt.Println(entry.Message)
	//fmt.Println(entry.Buffer)

	fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s]\n", 30, entry.Level.String(), entry.Time.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s]\n", 31, entry.Level.String(), entry.Time.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s]\n", 32, entry.Level.String(), entry.Time.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s]\n", 33, entry.Level.String(), entry.Time.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s]\n", 34, entry.Level.String(), entry.Time.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s]\n", 35, entry.Level.String(), entry.Time.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s]\n", 36, entry.Level.String(), entry.Time.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s]\n", 37, entry.Level.String(), entry.Time.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(b, "message=%s", entry.Message)

	b.WriteString("Meeting")
	b.WriteByte('=')
	b.WriteString("UP")

	b.WriteByte('\n')
	return b.Bytes(), nil
}
