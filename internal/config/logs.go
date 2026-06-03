package config

import (
	"log"
	"os"
)

func Logs(app any) {
	cfg := Load()

	logFile, err := os.OpenFile(cfg.LOGSFILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println(app)
}
