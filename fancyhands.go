package fancyhands

import (
    "io/ioutil"
    "net/http"
    "strconv"
    "time"
    "github.com/mrjones/oauth"
)

const (
    API_HOST = "https://www.fancyhands.com/api/v1/"
)

type Client struct {
    key          string
    secret       string
    test         bool
    token        *oauth.AccessToken
    oauth_client *oauth.Consumer
}

// Private
// Create a new Client object
func createClient(key string, secret string, test bool) *Client {
    dummyAccessToken := &oauth.AccessToken{}
    dummyServiceProvider := oauth.ServiceProvider{}

    return &Client{
        key:          key,
        secret:       secret,
        test:         test,
        token:        dummyAccessToken,
        oauth_client: oauth.NewConsumer(key, secret, dummyServiceProvider),
    }
}

// Create an API client
func NewClient(key string, secret string) *Client {
    return createClient(key, secret, false)
}

// Create a test API client. This will send API calls without actually
// sending tasks to Fancy Hands assistants.
func NewTestClient(key string, secret string) *Client {
    return createClient(key, secret, true)
}

// Send a string to the API and it will echo it back.
func (c *Client) Echo(value string) (code int, body string, err error) {
    data := make(map[string]string)
    data["echo"] = value

    return c.get("echo", data)
}

// Get all tasks
func (c *Client) GetAllTasks() (code int, body string, err error) {
    data := make(map[string]string)
    return c.get("request/custom", data)
}

// Get a specific task based on its key
func (c *Client) GetTask(key string) (code int, body string, err error) {
    data := make(map[string]string)
    data["key"] = key

    return c.get("request/standard", data)
}

// Get tasks filtered by status or by cursor. Pass in an empty string to omit one.
func (c *Client) GetTasks(status string, cursor string) (code int, body string, err error) {
    data := make(map[string]string)

    if status != "" {
        data["status"] = status
    }

    if cursor != "" {
        data["cursor"] = cursor
    }

    return c.get("request/standard", data)
}

// Create a task. All fields are required.
func (c *Client) CreateTask(title string, desc string, bid float64, expiration time.Time) (code int, body string, err error) {
    return c.CreateCustomTask(title, desc, bid, expiration, "")
}

// Create a custom task. The custom field must be a string formatted as JSON.
func (c *Client) CreateCustomTask(title string, desc string, bid float64, expiration time.Time, custom string) (code int, body string, err error) {
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
func (c *Client) CancelTask(key string) (code int, body string, err error) {
    data := make(map[string]string)
    data["key"] = key

    return c.post("request/custom/cancel", data)
}

// Add a message to a specific task based on the task's key.
func (c *Client) CreateMessage(key string, message string) (code int, body string, err error) {
    data := make(map[string]string)
    data["key"] = key
    data["message"] = message

    return c.post("request/standard/messages", data)
}

// Private
// Send a POST request to the API
func (c *Client) post(url string, data map[string]string) (code int, body string, err error) {
    return c.request(url, data, "post")
}

// Private
// Send a GET request to the API
func (c *Client) get(url string, data map[string]string) (code int, body string, err error) {
    return c.request(url, data, "get")
}

// Private
// Send a request to the API
func (c *Client) request(url string, data map[string]string, method string) (code int, body string, err error) {
    var response *http.Response
    var newErr error
    var response_body []byte

    if c.test == true {
        data["test"] = "true"
    }

    if method == "get" {
        response, newErr = c.oauth_client.Get(API_HOST + url, data, c.token)
    } else if method == "post" {
        response, newErr = c.oauth_client.Post(API_HOST + url, data, c.token)
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
