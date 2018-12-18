package golog

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

type Log_type uint8

const (
	ERR Log_type = iota
	INFO
	LOG
)

//Log_file accepts a file and uses that file for logging
//unless you are not doing anything custom this shouldn't be necessary
func Log_file(f io.Writer, data ...interface{}) { //{{{
	f.Write(log__format(data...))
	f.Write([]byte{'\n'})
} //}}}

//Log_file_by_name does exactly what Log_file does except accepts a filename
//unless you are not doing anything custom this shouldn't be necessary
func Log_file_by_name(fname string, data ...interface{}) { //{{{
	f, err := os.OpenFile(fname, opts, perm)
	if nil != err {
		fmt.Fprintln(os.Stderr, log__format(err.Error()))
		return
	}
	defer f.Close()
	Log_file(f, data...)
} //}}}
//Log_t logs the data by given log type (ERR INFO or LOG)
//you can use higher lever ERR(),INFO or LOG functions instead
func Log_t(t Log_type, data ...interface{}) { //{{{
	_, file, line, _ := runtime.Caller(1)
	f, ok := _files[t]
	if !ok {
		fmt.Fprintln(os.Stderr, log__format(file, line, "Unknown logger type"))
		return
	}
	a := append([]interface{}{}, fmt.Sprintf("%s:%d\t", file, line))
	a = append(a, data...)
	Log_file(f, a...)
} //}}}
//Err high level error logger
func Err(data ...interface{}) { //{{{
	Log_t(ERR, data...)
} //}}}
//Log high level log logger
func Log(data ...interface{}) { //{{{
	Log_t(LOG, data...)
} //}}}
//Info high level info logger
func Info(data ...interface{}) { //{{{
	Log_t(INFO, data...)
} //}}}
