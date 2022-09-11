package pkg

type ILogger interface {
	Error(message string)
}

type Logger struct {
}

func NewLogger() ILogger {
	return &Logger{}
}

func (l *Logger) Error(message string) {

}
