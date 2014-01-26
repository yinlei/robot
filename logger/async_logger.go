package logger



type AsyncLogger struct {
	buffer chan string
}


func NewAsyncLogger(size int) Logger {
	l := AsyncLogger {buffer: make(chan string, size)}
	//go func() {
		//for {
		//	log <- l.buffer
		//}
	//}

	return &l
}


func (al *AsyncLogger) Debug(info, message string) {
	al.buffer <- info
}

func (al *AsyncLogger) Info(info, message string) {
	al.buffer <- info
}

func (al *AsyncLogger) Warning(info, message string) {
	al.buffer <- info
}

func (al *AsyncLogger) Error(info, message string) {
	al.buffer <- info
}

func (al *AsyncLogger) Critical(info, message string) {
	al.buffer <- info
}

