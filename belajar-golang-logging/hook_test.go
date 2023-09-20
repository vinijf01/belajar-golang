package belajargolanglogging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

/*
HOOK adalah struct yang bisa ditambahakan ke dalam logger sebagai callback
yang akan dieksekusi ketika log untuk level tertentu

ex case : ketika ada log error, kita ingin mengirim notifikasi via chat ke programmar
*/

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook", entry.Level, entry.Message)
	return nil

}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})
	logger.SetLevel(logrus.TraceLevel)

	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
	logger.Debug("This is Debug")
}
