package util

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime/debug"
)

func NewError(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func Format(template string, args ...interface{}) string {
	return fmt.Sprintf(template, args...)
}

func Error(format string, args ...interface{}) error {
	if len(args) == 0 {
		return errors.New(format)
	}
	return fmt.Errorf(format, args...)
}

func AnyToPtr[T any](value T) *T {
	return &value
}

func GetAbsPath(directory string) (string, error) {
	return filepath.Abs(directory)
}

func RecoverPanic() {
	if r := recover(); r != nil {
		log.Printf("Recovered in function %v", r)
		var err error
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("unknown panic")
		}
		if err != nil {
			stackTrace := string(debug.Stack())
			log.Printf("error:%v | stacktrace from panic: \n [%s]", err, stackTrace)
		}
	}
}
