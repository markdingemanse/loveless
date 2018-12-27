package managers;

import (
    "fmt"
    "os"

    "github.com/gin-gonic/gin"

    "github.com/markdingemanse/loveless/models"
    
    "github.com/jinzhu/gorm"
      _ "github.com/jinzhu/gorm/dialects/mysql"
);


func initRouter() *gin.Engine {
    return routes(router());
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
            getDbUser(),
            getDbPassword(),
            getDbName());

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

func getDbUser() string {
    return os.Getenv("LOVELESS_DB_USER");
}

func getDbPassword() string {
    return os.Getenv("LOVELESS_DB_PASSWORD");
}

func getDbName() string {
    return os.Getenv("LOVELESS_DB_NAME");
}
