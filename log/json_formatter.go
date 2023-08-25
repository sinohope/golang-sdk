package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"sort"
	"time"

	"github.com/sirupsen/logrus"
)

/*
our json formatter
*/

// fieldMap allows customization of the key names for default fields.
type fieldMap map[string]string

func (f fieldMap) resolve(key string) string {
	if k, ok := f[key]; ok {
		return k
	}
	return key
}

// jsonFormatter formats logs into parsable json
type jsonFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	// The format to use is the same than for time.Format or time.Parse from the standard
	// library.
	// The standard Library already provides a set of predefined format.
	TimestampFormat string

	// DisableTimestamp allows disabling automatic timestamps in output
	DisableTimestamp bool

	// DisableHTMLEscape allows disabling html escaping in output
	DisableHTMLEscape bool

	// DataKey allows users to put all the log entry parameters into a nested dictionary at a given key.
	DataKey string

	// FieldMap allows users to customize the names of keys for default fields.
	// As an example:
	// formatter := &JSONFormatter{
	//   	FieldMap: FieldMap{
	// 		 FieldKeyTime:  "@timestamp",
	// 		 FieldKeyLevel: "@level",
	// 		 FieldKeyMsg:   "@message",
	// 		 FieldKeyFunc:  "@caller",
	//    },
	// }
	FieldMap fieldMap

	// CallerPrettyfier can be set by the user to modify the content
	// of the function and file keys in the json data when ReportCaller is
	// activated. If any of the returned value is the empty string the
	// corresponding key will be removed from json fields.
	CallerPrettyfier func(*runtime.Frame) (function string, file string)

	// PrettyPrint will indent all json logs
	PrettyPrint bool
}

var skip = map[string]struct{}{
	"time":    struct{}{},
	"level":   struct{}{},
	"file":    struct{}{},
	"message": struct{}{},
}

// Format renders a single log entry
func (f *jsonFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	if f.DataKey != "" {
		newData := make(logrus.Fields, 4)
		newData[f.DataKey] = data
		data = newData
	}
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339
	}
	if !f.DisableTimestamp {
		data[f.FieldMap.resolve(logrus.FieldKeyTime)] = entry.Time.Format(timestampFormat)
	}
	data[f.FieldMap.resolve(logrus.FieldKeyMsg)] = entry.Message
	level := entry.Level.String()
	if level == "warning" {
		level = "warn"
	}
	data[f.FieldMap.resolve(logrus.FieldKeyLevel)] = level
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		if f.CallerPrettyfier != nil {
			funcVal, fileVal = f.CallerPrettyfier(entry.Caller)
		}
		if funcVal != "" {
			data[f.FieldMap.resolve(logrus.FieldKeyFunc)] = funcVal
		}
		if fileVal != "" {
			data[f.FieldMap.resolve(logrus.FieldKeyFile)] = fileVal
		}
	}
	buffer := bytes.NewBufferString("{")
	if timeValue, ok := data["time"]; ok {
		buffer.WriteString(fmt.Sprintf(`"time":%q,`, timeValue))
	}
	if levelValue, ok := data["level"]; ok {
		buffer.WriteString(fmt.Sprintf(`"level":%q,`, levelValue))
	}
	for _, v := range keys {
		if _, ok := skip[v]; !ok {
			v2 := data[v]
			valueJSON, err := json.Marshal(v2)
			if err != nil {
				return nil, err
			}
			buffer.WriteString(fmt.Sprintf(`%q:%s,`, v, valueJSON))
		}
	}
	if messageValue, ok := data["message"]; ok {
		buffer.WriteString(fmt.Sprintf(`"message":%q`, messageValue))
	}
	buffer.WriteString("}")
	buffer.WriteByte('\n')
	return buffer.Bytes(), nil
}
