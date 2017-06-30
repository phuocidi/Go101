package main

import (
	"log/syslog"
)

/*
On a Mac, /var/log/system.log,
whereas some flavors of Linux write it to /var/log/messages.
*/
func main() {

	logger, err := syslog.New(syslog.LOG_LOCAL3, "Pete")

	if err != nil {
		panic("Cannot attach to syslog")
	}

	defer logger.Close()
	logger.Debug("Debug message")
	logger.Notice("Notice message")
	logger.Warning("Warning message")
	logger.Alert("Alert message")
}
