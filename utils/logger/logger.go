package logger

import (
	"log"
	"os"
)

var (
	Debug   = log.New(os.Stdout, "[Debug] ", log.LstdFlags|log.Lshortfile)
	Info    = log.New(os.Stdout, "[Info] ", log.LstdFlags)
	Warning = log.New(os.Stdout, "[Warning] ", log.LstdFlags)
	Error   = log.New(os.Stderr, "[Error] ", log.LstdFlags)
)
