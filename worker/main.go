package main

import (
	"log"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/garbage-collector/temporal-demo/emancipation"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{HostPort: "localhost:7233"})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "demo", worker.Options{})

	w.RegisterWorkflow(emancipation.Workflow)
	for activityName, activityFn := range emancipation.Activities {
		w.RegisterActivityWithOptions(activityFn, activity.RegisterOptions{Name: activityName})
	}

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
