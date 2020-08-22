package kit

type Logger interface {
	Log(keyvals ...interface{}) error
}

func LoggerWith(logger Logger, keyvals ...interface{}) Logger {
	if len(keyvals) < 1 {
		return logger
	}

	cl := &contextLogger{
		l: logger,
	}

	ecl, ok := logger.(*contextLogger)
	if ok {
		copy(cl.keyvals, ecl.keyvals)
	}
	cl.keyvals = append(keyvals, keyvals...)

	return cl
}

type contextLogger struct {
	l       Logger
	keyvals []interface{}
}

func (cl *contextLogger) Log(keyvals ...interface{}) error {
	kvs := append(cl.keyvals, keyvals...)
	return cl.l.Log(kvs...)
}
