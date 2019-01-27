package managers;

import (
    // "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/tidwall/gjson"
    // configs "github.com/markdingemanse/loveless/configs"
);

// Send a request.
func SendGetRequest(url string, userAgent string) (string, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "error", err;
    }
    req.Header.Set("User-Agent", userAgent)

    client := &http.Client{Timeout: 10 * time.Second};

    resp, err := client.Do(req);
    if err != nil {
        fmt.Printf("%g", err.Error());
    }

    body, err := ioutil.ReadAll(resp.Body);
    defer resp.Body.Close();

    // var data map[string]interface{};
    //
    // json.Unmarshal([]byte(body), &data);

    hash  := gjson.Get(string(body), "data.children.#.title");
    title := hash.Array()[0].String()

    return title, nil;
}
