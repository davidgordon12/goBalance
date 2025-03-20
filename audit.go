/*
I may make this into a standalone lib in the future for my self.
But for now it will be apart of this project. Just a log wrapper.
*/
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

type Audit struct {
	logger log.Logger
}

type AuditType int

func NewAudit() *Audit {
	a := new(Audit)
	a.logger = *log.New(os.Stderr, time.Now().UTC().Format("[2006-01-02 15:04:05] "), 0)
	return a
}

func (audit *Audit) addFile(path string) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Couldn't open file " + path)
		return
	}
	wrt := io.MultiWriter(os.Stdout, f)
	audit.logger.SetOutput(wrt)
}

func (audit *Audit) info(msg string) {
	logg("INFO", msg)
}

func (audit *Audit) warn(msg string) {
	logg("WARNING", msg)
}

func (audit *Audit) error(msg string) {
	/* Send some sort of alert here as well eventually */
	logg("ERROR", msg)
}

func logg(step, msg string) {
	pattern, _ := regexp.Compile("\r?\n")
	msg = pattern.ReplaceAllString(msg, " ")
	audit.logger.Printf("%s: %s", step, msg)
}
