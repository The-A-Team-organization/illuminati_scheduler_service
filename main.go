package main

import (
	"illuminati/go/microservice/callendpoints"

	"github.com/robfig/cron/v3"
)

func main() {
   	c := cron.New()

	
	//c.AddFunc("@hourly", func() { callendpoints.CloseVotes() })
    c.AddFunc("@every 10s", func() { callendpoints.CloseVotes() })

	
	c.AddFunc("@daily", func() { callendpoints.SetInquisitor() })

	
	c.AddFunc("0 20 * * *", func() { callendpoints.UnsetInquisitor() })

	c.Start()

	select {}
}
