package managers;

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/autotls"

    models "github.com/markdingemanse/loveless/models"

    "github.com/jinzhu/gorm"
      _ "github.com/jinzhu/gorm/dialects/mysql"
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

// Provisions the provided router with the needed routes.
func routes(router *gin.Engine) *gin.Engine {
    //basic ping test with a simple db check.
    router.GET("/", func(c *gin.Context) {
        dbURI := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
            Config("LOVELESS_DB_USER"),
            Config("LOVELESS_DB_PASSWORD"),
            Config("LOVELESS_DB_NAME"));

        db, err := gorm.Open("mysql", dbURI);

        if (err != nil) {
            fmt.Println("db error: ", err);
        }

        firstRss := models.Rss{};
        db.Table("rss").First(&firstRss);
        fmt.Printf("%v\n", "[TEST] The message of the first post seems to be: " + firstRss.FirstPost);

        c.String(200, "pong");
        db.Close();
    });

    return router;
}
