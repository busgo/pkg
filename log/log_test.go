package log

import (
	"testing"
	"time"
)

func TestDebug(t *testing.T) {

	Debugf("ok %s", "debug")

	time.Sleep(time.Second * 10)

	_ = NewLoggerSugar("gateway", "gateway.log", "debug")
	Debugf("ok %s", "debugf")
	Infof("ok %s", "infof")
	Warnf("ok %s", "warnf")
	Errorf("ok %s", "errorf")

}
