package logrus_mate

import (
	"github.com/45hur/logrus"
	"github.com/gogap/config"
)

type NullFormatter struct {
}

func init() {
	RegisterFormatter("nil", NewNullFormatter)
}

func NewNullFormatter(config config.Configuration) (formatter logrus.Formatter, err error) {
	formatter = &NullFormatter{}
	return
}

func (NullFormatter) Format(e *logrus.Entry) ([]byte, error) {
	return []byte{}, nil
}
