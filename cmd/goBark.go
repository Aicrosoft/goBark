package main

import (
	"flag"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/fatih/color"
)

var (
	// Version is current version of GoDNS.
	Version   = "0.1"
	DebugMode = "false"
	IsDebug   = false
)

func init() {

	IsDebug, _ = strconv.ParseBool(DebugMode)
	if IsDebug {
		log.SetOutput(os.Stdout)
		//log.SetFormatter(&diagnose.DebugFormatter{})
		log.SetFormatter(&log.TextFormatter{
			ForceColors:     true,
			TimestampFormat: "15:04:05",
			FullTimestamp:   true,
		})
		if IsDebug {
			log.SetLevel(log.DebugLevel)
		}
		log.Debug("Debug log")
		log.Info("Info log")
		log.Warn("Warn log")
		log.Error("Error log")
		//log.Fatal("Fatal log") //碰到这种Log会直接中断后续运行。NB！

	} else {
		log.SetOutput(os.Stdout)
	}

}

func main() {
	flag.Parse()

	if *optHelp {
		color.Cyan(util.Logo, Version, IsDebug)
		flag.Usage()
		return
	}

}
