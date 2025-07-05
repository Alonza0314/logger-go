package model

type LoggerInterface interface {
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Testf(format string, args ...interface{})

	Errorln(args ...interface{})
	Warnln(args ...interface{})
	Infoln(args ...interface{})
	Debugln(args ...interface{})
	Traceln(args ...interface{})
	Testln(args ...interface{})
}