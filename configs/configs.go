package configs;

import (
    "os"
    "fmt"
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

// print a like of info to the lain log.
func Info(line string) {
    f, err := os.OpenFile("../logs/lain.log", os.O_APPEND|os.O_WRONLY, 0644);

    if (err != nil) {
        fmt.Println(err);
        return
    }

    _, err = fmt.Fprintln(f, line);

    if (err != nil) {
        fmt.Println(err)
        f.Close()
        return
    }

    err = f.Close();

    if (err != nil) {
        fmt.Println(err);
        return
    }
}
