package main;

import (
    "fmt"
    "os"
    "github.com/gin-gonic/autotls"
    "github.com/gin-gonic/gin"

    "github.com/jinzhu/gorm"
      _ "github.com/jinzhu/gorm/dialects/mysql"
);

type Rss struct {
  gorm.Model
  FirstPost   string
}

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
    //basic ping test with a simple db check.
    router.GET("/", func(c *gin.Context) {
        dbURI := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
            getDbUser(),
            getDbPassword(),
            getDbName());

        // db, err := gorm.Open("mysql", "homestead:secret@/homestead?charset=utf8&parseTime=True&loc=Local");
        db, err := gorm.Open("mysql", dbURI);

        if (err != nil) {
            fmt.Println("db error: ", err);
        }

        firstRss := Rss{};
        db.Table("rss").First(&firstRss);
        fmt.Printf("%v\n", "[TEST] The message of the first post seems to be: " + firstRss.FirstPost);

        c.String(200, "pong");
        db.Close();
    });

    return router;
}

// Get the current app mode we are running under.
func getAppMode() string {
    return os.Getenv("LOVELESS_MODE");
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
