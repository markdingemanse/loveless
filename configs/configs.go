package configs;

import (
    "os"

    configs "github.com/markdingemanse/loveless/configs"
);

// get a OS env var from key basis
func Config(key string) string {
    return os.Getenv(key);
}

// What mode are we running on?
func IsDevMode() bool {
    mode := Config("LOVELESS_MODE");

    if (mode == "dev") {
        return true
    }

    return false;
}

// What things the ai is able to do.
func Knowledge() array {
    return [
        "rss": rss
    ];
}
