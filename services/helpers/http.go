package services;

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
// func Redify(url string, userAgent string, key string) (string) {
//     req, err := GetBaseRequest(url, userAgent);
//
//     if (err != nil) {
//         HandleError(err)
//     }
//
//     client := GetClient();
//     resp, err := HandleRequest(client, req);
//
//     if (err != nil) {
//         HandleError(err)
//     }
//
//     hash  := ParseResponse(resp, key);
//
//     return GetTitle(hash);
// }

// Get the title of the given Result.
// func GetTitle(hash gjson.Result) (string) {
//     return hash.Array()[0].String();
// }

//############### Generic http logic ############### //

// Handle the response of the request.
func ParseResponse(resp *http.Response, key string) (gjson.Result) {
    body, err := ioutil.ReadAll(resp.Body);

    if (err != nil) {
        fmt.Printf("%g", err.Error());
    }

    defer resp.Body.Close();

    return gjson.Get(string(body), key);
}

// Run a request.
func HandleRequest(client *http.Client, request *http.Request) (*http.Response, error) {
    resp, err := client.Do(request);

    if (err != nil) {
        return nil, err;
    }

    return resp, nil;
}

// Build a base request.
func GetBaseRequest(url string, userAgent string) (*http.Request, error) {
    req, err := http.NewRequest("GET", url, nil)

    if (err != nil) {
        return nil, err;
    }

    req.Header.Set("User-Agent", userAgent)

    return req, nil;
}

// get a base client.
func GetClient() (*http.Client) {
    return &http.Client{Timeout: 10 * time.Second};
}

// handler a error.
func HandleError(err error) {
    fmt.Printf("%g", err.Error());
}
