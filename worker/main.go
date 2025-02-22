package main

import (
	"log"
	"net/http"
	iplocate "temporal-tryout"

 	"go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, iplocate.TaskQueueName, worker.Options{})

	activities := &iplocate.IPActivities{
		HTTPClient: http.DefaultClient,
	}

	w.RegisterActivity(iplocate.GetAddressFromIP)
	w.RegisterActivity(activities)

	// Start the Worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Temporal worker", err)
	}
}