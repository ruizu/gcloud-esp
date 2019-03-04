# Google Cloud Endpoints - Extensible Service Proxy
This module decodes the user info header sent by Cloud Endpoints ESP.
Be sure to install Local ESP on your machine or send the `X-Endpoint-API-UserInfo` header to test.

## Usage
Get the source code using `go get`.
``` bash
$ go get -u github.com/ruizu/gcloud-esp/esp
```

### Example
``` go
package main

import (
	"log"
	"net/http"

	"github.com/ruizu/gcloud-esp/esp"
)

func main() {
	http.Handle("/hello", func(w http.ResponseWriter, r *http.Request) {
		user, err := esp.GetUserInfo(r)
		if err != nil {
			log.Panic(err)
		}
		fmt.Fprintf(w, "Hello, %s", user.ID)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
