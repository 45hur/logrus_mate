package logrus_mate

import (
	"github.com/45hur/logrus"
	"github.com/gogap/config"
)

func init() {
	RegisterFormatter("text", NewTextFormatter)
}

func NewTextFormatter(config config.Configuration) (formatter logrus.Formatter, err error) {

	f := &logrus.TextFormatter{}

	if config != nil {
		f.ForceColors = config.GetBoolean("force-colors")
		f.DisableColors = config.GetBoolean("disable-colors")
		f.DisableTimestamp = config.GetBoolean("disable-timestamp")
		f.FullTimestamp = config.GetBoolean("full-timestamp")
		f.TimestampFormat = config.GetString("timestamp-format")
		f.DisableSorting = config.GetBoolean("disable-sorting")
	}

	formatter = f

	return
}
