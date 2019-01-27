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
    // helpers.SendGetRequest("https://jsonplaceholder.typicode.com/todos/1");
    helpers.SendGetRequest("https://www.reddit.com/r/symphonicmetal/new.json?sort=new&limit=1", "loveless");
}
