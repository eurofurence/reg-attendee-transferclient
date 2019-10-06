package main

import (
	"github.com/jumpy-squirrel/rexis-go-transferclient/internal/repository/client"
	"github.com/jumpy-squirrel/rexis-go-transferclient/internal/repository/config"
	"log"
	"time"
)

func main() {
	config.LoadConfiguration()

	for {
		log.Println("DEBUG sleeping for 30 seconds between runs, press Ctrl-C to terminate this process")
		time.Sleep(30 * time.Second)

		regsysMax, err := client.RetrieveRegsysMaxId()
		if err == nil {
			serviceMax, err := client.RetrieveAttendeeServiceMaxId()
			if err == nil {
				if regsysMax >= serviceMax {
					log.Printf("INFO  nothing to do, regsys max id %v, service max id %v", regsysMax, serviceMax)
				} else if regsysMax == 0 || serviceMax == 0 {
					log.Printf("INFO  skipping transfer due to 0, regsys max id %v, service max id %v", regsysMax, serviceMax)
				} else {
					log.Printf("INFO  initiating transfers, regsys max id %v, service max id %v", regsysMax, serviceMax)
					for id := regsysMax + 1; id <= serviceMax; id++ {
						err := client.PerformTransfer(id)
						if err != nil {
							// abort this run after any errors
							break
						}
					}
				}
			}
		}
	}
}
