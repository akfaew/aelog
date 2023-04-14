package aelog

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/andyfusniak/stackdriver-gae-logrus-plugin"
	"github.com/sirupsen/logrus"
)

var trimprefix = ""

type Log struct {
	*logrus.Entry
}

func Init() {
	formatter := stackdriver.GAEStandardFormatter(
		stackdriver.WithProjectID(os.Getenv("GOOGLE_CLOUD_PROJECT")),
	)
	logrus.SetFormatter(formatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel) // Log the debug severity or above.
}

func New(ctx context.Context) *Log {
	return &Log{
		logrus.WithContext(ctx),
	}
}

// Set trimprefix to the path to the source code directory, so that we only log the filename and not the full path.
func init() {
	_, path, _, _ := runtime.Caller(1)

	trimprefix = path[:strings.LastIndex(path, "/")+1]
}

func logctx(skip int) (file string, line int) {
	_, file, line, _ = runtime.Caller(skip + 1)
	file = strings.TrimPrefix(file, trimprefix)

	return
}

func (log *Log) Debugf(format string, a ...interface{}) {
	log.Entry.Debugf(format, a...)
}

func (log *Log) Debugfd(format string, a ...interface{}) {
	file, line := logctx(1)

	log.Entry.Debugf("%s:%d %s", file, line, fmt.Sprintf(format, a...))
}

func (log *Log) DebugJSON(v interface{}) {
	if b, err := json.MarshalIndent(v, "", "\t"); err != nil {
		log.Entry.Debugf("json.MarshalIndent(): err=%v", err)
	} else {
		log.Entry.Debugf(string(b))
	}
}

func (log *Log) Infof(format string, a ...interface{}) {
	log.Entry.Infof(format, a...)
}

func (log *Log) Infofd(format string, a ...interface{}) {
	file, line := logctx(1)

	log.Entry.Infof("%s:%d %s", file, line, fmt.Sprintf(format, a...))
}

func (log *Log) Warningf(format string, a ...interface{}) {
	log.Entry.Warningf(format, a...)
}

func (log *Log) Warningfd(format string, a ...interface{}) {
	file, line := logctx(1)

	log.Entry.Warningf("%s:%d %s", file, line, fmt.Sprintf(format, a...))
}

func (log *Log) Errorf(format string, a ...interface{}) {
	log.Entry.Errorf(format, a...)
}

func (log *Log) Errorfd(format string, a ...interface{}) {
	file, line := logctx(1)

	log.Entry.Errorf("%s:%d %s", file, line, fmt.Sprintf(format, a...))
}
