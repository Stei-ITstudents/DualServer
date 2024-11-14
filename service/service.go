package service

import "github.com/sirupsen/logrus"

func ProcessRequest(data string) string {
	logrus.Infof("Processing data: %s", data)

	return "Processed: " + data
}
