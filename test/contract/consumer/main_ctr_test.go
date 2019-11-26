package main

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"log"
	"github.com/eurofurence/reg-attendee-transferclient/healthclient"
	"testing"
)

// contract test consumer side

func TestConsumer(t *testing.T) {
	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "DemoClient",
		Provider: "AttendeeService",
		Host:     "localhost",
	}
	defer pact.Teardown()

	// Pass in test case (consumer side)
	// This uses the repository on the consumer side to make the http call, should be as low level as possible
	var test = func() (err error) {
		reply, err := healthclient.RetrieveHealthInfo("http", "localhost", pact.Server.Port, "/info/health")
		if err != nil {
			return err
		}
		if reply != "OK" {
			return fmt.Errorf("unexpected reply %s", reply)
		}
		return nil
	}

	// Set up our expected interactions.
	pact.
		AddInteraction().
		// contrived example, not really needed. This is the identifier of the state handler that will be called on the other side
		Given("Attendee 1 exists").
		UponReceiving("A request to get health info").
		WithRequest(dsl.Request{
			Method:  "GET",
			Path:    dsl.String("/info/health"),
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("text/plain; charset=utf-8")},
			Body:    dsl.String("OK"),
		})

	// Run the test, verify it did what we expected and capture the contract (writes a test log to logs/pact.log)
	if err := pact.Verify(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}

	// now write out the contract json (by default it goes to subdirectory pacts)
	if err := pact.WritePact(); err != nil {
		log.Fatalf("Error on pact write: %v", err)
	}
}
