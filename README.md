FancyHands-Go
=============

FancyHands API Wrapper in Go!

Installation:
-------------
```
go get github.com/AMeng/fancyhands-go
```

Example:
--------

```golang
func main() {
    client := NewTestClient("YOUR KEY HERE", "YOUR SECRET HERE")

    status_code, response, err := client.Echo("hello world")

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(status_code, response)
    }
}
```
