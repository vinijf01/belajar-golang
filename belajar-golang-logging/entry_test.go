package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

//ini mempelajari cara membuat entry baru
func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithField("usrrname", "vini")
	entry.Info("Helo Entry")
}
