package main

import (
	"service/mailer"
	"service/scheduler"
)

func main() {
	
	
	//scheduler.StartCron("0 9 * * *", mailer.SendWordEmail)

	//  run every 10 minutes
	scheduler.StartCron("@every 10m", mailer.SendWordEmail)

	// Start the HTTP server for manual triggering

	scheduler.StartHTTPServer(mailer.SendWordEmail)

	// Keep the main thread alive 
	select {}
}
