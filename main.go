package main

import (
	"illuminati/go/microservice/controllers"
	"illuminati/go/microservice/routes"
	"illuminati/go/microservice/scheduler"
	"log"
	"net/http"
)

func main() {
	
	// Example cron expression: "0 9 * * *" runs daily at 9:00 AM
	//scheduler.StartCron("0 9 * * *", mailer.SendWordEmail)

	var mutexManager controllers.MutexManager

	routes.SetupAPI(mutexManager)
	log.Fatal(http.ListenAndServe(":1080", nil))

	//  run every 20 minutes
	scheduler.StartCron()

	// Start the HTTP server for manual triggering
	// Access via: http://localhost:8080/trigger
	//scheduler.StartHTTPServer(mailer.SendWordEmail)

	// Keep the main thread alive so cron and HTTP server continue running
	select {}
}
