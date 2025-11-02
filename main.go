package main

import (
	"illuminati/go/microservice/callendpoints"

	"github.com/robfig/cron/v3"
)

func main() {
   	c := cron.New()

	
	c.AddFunc("@hourly", func() { callendpoints.CloseVotes() })
    //c.AddFunc("@every 10s", func() { callendpoints.CloseVotes() })

	
	c.AddFunc("@daily", func() { callendpoints.SetInquisitor() })
    

	
	c.AddFunc("0 20 * * *", func() { callendpoints.UnsetInquisitor() })

    c.AddFunc("@every 15s", func() { callendpoints.BanArchitect() })
    //c.AddFunc("@every 42d", func() { callendpoints.BanArchitect() })
    
    c.AddFunc("@daily", func() { callendpoints.NewEntryPassword() })
    //c.AddFunc("@every 10s", func() { callendpoints.NewEntryPassword() })

	c.Start()

	select {}
}
