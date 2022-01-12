package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	info  = color.New(color.FgGreen).PrintlnFunc()
	err   = color.New(color.FgRed).PrintlnFunc()
	fatal = color.New(color.BgRed).PrintlnFunc()
)

func Info(contents ...interface{}) {
	content := formatMessage(contents...)
	info("~ TG-Builder > ", content)
}

func Error(contents ...interface{}) {
	content := formatMessage(contents...)
	err("~ TG-Builder > ", content)
}

func Fatal(contents ...interface{}) {
	content := formatMessage(contents...)
	fatal("~ TG-Builder > ", content)
}

func formatMessage(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	msg = strings.TrimRight(msg, " \n\r")
	return msg
}
