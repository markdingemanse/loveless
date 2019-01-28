package managers;

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/tidwall/gjson"
    // configs "github.com/markdingemanse/loveless/configs"
);

// ################### REDIFY ###################### //

// Send a request.
func Redify(url string, userAgent string, key string) (string) {
    req, err := GetBaseRequest(url userAgent);
    client := GetClient();
    resp, err := HandleRequest(client, req);
    hash  := ParseResponse(resp, key);

    return GetTitle(hash);
}

// Get the title of the given Result.
func GetTitle(hash *Result) (string) {
    return hash.Array()[0].String();
}

//############### Generic http logic ############### //

// Handle the response of the request.
func ParseResponse(resp *Response, key string) (*Result) {
    body, err := ioutil.ReadAll(resp.Body);

    if err != nil {
        fmt.Printf("%g", err.Error());
    }

    defer resp.Body.Close();

    return gjson.Get(string(body), key);
}

// Run a request.
func HandleRequest(client *Client, request *Request) (*Response, error) {
    resp, err := client.Do(request);

    if err != nil {
        return nil, err;
    }

    return resp, nil;
}

// Build a base request.
func GetBaseRequest(url string, userAgent string) (*Request, error) {
    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
        return nil, err;
    }

    req.Header.Set("User-Agent", userAgent)

    return req;
}

// get a base client.
func GetClient() (*Client) {
    return &http.Client{Timeout: 10 * time.Second};
}
