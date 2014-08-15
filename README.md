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

Instantiate the client:
-------------------------

**For testing:**
```go
client := fancyhands.NewTestClient("YOUR KEY HERE", "YOUR SECRET HERE")
```

**For creating actual tasks:**
```go
client := fancyhands.NewClient("YOUR KEY HERE", "YOUR SECRET HERE")
```

Using the client:
-----------------
All client methods return the same three values. The response code, the JSON response (as a string), and any error that was raised.
```go
(code int, body string, err error)
```

Client methods:
---------------

**Echo:** *Send a string and the API will return the message back.*
```go
func (c *Client) Echo(value string)
```

**GetAllTasks:** *Get all tasks you have created. You may receive a "cursor" for pagination.*
```go
func (c *Client) GetAllTasks()
```

**GetTask:** *Get a specific task based on its key.*
```go
func (c *Client) GetTask(key string)
```

**GetTasks:** *Filter tasks by status or by cursor.*
```go
func (c *Client) GetTasks(status string, cursor string)
```

**CreateTask:** *Create a task.*
```go
func (c *Client) CreateTask(title string, desc string, bid float64, expiration time.Time)
```

**CreateCustomTask:** *Create a custom task. Pass in a JSON string of custom fields. See [the API docs](https://github.com/fancyhands/api/wiki/fancyhands.request.Custom#custom_fields) for details.*
```go
func (c *Client) CreateCustomTask(title string, desc string, bid float64, expiration time.Time, custom string)
```

**CancelTask:** *Cancel a task based on its key.*
```go
func (c *Client) CancelTask(key string) (code int, body string, err error)
```
