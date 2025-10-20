package scheduler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/robfig/cron/v3"
)


func StartCron(spec string, task func()) {
	c := cron.New()

	_, err := c.AddFunc(spec, func() {
		log.Println("Task executed via cron")
		task()})
		if err != nil {
		log.Println("Failed to schedule cron task:", err)
		return
	}
	


	c.Start()
	log.Println("Cron task scheduled with spec:", spec)
}


func StartHTTPServer(task func()) {
	http.HandleFunc("/trigger", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			log.Println("Task triggered manually via HTTP")
			task()
		}()
		fmt.Fprintln(w, "Task triggered manually")
	})

	log.Println("HTTP server is running on :8080")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("HTTP server error:", err)
		}
	}()
}
