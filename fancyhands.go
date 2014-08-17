package fancyhands

import (
    "io/ioutil"
    "net/http"
    "strconv"
    "time"
    "github.com/mrjones/oauth"
)

const (
    api_host       = "https://www.fancyhands.com/api/v1/"
    STATUS_NEW     = 1
    STATUS_OPEN    = 5
    STATUS_CLOSED  = 20
    STATUS_EXPIRED = 21
)

type client struct {
    test  bool
    oauth *oauth.Consumer
}

// Create a new client object
func createClient(key string, secret string, test bool) *client {
    return &client{
        test:  test,
        oauth: oauth.NewConsumer(key, secret, oauth.ServiceProvider{}),
    }
}

// Create an API client
func NewClient(key string, secret string) *client {
    return createClient(key, secret, false)
}

// Create a test API client. This will send API calls without actually
// sending tasks to Fancy Hands assistants.
func NewTestClient(key string, secret string) *client {
    return createClient(key, secret, true)
}

// Send a string to the API and it will echo it back.
func (c *client) Echo(value string) (code int, body string, err error) {
    return c.get("echo", map[string]string{"echo": value})
}

// Get all tasks
func (c *client) GetAllTasks() (code int, body string, err error) {
    return c.get("request/custom", map[string]string{})
}

// Get a specific task based on its key
func (c *client) GetTask(key string) (code int, body string, err error) {
    return c.get("request/custom", map[string]string{"key": key})
}

// Get tasks filtered by status. Use the predefined status constants (fancyhands.STATUS_OPEN, etc).
func (c *client) GetTasksByStatus(status int) (code int, body string, err error) {
    return c.get("request/custom", map[string]string{"status": strconv.Itoa(status)})
}

// Get tasks filtered by cursor. The API may return a cursor for pagination.
func (c *client) GetTasksByCursor(cursor string) (code int, body string, err error) {
    return c.get("request/custom", map[string]string{"cursor": cursor})
}

// Create a task. All fields are required.
func (c *client) CreateTask(title string, desc string, bid float64, expiration time.Time) (code int, body string, err error) {
    return c.CreateCustomTask(title, desc, bid, expiration, "")
}

// Create a custom task. The custom field must be a string formatted as JSON.
func (c *client) CreateCustomTask(title string, desc string, bid float64, expiration time.Time, custom string) (code int, body string, err error) {
    // TODO: It would be nice to create a struct for 'custom'. Something like: map[int]map[string]string

    data := make(map[string]string)
    data["title"] = title
    data["description"] = desc
    data["bid"] = strconv.FormatFloat(bid, 'f', 2, 64)
    data["expiration_date"] = expiration.Format(time.RFC3339)

    if custom != "" {
        data["custom_fields"] = custom
    }

    return c.post("request/custom", data)
}

// Cancel a task based on its key
func (c *client) CancelTask(key string) (code int, body string, err error) {
    return c.post("request/custom/cancel", map[string]string{"key": key})
}

// Add a message to a specific task based on the task's key.`
func (c *client) CreateMessage(key string, message string) (code int, body string, err error) {
    return c.post("request/standard/messages", map[string]string{"key": key, "message": message})
}

// Create a task for a phone call. The conversation field must be a string formatted as JSON.
func (c *client) Call(phone string, conversation string) (code int, body string, err error) {
    return c.post("call", map[string]string{"phone": phone, "conversation": conversation})
}

// Send a POST request to the API
func (c *client) post(url string, data map[string]string) (code int, body string, err error) {
    return c.request(url, data, "post")
}

// Send a GET request to the API
func (c *client) get(url string, data map[string]string) (code int, body string, err error) {
    return c.request(url, data, "get")
}

// Send a request to the API
func (c *client) request(url string, data map[string]string, method string) (code int, body string, err error) {
    var response *http.Response
    var newErr error
    var response_body []byte

    if c.test == true {
        data["test"] = "true"
    }

    if method == "get" {
        response, newErr = c.oauth.Get(api_host + url, data, &oauth.AccessToken{})
    } else if method == "post" {
        response, newErr = c.oauth.Post(api_host + url, data, &oauth.AccessToken{})
    }

    if newErr != nil {
        return 0, "", newErr
    }

    response_body, newErr = ioutil.ReadAll(response.Body)

    if newErr != nil {
        return 0, "", newErr
    }

    return response.StatusCode, string(response_body), nil
}
