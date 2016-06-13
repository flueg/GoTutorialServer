package main

import (
	"flag"
	"fmt"

	log "github.com/cihub/seelog"
)

func InitLogSystem(confFile string) error {
	logger, err := log.LoggerFromConfigAsFile(confFile)
	if err != nil {
		return err
	}
	return log.ReplaceLogger(logger)
}

func main() {
	fmt.Println("Hello Flueg!")

	// Parse passed in parameters.
	//var appConf = flag.String("appConf", "../conf/interface.xml", "Server configuration file format (xml)")
	var logConf = flag.String("logConf", "../conf/tutorial_server_log.conf", "Log configuration file (seelog)")
	flag.Parse()

	// Initialize logger system.
	err := InitLogSystem((*logConf))
	if err != nil {
		fmt.Errorf("Failed to init log configuration %s[err:%s]", logConf, err)
	}
}
