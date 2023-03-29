package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type CustomFormatter struct {
}

func (c *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	fmt.Println("HHHHHHHHH")
	return nil, nil
}
