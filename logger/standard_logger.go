package logger

import (
	"sync"
	"io"
	"time"
)

type StandardLogger struct {
	mutex sync.Mutex
	out	  io.Writer
}

func NewStandardLogger(out io.Writer) Logger {
	return &StandardLogger{out: out}
}

func (sl *StandardLogger) Debug(info, message string) {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()

	io.WriteString(sl.out, "[D] ")
	io.WriteString(sl.out, time.Now().Format(timeFormat))
	io.WriteString(sl.out, " ")
	io.WriteString(sl.out, info)
	io.WriteString(sl.out, " ")
	io.WriteString(sl.out, message)
	io.WriteString(sl.out, "\n")
}

func (sl *StandardLogger) Info(info, message string) {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()

	io.WriteString(sl.out, "[I] ")
	io.WriteString(sl.out, time.Now().Format(timeFormat))
	io.WriteString(sl.out, " ")
	io.WriteString(sl.out, info)
	io.WriteString(sl.out, " ")
	io.WriteString(sl.out, message)
	io.WriteString(sl.out, "\n")
}

func (sl *StandardLogger) Warning(info, message string) {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()

	io.WriteString(sl.out, "[W] ")
	io.WriteString(sl.out, time.Now().Format(timeFormat))
	io.WriteString(sl.out, " ")
	io.WriteString(sl.out, info)
	io.WriteString(sl.out, " ")
	io.WriteString(sl.out, message)
	io.WriteString(sl.out, "\n")
}

func (sl *StandardLogger) Error(info, message string) {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()

	io.WriteString(sl.out, "[E] ")
	io.WriteString(sl.out, time.Now().Format(timeFormat))
	io.WriteString(sl.out, " ")
	io.WriteString(sl.out, info)
	io.WriteString(sl.out, " ")
	io.WriteString(sl.out, message)
	io.WriteString(sl.out, "\n")
}

func (sl *StandardLogger) Critical(info, message string) {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()

	io.WriteString(sl.out, "[C] ")
	io.WriteString(sl.out, time.Now().Format(timeFormat))
	io.WriteString(sl.out, " ")
	io.WriteString(sl.out, info)
	io.WriteString(sl.out, " ")
	io.WriteString(sl.out, message)
	io.WriteString(sl.out, "\n")
}

