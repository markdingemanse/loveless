package managers;

import (
    "fmt"

    configs "github.com/markdingemanse/loveless/configs"

    "github.com/jinzhu/gorm"
      _ "github.com/jinzhu/gorm/dialects/mysql"
);

// generate the mysql db url for gorm.
func DatabaseUrl() string {
    return fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
        configs.Config("LOVELESS_DB_USER"),
        configs.Config("LOVELESS_DB_PASSWORD"),
        configs.Config("LOVELESS_DB_NAME"));
}

// open up a connection for mysql database
func OpenDbConnection(uri string) *gorm.DB {
    db, err := gorm.Open("mysql", uri);

    if (err != nil) {
        fmt.Println("db error: ", err);
    }

    return db;
}

// select a table to perform actions upon
func SelectTable(table string, db *gorm.DB) *gorm.DB {
    return db.Table(table);
}
