package golog

import (
	"fmt"
	"io"
	"os"
	"time"
)

var file__error, file__info, file__log io.Writer

const opts = os.O_CREATE | os.O_WRONLY | os.O_APPEND
const perm = 0644

var _files map[Log_type]io.Writer

func init() {
	var err error
	file__error, err = os.OpenFile("error.log", opts, perm)
	__panic__err(err)
	file__info, err = os.OpenFile("info.log", opts, perm)
	__panic__err(err)
	file__log, err = os.OpenFile("log.log", opts, perm)
	__panic__err(err)
	_files = map[Log_type]io.Writer{
		ERR:  file__error,
		INFO: file__info,
		LOG:  file__log,
	}
}

func log__format(data ...interface{}) []byte {
	var _data = []interface{}{time.Now().Format("2006-01-02 03:04:05"), ">> "}
	return []byte(fmt.Sprint(append(_data, data...)...))
}

func __panic__err(err error) {
	if nil != err {
		panic(err)
	}
}
