package managers;

import (
    "os"
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
