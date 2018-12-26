package models;

import (
    "github.com/jinzhu/gorm"
);

type Rss struct {
  gorm.Model
  FirstPost   string
}
