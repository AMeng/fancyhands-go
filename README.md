FancyHands-Go
=============

[FancyHands API](https://www.fancyhands.com/developer) Wrapper in Go!

Installation:
-------------
```
go get github.com/ameng/fancyhands-go
```

Example:
--------

```go
package main

import (
    "fmt"
    "github.com/ameng/fancyhands-go"
)

func main() {
    client := fancyhands.NewTestClient("YOUR KEY HERE", "YOUR SECRET HERE")

    status_code, response, err := client.Echo("hello world")

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(status_code, response)
    }
}
```
