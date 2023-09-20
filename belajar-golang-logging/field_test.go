package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFiled(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Vini").Info("ini pake field")
	logger.WithField("username", "Vini").WithField("nama", "nama").Info("ini pake field")

	logger.Info("HELLO WORLD")
}

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "vini",
		"name":     "fitri",
	}).Info("Hello World")
}
