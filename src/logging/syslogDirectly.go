package main

import (
	"fmt"
	"log"
	"log/syslog"
)

/*
On a Mac, /var/log/system.log,
whereas some flavors of Linux write it to /var/log/messages.
*/
func main() {
	priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	flags := log.Ldate | log.Lshortfile
	logger, err := syslog.NewLogger(priority, flags)
	if err != nil {
		fmt.Printf("Can't attach to syslog: %s", err)
		return
	}

	logger.Println("This is a test log message.")
}
