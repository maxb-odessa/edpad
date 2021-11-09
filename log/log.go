package log

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var lck sync.Mutex

var debug bool

func SetDebug(d bool) {
	debug = d
}

// log error message
func genlog(level string, format string, args ...interface{}) {
	lck.Lock()
	fmt.Fprintf(os.Stderr, time.Now().Format("15:04:05 01-02-2006")+" edpad["+level+"] "+format, args...)
	lck.Unlock()
}

// Err logs ERROR messages to stderr
func Err(format string, args ...interface{}) {
	genlog("ERR", format, args...)
}

// Info logs INFO messages to stderr
func Info(format string, args ...interface{}) {
	genlog("INFO", format, args...)
}

// Warn logs WARNING messages to stderr
func Warn(format string, args ...interface{}) {
	genlog("WARN", format, args...)
}

// Fatal logs FATAL messages to stderr and calls os.Exit()
func Fatal(format string, args ...interface{}) {
	genlog("FATAL", format, args...)
	os.Exit(1)
}

// Debug prints some debug if ELDA_DEBUD env var defined
func Debug(format string, args ...interface{}) {
	if debug {
		genlog("DEBUG", format, args...)
	}
}
