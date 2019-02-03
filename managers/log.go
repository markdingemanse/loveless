package managers;

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
);

func ResetLog(log string) {
    err := os.Remove("../logs/lain.log");

    if err != nil {
        fmt.Println("[MAIN] error occured when deleting log: ", err.Error);
    }

    emptyFile, err := os.Create("../logs/lain.log");

    if err != nil {
        fmt.Println("[MAIN] error occured when creating log: ", err.Error);
    }
    emptyFile.Close();
}

func ReadLog(log string) []string {
    content, err := ioutil.ReadFile("../logs/lain.log");

    if err != nil {
        fmt.Println("[MAIN] error occured when reading log: ", err.Error);
    }

    return strings.Split(string(content), "\n");
}
