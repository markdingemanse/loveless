package models;

import (
    "github.com/jinzhu/gorm"
);

type Rss struct {
  gorm.Model
  FirstPost   string
}

// get the first_post from the model Rss
func (rss Rss) GetFirstPost() string {
    return rss.FirstPost
}
