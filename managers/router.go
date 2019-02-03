package managers;

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/autotls"

    configs "github.com/markdingemanse/loveless/configs"
    handlers "github.com/markdingemanse/loveless/services"
);

// create the app engine and provision the created router.
func InitRouter() *gin.Engine {
    return routes(router());
}

// start the router for the correct situation
func StartRouter(router *gin.Engine, dev bool) {
    if (dev) {
        fmt.Println("[MAIN] Running Gin in dev mode");
        router.Run();
    } else {
        fmt.Println("[MAIN] Running Gin in live mode with ssl");
        autotls.Run(router, "lovelesswired.com", "www.lovelesswired.com");
    }
}

// Bootstraps the basic gin router.
func router() *gin.Engine {
    return gin.Default();
}

// What mode are we running on? Proxy function to prevent dependency circle
func IsDevMode() bool {
    return configs.IsDevMode();
}

// Provisions the provided router with the needed routes.
func routes(router *gin.Engine) *gin.Engine {
    router.GET("/", func(c *gin.Context) {
        c.String(200, "I love you");
    });

    authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
       "love": configs.Secret(),
    }));

    authorized.GET("/question", func(c *gin.Context) {
        ResetLog(configs.Config("LOG_PATH"));

        f := handlers.Functions("rss");
        f();

        c.JSON(200, ReadLog(configs.Config("LOG_PATH")));
    });

    return router;
}
