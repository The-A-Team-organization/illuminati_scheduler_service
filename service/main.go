package main

import (
	"service/mailer"
	"service/scheduler"
)

func main() {
	
	// Example cron expression: "0 9 * * *" runs daily at 9:00 AM
	//scheduler.StartCron("0 9 * * *", mailer.SendWordEmail)

	//  run every 20 minutes
	scheduler.StartCron("@every 20m", mailer.SendWordEmail)

	// Start the HTTP server for manual triggering
	// Access via: http://localhost:8080/trigger
	scheduler.StartHTTPServer(mailer.SendWordEmail)

	// Keep the main thread alive so cron and HTTP server continue running
	select {}
}
