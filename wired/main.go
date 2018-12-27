package main;

import (
    managers "github.com/markdingemanse/loveless/managers"
);

func main() {
        router := managers.InitRouter();
        dev   := managers.IsDevMode();
        managers.StartRouter(router, dev);
}
