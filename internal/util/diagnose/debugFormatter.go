package diagnose

import (
	"bytes"
	"fmt"

	"github.com/sirupsen/logrus"
)

type DebugFormatter struct {
}

func (m *DebugFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("15:04:05")
	newLog := fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}
