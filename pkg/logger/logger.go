package logger

import "log"

// Пакеты, которые будут использоваться повторно, выносим в отдельную библиотеку

const (
	LevelDebug   = 1
	LevelInfo    = 3
	LevelWarning = 4
	LevelError   = 5
)

type Logger struct {
	level int
}

func New(lvl int) *Logger {
	return &Logger{
		level: lvl,
	}
}

func (l *Logger) Debug(msg string) {
	if l.level >= LevelDebug {
		log.Println("Debug: ", msg)
	}
}

func (l *Logger) Info(msg string) {
	if l.level >= LevelInfo {
		log.Println("Info: ", msg)
	}
}

func (l *Logger) Warning(msg string, err error) {
	if l.level >= LevelWarning {
		log.Println("Warning: ", msg, "(", err.Error(), ")")
	}
}
func (l *Logger) Error(msg string, err error) {
	if l.level >= LevelError {
		log.Println("Error: ", msg, "(", err.Error(), ")")
	}
}
