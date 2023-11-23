package main

import (
	"log"

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

	w := worker.New(c, "hello-world", worker.Options{})

	w.RegisterWorkflow(emancipation.Workflow)
	w.RegisterActivity(emancipation.BuyMyselfFlowers)
	w.RegisterActivity(emancipation.WriteMyNameInTheSand)
	w.RegisterActivity(emancipation.TalkToMyselfForHours)
	w.RegisterActivity(emancipation.SayThingsYouDontUnderstand)
	w.RegisterActivity(emancipation.TakeMyselfDancing)
	w.RegisterActivity(emancipation.HoldMyOwnHand)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
