# reg-attendee-transferclient

## Transfer Client Overview

Implements a simple transfer client that periodically gets the maximum assigned attendee id from the go attendee service

```GET http://localhost:9091/api/rest/v1/attendees/max-id```

then gets the maximum attendee id known to the classic regsys

```GET http://localhost:8080/regsys/service/max-regnum-api```

and then transfers any missing registrations to the regsys via its transfer api

```GET http://localhost:8080/regsys/service/transfer-api?id=...&token=...```

where all the URLs as well as the transfer api token are read from a file called ```config.yaml```.  

## Contract Tests Overview

A very minimal example for using pact-go for consumer driven contract tests.

This is the consumer side. 

See [rexis-go-attendee](https://github.com/jumpy-squirrel/rexis-go-attendee) for the producer side.

### Scenario

This simulates a microservice that wants to call another microservice, the Attendee Service.

Specifically, it wants to call its /info/health endpoint (a very contrived example for simplicity's sake).

### Solution Concept

In order to automatically verify that the interaction will work as expected, we have implemented the 
consumer side of a consumer driven contract test (see `test/contract/consumer/main_ctr_test.go`).

When the test suite of this "microservice" runs, the consumer side test is run, and a pact json is written out.

_Note how the consumer test calls into a very low level function, the one that uses a httpclient to make the call. 
So we are not testing the business logic, only the actual technical client code._

When the test suite of the producer runs, it reads the pact json and uses it to replay the interaction.

_Again, we use a mock service underneath the web controller to only test the technical interaction,
not the business logic. This is easy to do using a httptest server._

_In a more real world example you'd have some way to publish the generated pact jsons to a server and/or
check them into a repository. The producer side test can then use a URL on this server to get at the current
pact json._

### Installation of Pact

Download and install the pact command line tools and add them to your path as described in the
[pact-go manual](https://github.com/pact-foundation/pact-go#installation). This step is system
dependent.

### Go Dependencies

The regular go dependencies will be downloaded by go build / go test, because this project uses 
[go modules](https://blog.golang.org/using-go-modules).

_Note that in order for modules to kick in when building it, this source must reside OUTSIDE your go path.
At least it did for me._

### Run The Contract Tests

The main program isn't really of interest.

Instead, use

`go test -v github.com/jumpy-squirrel/rexis-go-democlient/...`

to run the consumer side test and generate the pact json file.

Then do the same in the producer project.

`go test -v github.com/jumpy-squirrel/rexis-go-attendee/...`

You should see output like this:

```
=== RUN   TestProvider
2019/09/15 18:59:03 [INFO] checking pact-mock-service within range >= 3.1.0, < 4.0.0
2019/09/15 18:59:04 [INFO] checking pact-provider-verifier within range >= 1.23.0, < 3.0.0
2019/09/15 18:59:05 [INFO] checking pact-broker within range >= 1.18.0
=== RUN   TestProvider/has_status_code_200
=== RUN   TestProvider/has_a_matching_body
=== RUN   TestProvider/"Content-Type"_which_equals_"text/plain;_charset=utf-8"
--- PASS: TestProvider (3.57s)
    --- PASS: TestProvider/has_status_code_200 (0.00s)
        pact.go:390: Verifying a pact between DemoClient and AttendeeService Given Attendee 1 exists A request to get health info with GET /info/health returns a response which has status code 200
    --- PASS: TestProvider/has_a_matching_body (0.00s)
        pact.go:390: Verifying a pact between DemoClient and AttendeeService Given Attendee 1 exists A request to get health info with GET /info/health returns a response which has a matching body
    --- PASS: TestProvider/"Content-Type"_which_equals_"text/plain;_charset=utf-8" (0.00s)
        pact.go:390: Verifying a pact between DemoClient and AttendeeService Given Attendee 1 exists A request to get health info with GET /info/health returns a response which includes headers "Content-Type" which equals "text/plain; charset=utf-8"
PASS
ok      rexis/rexis-go-attendee/test/contract/producer  4.101s
```
