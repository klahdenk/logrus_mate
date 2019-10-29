package logrus_mate

import (
	"github.com/gogap/config"
	"github.com/sirupsen/logrus"
)

const (
	FieldRefTime = "time"
)

type JSONFormatterConfig struct {
	TimestampFormat string `json:"timestamp_format"`
}

func init() {
	RegisterFormatter("json", NewJSONFormatter)
}

func NewJSONFormatter(config config.Configuration) (logrus.Formatter, error) {
	var format string
	var fieldMap logrus.FieldMap

	configuration := config.GetConfig("field-overrides")

	if config != nil {
		format = config.GetString("timestamp_format")
		if len(configuration.GetString(FieldRefTime)) > 0 {
			fieldMap[logrus.FieldKeyTime] = configuration.GetString(FieldRefTime)
		}
	}
	formatter := &logrus.JSONFormatter{TimestampFormat: format}
	if len(fieldMap) > 0 {
		formatter.FieldMap = fieldMap
	}


	return formatter, nil
}
