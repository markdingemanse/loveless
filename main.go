package main;

import (
    "fmt"
    "os"
    "github.com/gin-gonic/autotls"
    "github.com/gin-gonic/gin"
);

func main() {
        router := routes(router());
        mode   := getAppMode();

        fmt.Println("[MAIN] Running Gin in mode: " + mode);

        if (mode == "dev") {
            router.Run();
        } else {
            // log.Fatal(autotls.Run(r, "lovelesswired.com", "www.lovelesswired.com"));
            autotls.Run(router, "lovelesswired.com", "www.lovelesswired.com");
        }
}

// Bootstraps the basic gin router.
func router() *gin.Engine {
    return gin.Default();
}

// Provisions the provided router with the needed routes.
func routes(router *gin.Engine) *gin.Engine {
    router.GET("/", func(c *gin.Context) {
            c.String(200, "pong");
    });

    return router;
}

// Get the current app mode we are running under.
func getAppMode() string {
    return os.Getenv("LOVELESS_MODE");
}
