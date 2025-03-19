/* I may make this into a standalone lib in the future for my self.
But for now it will be apart of this project. Just a log wrapper. */

package main

import (
	"io"
	"log"
	"os"
	"regexp"
)

type Audit struct {
	logger log.Logger
}

type AuditType int

func NewAudit() *Audit {
	a := new(Audit)
	a.logger = *log.New(os.Stderr, "", 0)
	return a
}

func (audit *Audit) toFile(w io.Writer) {
	audit.logger.SetOutput(w)
}

func logg(step, msg string) {
	pattern, _ := regexp.Compile("\r?\n")
	msg = pattern.ReplaceAllString(msg, " ")
	audit.logger.Printf("%s: %s", step, msg)
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
