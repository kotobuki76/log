package log

import "testing"

func TestLevelDebug(t *testing.T) {

	a := "string"
	i := 1
	SetOutputLevel(LOG_LEVEL_DEBUG)
	Debugf("%v,%v", a, i)
	Infof("%v,%v", a, i)
	Warningf("%v,%v", a, i)
	Errorf("%v,%v", a, i)
	Criticalf("%v,%v", a, i)
}

func TestLevelInfo(t *testing.T) {

	a := "string"
	i := 1
	SetOutputLevel(LOG_LEVEL_INFO)
	Debugf("%v,%v", a, i)
	Infof("%v,%v", a, i)
	Warningf("%v,%v", a, i)
	Errorf("%v,%v", a, i)
	Criticalf("%v,%v", a, i)
}

func TestLevelWarning(t *testing.T) {

	a := "string"
	i := 1
	SetOutputLevel(LOG_LEVEL_WARNING)
	Debugf("%v,%v", a, i)
	Infof("%v,%v", a, i)
	Warningf("%v,%v", a, i)
	Errorf("%v,%v", a, i)
	Criticalf("%v,%v", a, i)
}

func TestLevelError(t *testing.T) {

	a := "string"
	i := 1
	SetOutputLevel(LOG_LEVEL_ERROR)
	Debugf("%v,%v", a, i)
	Infof("%v,%v", a, i)
	Warningf("%v,%v", a, i)
	Errorf("%v,%v", a, i)
	Criticalf("%v,%v", a, i)
}

func TestLevelCritical(t *testing.T) {

	a := "string"
	i := 1
	SetOutputLevel(LOG_LEVEL_CRITICAL)
	Debugf("%v,%v", a, i)
	Infof("%v,%v", a, i)
	Warningf("%v,%v", a, i)
	Errorf("%v,%v", a, i)
	Criticalf("%v,%v", a, i)
}
