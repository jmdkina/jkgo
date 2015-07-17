package jklog

func (l *Logger) Debug(v ...interface{}) {
	if LEVEL_DEBUG&l.level == 0 {
		return
	}

	l.print(l.name+logPrefixs[LEVEL_DEBUG]+" ", v...)
}

func (l *Logger) Debugln(v ...interface{}) {
	if LEVEL_DEBUG&l.level == 0 {
		return
	}

	l.println(l.name+logPrefixs[LEVEL_DEBUG]+" ", v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if LEVEL_DEBUG&l.level == 0 {
		return
	}

	l.printf(l.name+logPrefixs[LEVEL_DEBUG]+" ", format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	if LEVEL_INFO&l.level == 0 {
		return
	}

	l.print(l.name+logPrefixs[LEVEL_INFO]+" ", v...)
}

func (l *Logger) Infoln(v ...interface{}) {
	if LEVEL_INFO&l.level == 0 {
		return
	}

	l.println(l.name+logPrefixs[LEVEL_INFO]+" ", v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if LEVEL_INFO&l.level == 0 {
		return
	}

	l.printf(l.name+logPrefixs[LEVEL_INFO]+" ", format, v...)
}

func (l *Logger) Notice(v ...interface{}) {
	if LEVEL_NOTICE&l.level == 0 {
		return
	}

	l.print(l.name+logPrefixs[LEVEL_NOTICE]+" ", v...)
}

func (l *Logger) Noticeln(v ...interface{}) {
	if LEVEL_NOTICE&l.level == 0 {
		return
	}

	l.println(l.name+logPrefixs[LEVEL_NOTICE]+" ", v...)
}

func (l *Logger) Noticef(format string, v ...interface{}) {
	if LEVEL_NOTICE&l.level == 0 {
		return
	}

	l.printf(l.name+logPrefixs[LEVEL_NOTICE]+" ", format, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	if LEVEL_WARN&l.level == 0 {
		return
	}

	l.print(l.name+logPrefixs[LEVEL_WARN]+" ", v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	if LEVEL_WARN&l.level == 0 {
		return
	}

	l.println(l.name+logPrefixs[LEVEL_WARN]+" ", v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if LEVEL_WARN&l.level == 0 {
		return
	}

	l.printf(l.name+logPrefixs[LEVEL_WARN]+" ", format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	if LEVEL_ERROR&l.level == 0 {
		return
	}

	l.print(l.name+logPrefixs[LEVEL_ERROR]+" ", v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	if LEVEL_ERROR&l.level == 0 {
		return
	}

	l.println(l.name+logPrefixs[LEVEL_ERROR]+" ", v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if LEVEL_ERROR&l.level == 0 {
		return
	}

	l.printf(l.name+logPrefixs[LEVEL_ERROR]+" ", format, v...)
}

func (l *Logger) Critical(v ...interface{}) {
	if LEVEL_CRITICAL&l.level == 0 {
		return
	}

	l.print(l.name+logPrefixs[LEVEL_CRITICAL]+" ", v...)
}

func (l *Logger) Criticalln(v ...interface{}) {
	if LEVEL_CRITICAL&l.level == 0 {
		return
	}

	l.println(l.name+logPrefixs[LEVEL_CRITICAL]+" ", v...)
}

func (l *Logger) Criticalf(format string, v ...interface{}) {
	if LEVEL_CRITICAL&l.level == 0 {
		return
	}

	l.printf(l.name+logPrefixs[LEVEL_CRITICAL]+" ", format, v...)
}

func (l *Logger) Noneln(v ...interface{}) {
	l.noneprintln(v...)
}

func (l *Logger) Nonef(format string, v ...interface{}) {
	l.noneprintf(format, v...)
}

func (l *Logger) None(v ...interface{}) {
	l.noneprint(v...)
}
