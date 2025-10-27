package scheduler

import (
	"log"

	"github.com/robfig/cron/v3"
)



func StartCron() {
	c := cron.New()
	
	_, err := c.AddFunc("@every 24m", newPassword)
	if err != nil {
	log.Println("Failed to schedule cron task:", err)
	return
	}
	
	c.Start()
	log.Println("Cron task scheduled with spec:", "@every 20m")
}
