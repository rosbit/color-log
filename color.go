package clog

import (
	"github.com/gookit/color"
	"log"
	"io"
	"fmt"
)

const (
	l_info = iota
	l_warn
	l_error
)

var (
	cError = color.FgRed.Render
	cInfo  = color.New(color.FgWhite, color.BgGreen).Render
	cWarn  = color.FgYellow.Render

	levels = map[int]string{
		l_info: cInfo("info"),
		l_warn: cWarn("warn"),
		l_error: cError("error"),
	}

	defaultLog = &ColorLogger{l: log.Default()}
)

type ColorLogger struct {
	l *log.Logger
}

func SetOutput(w io.Writer) {
	defaultLog.l.SetOutput(w)
}

func Info(args ...interface{}) {
	defaultLog.log(l_info, args...)
}
func Infof(format string, args ...interface{}) {
	defaultLog.logf(l_info, format, args...)
}

func Warn(args ...interface{}) {
	defaultLog.log(l_warn, args...)
}
func Warnf(format string, args ...interface{}) {
	defaultLog.logf(l_warn, format, args...)
}

func Error(args ...interface{}) {
	defaultLog.log(l_error, args...)
}
func Errorf(format string, args ...interface{}) {
	defaultLog.logf(l_error, format, args...)
}

func NewColorLogger(l *log.Logger) *ColorLogger {
	return &ColorLogger{l:l}
}

func (cl *ColorLogger) SetOutput(w io.Writer) {
	cl.l.SetOutput(w)
}

func (cl *ColorLogger) Info(args ...interface{}) {
	cl.log(l_info, args...)
}
func (cl *ColorLogger) Infof(format string, args ...interface{}) {
	cl.logf(l_info, format, args...)
}

func (cl *ColorLogger) Warn(args ...interface{}) {
	cl.log(l_warn, args...)
}
func (cl *ColorLogger) Warnf(format string, args ...interface{}) {
	cl.logf(l_warn, format, args...)
}

func (cl *ColorLogger) Error(args ...interface{}) {
	cl.log(l_error, args...)
}
func (cl *ColorLogger) Errorf(format string, args ...interface{}) {
	cl.logf(l_error, format, args...)
}

func (cl *ColorLogger) log(level int, args ...interface{}) {
	colorLevel, _ := levels[level]
	cl.l.Println(fmt.Sprintf("[%s]", colorLevel), fmt.Sprint(args...))
}
func (cl *ColorLogger) logf(level int, format string, args ...interface{}) {
	colorLevel, _ := levels[level]
	cl.l.Printf(fmt.Sprintf("[%s] %s", colorLevel, format), args...)
}

