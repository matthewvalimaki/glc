package glc

import (
	"log"
	"fmt"
)

const (
	// see http://golang.org/pkg/log/#pkg-constants
	timeFormat = log.Ldate | log.Lmicroseconds | log.LUTC
)

func AddLog(packageName string, body string, arguments ...interface{}) {
	log.SetFlags(timeFormat)
	
	if arguments == nil {
		log.Printf("[%s] %s", packageName, body)
	} else {
		log.Printf("[%s] %s", packageName, fmt.Sprintf(body, arguments...))
	}
}

func AddErrorLog(packageName string, body error, arguments ...string) {
	log.SetFlags(timeFormat)
	
	if arguments == nil {
		log.Printf("[%s] %s", packageName, body)
	} else {
		log.Printf("[%s] %s", packageName, body, arguments)
	}
}

func AddUpdateLog(packageName string, body string, arguments ...interface{}) {
	log.SetFlags(timeFormat)
	
	log.Printf("[%s][UPDATE] %s", packageName, fmt.Sprintf(body, arguments...))
}