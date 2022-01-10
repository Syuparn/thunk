package thunk

import (
	log "github.com/sirupsen/logrus"
)

var (
	flagOutput  string
	printsDebug bool
)

func init() {
	Generator.Flags.StringVar(&flagOutput, "o", "", "output file name")
	Generator.Flags.BoolVar(&printsDebug, "debug", false, "print debug logs while generating")
}

func initDebugLog() {
	if printsDebug {
		log.SetLevel(log.DebugLevel)
	}
}
