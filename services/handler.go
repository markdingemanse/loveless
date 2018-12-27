package "services"

import (
    "fmt"
)

func Functions(key string) func {
    return map[string]func(int, int) int {
        "rss": rss,
    }[key];
}

func rss() {
    fmt.Println("[RSS_HANDLER] reached the handler rss function"));
}
