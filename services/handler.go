package services;

import (
    "fmt"

    helpers "github.com/markdingemanse/loveless/services/helpers"
)

// Gets the needed function to run dynamicly based on a k / v array.
func Functions(key string) func() {
    return map[string] func() {
        "rss": rss}[key];
}

// TODO::handle rss
func rss() {
    fmt.Println("[RSS_HANDLER] reached the handler rss function");

    parsed, err := helpers.SendGetRequest("https://www.reddit.com/r/symphonicmetal/new.json?sort=new&limit=1", "loveless");

    if err != nil {
        fmt.Println("[RSS_HANDLER] error: " + err.Error())
    }

    fmt.Println("[RSS_HANDLER] finished rss most recent post is: " + parsed);
}
