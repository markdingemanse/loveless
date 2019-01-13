package services;

import (
    "fmt"
)

// Gets the needed function to run dynamicly based on a k / v array.
func Functions(key string) func() {
    return map[string] func() {
        configs.Knowledge()}[key];
}

// TODO::handle rss
func rss() {
    fmt.Println("[RSS_HANDLER] reached the handler rss function");
}
