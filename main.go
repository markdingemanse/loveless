package main;

import (
        "log"
        "github.com/gin-gonic/autotls"
        "github.com/gin-gonic/gin"
);

func main() {
        r := gin.Default();

        // Ping handler
        r.GET("/", func(c *gin.Context) {
                c.String(200, "pong");
        })

        log.Fatal(autotls.Run(r, "lovelesswired.com", "www.lovelesswired.com"));
}
