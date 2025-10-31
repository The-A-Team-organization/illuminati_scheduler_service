package main

import (
	"illuminati/go/microservice/callendpoints"

	"github.com/robfig/cron/v3"
)

func main() {
    c := cron.New()
    
    
    c.AddFunc("@every 10s", func() { callendpoints.CloseVotes() })

    c.Start()

    select {} 
}
