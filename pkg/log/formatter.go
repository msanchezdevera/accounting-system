package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Formatter struct {
	host string
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {

	msg := fmt.Sprintf("%s %s [%s]: %s",
		entry.Time.Format("2006-01-02T15:04:05.999"),
		entry.Level,
		f.host,
		entry.Message)

	msg = msg + "\n"

	return []byte(msg), nil
}
