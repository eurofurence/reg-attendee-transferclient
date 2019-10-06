# rexis-go-transferclient

## Overview

Implements a simple transfer client that periodically gets the maximum assigned attendee id from the go attendee service

```GET http://localhost:9091/api/rest/v1/attendees/max-id```

then gets the maximum attendee id known to the classic regsys

```GET http://localhost:8080/regsys/service/max-regnum-api```

and then transfers any missing registrations to the regsys via its transfer api

```GET http://localhost:8080/regsys/service/transfer-api?id=...&token=...```

where all the URLs as well as the transfer api token are read from a file called ```config.yaml```.  
