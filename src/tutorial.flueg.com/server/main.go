package main

import (
	"flag"
	"fmt"

	log "github.com/cihub/seelog"
	"tutorial.flueg.com/config"
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
	var appConf = flag.String("appConf", "../conf/tutorial_server.xml", "Server configuration file format (xml)")
	var logConf = flag.String("logConf", "../conf/tutorial_server_log.conf", "Log configuration file (seelog)")
	flag.Parse()

	// Initialize logger system.
	err := InitLogSystem((*logConf))
	defer log.Flush()
	if err != nil {
		fmt.Errorf("Failed to init log configuration %s[err:%s]", *logConf, err)
	}
	log.Infof("Succeed to initialize log configuration file: %s", *logConf)

	// Initialize server configurations.
	myConf := config.MCSingleton()
	err = myConf.Load((*appConf))
	if err != nil {
		fmt.Errorf("Failed to init server configuration from file %s. err:%s", *appConf, err)
	}
	log.Infof("Succeed to initialize server configuration file: %s", *appConf)

}
