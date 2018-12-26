package main;

import (
    "fmt"
    "os"
    "github.com/gin-gonic/autotls"
    "github.com/gin-gonic/gin"
);

func main() {
        r := gin.Default();

        // Ping handler
        r.GET("/", func(c *gin.Context) {
                c.String(200, "pong");
        })

        var mode string;
        mode = os.Getenv("LOVELESS_MODE");

        fmt.Println("[MAIN] Running Gin in mode: " + mode);

        if (mode == "dev") {
            r.Run();
        } else {
            // log.Fatal(autotls.Run(r, "lovelesswired.com", "www.lovelesswired.com"));
            autotls.Run(r, "lovelesswired.com", "www.lovelesswired.com");
        }
}
