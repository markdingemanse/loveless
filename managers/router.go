package managers;

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/autotls"

    // behaviourModels "github.com/markdingemanse/loveless/models/behaviour"
    configs "github.com/markdingemanse/loveless/configs"
    handlers "github.com/markdingemanse/loveless/services"
    helpers "github.com/markdingemanse/loveless/services/helpers"
    knowledgeModels "github.com/markdingemanse/loveless/models/knowledge"

    // "github.com/jinzhu/gorm"
    //   _ "github.com/jinzhu/gorm/dialects/mysql"
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
    //basic ping test with a simple db check.
    router.GET("/", func(c *gin.Context) {
        uri := helpers.DatabaseUrl();
        db := helpers.OpenDbConnection(uri);

        model := knowledgeModels.CreateModel();
        helpers.SelectTable("rss", db).First(&model);
        fmt.Printf("%v\n", "[TEST] The message of the first post seems to be: " + model.GetFirstPost());

        f := handlers.Functions("rss");
        f();

        c.String(200, "pong");
        db.Close();
    });

    return router;
}
